package database

import (
	"fmt"
	"strings"
)

var ErrorNotFound = fmt.Errorf("NotFound")

func (h *Handler) GetChannels(Keyword string) (channelIDs []string, err error) {
	rows, err := h.db.Query(
		"SELECT ID, ChannelIDs, Keyword FROM channels WHERE Keyword = ?",
		Keyword,
	)

	if err != nil {
		return nil, fmt.Errorf("Query: %w", err)
	}
	defer rows.Close()

	channelIDs = make([]string, 0)

	for rows.Next() {
		var id int
		var channelIDstring string
		var Keyword string
		err = rows.Scan(&id, &channelIDstring, &Keyword)
		if err != nil {
			return nil, fmt.Errorf("Scan: %w", err)
		}

		return strings.Split(channelIDstring, channelIDseparater), nil
	}

	return nil, ErrorNotFound
}
