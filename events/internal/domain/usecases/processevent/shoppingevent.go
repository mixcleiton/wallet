package processevent

import (
	"fmt"

	"br.com.cleiton/events/internal/adapters/output/database"
	"br.com.cleiton/events/internal/domain/entities"
)

type shoppingEvent struct {
	eventDatabase  database.EventDatabaseInterface
	walletDatabase database.WalletDatabaseInterface
}

func NewShoppingEvent(eventDatabase database.EventDatabaseInterface,
	walletDatabase database.WalletDatabaseInterface) *shoppingEvent {
	return &shoppingEvent{eventDatabase: eventDatabase, walletDatabase: walletDatabase}
}

func (s *shoppingEvent) ProcessEvent(event entities.Event) error {
	tx, err := s.eventDatabase.GetDB().Begin()
	if err != nil {
		return fmt.Errorf("erro ao abrir transação no banco de dados, %w", err)
	}

	eventValueNegative := event.Value * -1

	newValue, err := s.walletDatabase.UpdateWalletValueById(tx, eventValueNegative, event.WalletId)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("erro ao realizar atualização da carteira, erro %w", err)
	}

	if newValue < 0 {
		tx.Rollback()
		return fmt.Errorf("dando rollback na transação, pois a carteira %d, ficaria com valor negativo", event.WalletId)
	}

	err = s.eventDatabase.UpdateEventStatusByID(tx, event.Id, int(entities.PROCESSED))
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
