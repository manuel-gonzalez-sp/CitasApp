package internal

import (
	"citasapp/internal/application/command"
	"citasapp/internal/infra/persistence"
	"net/http"
)

type CitasApp struct {
	Command command.CommandHandler
}

func NewCitasApp() (*CitasApp, error) {

	userRepo := persistence.NewSQLiteUserRepository()

	command := &command.DefaultCommandHandler{
		UserRepo: userRepo,
	}

	app := &CitasApp{
		Command: command,
	}

	return app, nil
}

func (app *CitasApp) Run(router http.Handler) error {
	http.ListenAndServe(":3000", router)
	return nil
}
