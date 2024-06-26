package database

import (
	"fmt"
	"strings"
)

func (h *Handler) AddChannels(Keyword string, channelIDs ...string) error {
	previousChannelIDs, isExistErr := h.GetChannels(Keyword)

	var newChannelIDs = previousChannelIDs
	for _, newID := range channelIDs {
		found := false
		for _, id := range previousChannelIDs {
			if id == newID {
				found = true
			}
		}
		if !found {
			newChannelIDs = append(newChannelIDs, newID)
		}
	}

	// Keyword already exist
	if isExistErr == nil {
		query := "UPDATE channels SET ChannelIDs = ? WHERE Keyword = ?"
		_, err := h.db.Exec(query, strings.Join(newChannelIDs, channelIDseparater), Keyword)
		if err != nil {
			return fmt.Errorf("Exec: %w", err)
		}
		return nil
	}

	// Keyword not exist
	query := `INSERT INTO channels (ChannelIDs, Keyword) VALUES (?, ?)`
	_, err := h.db.Exec(query, strings.Join(channelIDs, channelIDseparater), Keyword)
	if err != nil {
		return fmt.Errorf("Exec: %w", err)
	}

	return nil
}
