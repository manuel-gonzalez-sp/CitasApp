package persistence

import (
	"citasapp/internal/infra/utils"

	"github.com/upper/db/v4"
	"github.com/upper/db/v4/adapter/sqlite"
)

var settings = sqlite.ConnectionURL{
	Database: `citasapp.db`,
}

var Session = defaultSession()

func defaultSession() db.Session {
	sess, err := sqlite.Open(settings)
	if err != nil {
		utils.Logger.Fatalf("db.Open(): %q\n", err)
	}
	defer sess.Close()
	return sess
}
