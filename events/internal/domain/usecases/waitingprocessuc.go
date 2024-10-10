package usecases

import (
	"errors"
	"fmt"

	"br.com.cleiton/events/internal/adapters/output/database"
	"br.com.cleiton/events/internal/domain/entities"
	"br.com.cleiton/events/internal/domain/usecases/processevent"
)

type waitingProcessUC struct {
	eventDatabase   database.EventDatabaseInterface
	walletDatabase  database.WalletDatabaseInterface
	extractDatabase database.ExtractDatabaseInterface
}

var (
	eventProcessMap = make(map[entities.Type]processevent.IStatusEventProcess)
)

func NewWaitingProcessUC(eventDatabase database.EventDatabaseInterface,
	walletDatabase database.WalletDatabaseInterface,
	extractDatabase database.ExtractDatabaseInterface) *waitingProcessUC {
	eventProcessMap[entities.ADD] = processevent.NewAddEvent(eventDatabase, walletDatabase, extractDatabase)
	eventProcessMap[entities.CANCELLATION] = processevent.NewCancellationEvent(eventDatabase, walletDatabase, extractDatabase)
	eventProcessMap[entities.REVERSAL] = processevent.NewReversalEvent(eventDatabase, walletDatabase, extractDatabase)
	eventProcessMap[entities.SHOPPING] = processevent.NewShoppingEvent(eventDatabase, walletDatabase, extractDatabase)
	eventProcessMap[entities.WITHDRAWAL] = processevent.NewWithdrawalEvent(eventDatabase, walletDatabase, extractDatabase)

	return &waitingProcessUC{eventDatabase: eventDatabase, walletDatabase: walletDatabase}
}

func (w *waitingProcessUC) ProcessEvent(event *entities.Event) error {

	method, exists := eventProcessMap[entities.Type(event.Type)]
	if !exists {
		message := fmt.Sprintf("tipo do evento é inválido, id %d ", event.Id)
		return errors.New(message)
	}

	err := method.ProcessEvent(*event)
	if err != nil {
		return err
	}

	return nil
}
