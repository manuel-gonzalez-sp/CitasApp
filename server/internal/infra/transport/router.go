package transport

import (
	"citasapp/internal"
	"citasapp/internal/infra/transport/handler"
	"citasapp/internal/infra/transport/middleware"
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

	router.Handle("/login",
		middleware.JWTAuthenticationMiddleware(handler.LogInHandler(app)),
	).Methods(http.MethodPost)

	router.Handle("/user", handler.CreateUserHandler(app)).Methods(http.MethodPost)
	router.Handle("/user", handler.ListUsersHandler(app)).Methods(http.MethodGet)
	router.Handle("/user/{id}", handler.GetUserHandler(app)).Methods(http.MethodGet)
	router.Handle("/user/{id}", handler.UpdateUserHandler(app)).Methods(http.MethodPatch)

	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	return router
}
