package commands

import (
	"fmt"
	"time"

	"github.com/galenguyer/genericbot/entities"
	"github.com/tj/go-naturaldate"
)

var Time = &entities.Command{
	Name:        "time",
	Description: "Take a human readable time string and return a Discord Time object",
	Usage:       "<time>",
	Action: func(c entities.Context) error {
		value, err := naturaldate.Parse(c.ParsedCommand.ParameterString, time.Now().UTC(), naturaldate.WithDirection(naturaldate.Future))
		if err != nil {
			c.Reply("An error occured parsing that time string")
			return err
		}
		_, err = c.Reply(fmt.Sprintf(
			"Parsed \"%s\" as <t:%d>\n*(If this looks wrong, sorry! I don't know about time zones and parsing time is hard)*",
			c.ParsedCommand.ParameterString,
			value.UTC().Unix()),
		)
		return err
	},
}
