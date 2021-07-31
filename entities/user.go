package entities

import (
	"fmt"
	"strings"
	"time"
)

type User struct {
	Id              string `bson:"_id"`
	Usernames       []string
	Nicknames       []string
	Warnings        []string
	RoleStore       []string
	Points          int
	LastPointsAdded time.Time
	Messages        int
}

func (u User) String() string {
	var output string

	output += fmt.Sprintf("**User:** <@!%s>\n", u.Id)
	output += fmt.Sprintf("**Id:** `%s`\n", u.Id)

	usernames := "**Usernames:** "
	if len(u.Usernames) > 0 {
		for _, un := range u.Usernames {
			usernames += fmt.Sprintf("`%s`, ", strings.ReplaceAll(un, "`", "'"))
		}
		usernames = strings.TrimRight(usernames, ", ")
	} else {
		usernames += "None Found"
	}
	output += usernames + "\n"

	nicknames := "**Nicknames:** "
	if len(u.Nicknames) > 0 {
		for _, nn := range u.Nicknames {
			nicknames += fmt.Sprintf("`%s`, ", strings.ReplaceAll(nn, "`", "'"))
		}
		nicknames = strings.TrimRight(nicknames, ", ")
	} else {
		nicknames += "None Found"
	}
	output += nicknames + "\n"

	warnings := "**Warnings:** "
	if len(u.Warnings) > 0 {
		for _, w := range u.Warnings {
			warnings += fmt.Sprintf("%s, ", w)
		}
		warnings = strings.TrimRight(warnings, ", ")
		output += warnings
	}

	return output
}
