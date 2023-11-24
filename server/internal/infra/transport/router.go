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
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func HttpRouter(app *internal.CitasApp) *mux.Router {
	router := mux.NewRouter()

	router.Handle("/login", handler.LogInHandler(app)).Methods(http.MethodPost)
	router.Handle("/signup", handler.SignUpHandler(app)).Methods(http.MethodPost)

	router.Handle("/user", middleware.JWTAuthenticationMiddleware(handler.CreateUserHandler(app))).Methods(http.MethodPost)
	router.Handle("/user", middleware.JWTAuthenticationMiddleware(handler.ListUsersHandler(app))).Methods(http.MethodGet)
	router.Handle("/user/{id}", middleware.JWTAuthenticationMiddleware(handler.GetUserHandler(app))).Methods(http.MethodGet)
	router.Handle("/user/{id}", middleware.JWTAuthenticationMiddleware(handler.UpdateUserHandler(app))).Methods(http.MethodPatch)

	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	return router
}
