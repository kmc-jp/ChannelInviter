package mentionhandler

import (
	"fmt"
	"strings"

	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
)

func (h *Handler) addChannels(ev *slackevents.AppMentionEvent) error {
	submatch := addChannelRegExp.FindAllStringSubmatch(ev.Text, 1)
	Keyword := submatch[0][1]
	newChannelIDs := []string{}

	for _, rkey := range ReservedWords {
		if Keyword == rkey {
			return fmt.Errorf("Reserved")
		}
	}

	for _, channel := range strings.Split(submatch[0][2], "\n") {
		if channelIDRegExp.MatchString(channel) {
			newChannelIDs = append(newChannelIDs, channelIDRegExp.FindAllStringSubmatch(channel, 1)[0][1])
		}
	}

	err := h.db.AddChannels(Keyword, newChannelIDs...)
	if err != nil {
		return fmt.Errorf("AddChannels: %w", err)
	}

	err = h.sendSetChannels(Keyword, ev.Channel)
	if err != nil {
		return fmt.Errorf("getChannels: %w", err)
	}

	h.api.PostMessage(ev.Channel,
		slack.MsgOptionText(
			fmt.Sprintf("Hint: You have to invite <@%s> to the channels before using this.", h.userID),
			false,
		),
	)

	return nil
}
