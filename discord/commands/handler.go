package commands

import (
	"fmt"
	"strings"
	"time"

	"github.com/Meonako/Aniko/config"
	"github.com/Meonako/Aniko/model"
	"github.com/Meonako/Aniko/utils"

	"github.com/Meonako/go-logger/v2"

	"github.com/bwmarrin/discordgo"
)

func Generate(discord *discordgo.Session, interaction *discordgo.InteractionCreate) {
	err := discord.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseDeferredChannelMessageWithSource,
	})

	if err != nil {
		logger.ToTerminalRed(err)
		return
	}

	rawOptions := interaction.ApplicationCommandData().Options
	// options := map[string]any{}
	StringOptions := map[string]string{}
	IntergerOptions := map[string]int{}
	FloatOptions := map[string]float64{}
	for _, opt := range rawOptions {
		logger.ToTerminal("RECEIVED: ", opt.Name, " | ", opt.Value, " | ", opt.Type.String(), " | ")
		// options[opt.Name] = opt
		switch opt.Type.String() {
		case "String":
			StringOptions[opt.Name] = opt.StringValue()
		case "Integer":
			IntergerOptions[opt.Name] = int(opt.IntValue())
		case "Number":
			FloatOptions[opt.Name] = opt.FloatValue()
		default:
			logger.ToTerminalRed("Unexpected Type Received: ", opt.Type.String())
		}
	}

	API := &model.Txt2ImgAPI{
		Prompt:         model.GetDefault[string](model.PROMPT, StringOptions[model.PROMPT]),
		NegativePrompt: model.GetDefault[string](model.NEGATIVE_PROMPT, StringOptions[model.NEGATIVE_PROMPT]),
		SamplingMethod: model.GetDefault[string](model.SAMPLING_METHOD, StringOptions[model.SAMPLING_METHOD]),
		SamplingSteps:  model.GetDefault[int](model.SAMPLING_STEPS, IntergerOptions[model.SAMPLING_STEPS]),
		Width:          model.GetDefault[int](model.WIDTH, IntergerOptions[model.WIDTH]),
		Height:         model.GetDefault[int](model.HEIGHT, IntergerOptions[model.HEIGHT]),
		CFGScale:       model.GetDefault[float64](model.CFG_SCALE, FloatOptions[model.CFG_SCALE]),
		Seed:           model.GetDefault[int](model.SEED, IntergerOptions[model.SEED]),
		Count:          model.GetDefault[int](model.COUNT, IntergerOptions[model.COUNT]),
	}

	seedText := ""
	if API.Seed == 0 {
		seedText = "Random"
	} else {
		seedText = fmt.Sprint(API.Seed)
	}

	_, err = discord.FollowupMessageCreate(interaction.Interaction, false, &discordgo.WebhookParams{
		Embeds: []*discordgo.MessageEmbed{
			utils.CreateAdvanceEmbed(
				"Creating image(s) with your parameters...",
				[]*discordgo.MessageEmbedField{
					{Name: model.ToReadable(model.PROMPT), Value: API.Prompt},
					{Name: model.ToReadable(model.NEGATIVE_PROMPT), Value: API.NegativePrompt},

					{Name: model.ToReadable(model.SAMPLING_STEPS), Value: fmt.Sprint(API.SamplingSteps), Inline: true},
					{Name: model.ToReadable(model.SAMPLING_METHOD), Value: API.SamplingMethod, Inline: true},
					{Name: "\u200B", Value: "\u200B", Inline: true}, // EMPTY FIELD

					{Name: model.ToReadable(model.WIDTH), Value: fmt.Sprint(API.Width), Inline: true},
					{Name: model.ToReadable(model.HEIGHT), Value: fmt.Sprint(API.Height), Inline: true},
					{Name: "\u200B", Value: "\u200B", Inline: true}, // EMPTY FIELD

					{Name: model.ToReadable(model.CFG_SCALE), Value: fmt.Sprint(API.CFGScale), Inline: true},
					{Name: model.ToReadable(model.SEED), Value: seedText, Inline: true},
					{Name: model.ToReadable(model.COUNT), Value: fmt.Sprint(API.Count), Inline: true},
				},
			),
		},
	})

	if err != nil {
		logger.ToTerminalRed(err)
		_, err1 := discord.FollowupMessageCreate(interaction.Interaction, false, &discordgo.WebhookParams{
			Embeds: []*discordgo.MessageEmbed{
				{
					Title:       "ERROR",
					Color:       config.Conf.EMBED_ERROR_COLOR,
					Description: err.Error(),
					Footer:      config.Conf.EMBED_FOOTER,
				},
			},
		})
		logger.ToTerminalRedIfError(err1)
	}

	// API := generateInfo.GenerateAPI()
	resp := API.SendTXT2IMG()
	webhook := &discordgo.WebhookParams{}
	if resp == nil || resp.IsEmpty() {
		webhook.Embeds = []*discordgo.MessageEmbed{
			{
				Title:       "ERROR",
				Color:       config.Conf.EMBED_ERROR_COLOR,
				Description: strings.Replace(resp.ERROR, "ERROR: ", "", 1),
				Footer:      config.Conf.EMBED_FOOTER,
			},
		}
	} else {
		webhook.Embeds = []*discordgo.MessageEmbed{
			utils.CreateAdvanceEmbed(
				"Success",
				[]*discordgo.MessageEmbedField{},
				utils.EmbedOptionalField{
					Description: "All yours image(s) :)",
				},
			),
		}
		webhook.Files = resp.GenerateDiscordFile()
	}

	_, err = discord.FollowupMessageCreate(interaction.Interaction, false, webhook)
	logger.ToTerminalRedIfError(err)
}

