package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

const channelIDseparater = ", "

type Handler struct {
	settings Settings
	db       *sql.DB
}

func New(settings Settings) *Handler {
	return &Handler{
		settings: settings,
	}
}
