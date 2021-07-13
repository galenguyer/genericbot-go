package commands

import (
	"github.com/bwmarrin/discordgo"
	"github.com/galenguyer/genericbot/entities"
)

var Ping = &entities.Command{
	Name: "ping",
	Execute: func(c entities.Context) error {
		c.Session.ChannelMessageSendComplex(c.Message.ChannelID, &discordgo.MessageSend{
			Content:         "pong!",
			TTS:             false,
			AllowedMentions: &discordgo.MessageAllowedMentions{},
			Reference:       c.Message.Reference(),
		})
		return nil
	},
}
