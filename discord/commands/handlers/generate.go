package handlers

import (
	"fmt"
	"strings"

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
