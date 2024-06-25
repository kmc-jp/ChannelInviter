package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/kmc-jp/ChannelInviter/database"
	"github.com/kmc-jp/ChannelInviter/slack"
)

func main() {
	settings, err := ReadSettings()
	if err != nil {
		fmt.Println("ReadSettings: ", err)
		return
	}

	fmt.Printf("%+v\n", settings)

	var db = database.New(settings.Database)
	err = db.Initialize()
	if err != nil {
		fmt.Println("Initialize: ", err)
		return
	}

	slackHandler := slack.New(settings.Slack, db)
	err = slackHandler.Start()
	if err != nil {
		fmt.Println("SlackHandlerError: ", err)
		return
	}

	fmt.Println("App started successfully")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)

	<-sc
}
