package handlers

import (
	"strings"

	linq "github.com/ahmetb/go-linq/v3"
	"github.com/bwmarrin/discordgo"
	"github.com/galenguyer/genericbot/commands"
	"github.com/galenguyer/genericbot/config"
	"github.com/galenguyer/genericbot/entities"
	"github.com/galenguyer/genericbot/logging"
	"github.com/galenguyer/genericbot/permissions"
	"github.com/sirupsen/logrus"
)

var (
	Commands []*entities.Command
)

func init() {
	Commands = append(Commands, commands.Config)
	Commands = append(Commands, commands.Echo)
	Commands = append(Commands, commands.Find)
	Commands = append(Commands, commands.Import)
	Commands = append(Commands, commands.Mock)
	Commands = append(Commands, commands.Ping)
	Commands = append(Commands, commands.Time)
	logging.Logger.WithFields(logrus.Fields{"module": "handlers", "method": "init"}).Infof("loaded %d commands", len(Commands))
}

func OnMessageRecieved(s *discordgo.Session, m *discordgo.MessageCreate, config *config.Config) {
	command := parseCommand(m.Message.Content, config)
	if command != nil {
		commandToExecute := linq.From(Commands).FirstWith(func(i interface{}) bool {
			return strings.EqualFold(i.(*entities.Command).Name, command.Name) || linq.From(i.(*entities.Command).Aliases).AnyWith(func(x interface{}) bool {
				return strings.EqualFold(x.(string), command.Name)
			})
		})

		if commandToExecute != nil {
			commandToExecute.(*entities.Command).Execute(entities.Context{
				Session:       s,
				GuildId:       m.GuildID,
				ChannelId:     m.ChannelID,
				Message:       *m.Message,
				Config:        *config,
				ParsedCommand: *command,
				Permissions:   getPermissions(s, m, config),
			})
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

func getPermissions(s *discordgo.Session, m *discordgo.MessageCreate, config *config.Config) permissions.PermissionLevel {
	if config.BotConfig.OwnerId == m.Author.ID {
		return permissions.BotOwner
	}

	guild, err := s.Guild(m.GuildID)
	if err != nil {
		return permissions.User
	}
	if guild.OwnerID == m.Author.ID {
		return permissions.GuildOwner
	}

	member, err := s.GuildMember(m.GuildID, m.Author.ID)
	if err != nil {
		return permissions.User
	}
	if member.Permissions&discordgo.PermissionAdministrator == discordgo.PermissionAdministrator {
		return permissions.Administrator
	}

	return permissions.User
}
