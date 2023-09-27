package transport

import (
	"citasapp/internal"
	"net/http"

	"github.com/gorilla/mux"
)

func HttpRouter(app *internal.CitasApp) *mux.Router {
	router := mux.NewRouter()

	router.Handle("/user", createUserHandler(app)).Methods(http.MethodPost)
	router.Handle("/user", listUsersHandler(app)).Methods(http.MethodGet)
	router.Handle("/user/{id}", getUserHandler(app)).Methods(http.MethodGet)

	return router
}
