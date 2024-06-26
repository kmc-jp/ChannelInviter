package slack

import (
	"github.com/kmc-jp/ChannelInviter/database"
	mentionhandler "github.com/kmc-jp/ChannelInviter/slack/mention_handler"
	slashcommandhandler "github.com/kmc-jp/ChannelInviter/slack/slashcommand_handler"
	"github.com/slack-go/slack"
	"github.com/slack-go/slack/socketmode"
)

type Handler struct {
	api      *slack.Client
	scm      *socketmode.Client
	userID   string
	settings Settings
	db       *database.Handler

	mentionHandler      *mentionhandler.Handler
	slashcommandHandler *slashcommandhandler.Handler
}

func New(settings Settings, db *database.Handler) *Handler {
	var api = slack.New(
		settings.Token,
		slack.OptionAppLevelToken(settings.AppLevelToken),
	)

	var handler = Handler{
		settings: settings,
		api:      api,
		db:       db,
	}

	return &handler
}
