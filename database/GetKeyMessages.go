package database

import (
	"fmt"
)

func (h *Handler) GetKeywords() (Keywords []string, err error) {
	rows, err := h.db.Query(
		"SELECT Keyword FROM channels",
	)
	if err != nil {
		return nil, fmt.Errorf("Query: %w", err)
	}
	defer rows.Close()

	Keywords = make([]string, 0)

	for rows.Next() {
		var Keyword string
		err = rows.Scan(&Keyword)
		if err != nil {
			return nil, fmt.Errorf("Scan: %w", err)
		}
		Keywords = append(Keywords, Keyword)
	}

	return Keywords, nil
}