func GetStyles(discord *discordgo.Session, interaction *discordgo.InteractionCreate) {
	err := discord.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseDeferredChannelMessageWithSource,
	})

	if err != nil {
		logger.ToTerminalRed(err)
		return
	}

	webhook := &discordgo.WebhookParams{}
	styles, serr := model.SendStyles()
	if serr != "" {
		webhook.Embeds = []*discordgo.MessageEmbed{
			{
				Title:       "ERROR",
				Description: serr,
				Color:       config.Conf.EMBED_ERROR_COLOR,
			},
		}
	} else {
		fields := []*discordgo.MessageEmbedField{}
		for _, style := range styles {
			fields = append(fields, &discordgo.MessageEmbedField{
				Name:  style.Name,
				Value: style.Prompt,
			})
		}

		webhook.Embeds = []*discordgo.MessageEmbed{
			utils.CreateAdvanceEmbed(
				"All Styles",
				fields,
			),
		}
	}

	_, err = discord.FollowupMessageCreate(interaction.Interaction, false, webhook)
	logger.ToTerminalRedIfError(err)
}

func GeneratePreset(discord *discordgo.Session, interaction *discordgo.InteractionCreate) {
	err := discord.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseDeferredChannelMessageWithSource,
	})

	if err != nil {
		logger.ToTerminalRed(err)
		return
	}

	API := model.Txt2ImgAPI{
		Styles: []string{
			interaction.ApplicationCommandData().Options[0].StringValue(),
		},
	}
	res := API.SendTXT2IMG()
	webhook := &discordgo.WebhookParams{}
	if res == nil || res.IsEmpty() {
		webhook.Embeds = []*discordgo.MessageEmbed{
			{
				Title:       "ERROR",
				Color:       config.Conf.EMBED_ERROR_COLOR,
				Description: strings.Replace(res.ERROR, "ERROR: ", "", 1),
			},
		}
	} else {
		webhook.Embeds = []*discordgo.MessageEmbed{
			utils.CreateAdvanceEmbed(
				"Sucess",
				[]*discordgo.MessageEmbedField{},
				utils.EmbedOptionalField{
					Description: "All yours image(s) :)",
				},
			),
		}
		webhook.Files = res.GenerateDiscordFile()
	}

	_, err = discord.FollowupMessageCreate(interaction.Interaction, false, webhook)
	logger.ToTerminalRedIfError(err)
}

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

	discord.FollowupMessageCreate(interaction.Interaction, true, &discordgo.WebhookParams{
		Embeds: []*discordgo.MessageEmbed{
			{
				Title:       "Success",
				Description: fmt.Sprintf("Deleted %v messages", len(msgId)),
			},
		},
	})
}
