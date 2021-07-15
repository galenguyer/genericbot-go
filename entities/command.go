package entities

import (
	"github.com/galenguyer/genericbot/logging"
	"github.com/galenguyer/genericbot/permissions"
	"github.com/sirupsen/logrus"
)

type Command struct {
	Name        string
	Description string
	Usage       string
	Permissions permissions.PermissionLevel
	Action      func(Context) error
}

func (c Command) Execute(ctx Context) {
	logging.Logger.WithFields(logrus.Fields{
		"module":  "handlers",
		"method":  "OnMessageRecieved",
		"guild":   ctx.Message.GuildID,
		"channel": ctx.Message.ChannelID,
		"author":  ctx.Message.Author.ID,
		"command": c.Name,
	}).Info("got command " + c.Name)

	// if the command's permissions are greater than the user's permissions, bail early
	if c.Permissions > ctx.Permissions {
		return
	}

	if err := c.Action(ctx); err != nil {
		if err != nil {
			logging.Logger.WithFields(logrus.Fields{
				"error":   err,
				"module":  "command",
				"method":  "Execute",
				"guild":   ctx.Message.GuildID,
				"channel": ctx.Message.ChannelID,
				"author":  ctx.Message.Author.ID,
				"command": c.Name,
			}).Error("error executing command " + c.Name)
		}
	}
}
