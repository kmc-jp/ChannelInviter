package database

import (
	"fmt"
	"strings"
)

func (h *Handler) DeleteChannels(KeyMessage string, channelIDs ...string) error {
	previousChannelIDs, err := h.GetChannels(KeyMessage)
	if err != nil {
		return fmt.Errorf("GetChannels: %w", err)
	}

	var newChannelIDs = []string{}

	var found bool
	for _, id := range previousChannelIDs {
		for _, deleteID := range channelIDs {
			if id == deleteID {
				found = true
			}
		}
		if !found {
			newChannelIDs = append(newChannelIDs, id)
		}
	}

	if len(newChannelIDs) == 0 {
		query := `DELETE FROM channels WHERE KeyMessage = ?`
		_, err = h.db.Exec(query, KeyMessage)
		if err != nil {
			return fmt.Errorf("Exec: %w", err)
		}
		return nil
	}

	query := "UPDATE channels SET ChannelIDs = ? WHERE KeyMessage = ?"
	_, err = h.db.Exec(query, strings.Join(newChannelIDs, channelIDseparater), KeyMessage)
	if err != nil {
		return fmt.Errorf("Exec: %w", err)
	}

	return nil
}
