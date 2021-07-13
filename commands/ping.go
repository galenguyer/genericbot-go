package commands

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/galenguyer/genericbot/entities"
)

var Ping = &entities.Command{
	Name: "ping",
	Action: func(c entities.Context) error {
		msg, err := c.Session.ChannelMessageSendComplex(c.Message.ChannelID, &discordgo.MessageSend{
			Content:         "pong!",
			TTS:             false,
			AllowedMentions: &discordgo.MessageAllowedMentions{},
			Reference:       c.Message.Reference(),
		})
		if err != nil {
			return err
		}

		cmdCreation, _ := discordgo.SnowflakeTimestamp(c.Message.ID)
		rplCreation, _ := discordgo.SnowflakeTimestamp(msg.ID)
		timeDiff := rplCreation.Sub(cmdCreation)
		diffString := fmt.Sprintf("pong! time taken: `%dms`", timeDiff.Milliseconds())

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
