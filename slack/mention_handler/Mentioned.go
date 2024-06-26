package mentionhandler

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
)

var addChannelRegExp = regexp.MustCompile(`AddChannels?\s+(\S+)\n+((<#\S+>\n?)+)`)
var removeChannelRegExp = regexp.MustCompile(`RemoveChannels?\s+(\S+)\n+((<#\S+>\n?)+)`)
var getChannelRegExp = regexp.MustCompile(`GetChannels?\s+(\S+)`)

var channelIDRegExp = regexp.MustCompile(`<(#[A-Z0-9]+)[^\s]*>`)
var userIDRegExp = regexp.MustCompile(`<@([A-Z0-9]+)>`)

var ReservedWords = []string{"AddChannels", "RemoveChannels", "GetChannels"}

func (h *Handler) Mentioned(ev *slackevents.AppMentionEvent) {
	if addChannelRegExp.MatchString(ev.Text) {
		err := h.addChannels(ev)
		if err != nil {
			h.api.PostMessage(ev.Channel,
				slack.MsgOptionText(
					fmt.Sprintf("Error: %v", err),
					true,
				),
			)
			return
		}
		return
	}

	if removeChannelRegExp.MatchString(ev.Text) {
		err := h.removeChannels(ev)
		if err != nil {
			h.api.PostMessage(ev.Channel,
				slack.MsgOptionText(
					fmt.Sprintf("Error: %v", err),
					true,
				),
			)
			return
		}
		return
	}

	if getChannelRegExp.MatchString(ev.Text) {
		err := h.sendSetChannels(getChannelRegExp.FindAllStringSubmatch(ev.Text, 1)[0][1], ev.Channel)
		if err != nil {
			h.api.PostMessage(ev.Channel,
				slack.MsgOptionText(
					fmt.Sprintf("Error: %v", err),
					true,
				),
			)
			return
		}
		return
	}

	Keywords, err := h.db.GetKeywords()
	if err != nil {
		h.api.PostMessage(ev.Channel,
			slack.MsgOptionText(
				fmt.Sprintf("Error: %v", err),
				true,
			),
		)
		return
	}

	for _, phrase := range strings.Split(ev.Text, " ") {
		for _, key := range Keywords {
			if key == phrase {
				channels, err := h.db.GetChannels(key)
				if err != nil {
					h.api.PostMessage(ev.Channel,
						slack.MsgOptionText(
							fmt.Sprintf("Error: GetChannels %v", err),
							true,
						),
					)
					return
				}

				usernum := 0
				for _, match := range userIDRegExp.FindAllStringSubmatch(ev.Text, -1) {
					if match[1] == h.userID {
						continue
					}
					usernum++
					err := h.inviteUser(match[1], channels)
					if err != nil {
						h.api.PostMessage(ev.Channel,
							slack.MsgOptionText(
								fmt.Sprintf("Error: inviteUsers(<@%s>): %v", match[1], err),
								false,
							),
						)
						return
					}
				}
				if usernum == 0 {
					err := h.inviteUser(ev.User, channels)
					if err != nil {
						h.api.PostMessage(ev.Channel,
							slack.MsgOptionText(
								fmt.Sprintf("Error: inviteUser %v", err),
								true,
							),
						)
						return
					}
				}

			}
		}
	}
}
