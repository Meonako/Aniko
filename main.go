package main

import (
	"os"
	"os/signal"

	"github.com/Meonako/Aniko/config"
	"github.com/Meonako/Aniko/discord/commands"
	"github.com/Meonako/Aniko/discord/event"
	"github.com/Meonako/Aniko/env"
	_ "github.com/Meonako/Aniko/extension"

	"github.com/Meonako/go-logger/v2"

	"github.com/bwmarrin/discordgo"
	"github.com/fatih/color"
)

func init() {
	err := env.Load()
	logger.ToTerminalRedAndExitIfError(err, "Cannot load ENV: ", err)
}

func main() {
	discord, err := discordgo.New("Bot " + os.Getenv("TOKEN"))
	logger.ToTerminalRedAndExitIfError(err, "Invalid bot parameter: ", err)

	event.Listen(discord)

	err = discord.Open()
	logger.ToTerminalRedAndExitIfError(err, "Cannot open the session: ", err)

	commands.Register(discord)

	defer discord.Close()

	logger.ToTerminal("Current API BASE URL: ", logger.Cyan(config.Conf.BASE_URL))
	logger.ToTerminal(logger.Yellow("Ready.") + " Press " + logger.Colorize(" CTRL + C ", color.FgHiMagenta, color.BgBlack, color.Underline, color.Bold) + " to exit.")
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	<-interrupt

	logger.ToTerminal(logger.Yellow("Cleaning up..."))
	config.Save()

	registeredCommands, err := discord.ApplicationCommands(discord.State.User.ID, config.Conf.REGISTER_GUILD_ID)
	logger.ToTerminalRedAndExitIfError(err, "Cannot get commands: ", err)

	for _, v := range registeredCommands {
		err = discord.ApplicationCommandDelete(discord.State.User.ID, config.Conf.REGISTER_GUILD_ID, v.ID)
		logger.ToTerminalRedAndExitFormatIfError(err,
			"Cannot delete '%v' command: %v",
			v.Name, err,
		)
		logger.ToTerminal(logger.Green("Unregistered command: "), logger.Cyan(v.Name))
	}

	logger.ToTerminal(logger.Yellow("Gracefully shutting down."))
}
