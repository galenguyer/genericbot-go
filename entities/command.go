package entities

import "github.com/galenguyer/genericbot/permissions"

type Command struct {
	Name        string
	Permissions permissions.PermissionLevel
	Action      func(Context) error
}

func (c Command) Execute(ctx Context) error {
	// if the command's permissions are greater than the user's permissions, bail early
	if c.Permissions > ctx.Permissions {
		return nil
	}
	return c.Action(ctx)
}
