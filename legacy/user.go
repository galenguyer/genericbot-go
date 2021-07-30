package legacy

import "time"

type User struct {
	Id              uint64
	Usernames       []string
	Nicknames       []string
	Warnings        []string
	RoleStore       []uint64
	Points          int
	LastPointsAdded time.Time
	Messages        int
}
