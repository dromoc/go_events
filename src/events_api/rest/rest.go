package rest

import (
	"net/http"
	"github.com/gorilla/mux"
    "persistence"
)

func ServeAPI(endpoint string, databasehandler persistence.DatabaseHandler) error {
    r := mux.NewRouter()

	return http.ListenAndServe(endpoint, r)
}
