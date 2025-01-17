package rest

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
//	"time"
	
	"github.com/gorilla/mux"
	
	"contracts"
	"persistence"
	"message_queue"

)

type eventServiceHandler struct {
	dbhandler persistence.DatabaseHandler
	eventEmitter message_queue.EventEmitter
}

func NewEventHandler(databasehandler persistence.DatabaseHandler, eventEmitter message_queue.EventEmitter) *eventServiceHandler {
	return &eventServiceHandler{
		dbhandler: databasehandler,
		eventEmitter: eventEmitter,
	}
}

func (eh *eventServiceHandler) FindEventHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	criteria, ok := vars["SearchCriteria"]
	if !ok {
		w.WriteHeader(400)
		fmt.Fprint(w, `{"error": "No search criteria found, you can either
						search by id via /id/4
						to search by name via /name/coldplayconcert}"`)
		return
	}

	searchkey, ok := vars["search"]
	if !ok {
		w.WriteHeader(400)
		fmt.Fprint(w, `{"error": "No search keys found, you can either search
						by id via /id/4
						to search by name via /name/coldplayconcert}"`)
		return
	}

	var event persistence.Event
	var err error
	switch strings.ToLower(criteria) {
	case "name":
		event, err = eh.dbhandler.FindEventByName(searchkey)
	case "id":
		id, err := hex.DecodeString(searchkey)
		if err == nil {
			event, err = eh.dbhandler.FindEvent(id)
		}
	}
	if err != nil {
		fmt.Fprintf(w, `{"error": "%s"}`, err)
		return
	}
	w.Header().Set("Content-Type", "application/json;charset=utf8")
	json.NewEncoder(w).Encode(&event)
}


func (eh *eventServiceHandler) AllEventHandler(w http.ResponseWriter, r *http.Request) {
	events, err := eh.dbhandler.FindAllAvailableEvents()
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, `{"error": "Error occured while trying to find all available events %s"}`, err)
		return
	}
	w.Header().Set("Content-Type", "application/json;charset=utf8")
	err = json.NewEncoder(w).Encode(&events)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, `{"error": "Error occured while trying encode events to JSON %s"}`, err)
	}
}

func (eh *eventServiceHandler) NewEventHandler(w http.ResponseWriter, r *http.Request) {
	event := persistence.Event{}
	err := json.NewDecoder(r.Body).Decode(&event)
	if nil != err {
		w.WriteHeader(500)
		fmt.Fprintf(w, "error occured while decoding event data %s", err)
		return
	}
	id, err := eh.dbhandler.AddEvent(event)
	if nil != err {
		w.WriteHeader(500)
		fmt.Fprintf(w, "error occured while persisting event %s", err)
		return
	}

	msg := contracts.EventCreatedEvent{
		ID:         hex.EncodeToString(id),
		Name:       event.Name,
		//Start:      time.Unix(event.StartDate, 0),
		//End:        time.Unix(event.EndDate, 0),
		LocationID: string(event.Location.ID),
	}
	eh.eventEmitter.Emit(&msg)

	w.Header().Set("Content-Type", "application/json;charset=utf8")

	w.WriteHeader(201)
	json.NewEncoder(w).Encode(&event)
}