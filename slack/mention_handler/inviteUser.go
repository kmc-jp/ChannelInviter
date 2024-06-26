package mentionhandler

import (
	"fmt"
	"strings"
)

func (h *Handler) inviteUser(userid string, channels []string) error {
	for _, channel := range channels {
		_, err := h.api.InviteUsersToConversation(strings.TrimLeft(channel, "#"), userid)
		if err != nil {
			return fmt.Errorf("InviteUsersToConversation: %w", err)
		}
	}
	return nil
}
