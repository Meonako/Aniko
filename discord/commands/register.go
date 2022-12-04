package commands

import (
	"github.com/Meonako/Aniko/config"
	"github.com/Meonako/go-logger/v2"

	"github.com/bwmarrin/discordgo"
)

func Register(discord *discordgo.Session) {
	LogMsg := "Registered command"
	if config.Conf.REGISTER_GUILD_ID != "" {
		LogMsg += " for " + config.Conf.REGISTER_GUILD_ID
	}
	LogMsg += ": "
	for _, v := range CommandsList {
		_, err := discord.ApplicationCommandCreate(discord.State.User.ID, config.Conf.REGISTER_GUILD_ID, v)
		if err != nil {
			logger.ToTerminalRedAndExitFormatIfError(err, "Cannot create '%v' commands: %v", v.Name, err)
		}
		logger.ToTerminal(logger.Yellow(LogMsg), logger.Cyan(v.Name))
	}
}
