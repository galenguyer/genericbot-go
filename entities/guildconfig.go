package entities

import "time"

type GuildConfig struct {
	GuildId                        string               `bson:"_id" json:"GuildId"`
	Prefix                         string               `bson:"Prefix" json:"Prefix"`
	AdminRoleIds                   []string             `bson:"AdminRoleIds" json:"AdminRoleIds"`
	ModRoleIds                     []string             `bson:"ModRoleIds" json:"ModRoleIds"`
	UserRoleIds                    map[string][]string  `bson:"UserRoleIds" json:"UserRoleIds"`
	RequiresRoles                  map[string][]string  `bson:"RequiresRoles" json:"RequiresRoles"` // {"":["",""]}
	MutedRoleId                    string               `bson:"MutedRoleId" json:"MutedRoleId"`
	MutedUsers                     map[string]time.Time `bson:"MutedUsers" json:"MutedUsers"` //{"":Date}
	AutoRoleIds                    []string             `bson:"AutoRoleIds" json:"AutoRoleIds"`
	MessageLoggingChannelId        string               `bson:"MessageLoggingChannelId" json:"MessageLoggingChannelId"`
	UserLoggingChannelId           string               `bson:"UserLoggingChannelId" json:"UserLoggingChannelId"`
	MessageLoggingIgnoreChannelIds []string             `bson:"MessageLoggingIgnoreChannelIds" json:"MessageLoggingIgnoreChannelIds"`
	VerifiedRoleId                 string               `bson:"VerifiedRoleId" json:"VerifiedRoleId"`
	VerificationMessage            string               `bson:"VerificationMessage" json:"VerificationMessage"`
	JoinMessage                    string               `bson:"JoinMessage" json:"JoinMessage"`
	JoinMessageChannelId           string               `bson:"JoinMessageChannelId" json:"JoinMessageChannelId"`
	PointsEnabled                  bool                 `bson:"PointsEnabled" json:"PointsEnabled"`
	TrustedRoleId                  string               `bson:"TrustedRoleId" json:"TrustedRoleId"`
	TrustedRolePointsThreshold     int                  `bson:"TrustedRolePointsThreshold" json:"TrustedRolePointsThreshold"`
}
