package internal

import (
	"citasapp/internal/application/command"
	"citasapp/internal/infra/persistence"
	"net/http"

	"github.com/gorilla/handlers"
)

type CitasApp struct {
	command.Commander
}

func NewCitasApp() (*CitasApp, error) {
	userRepo := persistence.NewSQLiteUserRepository(persistence.Database)

	app := &CitasApp{
		Commander: &command.DefaultCommander{
			UserRepo: userRepo,
		},
	}
	return app, nil
}

func (app *CitasApp) Run(router http.Handler) error {
	corsOpts := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "X-Requested-With"}),
		handlers.AllowCredentials(),
	)
	http.ListenAndServe(":3000", corsOpts(router))
	return nil
}
