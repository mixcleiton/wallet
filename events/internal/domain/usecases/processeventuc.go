package usecases

import (
	"encoding/json"
	"fmt"
	"log"

	"br.com.cleiton/events/internal/adapters/input/kafkainterfaces"
	"br.com.cleiton/events/internal/adapters/output/database"
	"br.com.cleiton/events/internal/domain/entities"
)

var (
	statusEventProcessMap = make(map[entities.Status]StatusEventProcess)
)

type ProcessEvent interface {
	ProcessEvent(event *entities.Event) error
}

type StatusEventProcess interface {
	ProcessEvent(event *entities.Event) error
}

type processEventUC struct {
	eventDatabase  database.EventDatabaseInterface
	walletDatabase database.WalletDatabaseInterface
	kafkaProducer  kafkainterfaces.KafkaProducerInterface
}

func NewProcessEventUC(eventDatabase database.EventDatabaseInterface, walletDatabase database.WalletDatabaseInterface,
	kafkaProducer kafkainterfaces.KafkaProducerInterface) processEventUC {
	statusEventProcessMap[entities.WAITING_PROCESS] = NewWaitingProcessUC(eventDatabase, walletDatabase)
	statusEventProcessMap[entities.PROCESSED] = NewCompletedEventUC(eventDatabase, walletDatabase)

	return processEventUC{eventDatabase: eventDatabase, walletDatabase: walletDatabase, kafkaProducer: kafkaProducer}
}

func (p *processEventUC) ProcessEvent(event *entities.Event) error {
	event, err := p.eventDatabase.GetEventByUUID(event.IdUUID)
	if err != nil {
		log.Printf("não foi possível buscar o evento pelo id %d", &event.IdUUID)
		return err
	}

	processEvent, exists := statusEventProcessMap[entities.Status(event.Status)]
	if !exists {
		log.Printf("status do event para processar é inválido, %d", event.Id)
		return err
	}

	err = processEvent.ProcessEvent(event)
	if err != nil {
		log.Printf("não foi possível processar o event de id %d, erro: %w", event.EventID, err)
		return err
	}

	if event.Status != int(entities.COMPLETED) || event.Status != int(entities.CANCELED) {
		body, err := json.Marshal(event)
		if err != nil {
			return fmt.Errorf("erro ao converter para json o evento, erro %w", err)
		}

		p.kafkaProducer.Producer("event-process", string(body))
	}

	return nil
}
