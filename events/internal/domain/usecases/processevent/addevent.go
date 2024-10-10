package processevent

import (
	"fmt"

	"br.com.cleiton/events/internal/adapters/output/database"
	"br.com.cleiton/events/internal/domain/entities"
)

type addEvent struct {
	eventDatabase   database.EventDatabaseInterface
	walletDatabase  database.WalletDatabaseInterface
	extractDatabase database.ExtractDatabaseInterface
}

func NewAddEvent(eventDatabase database.EventDatabaseInterface,
	walletDatabase database.WalletDatabaseInterface,
	extractDatabase database.ExtractDatabaseInterface) *addEvent {
	return &addEvent{eventDatabase: eventDatabase, walletDatabase: walletDatabase, extractDatabase: extractDatabase}
}

func (a *addEvent) ProcessEvent(event entities.Event) error {

	tx, err := a.eventDatabase.GetDB().Begin()
	if err != nil {
		return fmt.Errorf("erro ao abrir transação no banco de dados, %w", err)
	}

	newValue, err := a.walletDatabase.UpdateWalletValueById(tx, event.Value, event.WalletId)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("erro ao realizar atualização da carteira, erro %w", err)
	}

	if newValue < 0 {
		tx.Rollback()
		return fmt.Errorf("dando rollback na transação, pois a carteira %d, ficaria com valor negativo", event.WalletId)
	}

	err = a.eventDatabase.UpdateEventStatusByID(tx, event.Id, int(entities.PROCESSED))
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("erro ao alterar o status do event para processado. erro %w", err)
	}

	err = a.extractDatabase.CreateExtract(tx, event.Id)
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
