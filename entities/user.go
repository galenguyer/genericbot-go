package entities

import "time"

type User struct {
	Id              string `bson:"_id"`
	Usernames       []string
	Nicknames       []string
	Warnings        []string
	Rolestore       []string
	Points          int
	LastPointsAdded time.Time
	Messages        int
}
