package commands

import (
	"github.com/galenguyer/genericbot/database"
	"github.com/galenguyer/genericbot/entities"
	"github.com/galenguyer/genericbot/permissions"
)

var Migrate = &entities.Command{
	Name:        "migrate",
	Description: "Migrate all guild structs to the new format",
	Permissions: permissions.BotOwner,
	Action: func(c entities.Context) error {
		conf, err := database.ConvertLegacyGuildConfig(c.GuildId)
		if err != nil {
			c.Reply("An error occured retrieving the configuration for your server")
			return err
		}

		err = database.SaveGuildConfig(c.GuildId, c.Message.ID, *conf)
		if err != nil {
			c.Reply("An error occured saving the configuration for your server")
			return err
		}

		_, err = c.Reply("migrated guildconfig")
		return err
	},
}
