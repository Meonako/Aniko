package commands

import (
	"github.com/bwmarrin/discordgo"
)

var CommandsHandler = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
	"generate":             Generate,
	"progress":             Progress,
	"preset":               GetStyles,
	"generate-from-preset": GeneratePreset,
	"set-url":              SetURL,
	"clear":                Clear,
}
