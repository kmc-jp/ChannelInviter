package database

import (
	"database/sql"
	"fmt"
)

func (h *Handler) Initialize() error {
	if h.settings.Directory == "" {
		h.settings.Directory = "channels.db"
	}
	db, err := sql.Open("sqlite3", h.settings.Directory)
	if err != nil {
		return fmt.Errorf("Open: %w", err)
	}

	h.db = db

	err = h.CreateTable()
	if err != nil {
		return fmt.Errorf("CreateTable: %w", err)
	}

	return nil
}

func (h *Handler) CreateTable() error {
	query := `CREATE TABLE IF NOT EXISTS channels (
        ID INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
        ChannelIDs TEXT,
		Keyword TEXT
	);`

	_, err := h.db.Exec(query)
	if err != nil {
		return fmt.Errorf("Exec: %w", err)
	}

	return nil
}
