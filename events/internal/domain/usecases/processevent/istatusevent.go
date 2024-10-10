package processevent

import "br.com.cleiton/events/internal/domain/entities"

type IStatusEventProcess interface {
	ProcessEvent(event entities.Event) error
}
