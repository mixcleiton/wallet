package databasemocks

import (
	// Se a sua interface utilizar contexto
	"database/sql"

	"br.com.cleiton/events/internal/domain/entities"
	"github.com/stretchr/testify/mock"
)

// MockEventDatabase é a estrutura que implementa o mock
type MockEventDatabase struct {
	mock.Mock
}

// GetEventByUUID é a implementação do mock para o método correspondente
func (m *MockEventDatabase) GetEventByUUID(uuid string) (*entities.Event, error) {
	args := m.Called(uuid)
	return args.Get(0).(*entities.Event), args.Error(1)
}

// UpdateEventStatusByID é a implementação do mock para o método correspondente
func (m *MockEventDatabase) UpdateEventStatusByID(tx *sql.Tx, id int, status int) error {
	args := m.Called(tx, id, status)
	return args.Error(0)
}

func (m *MockEventDatabase) CreateEvent(event entities.Event) error {
	args := m.Called(event)
	return args.Error(0)
}

func (m *MockEventDatabase) GetEventByID(eventId int) (*entities.Event, error) {
	args := m.Called(eventId)
	return args.Get(0).(*entities.Event), args.Error(1)
}

func (m *MockEventDatabase) GetDB() *sql.DB {
	args := m.Called()
	return args.Get(0).(*sql.DB)
}

func (m *MockEventDatabase) GetTxEventByID(tx *sql.Tx, eventId int) (*entities.Event, error) {
	args := m.Called(tx, eventId)
	return args.Get(0).(*entities.Event), args.Error(1)
}

func (m *MockEventDatabase) CreateEventWithoutEventId(event entities.Event) error {
	args := m.Called(event)
	return args.Error(0)
}
