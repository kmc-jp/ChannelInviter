package slashcommandhandler

import (
	"fmt"
	"strings"

	"github.com/kmc-jp/ChannelInviter/database"
	"github.com/slack-go/slack"
)

func (h *Handler) sendSetChannels(Keyword string, channel string, userid string) error {
	channelIDs, err := h.db.GetChannels(Keyword)
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

	h.api.PostEphemeral(channel, userid,
		slack.MsgOptionText(
			fmt.Sprintf("These are the channels set to the keyword: %s\n%s",
				Keyword,
				fmt.Sprintf("<%s>", strings.Join(channelIDs, ">\n<")),
			),
			false,
		),
	)

	return nil
}
