package commands

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/galenguyer/genericbot/database"
	"github.com/galenguyer/genericbot/entities"
	"github.com/galenguyer/genericbot/legacy"
	"github.com/galenguyer/genericbot/logging"
	"github.com/galenguyer/genericbot/permissions"
	"github.com/sirupsen/logrus"
)

var Import = &entities.Command{
	Name:        "import",
	Description: "Import data from json",
	Permissions: permissions.BotOwner,
	Action: func(c entities.Context) error {
		var guilds []string

		// read guilds file
		guildsFile, err := os.Open("dump/guilds.json")
		if err != nil {
			return err
		}
		guildsBytes, err := ioutil.ReadAll(guildsFile)
		if err != nil {
			return err
		}

		err = json.Unmarshal(guildsBytes, &guilds)
		if err != nil {
			return err
		}

		// import configs
		for _, guild := range guilds {
			configFile, err := os.Open("dump/" + guild + "-config.json")
			if err != nil {
				logging.Logger.WithFields(logrus.Fields{
					"error":    err,
					"guild":    guild,
					"messsage": c.Message.ID,
					"command":  "import",
				}).Error("could not open config file")
				continue
			}
			configBytes, err := ioutil.ReadAll(configFile)
			if err != nil {
				logging.Logger.WithFields(logrus.Fields{
					"error":    err,
					"guild":    guild,
					"messsage": c.Message.ID,
					"command":  "import",
				}).Error("could not read config file")
				continue
			}
			var legConf legacy.GuildConfig
			err = json.Unmarshal(configBytes, &legConf)
			if err != nil {
				logging.Logger.WithFields(logrus.Fields{
					"error":    err,
					"guild":    guild,
					"messsage": c.Message.ID,
					"command":  "import",
				}).Error("could not parse config file")
				continue
			}

			var newAdminRoleIds []string
			for _, role := range legConf.AdminRoleIds {
				newAdminRoleIds = append(newAdminRoleIds, fmt.Sprint(role))
			}
			var newModRoleIds []string
			for _, role := range legConf.ModRoleIds {
				newModRoleIds = append(newModRoleIds, fmt.Sprint(role))
			}
			var newUserRoleIds = make(map[string][]string)
			for key, group := range legConf.UserRoles {
				if key == "" {
					key = "Ungrouped"
				}
				newUserRoleIds[key] = make([]string, len(group))
				for i, role := range group {
					newUserRoleIds[key][i] = fmt.Sprint(role)
				}
			}
			var newRequiresRoles = make(map[string][]string)
			var newAutoRoleIds []string
			for _, role := range legConf.AutoRoleIds {
				newAutoRoleIds = append(newAutoRoleIds, fmt.Sprint(role))
			}
			var newMessageLoggingIgnoreChannelIds []string
			for _, channel := range legConf.MessageLoggingIgnoreChannels {
				newMessageLoggingIgnoreChannelIds = append(newMessageLoggingIgnoreChannelIds, fmt.Sprint(channel))
			}

			fmt.Println(fmt.Sprint(legConf.TrustedRolePointsThreshold))

			conf := entities.GuildConfig{
				GuildId:                        guild,
				Prefix:                         legConf.Prefix,
				AdminRoleIds:                   newAdminRoleIds,
				ModRoleIds:                     newModRoleIds,
				UserRoleIds:                    newUserRoleIds,
				RequiresRoles:                  newRequiresRoles,
				MutedRoleId:                    fmt.Sprint(legConf.MutedRoleId),
				MutedUsers:                     make(map[string]time.Time),
				AutoRoleIds:                    newAutoRoleIds,
				MessageLoggingChannelId:        fmt.Sprint(legConf.LoggingChannelid),
				MessageLoggingIgnoreChannelIds: newMessageLoggingIgnoreChannelIds,
				VerifiedRoleId:                 fmt.Sprint(legConf.VerifiedRole),
				VerificationMessage:            legConf.VerifiedMessage,
				JoinMessage:                    legConf.JoinMessage,
				JoinMessageChannelId:           fmt.Sprint(legConf.JoinMessageChannelId),
				PointsEnabled:                  legConf.PointsEnabled,
				TrustedRoleId:                  fmt.Sprint(legConf.TrustedRoleId),
				TrustedRolePointsThreshold:     legConf.TrustedRolePointsThreshold,
			}

			if err != nil {
				c.Reply("An error occured retrieving the configuration for your server")
				continue
			}

			err = database.SaveGuildConfig(guild, c.Message.ID, conf)
			if err != nil {
				c.Reply("An error occured saving the configuration for your server")
				continue
			}

			_, err = c.Reply("imported guildconfig for " + guild)
			if err != nil {
				continue
			}
		}
		// import bans
		for _, guild := range guilds {
			banFile, err := os.Open("dump/" + guild + "-bans.json")
			if err != nil {
				logging.Logger.WithFields(logrus.Fields{
					"error":    err,
					"guild":    guild,
					"messsage": c.Message.ID,
					"command":  "import",
				}).Error("could not open ban file")
				continue
			}
			banBytes, err := ioutil.ReadAll(banFile)
			if err != nil {
				logging.Logger.WithFields(logrus.Fields{
					"error":    err,
					"guild":    guild,
					"messsage": c.Message.ID,
					"command":  "import",
				}).Error("could not read ban file")
				continue
			}
			var legBans []legacy.Ban
			err = json.Unmarshal(banBytes, &legBans)
			if err != nil {
				logging.Logger.WithFields(logrus.Fields{
					"error":    err,
					"guild":    guild,
					"messsage": c.Message.ID,
					"command":  "import",
				}).Error("could not parse ban file")
				continue
			}

			for _, ban := range legBans {
				database.SaveBan(guild, entities.Ban{
					UserId: fmt.Sprint(ban.Id),
					Reason: ban.Reason,
					Until:  ban.BannedUntil,
				})
			}
			c.Reply(fmt.Sprintf("imported %d bans for %s", len(legBans), guild))
		}
		// import custom commands
		for _, guild := range guilds {
			ccFile, err := os.Open("dump/" + guild + "-customCommands.json")
			if err != nil {
				logging.Logger.WithFields(logrus.Fields{
					"error":    err,
					"guild":    guild,
					"messsage": c.Message.ID,
					"command":  "import",
				}).Error("could not open customCommands file")
				continue
			}
			ccbytes, err := ioutil.ReadAll(ccFile)
			if err != nil {
				logging.Logger.WithFields(logrus.Fields{
					"error":    err,
					"guild":    guild,
					"messsage": c.Message.ID,
					"command":  "import",
				}).Error("could not read customCommands file")
				continue
			}
			var customCommands []entities.CustomCommand
			err = json.Unmarshal(ccbytes, &customCommands)
			if err != nil {
				logging.Logger.WithFields(logrus.Fields{
					"error":    err,
					"guild":    guild,
					"messsage": c.Message.ID,
					"command":  "import",
				}).Error("could not parse customCommands file")
				continue
			}

			for _, command := range customCommands {
				database.SaveCustomCommand(guild, command)
			}
			c.Reply(fmt.Sprintf("imported %d customCommands for %s", len(customCommands), guild))
		}
		return nil
	},
}
