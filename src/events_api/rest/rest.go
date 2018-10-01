package rest

import (
	"net/http"
	"github.com/gorilla/mux"
    "github.com/gorilla/handlers"
    "persistence"
    "message_queue"
)

func ServeAPI(endpoint string, databasehandler persistence.DatabaseHandler, eventEmitter message_queue.EventEmitter) error {
    handler := NewEventHandler(databasehandler, eventEmitter)
    r := mux.NewRouter()
    eventsrouter := r.PathPrefix("/events").Subrouter()

    eventsrouter.Methods("GET").Path("/{SearchCriteria}/{search}").HandlerFunc(handler.FindEventHandler)
    eventsrouter.Methods("GET").Path("").HandlerFunc(handler.AllEventHandler)
    eventsrouter.Methods("POST").Path("").HandlerFunc(handler.NewEventHandler)

    server := handlers.CORS()(r)
	return http.ListenAndServe(endpoint, server)
}
