package commands

import (
	"fmt"
	"math/rand"
	"unicode"

	"github.com/bwmarrin/discordgo"
	"github.com/galenguyer/genericbot/entities"
	"github.com/galenguyer/genericbot/logging"
	"github.com/sirupsen/logrus"
)

var Mock = &entities.Command{
	Name:        "mock",
	Description: "Mock someone ruthlessly",
	Usage:       "<text>",
	Action: func(c entities.Context) error {
		if c.Message.MessageReference != nil {
			message, err := c.Session.ChannelMessage(c.ChannelId, c.Message.MessageReference.MessageID)
			if err != nil {
				c.Reply("Sorry, I couldn't find the message you were replying to")
				return err
			}

			// delete the original message and log failures
			err = c.Session.ChannelMessageDelete(c.ChannelId, c.Message.ID)
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

			_, err = c.Session.ChannelMessageSendComplex(c.ChannelId, &discordgo.MessageSend{
				Content:         mockText(message.Content),
				TTS:             false,
				AllowedMentions: &discordgo.MessageAllowedMentions{},
				Reference:       message.Reference(),
			})
			return err
		} else {
			_, err := c.Reply(mockText(c.ParsedCommand.ParameterString))
			return err
		}
	},
}

func mockText(input string) string {
	output := ""
	ignore := false
	for _, char := range input {
		if ignore {
			if char == ':' || char == '>' {
				ignore = false
			}
			output += fmt.Sprintf("%c", char)
		} else {
			if char == ':' || char == '<' {
				ignore = true
				output += fmt.Sprintf("%c", char)
			} else if ignore {
				output += fmt.Sprintf("%c", char)
			} else {
				if rand.Intn(2) == 0 {
					char = unicode.ToUpper(char)
				} else {
					char = unicode.ToLower(char)
				}
				output += fmt.Sprintf("%c", char)
			}
		}
	}
	return output
}
