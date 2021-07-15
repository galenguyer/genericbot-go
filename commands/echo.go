package commands

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/galenguyer/genericbot/entities"
	"github.com/galenguyer/genericbot/logging"
	"github.com/galenguyer/genericbot/permissions"
	"github.com/sirupsen/logrus"
)

var Echo = &entities.Command{
	Name:        "echo",
	Description: "Make the bot say something",
	Usage:       "<message>",
	Permissions: permissions.BotAdministrator,
	Action: func(c entities.Context) error {
		// delete the original message and log failures
		err := c.Session.ChannelMessageDelete(c.ChannelId, c.Message.ID)
		if err != nil {
			logging.Logger.WithFields(logrus.Fields{
				"error":    err,
				"guild":    c.GuildId,
				"channel":  c.ChannelId,
				"author":   c.Message.Author.ID,
				"messsage": c.Message.ID,
				"command":  "ping",
			}).Error("could not delete message")
		}

		if len(c.ParsedCommand.ParameterString) > 0 {
			// echo the user's message back to them
			_, err = c.Session.ChannelMessageSendComplex(c.ChannelId,
				&discordgo.MessageSend{
					Content:         c.ParsedCommand.ParameterString,
					AllowedMentions: &discordgo.MessageAllowedMentions{},
				},
			)
			return err
		} else {
			dmChannel, err := c.Session.UserChannelCreate(c.Message.Author.ID)
			if err != nil {
				logging.Logger.WithFields(logrus.Fields{
					"error":    err,
					"guild":    c.GuildId,
					"channel":  c.ChannelId,
					"author":   c.Message.Author.ID,
					"messsage": c.Message.ID,
					"command":  "ping",
				}).Error("could not inform user of emtpy message")
			} else {
				_, err = c.Session.ChannelMessageSend(dmChannel.ID, fmt.Sprintf("You asked me to echo a message in <#%s>, but I can't echo an empty message!", c.ChannelId))
				if err != nil {
					logging.Logger.WithFields(logrus.Fields{
						"error":    err,
						"guild":    c.GuildId,
						"channel":  c.ChannelId,
						"author":   c.Message.Author.ID,
						"messsage": c.Message.ID,
						"command":  "ping",
					}).Error("could not inform user of emtpy message")
				}
			}
		}
		return nil
	},
}
