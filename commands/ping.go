package commands

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/galenguyer/genericbot/entities"
)

var Ping = &entities.Command{
	Name:        "ping",
	Description: "Get the time taken for the bot to reply to a message",
	Action: func(c entities.Context) error {
		msg, err := c.Reply("Pong!")
		if err != nil {
			return err
		}

		cmdCreation, _ := discordgo.SnowflakeTimestamp(c.Message.ID)
		rplCreation, _ := discordgo.SnowflakeTimestamp(msg.ID)
		timeDiff := rplCreation.Sub(cmdCreation)
		diffString := fmt.Sprintf("Pong! Time taken: `%dms`", timeDiff.Milliseconds())

		_, err = c.Session.ChannelMessageEditComplex(
			&discordgo.MessageEdit{
				Content:         &diffString,
				AllowedMentions: &discordgo.MessageAllowedMentions{},

				ID:      msg.ID,
				Channel: msg.ChannelID,
			},
		)
		return err
	},
}
