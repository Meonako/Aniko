package event

import (
	"time"

	"github.com/Meonako/Aniko/config"
	"github.com/Meonako/Aniko/discord/commands"
	"github.com/Meonako/Aniko/utils"

	"github.com/Meonako/go-logger/v2"

	"github.com/bwmarrin/discordgo"
	"github.com/fatih/color"
)

var (
	AlreadyInGuild = []string{}
)

func Ready(discord *discordgo.Session, ready *discordgo.Ready) {
	logger.ToTerminal("Logged in as:", logger.Colorize(
		utils.GetBotFullUsername(discord), color.FgHiBlue, color.Underline),
	)

	for _, guild := range discord.State.Guilds {
		AlreadyInGuild = append(AlreadyInGuild, guild.ID)
	}

	go func() {
		owner, err := discord.User(config.Conf.OWNER_ID)
		if err != nil {
			logger.ToTerminalRed(err)
			return
		}

		config.Conf.OWNER = owner

		config.Conf.EMBED_AUTHOR = &discordgo.MessageEmbedAuthor{
			Name:    owner.Username,
			URL:     "https://github.com/Meonako",
			IconURL: owner.AvatarURL(""),
		}

		config.Conf.EMBED_FOOTER = &discordgo.MessageEmbedFooter{
			Text:    "Need help? Contact: " + utils.GetFullUsername(owner),
			IconURL: owner.AvatarURL(""),
		}
	}()
}

func InteractionCreate(discord *discordgo.Session, interaction *discordgo.InteractionCreate) {
	switch interaction.Type {
	case discordgo.InteractionApplicationCommand:
		user := utils.GetUser(interaction)

		logger.ToTerminalFormat("%v used : %v",
			logger.Green(utils.GetFullUsername(user)),
			logger.Cyan(interaction.ApplicationCommandData().Name),
		)

		if handler, ok := commands.CommandsHandler[interaction.ApplicationCommandData().Name]; ok {
			handler(discord, interaction)
		}
	}
}

func GuildCreate(discord *discordgo.Session, guildCreate *discordgo.GuildCreate) {
	time.Sleep(1 * time.Second) // Delayed a little so that Ready event can finish adding guilds
	if utils.Contains(AlreadyInGuild, guildCreate.ID) {
		return
	}

	logger.ToTerminal("JOINED: ", logger.Green(guildCreate.ID), ":", logger.Green(guildCreate.Name))
	if guildCreate.SystemChannelID != "" {
		_, err := discord.ChannelMessageSend(guildCreate.SystemChannelID, "I have ARRIVED!")
		if err != nil {
			logger.ToTerminal("Cannot send message to channel", guildCreate.SystemChannelID, ":", err)
		}
	}
}
