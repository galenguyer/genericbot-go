package entities

import (
	"io/ioutil"
	"os"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/galenguyer/genericbot/config"
	"github.com/galenguyer/genericbot/logging"
	"github.com/galenguyer/genericbot/permissions"
	"github.com/sirupsen/logrus"
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
	return ctx.ReplySplit(message, " ")
}

func (ctx Context) ReplyFile(message string) (*discordgo.Message, error) {
	if len(message) > 2000 {
		tmpFile, err := ioutil.TempFile(os.TempDir(), "reply-*.txt")
		if err != nil {
			logging.Logger.WithFields(logrus.Fields{"error": err, "module": "context", "method": "Reply"}).Errorln("could not create temp file")
			return ctx.Reply("Whoops, it looks like something went wrong - my reply was too long and I couldn't put it in a file for you.")
		}
		if _, err = tmpFile.Write([]byte(strings.TrimSpace(strings.Trim(message, "` ")))); err != nil {
			logging.Logger.WithFields(logrus.Fields{"error": err, "module": "context", "method": "Reply"}).Errorln("could not write temp file")
			return ctx.Reply("Whoops, it looks like something went wrong - my reply was too long and I couldn't put it in a file for you.")
		}
		if err := tmpFile.Close(); err != nil {
			logging.Logger.WithFields(logrus.Fields{"error": err, "module": "context", "method": "Reply"}).Errorln("could not close temp file")
			return ctx.Reply("Whoops, it looks like something went wrong - my reply was too long and I couldn't put it in a file for you.")
		}
		reader, err := os.Open(tmpFile.Name())
		if err != nil {
			logging.Logger.WithFields(logrus.Fields{"error": err, "module": "context", "method": "Reply"}).Errorln("could not open temp file")
			return ctx.Reply("Whoops, it looks like something went wrong - my reply was too long and I couldn't put it in a file for you.")
		}
		defer reader.Close()
		defer os.Remove(tmpFile.Name())
		return ctx.Session.ChannelMessageSendComplex(ctx.Message.ChannelID, &discordgo.MessageSend{
			Content:         "My reply was really long, so I've had to put it in this file for you!",
			TTS:             false,
			AllowedMentions: &discordgo.MessageAllowedMentions{},
			Reference:       ctx.Message.Reference(),
			File: &discordgo.File{
				Name:        tmpFile.Name(),
				ContentType: "txt",
				Reader:      reader,
			},
		})
	} else {
		return ctx.Session.ChannelMessageSendComplex(ctx.Message.ChannelID, &discordgo.MessageSend{
			Content:         message,
			TTS:             false,
			AllowedMentions: &discordgo.MessageAllowedMentions{},
			Reference:       ctx.Message.Reference(),
		})
	}
}

func (ctx Context) ReplySplit(message, delimiter string) (*discordgo.Message, error) {
	if len(message) > 2000 {
		messages := splitMessage(message, delimiter)
		var finalMessage *discordgo.Message
		for _, msg := range messages {
			var err error
			finalMessage, err = ctx.ReplyFile(msg)
			if err != nil {
				return nil, err
			}
		}
		return finalMessage, nil
	} else {
		return ctx.Session.ChannelMessageSendComplex(ctx.Message.ChannelID, &discordgo.MessageSend{
			Content:         message,
			TTS:             false,
			AllowedMentions: &discordgo.MessageAllowedMentions{},
			Reference:       ctx.Message.Reference(),
		})
	}
}

// TODO: Markdown handling
func splitMessage(message, delimiter string) []string {
	var output []string
	//fallbackDelimiter := " "
	maxLength := 1800
	components := strings.Split(message, delimiter)

	var aggregator = ""
	for i, component := range components {
		if len(aggregator)+len(component) > maxLength {
			output = append(output, aggregator)
			aggregator = component
		} else {
			aggregator += component
			if i < len(components)-1 {
				aggregator += delimiter
			}
		}
	}

	output = append(output, aggregator)
	return output
}
