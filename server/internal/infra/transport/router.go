package transport

import (
	"citasapp/internal"
	"net/http"

	_ "citasapp/docs"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title			CitasApp API
// @version		1.0
// @description	This is a HTTP API for CitasApp.
func HttpRouter(app *internal.CitasApp) *mux.Router {
	router := mux.NewRouter()

	router.Handle("/user", createUserHandler(app)).Methods(http.MethodPost)
	router.Handle("/user", listUsersHandler(app)).Methods(http.MethodGet)
	router.Handle("/user/{id}", getUserHandler(app)).Methods(http.MethodGet)
	router.Handle("/user/{id}", updateUserHandler(app)).Methods(http.MethodPatch)

	router.Handle("/message", listMessagesHandler(app)).Methods(http.MethodGet)
	router.Handle("/message/send", sendMessageHandler(app)).Methods(http.MethodPost)
	router.Handle("/message/{id}/read", readMessageHandler(app)).Methods(http.MethodPatch)

	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	return router
}
