package mentionhandler

import (
	"fmt"
	"strings"

	"github.com/slack-go/slack/slackevents"
)

func (h *Handler) removeChannels(ev *slackevents.AppMentionEvent) error {
	submatch := removeChannelRegExp.FindAllStringSubmatch(ev.Text, 1)
	Keyword := submatch[0][1]
	delChannelIDs := []string{}

	for _, channel := range strings.Split(submatch[0][2], "\n") {
		if channelIDRegExp.MatchString(channel) {
			delChannelIDs = append(delChannelIDs, channelIDRegExp.FindAllStringSubmatch(channel, 1)[0][1])
		}
	}

	err := h.db.RemoveChannels(Keyword, delChannelIDs...)
	if err != nil {
		return fmt.Errorf("removeChannels: %w", err)
	}

	err = h.sendSetChannels(Keyword, ev.Channel)
	if err != nil {
		return fmt.Errorf("getChannels: %w", err)
	}

	return nil
}
