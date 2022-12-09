package commands

import (
	"github.com/Meonako/Aniko/discord/commands/handlers"
	"github.com/bwmarrin/discordgo"
)

var CommandsHandler = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
	"generate": handlers.Generate,
	"progress": handlers.Progress,
	"set-url":  handlers.SetURL,
	"clear":    handlers.Clear,
}
