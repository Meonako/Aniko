package commands

import (
	"strconv"

	"github.com/Meonako/Aniko/model"

	"github.com/bwmarrin/discordgo"
)

var CommandsList = []*discordgo.ApplicationCommand{
	{
		Name:        "generate",
		Description: "Generate image based on options",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Name: model.PROMPT,
				NameLocalizations: map[discordgo.Locale]string{
					discordgo.Thai: "คำกระตุ้น",
				},
				Description: "What you want the AI to think about.",
				DescriptionLocalizations: map[discordgo.Locale]string{
					discordgo.Thai: "อยากให้บอทคิดเกี่ยวกับอะไร. **ไม่ได้ทดสอบในภาษาไทย**",
				},
				Type:     discordgo.ApplicationCommandOptionString,
				Required: true,
			},
			{
				Name:        "negative-prompt",
				Description: "What you want the AI to avoid. At the very least \"TRY TO\".",
				Type:        discordgo.ApplicationCommandOptionString,
			},
			{
				Name:        "sampling-steps",
				Description: "How long do you want AI to spends working on the image(s). Default: 28",
				Type:        discordgo.ApplicationCommandOptionInteger,
				// MaxValue:    float64(model.DefaultValue["sampling-steps-Max"]),
				MaxValue: float64(model.DefaultValue["sampling-steps-Max"].(int)),
			},
			{
				Name:        model.SAMPLING_METHOD,
				Description: "How the AI thinks about your image. You can randomly select. Default: Euler",
				Type:        discordgo.ApplicationCommandOptionString,
				Choices: []*discordgo.ApplicationCommandOptionChoice{
					{
						Name:  "Euler a",
						Value: "Euler a",
					},
					{
						Name:  "Euler",
						Value: "Euler",
					},
					{
						Name:  "DDIM",
						Value: "DDIM",
					},
				},
			},
			{
				Name:        "width",
				Description: "Width of the image. Default: 512",
				Type:        discordgo.ApplicationCommandOptionInteger,
				Choices:     sizeChoices(),
			},
			{
				Name:        "height",
				Description: "Height of the image. Default: 512",
				Type:        discordgo.ApplicationCommandOptionInteger,
				Choices:     sizeChoices(),
			},
			{
				Name:        "cfg-scale",
				Description: "How \"focused\" you want the AI to be on your prompt. Default: 12.0",
				Type:        discordgo.ApplicationCommandOptionNumber,
				MaxValue:    model.DefaultValue["cfg-scale-Max"].(float64),
			},
			{
				Name:        "seed",
				Description: "A value that determines the output of random number generator. Default: Random",
				Type:        discordgo.ApplicationCommandOptionInteger,
			},
			{
				Name:        "count",
				Description: "How many sequential batches to run. Default: 1. Max: " + strconv.Itoa(model.DefaultValue["count-Max"].(int)),
				Type:        discordgo.ApplicationCommandOptionInteger,
				MaxValue:    float64(model.DefaultValue["count-Max"].(int)),
			},
		},
	},
	{
		Name:        "progress",
		Description: "Get Generation Progress",
	},
	{
		Name:        "set-url",
		Description: "Don't bother with it. It's for me (dev) only.",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Name:        "url",
				Description: "Links",
				Type:        discordgo.ApplicationCommandOptionString,
				Required:    true,
			},
		},
	},
	{
		Name:        "clear",
		Description: "Delete 100 message from channel",
	},
}

func sizeChoices() []*discordgo.ApplicationCommandOptionChoice {
	choices := []*discordgo.ApplicationCommandOptionChoice{}
	for i := 192; i < 1089; i += 64 {
		choices = append(choices, &discordgo.ApplicationCommandOptionChoice{
			Name:  strconv.Itoa(i),
			Value: i,
		})
	}
	return choices
}
