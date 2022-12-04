package event

import (
	"github.com/bwmarrin/discordgo"
)

var EventList = []any{
	Ready,
	InteractionCreate,
	GuildCreate,
}

func Listen(discord *discordgo.Session) {
	for _, event := range EventList {
		discord.AddHandler(event)
	}
}
