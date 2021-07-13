package commands

import (
	"github.com/bwmarrin/discordgo"
	"github.com/galenguyer/genericbot/entities"
	"github.com/galenguyer/genericbot/permissions"
)

var Ping = &entities.Command{
	Name:        "ping",
	Permissions: permissions.BotOwner,
	Action: func(c entities.Context) error {
		_, err := c.Session.ChannelMessageSendComplex(c.Message.ChannelID, &discordgo.MessageSend{
			Content:         "pong!",
			TTS:             false,
			AllowedMentions: &discordgo.MessageAllowedMentions{},
			Reference:       c.Message.Reference(),
		})
		return err
	},
}
