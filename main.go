package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"sse-play/events"
	"sse-play/sse"
)

func main() {
	router := chi.NewRouter()

	sseService := sse.NewSSEService()
	er := events.NewEventRouter(sseService)

	router.Mount("/events", er.Routes())

	http.ListenAndServe(":8080", router)
}
