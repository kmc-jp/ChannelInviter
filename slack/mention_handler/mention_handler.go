package mentionhandler

import (
	"github.com/kmc-jp/ChannelInviter/database"
	"github.com/slack-go/slack"
	"github.com/slack-go/slack/socketmode"
)

type Handler struct {
	api    *slack.Client
	scm    *socketmode.Client
	userID string
	db     *database.Handler
}

func New(api *slack.Client, scm *socketmode.Client, db *database.Handler) *Handler {
	return &Handler{
		api: api,
		scm: scm,
		db:  db,
	}
}

func (h *Handler) SetUserID(userid string) {
	h.userID = userid
}
