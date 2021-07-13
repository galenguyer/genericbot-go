package permissions

type PermissionLevel int

const (
	User PermissionLevel = iota
	Moderator
	Administrator
	GuildOwner
	BotAdministrator
	BotOwner
)

func (p PermissionLevel) String() string {
	return []string{"User", "Moderator", "Admin", "GuildOwner", "BotAdmin", "BotOwner"}[p]
}
