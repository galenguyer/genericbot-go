package entities

import (
	"github.com/bwmarrin/discordgo"
	"github.com/galenguyer/genericbot/config"
)

type Context struct {
	Session       *discordgo.Session
	GuildId       string
	ChannelId     string
	Message       discordgo.Message
	Config        config.Config
	ParsedCommand ParsedCommand
}
