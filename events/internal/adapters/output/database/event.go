package database

import (
	"database/sql"

	"br.com.cleiton/events/internal/domain/entities"
)

type EventDatabaseInterface interface {
	CreateEvent(event entities.Event) error
	GetEventByUUID(eventUUID string) (*entities.Event, error)
	GetEventByID(eventId int) (*entities.Event, error)
	GetDB() *sql.DB
	UpdateEventStatusByID(tx *sql.Tx, id int, status int) error
	GetTxEventByID(tx *sql.Tx, eventId int) (*entities.Event, error)
}

type eventDatabase struct {
	db *sql.DB
}

func NewEventDatabase(db *sql.DB) eventDatabase {
	return eventDatabase{db: db}
}

func (e *eventDatabase) GetDB() *sql.DB {
	return e.db
}

func (e *eventDatabase) UpdateEventStatusByID(tx *sql.Tx, id int, status int) error {
	_, err := tx.Exec("UPDATE event e SET e.status_id = $1 WHERE e.id = $2 ", status, id)
	if err != nil {
		return err
	}

	return err
}

func (e *eventDatabase) CreateEvent(event entities.Event) error {
	_, err := e.db.Exec("INSERT INTO event(id_uuid, status_id, value, created_at, type_id, wallet_id, description, event_id) "+
		"VALUES ($1, $2, $3, $4, $5, $6, $7)", event.IdUUID, event.Status, event.Value, event.CreateAt, event.Type, event.WalletId, event.Description, event.EventID)
	if err != nil {
		return err
	}

	return nil
}

func (e *eventDatabase) GetEventByUUID(eventUUID string) (*entities.Event, error) {
	row := e.db.QueryRow("SELECT e.id, e.wallet_id, e.created_at, e.updated_at, e.status_id, e.type_id, e.description, e.id_uuid, e.value, e.event_id FROM event e WHERE e.id_uuid = $1", eventUUID)

	var event entities.Event
	err := row.Scan(&event.Id, &event.WalletId, &event.CreateAt, &event.UpdatedAt, &event.Status, &event.Type, &event.Description, &event.IdUUID, &event.Value, &event.EventID)
	if err != nil {
		return nil, err
	}

	return &event, err
}

func (e *eventDatabase) GetEventByID(eventId int) (*entities.Event, error) {
	row := e.db.QueryRow("SELECT e.id, e.wallet_id, e.created_at, e.updated_at, e.status_id, e.type_id, e.description, e.id_uuid, e.value, e.event_id FROM event e WHERE e.id = $1", eventId)

	var event entities.Event
	err := row.Scan(&event.Id, &event.WalletId, &event.CreateAt, &event.UpdatedAt, &event.Status, &event.Type, &event.Description, &event.IdUUID, &event.Value, &event.EventID)
	if err != nil {
		return nil, err
	}

	return &event, err
}

func (e *eventDatabase) GetTxEventByID(tx *sql.Tx, eventId int) (*entities.Event, error) {
	row := e.db.QueryRow("SELECT e.id, e.wallet_id, e.created_at, e.updated_at, e.status_id, e.type_id, e.description, e.id_uuid, e.value, e.event_id FROM event e WHERE e.id = $1", eventId)

	var event entities.Event
	err := row.Scan(&event.Id, &event.WalletId, &event.CreateAt, &event.UpdatedAt, &event.Status, &event.Type, &event.Description, &event.IdUUID, &event.Value, &event.EventID)
	if err != nil {
		return nil, err
	}

	return &event, err
}
