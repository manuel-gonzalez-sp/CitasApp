package main

import (
	"citasapp/internal"
	"citasapp/internal/infra/transport"

	"github.com/rs/zerolog/log"
)

func main() {

	app, err := internal.NewCitasApp()
	if err != nil {
		log.Fatal().Msgf("internal.NewCitasApp errored", "err", err)
	}

	router := transport.HttpRouter(app)

	if err := app.Run(router); err != nil {
		log.Fatal().Msgf("service.Run errored", "err", err)
	}
}
