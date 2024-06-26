package slashcommandhandler

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/slack-go/slack"
)

var KeywordAndChannelsRegExp = regexp.MustCompile(`(\S+)\s+((<#\S+>\s*)+)`)

var channelIDRegExp = regexp.MustCompile(`<(#[A-Z0-9]+)[^\s]*>`)
var userIDRegExp = regexp.MustCompile(`<@([A-Z0-9]+)>`)

var ReservedWords = []string{"AddChannels", "DeleteChannels", "GetChannels"}

func (h *Handler) Executed(e slack.SlashCommand) {
	fmt.Printf("%+v\n", e)
	switch e.Command {
	case "/inviteraddchannels":
		if !KeywordAndChannelsRegExp.MatchString(e.Text) {
			h.api.PostEphemeral(
				e.ChannelID, e.UserID,
				slack.MsgOptionText("Error: invalid format", false),
			)
			return
		}

		err := h.addChannels(e)
		if err != nil {
			h.api.PostEphemeral(
				e.ChannelID, e.UserID,
				slack.MsgOptionText(fmt.Sprintf("Error: addChannels: %s", err), false),
			)
			return
		}
	case "/inviterremovechannels":
		if !KeywordAndChannelsRegExp.MatchString(e.Text) {
			h.api.PostEphemeral(
				e.ChannelID, e.UserID,
				slack.MsgOptionText("Error: invalid format", false),
			)
			return
		}

		err := h.removeChannels(e)
		if err != nil {
			h.api.PostEphemeral(
				e.ChannelID, e.UserID,
				slack.MsgOptionText(fmt.Sprintf("Error: removeChannels: %s", err), false),
			)
			return
		}
	case "/inviterjoin":
		Keyword := strings.Split(strings.TrimSpace(e.Text), " ")[0]
		channels, err := h.db.GetChannels(Keyword)
		if err != nil {
			h.api.PostEphemeral(
				e.ChannelID, e.UserID,
				slack.MsgOptionText(fmt.Sprintf("Error: GetChannels: %s", err), false),
			)
			return
		}
		for _, channel := range channels {
			h.api.InviteUsersToConversation(strings.TrimLeft(channel, "#"), e.UserID)
		}

		h.api.PostEphemeral(
			e.ChannelID, e.UserID,
			slack.MsgOptionText("done.", false),
		)
	case "/inviterinvite":
		Keyword := strings.Split(strings.TrimSpace(e.Text), " ")[0]
		channels, err := h.db.GetChannels(Keyword)
		if err != nil {
			h.api.PostEphemeral(
				e.ChannelID, e.UserID,
				slack.MsgOptionText(fmt.Sprintf("Error: GetChannels: %s", err), false),
			)
			return
		}

		for _, u := range userIDRegExp.FindAllStringSubmatch(e.Text, -1) {
			for _, channel := range channels {
				h.api.InviteUsersToConversation(strings.TrimLeft(channel, "#"), u[1])
			}
		}

		h.api.PostEphemeral(
			e.ChannelID, e.UserID,
			slack.MsgOptionText("done.", false),
		)
	}
}
