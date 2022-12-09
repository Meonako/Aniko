package handlers

import (
	"github.com/Meonako/Aniko/config"
	"github.com/Meonako/Aniko/utils"

	"github.com/bwmarrin/discordgo"
)

func SetURL(discord *discordgo.Session, interaction *discordgo.InteractionCreate) {
	user := utils.GetUser(interaction)
	var response *discordgo.MessageEmbed
	if user.ID != config.Conf.OWNER.ID {
		response = &discordgo.MessageEmbed{
			Title: "You are not allow to. :)",
			Color: config.Conf.EMBED_ERROR_COLOR,
		}
	} else {
		config.Conf.BASE_URL = interaction.ApplicationCommandData().Options[0].StringValue()
		response = &discordgo.MessageEmbed{
			Title:       "Success!",
			Description: "Base URL is now: " + config.Conf.BASE_URL,
		}
	}

	discord.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Flags:  discordgo.MessageFlagsEphemeral,
			Embeds: []*discordgo.MessageEmbed{response},
		},
	})
}
