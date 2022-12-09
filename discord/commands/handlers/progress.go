package handlers

import (
	"fmt"

	"github.com/Meonako/Aniko/model"
	"github.com/Meonako/Aniko/utils"

	"github.com/Meonako/go-logger/v2"

	"github.com/bwmarrin/discordgo"
)

func Progress(discord *discordgo.Session, interaction *discordgo.InteractionCreate) {
	discord.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseDeferredChannelMessageWithSource,
	})

	progress := model.SendProgress()
	Webhook := &discordgo.WebhookParams{
		Embeds: []*discordgo.MessageEmbed{
			utils.CreateAdvanceEmbed(
				"Current Progress",
				[]*discordgo.MessageEmbedField{
					{
						Name:   "Overall Progress",
						Value:  progress.GetProgress(),
						Inline: true,
					},
					{
						Name: "Overall Sampling Steps",
						Value: fmt.Sprint(
							progress.State.JobNo*progress.State.TargetSamplingSteps+progress.State.CurrentStep,
							"/",
							progress.State.JobCount*progress.State.TargetSamplingSteps),
						Inline: true,
					},
					{
						Name:   "ETA",
						Value:  progress.GetETA(),
						Inline: true,
					},
					{
						Name:   "Current Image",
						Value:  fmt.Sprintf("%v / %v", progress.State.JobNo+1, progress.State.JobCount),
						Inline: true,
					},
					{
						Name:   "Current Image Sampling Step",
						Value:  fmt.Sprint(progress.State.CurrentStep, "/", progress.State.TargetSamplingSteps),
						Inline: true,
					},
				},
				utils.EmbedOptionalField{
					Description: "Keep in mind that some data might not be accurate. This is not my fault. These data are transfer directly from the API. There's nothing I can do",
				},
			),
		},
	}

	if image, ok := progress.GetCurrentImage(); ok {
		Webhook.Files = []*discordgo.File{
			{
				Name:        "current_image.png",
				ContentType: "image/png",
				Reader:      image,
			},
		}
	}

	_, err := discord.FollowupMessageCreate(interaction.Interaction, false, Webhook)
	logger.ToTerminalRedIfError(err)
}