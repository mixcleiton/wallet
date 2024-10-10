package processevent

import (
	"fmt"

	"br.com.cleiton/events/internal/adapters/output/database"
	"br.com.cleiton/events/internal/domain/entities"
)

type cancellationEvent struct {
	eventDatabase   database.EventDatabaseInterface
	walletDatabase  database.WalletDatabaseInterface
	extractDatabase database.ExtractDatabaseInterface
}

func NewCancellationEvent(eventDatabase database.EventDatabaseInterface,
	walletDatabase database.WalletDatabaseInterface,
	extractDatabase database.ExtractDatabaseInterface) *cancellationEvent {
	return &cancellationEvent{eventDatabase: eventDatabase, walletDatabase: walletDatabase, extractDatabase: extractDatabase}
}

func (c *cancellationEvent) ProcessEvent(event entities.Event) error {
	tx, err := c.eventDatabase.GetDB().Begin()
	if err != nil {
		return fmt.Errorf("erro ao abrir transação no banco de dados, %w", err)
	}

	originEvent, err := c.eventDatabase.GetTxEventByID(tx, event.EventID)
	if err != nil {
		return fmt.Errorf("erro ao pegar o evento original que o cancelamento atual afetaria, erro %w", originEvent)
	}

	if originEvent.Status != int(entities.WAITING_PROCESS) {
		tx.Rollback()
		return fmt.Errorf("o evento original não está mais pendente ou em andamento e por isso não pode ser cancelado")
	}

	eventValue := originEvent.Value
	if originEvent.Type != int(entities.ADD) {
		eventValue = eventValue * -1
	}

	newValue, err := c.walletDatabase.UpdateWalletValueById(tx, eventValue, event.WalletId)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("erro ao realizar atualização da carteira, erro %w", err)
	}

	if newValue < 0 {
		tx.Rollback()
		return fmt.Errorf("dando rollback na transação, pois a carteira %d, ficaria com valor negativo", event.WalletId)
	}

	err = c.eventDatabase.UpdateEventStatusByID(tx, event.Id, int(entities.PROCESSED))
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("erro ao alterar o status do event para processado. erro %w", err)
	}

	err = c.extractDatabase.CreateExtract(tx, event.Id)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("erro ao criar extrato, %w", err)
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("erro ao realizar o commit da operação, erro %w", err)
	}

	return nil
}
