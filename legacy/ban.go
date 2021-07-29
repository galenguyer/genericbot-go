package legacy

import "time"

type Ban struct {
	Id          uint64    `json:"Id"`
	BannedUntil time.Time `json:"BannedUntil"`
	Reason      string    `json:"Reason"`
}
