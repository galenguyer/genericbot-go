package commands

import (
	"fmt"
	"math/rand"
	"unicode"

	"github.com/galenguyer/genericbot/entities"
)

var Mock = &entities.Command{
	Name:        "mock",
	Description: "Mock someone ruthlessly",
	Usage:       "<text>",
	Action: func(c entities.Context) error {
		_, err := c.Reply(mockText(c.ParsedCommand.ParameterString))
		return err
	},
}

func mockText(input string) string {
	output := ""
	ignore := false
	for _, char := range input {
		if ignore {
			if char == ':' || char == '>' {
				ignore = false
			}
			output += fmt.Sprintf("%c", char)
		} else {
			if char == ':' || char == '<' {
				ignore = true
				output += fmt.Sprintf("%c", char)
			} else if ignore {
				output += fmt.Sprintf("%c", char)
			} else {
				if rand.Intn(2) == 0 {
					char = unicode.ToUpper(char)
				} else {
					char = unicode.ToLower(char)
				}
				output += fmt.Sprintf("%c", char)
			}
		}
	}
	return output
}
