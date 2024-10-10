package processeventmocks

import (
	"br.com.cleiton/events/internal/domain/entities"
	"github.com/stretchr/testify/mock"
)

type MockStatusEventProcess struct {
	mock.Mock
}

func (m *MockStatusEventProcess) ProcessEvent(event *entities.Event) error {
	args := m.Called(event)
	return args.Error(0)
}
