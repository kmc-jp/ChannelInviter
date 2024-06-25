package database

import (
	"fmt"
	"strings"
)

var ErrorNotFound = fmt.Errorf("NotFound")

func (h *Handler) GetChannels(KeyMessage string) (channelIDs []string, err error) {
	rows, err := h.db.Query(
		"SELECT ID, ChannelIDs, KeyMessage FROM channels WHERE KeyMessage = ?",
		KeyMessage,
	)

	if err != nil {
		return nil, fmt.Errorf("Query: %w", err)
	}
	defer rows.Close()

	channelIDs = make([]string, 0)

	for rows.Next() {
		var id int
		var channelIDstring string
		var KeyMessage string
		err = rows.Scan(&id, &channelIDstring, &KeyMessage)
		if err != nil {
			return nil, fmt.Errorf("Scan: %w", err)
		}

		return strings.Split(channelIDstring, channelIDseparater), nil
	}

	return nil, ErrorNotFound
}
