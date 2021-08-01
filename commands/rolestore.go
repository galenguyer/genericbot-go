package commands

import (
	"github.com/galenguyer/genericbot/entities"
)

var RoleStore = &entities.Command{
	Name:        "rolestore",
	Description: "Save or restore your userroles",
	Usage:       "[save]restore]",
	Action: func(c entities.Context) error {
		_, err := c.Reply("This command is not currently implemented")
		return err
	},
}
