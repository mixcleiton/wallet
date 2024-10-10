package usecases

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"time"

	"br.com.cleiton/events/internal/adapters/input/kafkainterfaces"
	"br.com.cleiton/events/internal/adapters/output/database"
	"br.com.cleiton/events/internal/domain/entities"
)

type CreateEventInterface interface {
	CreateEvent(event entities.Event) error
}

type createEventUC struct {
	eventDatabase  database.EventDatabaseInterface
	walletDatabase database.WalletDatabaseInterface
	kafkaProducer  kafkainterfaces.KafkaProducerInterface
}

func NewCreateEventUC(eventDatabase database.EventDatabaseInterface,
	kafkaProducer kafkainterfaces.KafkaProducerInterface,
	walletDatabase database.WalletDatabaseInterface) createEventUC {
	return createEventUC{eventDatabase: eventDatabase, kafkaProducer: kafkaProducer, walletDatabase: walletDatabase}
}

func (c *createEventUC) CreateEvent(event entities.Event) error {

	if event.IdUUID == "" {
		return errors.New("identificação da carteira é obrigatório")
	}

	if event.WalletUUID == "" {
		return errors.New("identificação da carteira é obrigatório")
	}

	if event.Value <= 0 {
		return errors.New("valor do evento não poder ser menor ou igual a 0")
	}

	if event.Type <= 0 {
		return errors.New("valor para o tipo de evento é inválido")
	}

	if event.Type == int(entities.CANCELLATION) || event.Type == int(entities.REVERSAL) {
		originEvent, err := c.eventDatabase.GetEventByUUID(event.EventUUID)
		if err != nil {
			return err
		}

		event.EventID = originEvent.Id
	}

	wallet, err := c.walletDatabase.GetWalletByUUID(event.WalletUUID)
	if err != nil {
		return err
	}

	event.WalletId = wallet.Id
	event.CreateAt = time.Now()
	event.Status = 1

	log.Println("eventId: ", event.EventID)

	if event.EventID > 0 {
		err := c.eventDatabase.CreateEvent(event)
		if err != nil {
			return fmt.Errorf("erro ao salvar evento, erro %w", err)
		}
	} else {
		err := c.eventDatabase.CreateEventWithoutEventId(event)
		if err != nil {
			return fmt.Errorf("erro ao salvar evento, erro %w", err)
		}
	}

	body, err := json.Marshal(event)
	if err != nil {
		return fmt.Errorf("erro ao converter para json o evento, erro %w", err)
	}

	c.kafkaProducer.Producer("event-process", string(body))
	return nil
}
