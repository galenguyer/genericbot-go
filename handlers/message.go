package handlers

import (
	"strings"

	linq "github.com/ahmetb/go-linq/v3"
	"github.com/bwmarrin/discordgo"
	"github.com/galenguyer/genericbot/config"
	"github.com/galenguyer/genericbot/entities"
	"github.com/galenguyer/genericbot/loaders"
	"github.com/galenguyer/genericbot/permissions"
)

func OnMessageRecieved(s *discordgo.Session, m *discordgo.MessageCreate, config *config.Config) {
	command := parseCommand(m.Message.Content, config)
	if command != nil {
		commandToExecute := linq.From(loaders.Commands).FirstWith(func(i interface{}) bool {
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
	if err != nil {
		return permissions.User
	}

	if hasPermission(s, m, discordgo.PermissionAdministrator) {
		return permissions.Administrator
	}

	return permissions.User
}

func hasPermission(session *discordgo.Session, message *discordgo.MessageCreate, permission int64) bool {
	member, _ := session.GuildMember(message.GuildID, session.State.User.ID)
	guildRoles, _ := session.GuildRoles(message.GuildID)
	for _, role := range guildRoles {
		for _, roleID := range member.Roles {
			if role.ID == roleID {
				if role.Permissions&permission != 0 {
					return true
				}
			}
		}
	}
	return false
}
