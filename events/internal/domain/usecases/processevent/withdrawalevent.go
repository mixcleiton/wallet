package processevent

import (
	"fmt"

	"br.com.cleiton/events/internal/adapters/output/database"
	"br.com.cleiton/events/internal/domain/entities"
)

type withdrawalEvent struct {
	eventDatabase   database.EventDatabaseInterface
	walletDatabase  database.WalletDatabaseInterface
	extractDatabase database.ExtractDatabaseInterface
}

func NewWithdrawalEvent(eventDatabase database.EventDatabaseInterface,
	walletDatabase database.WalletDatabaseInterface,
	extractDatabase database.ExtractDatabaseInterface) *withdrawalEvent {
	return &withdrawalEvent{eventDatabase: eventDatabase, walletDatabase: walletDatabase, extractDatabase: extractDatabase}
}

func (w *withdrawalEvent) ProcessEvent(event entities.Event) error {
	tx, err := w.eventDatabase.GetDB().Begin()
	if err != nil {
		return fmt.Errorf("erro ao abrir transação no banco de dados, %w", err)
	}

	eventValueNegative := event.Value * -1
	newValue, err := w.walletDatabase.UpdateWalletValueById(tx, eventValueNegative, event.WalletId)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("erro ao realizar atualização da carteira, erro %w", err)
	}

	if newValue < 0 {
		tx.Rollback()
		return fmt.Errorf("dando rollback na transação, pois a carteira %d, ficaria com valor negativo", event.WalletId)
	}

	err = w.eventDatabase.UpdateEventStatusByID(tx, event.Id, int(entities.PROCESSED))
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("erro ao alterar o status do event para processado. erro %w", err)
	}

	err = w.extractDatabase.CreateExtract(tx, event.Id)
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
