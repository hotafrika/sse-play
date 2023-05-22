package events

import (
	"github.com/go-chi/chi/v5"

	"sse-play/sse"
)

type EventRouter struct {
	SSE sse.SSEService
}

func NewEventRouter(sse sse.SSEService) *EventRouter {
	return &EventRouter{SSE: sse}
}

func (er EventRouter) Routes() chi.Router {
	r := chi.NewRouter()

	r.Handle("/users/{userID}", er.SSE)

	return r
}
