package processevent

import (
	"fmt"

	"br.com.cleiton/events/internal/adapters/output/database"
	"br.com.cleiton/events/internal/domain/entities"
)

type reversalEvent struct {
	eventDatabase  database.EventDatabaseInterface
	walletDatabase database.WalletDatabaseInterface
}

func NewReversalEvent(eventDatabase database.EventDatabaseInterface,
	walletDatabase database.WalletDatabaseInterface) *reversalEvent {
	return &reversalEvent{eventDatabase: eventDatabase, walletDatabase: walletDatabase}
}

func (r *reversalEvent) ProcessEvent(event entities.Event) error {
	tx, err := r.eventDatabase.GetDB().Begin()
	if err != nil {
		return fmt.Errorf("erro ao abrir transação no banco de dados, %w", err)
	}

	newValue, err := r.walletDatabase.UpdateWalletValueById(tx, event.Value, event.WalletId)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("erro ao realizar atualização da carteira, erro %w", err)
	}

	if newValue < 0 {
		tx.Rollback()
		return fmt.Errorf("dando rollback na transação, pois a carteira %d, ficaria com valor negativo", event.WalletId)
	}

	err = r.eventDatabase.UpdateEventStatusByID(tx, event.Id, int(entities.PROCESSED))
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("erro ao alterar o status do event para processado. erro %w", err)
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("erro ao realizar o commit da operação, erro %w", err)
	}

	return nil
}
