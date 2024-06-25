package slack

import (
	"fmt"

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
						go h.mentionHandler(ev)
					}
				}
			}
		}
	}()

	return nil
}
