package slack

import (
	"fmt"
	"strings"

	"github.com/kmc-jp/ChannelInviter/database"
	"github.com/slack-go/slack"
)

func (h *Handler) sendSetChannels(KeyMessage string, channel string) error {
	channelIDs, err := h.db.GetChannels(KeyMessage)
	if err == database.ErrorNotFound {
		h.api.PostMessage(channel,
			slack.MsgOptionText(
				"There are no channels",
				false,
			),
		)
		return nil
	}
	if err != nil {
		return fmt.Errorf("GetChannels: %w", err)
	}

	h.api.PostMessage(channel,
		slack.MsgOptionText(
			fmt.Sprintf("These are the channels set to the key message: \n%s",
				fmt.Sprintf("<%s>", strings.Join(channelIDs, ">\n<")),
			),
			false,
		),
	)

	return nil
}
