package loaders

import (
	"fmt"
	"sort"
	"strings"

	"github.com/galenguyer/genericbot/commands"
	"github.com/galenguyer/genericbot/entities"
	"github.com/galenguyer/genericbot/logging"
	"github.com/galenguyer/genericbot/permissions"
	"github.com/sirupsen/logrus"
)

var (
	Commands []*entities.Command
)

var Help = &entities.Command{
	Name:        "help",
	Aliases:     []string{"halp"},
	Description: "Get a list of available commands or the description and useage for a specific command",
	Usage:       "<command?>",
	Action: func(c entities.Context) error {
		if len(c.ParsedCommand.ParameterString) != 0 {
			_, err := c.Reply("Sorry, searching for commands is not yet implemented")
			return err
		} else {
			message := ""

			for level := 0; level <= int(c.Permissions); level++ {
				message += permissions.PermissionLevel(level).String() + ": "

				for _, command := range Commands {
					if command.Permissions == permissions.PermissionLevel(level) {
						message += fmt.Sprintf("`%s`, ", command.Name)
					}
				}

				message = strings.TrimRight(message, ", ")
				message += "\n"
			}

			_, err := c.ReplySplit(message)
			return err
		}
	},
}

func init() {
	Commands = append(Commands, commands.Config)
	Commands = append(Commands, commands.Echo)
	Commands = append(Commands, commands.Find)
	Commands = append(Commands, commands.GitHub)
	Commands = append(Commands, Help)
	Commands = append(Commands, commands.Import)
	Commands = append(Commands, commands.Mock)
	Commands = append(Commands, commands.Ping)
	Commands = append(Commands, commands.RoleStore)
	Commands = append(Commands, commands.Time)

	sort.Slice(Commands, func(i, j int) bool {
		return Commands[i].Name > Commands[j].Name
	})

	logging.Logger.WithFields(logrus.Fields{"module": "handlers", "method": "init"}).Infof("loaded %d commands", len(Commands))
}
