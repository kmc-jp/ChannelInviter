package slack

import (
	"fmt"
	"strings"

	"github.com/slack-go/slack/slackevents"
)

func (h *Handler) deleteChannels(ev *slackevents.AppMentionEvent) error {
	submatch := deleteChannelRegExp.FindAllStringSubmatch(ev.Text, 1)
	KeyMessage := submatch[0][1]
	delChannelIDs := []string{}

	for _, channel := range strings.Split(submatch[0][2], "\n") {
		if channelIDRegExp.MatchString(channel) {
			delChannelIDs = append(delChannelIDs, channelIDRegExp.FindAllStringSubmatch(channel, 1)[0][1])
		}
	}

	err := h.db.DeleteChannels(KeyMessage, delChannelIDs...)
	if err != nil {
		return fmt.Errorf("DeleteChannels: %w", err)
	}

	err = h.sendSetChannels(KeyMessage, ev.Channel)
	if err != nil {
		return fmt.Errorf("getChannels: %w", err)
	}

	return nil
}
