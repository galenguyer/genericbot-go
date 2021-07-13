package commands

import "github.com/galenguyer/genericbot/entities"

var Ping = &entities.Command{
	Name: "ping",
	Execute: func(c entities.Context) error {
		c.Session.ChannelMessageSendReply(c.Message.ChannelID, "pong!", c.Message.MessageReference)
		return nil
	},
}
