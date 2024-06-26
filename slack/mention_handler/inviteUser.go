package mentionhandler

import (
	"strings"
)

func (h *Handler) inviteUser(userid string, channels []string) {
	for _, channel := range channels {
		h.api.InviteUsersToConversation(strings.TrimLeft(channel, "#"), userid)
	}
}
