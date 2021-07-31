package commands

import (
	"github.com/galenguyer/genericbot/entities"
)

var GitHub = &entities.Command{
	Name:        "github",
	Aliases:     []string{"source"},
	Description: "Send the GitHub link for the bot",
	Action: func(c entities.Context) error {
		_, err := c.Reply("You can find me on GitHub at <https://github.com/galenguyer/genericbot>!")
		return err
	},
}
