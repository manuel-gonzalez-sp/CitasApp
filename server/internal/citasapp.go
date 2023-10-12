package internal

import (
	"citasapp/internal/application/command"
	"citasapp/internal/infra/persistence"
	"net/http"
)

type CitasApp struct {
	command.Commander
}

func NewCitasApp() (*CitasApp, error) {
	userRepo := persistence.NewSQLiteUserRepository(persistence.Database)
	messageRepo := persistence.NewSQLiteMessageRepository(persistence.Database)

	app := &CitasApp{
		Commander: &command.DefaultCommander{
			UserRepo:    userRepo,
			MessageRepo: messageRepo,
		},
	}
	return app, nil
}

func (app *CitasApp) Run(router http.Handler) error {
	http.ListenAndServe(":3000", router)
	return nil
}

// JWT Auth
// Messages
// Members
// Lists

// Routes
// "": HomeComponent
// members : MemberList
// members/:id : MemberDetail
// lists : ListsComponent
// messages : MessagesComponent
// ** : HomeComponent, pathMatch: full
