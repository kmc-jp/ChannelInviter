package database

import (
	"fmt"
)

func (h *Handler) GetKeyMessages() (KeyMessages []string, err error) {
	rows, err := h.db.Query(
		"SELECT KeyMessage FROM channels",
	)
	if err != nil {
		return nil, fmt.Errorf("Query: %w", err)
	}
	defer rows.Close()

	KeyMessages = make([]string, 0)

	for rows.Next() {
		var KeyMessage string
		err = rows.Scan(&KeyMessage)
		if err != nil {
			return nil, fmt.Errorf("Scan: %w", err)
		}
		KeyMessages = append(KeyMessages, KeyMessage)
	}

	return KeyMessages, nil
}
