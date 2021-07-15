package commands

import (
	"encoding/json"

	"github.com/bwmarrin/discordgo"
	"github.com/galenguyer/genericbot/database"
	"github.com/galenguyer/genericbot/entities"
)

var Config = &entities.Command{
	Name:        "config",
	Description: "Get or set the server configuration",
	Usage:       "<option> <value>",
	Action: func(c entities.Context) error {
		conf, err := database.GetGuildConfig(c.GuildId)
		if err != nil {
			c.Session.ChannelMessageSendComplex(c.Message.ChannelID, &discordgo.MessageSend{
				Content:         "An error occured retrieving the configuration for your server",
				TTS:             false,
				AllowedMentions: &discordgo.MessageAllowedMentions{},
				Reference:       c.Message.Reference(),
			})
			return err
		}
		jsonConf, _ := json.MarshalIndent(conf, "", "    ")
		_, err = c.Session.ChannelMessageSendComplex(c.Message.ChannelID, &discordgo.MessageSend{
			Content:         "```\n" + string(jsonConf) + "\n```",
			TTS:             false,
			AllowedMentions: &discordgo.MessageAllowedMentions{},
			Reference:       c.Message.Reference(),
		})
		return err
	},
}

var Migrate = &entities.Command{
	Name: "migrate",
	Action: func(c entities.Context) error {
		conf, err := database.ConvertLegacyGuildConfig(c.GuildId)
		if err != nil {
			c.Session.ChannelMessageSendComplex(c.Message.ChannelID, &discordgo.MessageSend{
				Content:         "An error occured retrieving the configuration for your server",
				TTS:             false,
				AllowedMentions: &discordgo.MessageAllowedMentions{},
				Reference:       c.Message.Reference(),
			})
			return err
		}

		err = database.SaveGuildConfig(c.GuildId, *conf)
		if err != nil {
			c.Session.ChannelMessageSendComplex(c.Message.ChannelID, &discordgo.MessageSend{
				Content:         "An error occured saving the configuration for your server",
				TTS:             false,
				AllowedMentions: &discordgo.MessageAllowedMentions{},
				Reference:       c.Message.Reference(),
			})
			return err
		}

		_, err = c.Session.ChannelMessageSendComplex(c.Message.ChannelID, &discordgo.MessageSend{
			Content:         "migrated guildconfig",
			TTS:             false,
			AllowedMentions: &discordgo.MessageAllowedMentions{},
			Reference:       c.Message.Reference(),
		})
		return err
	},
}
