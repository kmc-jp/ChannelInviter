package slack

import (
	"fmt"

	mentionhandler "github.com/kmc-jp/ChannelInviter/slack/mention_handler"
	slashcommandhandler "github.com/kmc-jp/ChannelInviter/slack/slashcommand_handler"
	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
	"github.com/slack-go/slack/socketmode"
)

func (h *Handler) Start() error {
	h.scm = socketmode.New(h.api)

	res, _ := h.scm.AuthTest()
	h.userID = res.UserID

	go func() {
		err := h.scm.Run()
		if err != nil {
			fmt.Println(err)
		}
	}()

	h.mentionHandler = mentionhandler.New(h.api, h.scm, h.db)
	h.mentionHandler.SetUserID(h.userID)

	h.slashcommandHandler = slashcommandhandler.New(h.api, h.scm, h.db)
	h.slashcommandHandler.SetUserID(h.userID)

	go func() {
		for evt := range h.scm.Events {
			switch evt.Type {
			case socketmode.EventTypeEventsAPI:
				e, ok := evt.Data.(slackevents.EventsAPIEvent)
				if !ok {
					continue
				}

				h.scm.Ack(*evt.Request)

				switch e.Type {
				case slackevents.CallbackEvent:
					innerEvent := e.InnerEvent
					switch ev := innerEvent.Data.(type) {
					case *slackevents.AppMentionEvent:
						go h.mentionHandler.Mentioned(ev)
					}
				}
			case socketmode.EventTypeSlashCommand:
				e, ok := evt.Data.(slack.SlashCommand)
				if !ok {
					continue
				}
				h.scm.Ack(*evt.Request)
				h.slashcommandHandler.Executed(e)
			}
		}
	}()

	return nil
}
