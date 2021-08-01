package loaders

import (
	"github.com/galenguyer/genericbot/commands"
	"github.com/galenguyer/genericbot/entities"
	"github.com/galenguyer/genericbot/logging"
	"github.com/sirupsen/logrus"
)

var (
	Commands []*entities.Command
)

func init() {
	Commands = append(Commands, commands.Config)
	Commands = append(Commands, commands.Echo)
	Commands = append(Commands, commands.Find)
	Commands = append(Commands, commands.GitHub)
	Commands = append(Commands, commands.Import)
	Commands = append(Commands, commands.Mock)
	Commands = append(Commands, commands.Ping)
	Commands = append(Commands, commands.RoleStore)
	Commands = append(Commands, commands.Time)
	logging.Logger.WithFields(logrus.Fields{"module": "handlers", "method": "init"}).Infof("loaded %d commands", len(Commands))
}
