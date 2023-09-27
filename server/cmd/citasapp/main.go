package main

import (
	"citasapp/internal"
	"citasapp/internal/infra/transport"
	"citasapp/internal/infra/utils"
)

func main() {
	app, err := internal.NewCitasApp()
	if err != nil {
		utils.Logger.Fatalf("unable to create CitasApp: %v\n", err)
	}

	router := transport.HttpRouter(app)

	if err := app.Run(router); err != nil {
		utils.Logger.Fatalf("unable to Run CitasApp: %v\n", err)
	}
}
