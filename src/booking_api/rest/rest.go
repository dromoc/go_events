package rest

import (
	"net/http"
	"time"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	
	"message_queue"
	"persistence"
)

func ServeAPI(listenAddr string, database persistence.DatabaseHandler, eventEmitter message_queue.EventEmitter) {
	r := mux.NewRouter()
	r.Methods("post").Path("/events/{eventID}/bookings").Handler(&CreateBookingHandler{eventEmitter, database})

	srv := http.Server{
		Handler:      handlers.CORS()(r),
		Addr:         listenAddr,
		WriteTimeout: 2 * time.Second,
		ReadTimeout:  1 * time.Second,
	}

	srv.ListenAndServe()
}