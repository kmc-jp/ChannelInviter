package slashcommandhandler

import (
	"fmt"

	"github.com/slack-go/slack"
)

func (h *Handler) addChannels(ev slack.SlashCommand) error {
	submatch := KeywordAndChannelsRegExp.FindAllStringSubmatch(ev.Text, 1)
	Keyword := submatch[0][1]
	newChannelIDs := []string{}

	for _, rkey := range ReservedWords {
		if Keyword == rkey {
			return fmt.Errorf("Reserved")
		}
	}

	for _, c := range channelIDRegExp.FindAllStringSubmatch(ev.Text, -1) {
		channel := c[1]
		newChannelIDs = append(newChannelIDs, channel)
	}

	err := h.db.AddChannels(Keyword, newChannelIDs...)
	if err != nil {
		return fmt.Errorf("AddChannels: %w", err)
	}

	err = h.sendSetChannels(Keyword, ev.ChannelID, ev.UserID)
	if err != nil {
		return fmt.Errorf("getChannels: %w", err)
	}

	h.api.PostEphemeral(ev.ChannelID, ev.UserID,
		slack.MsgOptionText(
			fmt.Sprintf("Hint: You have to invite <@%s> to the channels before using this.", h.userID),
			false,
		),
	)

	return nil
}
