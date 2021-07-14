package entities

import "time"

type GuildConfig struct {
	Prefix                       string               `json:"Prefix"`
	AdminRoleIds                 []string             `json:"AdminRoleIds"`
	ModRoleIds                   []string             `json:"ModRoleIds"`
	UserRoleIds                  []string             `json:"UserRoleIds"`
	RequiresRoles                map[string][]string  `json:"RequiresRoles"` // {"":["",""]}
	MutedRoleId                  string               `json:"MutedRoleId"`
	MutedUsers                   map[string]time.Time `json:"MutedUsers"` //{"":Date}
	AutoRoleIds                  []string             `json:"AutoRoleIds"`
	MessageLoggingChannelId      string               `json:"MessageLoggingChannelId"`
	UserLoggingChannelId         string               `json:"UserLoggingChannelId"`
	MessageLoggingIgnoreChannels []string             `json:"MessageLoggingIgnoreChannels"`
	VerifiedRoleId               string               `json:"VerifiedRoleId"`
	VerificationMessage          string               `json:"VerificationMessage"`
	JoinMessage                  string               `json:"JoinMessage"`
	JoinMessageChannelId         string               `json:"JoinMessageChannelId"`
	PointsEnabled                bool                 `json:"PointsEnabled"`
	TrustedRoleId                string               `json:"TrustedRoleId"`
	TrustedRolePointsThreshold   int                  `json:"TrustedRolePointsThreshold"`
}
