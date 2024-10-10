package controller

import (
	"encoding/json"
	"io"
	"net/http"

	"br.com.cleiton/events/internal/adapters/input/controller/request"
	"br.com.cleiton/events/internal/domain/entities"
	"br.com.cleiton/events/internal/domain/usecases"
	"github.com/labstack/echo/v4"
)

type eventController struct {
	createEventUC usecases.CreateEventInterface
}

func NewEventController(createEventUC usecases.CreateEventInterface) eventController {
	return eventController{createEventUC: createEventUC}
}

func (e *eventController) CreateEvent(c echo.Context) error {
	body, err := io.ReadAll(c.Request().Body)
	if err != nil {
		return err
	}

	var eventRequest request.EventRequest
	err = json.Unmarshal(body, &eventRequest)
	if err != nil {
		return err
	}

	event := entities.Event{
		IdUUID:      eventRequest.UUID,
		WalletUUID:  eventRequest.WalletUUID,
		Type:        eventRequest.Type,
		Value:       eventRequest.Value,
		Description: eventRequest.Description,
		EventUUID:   eventRequest.EventId,
	}

	err = e.createEventUC.CreateEvent(event)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusCreated, "evento criado com sucesso")
}
