package slashcommandhandler

import (
	"fmt"

	"github.com/slack-go/slack"
)

func (h *Handler) removeChannels(ev slack.SlashCommand) error {
	submatch := KeywordAndChannelsRegExp.FindAllStringSubmatch(ev.Text, 1)
	Keyword := submatch[0][1]
	delChannelIDs := []string{}

	for _, c := range channelIDRegExp.FindAllStringSubmatch(ev.Text, -1) {
		channel := c[1]
		delChannelIDs = append(delChannelIDs, channel)
	}

	err := h.db.RemoveChannels(Keyword, delChannelIDs...)
	if err != nil {
		return fmt.Errorf("RemoveChannels: %w", err)
	}

	err = h.sendSetChannels(Keyword, ev.ChannelID, ev.UserID)
	if err != nil {
		return fmt.Errorf("getChannels: %w", err)
	}

	return nil
}
