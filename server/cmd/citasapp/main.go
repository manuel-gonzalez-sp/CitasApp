package main

import (
	"citasapp/internal"
	"citasapp/internal/infra/transport"
	"citasapp/internal/infra/utils"
)

func main() {
	app, err := internal.NewCitasApp()
	if err != nil {
		utils.Logger.Fatal("internal.NewCitasApp errored", "err", err)
	}

	router := transport.HttpRouter(app)

	if err := app.Run(router); err != nil {
		utils.Logger.Fatal("service.Run errored", "err", err)
	}
}
