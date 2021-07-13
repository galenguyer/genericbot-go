package entities

import "github.com/galenguyer/genericbot/config"

type Context struct {
	Config        config.Config
	ParsedCommand ParsedCommand
}
