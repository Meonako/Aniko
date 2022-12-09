package handlers

import (
	"fmt"
	"time"

	"github.com/Meonako/Aniko/config"

	"github.com/Meonako/go-logger/v2"

	"github.com/bwmarrin/discordgo"
)

func Clear(discord *discordgo.Session, interaction *discordgo.InteractionCreate) {
	err := discord.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{
				{
					Title: "Deleting messages...",
				},
			},
		},
	})

	if err != nil {
		logger.ToTerminalRed(err)
		return
	}

	messages, err := discord.ChannelMessages(interaction.ChannelID, 100, "", "", "")
	if err != nil {
		logger.ToTerminalRed(err)
		_, err1 := discord.FollowupMessageCreate(interaction.Interaction, true, &discordgo.WebhookParams{
			Embeds: []*discordgo.MessageEmbed{
				{
					Title: "ERROR",
					Color: config.Conf.EMBED_ERROR_COLOR,
				},
			},
		})
		logger.ToTerminalRedIfError(err1)
		return
	}

	msgId := []string{}
	for _, message := range messages {
		if message.Timestamp.Before(time.Now().Add(-(14 * 24 * time.Hour))) {
			continue
		}

		msgId = append(msgId, message.ID)
	}

	err = discord.ChannelMessagesBulkDelete(interaction.ChannelID, msgId)
	if err != nil {
		logger.ToTerminalRed(err)
		_, err1 := discord.FollowupMessageCreate(interaction.Interaction, true, &discordgo.WebhookParams{
			Embeds: []*discordgo.MessageEmbed{
				{
					Title: "ERROR",
					Color: config.Conf.EMBED_ERROR_COLOR,
				},
			},
		})
		logger.ToTerminalRedIfError(err1)
		return
	}

	_, err = discord.FollowupMessageCreate(interaction.Interaction, true, &discordgo.WebhookParams{
		Embeds: []*discordgo.MessageEmbed{
			{
				Title:       "Success",
				Description: fmt.Sprintf("Deleted %v messages", len(msgId)),
			},
		},
	})
	logger.ToTerminalRedIfError(err)
}
