package commands

import (
	"strconv"

	"github.com/galenguyer/genericbot/database"
	"github.com/galenguyer/genericbot/entities"
	"github.com/galenguyer/genericbot/logging"
	"github.com/galenguyer/genericbot/permissions"
	"github.com/sirupsen/logrus"
)

var Find = &entities.Command{
	Name:        "find",
	Description: "Search for a user from the database",
	Usage:       "<userid>",
	Permissions: permissions.Moderator,
	Action: func(c entities.Context) error {
		if len(c.ParsedCommand.ParameterString) == 0 {
			c.Reply("Please give me a user to search for")
			return nil
		}
		// check if we were given a valid user id
		id, err := strconv.ParseUint(c.ParsedCommand.ParameterString, 10, 64)
		if err == nil && id > 10000000000000000 {
			// if we're given a valid userid
			user, err := database.GetUser(c.GuildId, c.ParsedCommand.ParameterString)
			if err != nil {
				logging.Logger.WithFields(logrus.Fields{
					"error":    err,
					"guild":    c.GuildId,
					"channel":  c.ChannelId,
					"author":   c.Message.Author.ID,
					"messsage": c.Message.ID,
					"command":  "find",
				}).Error("error finding user")
				c.Reply("An error occured while looking up a user")
				return err
			} else if user == nil {
				c.Reply("I couldn't find a user with that ID")
			} else {
				c.ReplySplit(user.String(), "`,")
			}
		} else {
			// if we're given a string to search for
			c.Reply("Searching for users by string is not yet implemented")
		}
		return nil
	},
}
