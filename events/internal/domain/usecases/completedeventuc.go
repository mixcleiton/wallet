package usecases

import (
	"fmt"

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

	err = c.eventDatabase.UpdateEventStatusByID(tx, event.Id, int(entities.COMPLETED))
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("erro para atualizar o status do evento id %d, erro %w", event.Id, err)
	}

	return nil
}
