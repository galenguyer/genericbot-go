package handlers

import (
	"strings"

	linq "github.com/ahmetb/go-linq/v3"
	"github.com/bwmarrin/discordgo"
	"github.com/galenguyer/genericbot/commands"
	"github.com/galenguyer/genericbot/config"
	"github.com/galenguyer/genericbot/entities"
	"github.com/galenguyer/genericbot/logging"
	"github.com/sirupsen/logrus"
)

var (
	Commands []*entities.Command
)

func init() {
	Commands = append(Commands, commands.Ping)
}

func OnMessageRecieved(s *discordgo.Session, m *discordgo.MessageCreate, config *config.Config) {
	command := parseCommand(m.Message.Content, config)
	if command != nil {
		commandToExecute := linq.From(Commands).FirstWith(func(i interface{}) bool {
			return i.(*entities.Command).Name == command.Name
		})

		if commandToExecute != nil {
			logging.Logger.WithFields(logrus.Fields{
				"module":  "handlers",
				"method":  "OnMessageRecieved",
				"guild":   m.GuildID,
				"channel": m.ChannelID,
				"command": commandToExecute.(*entities.Command).Name,
			}).Info("got command " + commandToExecute.(*entities.Command).Name)
			err := commandToExecute.(*entities.Command).Execute(entities.Context{
				Session:       s,
				GuildId:       m.GuildID,
				ChannelId:     m.ChannelID,
				Message:       *m.Message,
				Config:        *config,
				ParsedCommand: *command,
			})
			if err != nil {
				logging.Logger.WithFields(logrus.Fields{
					"error":   err,
					"module":  "handlers",
					"method":  "OnMessageRecieved",
					"guild":   m.GuildID,
					"channel": m.ChannelID,
					"command": commandToExecute.(*entities.Command).Name,
				}).Error("error executing command " + commandToExecute.(*entities.Command).Name)
			}
		}
	}
}

func parseCommand(message string, config *config.Config) *entities.ParsedCommand {
	// first check to make sure we have the prefix
	if len(message) <= len(config.BotConfig.Prefix) {
		return nil
	} else if !strings.HasPrefix(message, config.BotConfig.Prefix) {
		return nil
	} else {
		// next let's extract the command
		messageParts := strings.Split(message, " ")
		command := messageParts[0][len(config.BotConfig.Prefix):]

		var paramString string
		if len(messageParts) > 1 {
			paramString = strings.SplitN(message, " ", 2)[1]
		} else {
			paramString = ""
		}

		return &entities.ParsedCommand{
			Name:            command,
			ParameterList:   messageParts[1:],
			ParameterString: paramString,
		}
	}
}
