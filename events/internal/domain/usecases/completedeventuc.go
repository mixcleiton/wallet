package usecases

import (
	"fmt"
	"log"

	"br.com.cleiton/events/internal/adapters/output/database"
	"br.com.cleiton/events/internal/domain/entities"
)

type completedEventUC struct {
	eventDatabase  database.EventDatabaseInterface
	walletDatabase database.WalletDatabaseInterface
}

func NewCompletedEventUC(eventDatabase database.EventDatabaseInterface,
	walletDatabase database.WalletDatabaseInterface) *completedEventUC {
	return &completedEventUC{
		eventDatabase: eventDatabase, walletDatabase: walletDatabase,
	}
}

func (c *completedEventUC) ProcessEvent(event *entities.Event) error {

	/*Poderiamos realizar outras operações ao completar o processamento de um evento
	exemplo enviar um email ao dono da carteira.
	*/
	tx, err := c.eventDatabase.GetDB().Begin()
	if err != nil {
		return fmt.Errorf("erro ao abrir transação no banco, erro %w", err)
	}

	log.Println("passou aqui", event.IdUUID)
	eventToStatus, err := c.eventDatabase.GetEventByUUID(event.IdUUID)
	if err != nil {
		return fmt.Errorf("erro para recuperar o evento pelo UUID")
	}

	log.Println("passou aqui", eventToStatus.Id)
	err = c.eventDatabase.UpdateEventStatusByID(tx, eventToStatus.Id, int(entities.COMPLETED))
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("erro para atualizar o status do evento id %d, erro %w", event.Id, err)
	}

	tx.Commit()
	return nil
}
