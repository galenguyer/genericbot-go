package commands

import (
	"github.com/galenguyer/genericbot/database"
	"github.com/galenguyer/genericbot/entities"
	"github.com/galenguyer/genericbot/json"
)

var Config = &entities.Command{
	Name:        "config",
	Description: "Get or set the server configuration",
	Usage:       "<option> <value>",
	Action: func(c entities.Context) error {
		conf, err := database.GetGuildConfig(c.GuildId)
		if err != nil {
			c.Reply("An error occured retrieving the configuration for your server")
			return err
		}
		jsonConf, _ := json.JSONMarshalIndented(conf, "", "    ")
		_, err = c.Reply("```\n" + string(jsonConf) + "\n```")
		return err
	},
}

var Migrate = &entities.Command{
	Name: "migrate",
	Action: func(c entities.Context) error {
		conf, err := database.ConvertLegacyGuildConfig(c.GuildId)
		if err != nil {
			c.Reply("An error occured retrieving the configuration for your server")
			return err
		}

		err = database.SaveGuildConfig(c.GuildId, *conf)
		if err != nil {
			c.Reply("An error occured saving the configuration for your server")
			return err
		}

		_, err = c.Reply("migrated guildconfig")
		return err
	},
}
