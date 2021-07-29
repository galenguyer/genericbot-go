package legacy

import "time"

type Ban struct {
	Id          uint64    `json:"Id"`
	BannedUntli time.Time `json:"BannedUntil"`
	Reason      string    `json:"Reason"`
}
