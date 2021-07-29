package entities

import "time"

type Ban struct {
	UserId string    `bson:"_id" json:"UserId"`
	Until  time.Time `bson:"Until" json:"Until"`
	Reason string    `bson:"Reason" json:"Reason"`
}
