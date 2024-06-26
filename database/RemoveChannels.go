package database

import (
	"fmt"
	"strings"
)

func (h *Handler) RemoveChannels(Keyword string, channelIDs ...string) error {
	previousChannelIDs, err := h.GetChannels(Keyword)
	if err != nil {
		return fmt.Errorf("GetChannels: %w", err)
	}

	var newChannelIDs = []string{}

	for _, id := range previousChannelIDs {
		found := false
		for _, removeID := range channelIDs {
			if id == removeID {
				found = true
			}
		}
		if !found {
			newChannelIDs = append(newChannelIDs, id)
		}
	}

	if len(newChannelIDs) == 0 {
		query := `DELETE FROM channels WHERE Keyword = ?`
		_, err = h.db.Exec(query, Keyword)
		if err != nil {
			return fmt.Errorf("Exec: %w", err)
		}
		return nil
	}

	query := "UPDATE channels SET ChannelIDs = ? WHERE Keyword = ?"
	_, err = h.db.Exec(query, strings.Join(newChannelIDs, channelIDseparater), Keyword)
	if err != nil {
		return fmt.Errorf("Exec: %w", err)
	}

	return nil
}
