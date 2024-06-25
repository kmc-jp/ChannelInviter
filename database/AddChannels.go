package database

import (
	"fmt"
	"strings"
)

func (h *Handler) AddChannels(KeyMessage string, channelIDs ...string) error {
	previousChannelIDs, isExistErr := h.GetChannels(KeyMessage)

	var newChannelIDs = previousChannelIDs
	var found bool
	for _, newID := range channelIDs {
		for _, id := range previousChannelIDs {
			if id == newID {
				found = true
			}
		}
		if !found {
			newChannelIDs = append(newChannelIDs, newID)
		}
	}

	// KeyMessage already exist
	if isExistErr == nil {
		query := "UPDATE channels SET ChannelIDs = ? WHERE KeyMessage = ?"
		_, err := h.db.Exec(query, strings.Join(newChannelIDs, channelIDseparater), KeyMessage)
		if err != nil {
			return fmt.Errorf("Exec: %w", err)
		}
		return nil
	}

	// KeyMessage not exist
	query := `INSERT INTO channels (ChannelIDs, KeyMessage) VALUES (?, ?)`
	_, err := h.db.Exec(query, strings.Join(channelIDs, channelIDseparater), KeyMessage)
	if err != nil {
		return fmt.Errorf("Exec: %w", err)
	}

	return nil
}
