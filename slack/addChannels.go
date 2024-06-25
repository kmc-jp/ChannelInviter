package slack

import (
	"fmt"
	"strings"

	"github.com/slack-go/slack/slackevents"
)

func (h *Handler) addChannels(ev *slackevents.AppMentionEvent) error {
	submatch := addChannelRegExp.FindAllStringSubmatch(ev.Text, 1)
	KeyMessage := submatch[0][1]
	newChannelIDs := []string{}

	for _, rkey := range ReservedWords {
		if KeyMessage == rkey {
			return fmt.Errorf("Reserved")
		}
	}

	for _, channel := range strings.Split(submatch[0][2], "\n") {
		if channelIDRegExp.MatchString(channel) {
			newChannelIDs = append(newChannelIDs, channelIDRegExp.FindAllStringSubmatch(channel, 1)[0][1])
		}
	}

	err := h.db.AddChannels(KeyMessage, newChannelIDs...)
	if err != nil {
		return fmt.Errorf("AddChannels: %w", err)
	}

	err = h.sendSetChannels(KeyMessage, ev.Channel)
	if err != nil {
		return fmt.Errorf("getChannels: %w", err)
	}

	return nil
}
