package legacy

type GuildConfig struct {
	Prefix                       string
	AdminRoleIds                 []uint64
	ModRoleIds                   []uint64
	UserRoles                    map[string][]uint64 `json:"UserRoles"`
	MutedRoleId                  uint64
	MutedUsers                   []uint64
	AutoRoleIds                  []uint64
	LoggingChannelid             uint64
	MessageLoggingIgnoreChannels []uint64
	VerifiedRole                 uint64
	VerifiedMessage              string
	JoinMessage                  string
	JoinMessageChannelId         uint64
	PointsEnabled                bool
	TrustedRoleId                uint64
	TrustedRolePointsThreshold   int
}
