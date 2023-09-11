package transport

import (
	"citasapp/internal"
	"net/http"
)

func getUserHandler(app *internal.CitasApp) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Workeanding again"))
	}
}
