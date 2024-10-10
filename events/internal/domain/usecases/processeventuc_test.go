package usecases

import (
	"testing"

	"br.com.cleiton/events/internal/adapters/input/kafkamessage/kafkamock"
	"br.com.cleiton/events/internal/adapters/output/database/databasemocks"
	"br.com.cleiton/events/internal/domain/entities"
	"br.com.cleiton/events/internal/domain/usecases/processevent/processeventmocks"
	"github.com/stretchr/testify/assert"
)

func TestProcessEvent_ValidEvent(t *testing.T) {
	// Mock interfaces
	mockEventDB := new(databasemocks.MockEventDatabase)
	mockProcess := new(processeventmocks.MockStatusEventProcess)
	mockKafkaProducer := new(kafkamock.MockKafkaProducer)

	// Expected event data
	expectedEvent := &entities.Event{
		IdUUID: "valid-uuid",
		Status: int(entities.WAITING_PROCESS),
	}

	mockEventDB.On("GetEventByUUID", expectedEvent.IdUUID).Return(expectedEvent, nil)

	mockProcess.On("ProcessEvent", expectedEvent).Return(nil)

	uc := processEventUC{
		eventDatabase:   mockEventDB,
		walletDatabase:  nil,               // Replace with mocks for other dependencies if needed
		kafkaProducer:   mockKafkaProducer, // You might not need a mock here depending on the test setup
		extractDatabase: nil,
	}

	statusEventProcessMap[entities.WAITING_PROCESS] = mockProcess

	err := uc.ProcessEvent(expectedEvent)

	assert.NoError(t, err, "Unexpected error during event processing")
	mockEventDB.AssertExpectations(t) // Verify mock interactions
	mockProcess.AssertExpectations(t) // Verify delegated process call
}
