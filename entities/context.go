package entities

import (
	"github.com/bwmarrin/discordgo"
	"github.com/galenguyer/genericbot/config"
	"github.com/galenguyer/genericbot/permissions"
)

type Context struct {
	Session       *discordgo.Session
	GuildId       string
	ChannelId     string
	Message       discordgo.Message
	Config        config.Config
	ParsedCommand ParsedCommand
	Permissions   permissions.PermissionLevel
}

func (ctx Context) Reply(message string) (*discordgo.Message, error) {
	return ctx.Session.ChannelMessageSendComplex(ctx.Message.ChannelID, &discordgo.MessageSend{
		Content:         message,
		TTS:             false,
		AllowedMentions: &discordgo.MessageAllowedMentions{},
		Reference:       ctx.Message.Reference(),
	})

}
