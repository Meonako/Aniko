package utils

import (
	// "fmt"

	"github.com/Meonako/Aniko/config"

	"github.com/bwmarrin/discordgo"
)

// Get Discord Full Username (e.g. test#1234)
func GetFullUsername(user *discordgo.User) string {
	return user.Username + "#" + user.Discriminator
}

// Get Discord Full Username of self (e.g. test#1234)
func GetBotFullUsername(bot *discordgo.Session) string {
	return bot.State.User.Username + "#" + bot.State.User.Discriminator
}

func GetUser(interaction *discordgo.InteractionCreate) *discordgo.User {
	var user *discordgo.User
	if interaction.User != nil {
		user = interaction.User
	} else {
		user = interaction.Member.User
	}

	return user
}

func CreateAdvanceEmbed(title string, fileds []*discordgo.MessageEmbedField, optional ...EmbedOptionalField) *discordgo.MessageEmbed {
	embed := &discordgo.MessageEmbed{
		Title: title,
		Author: &discordgo.MessageEmbedAuthor{
			Name: config.Conf.OWNER.Username,
			URL:  "https://github.com/Meonako",
		},
		Footer: &discordgo.MessageEmbedFooter{
			Text:    "Need help? Contact: ᴍᴇᴏɴᴀᴋᴏ#7724",
			IconURL: config.Conf.OWNER.AvatarURL(""),
		},
		Color: 0x0091f9,
	}

	if len(fileds) > 0 {
		embed.Fields = fileds
	}

	if len(optional) > 0 {
		Data := optional[0]
		if Data.Description != "" {
			embed.Description = Data.Description
		}

		if Data.Image != "" {
			embed.Image = &discordgo.MessageEmbedImage{
				URL: Data.Image,
			}
		}

		if Data.Thumbnail != "" {
			embed.Thumbnail = &discordgo.MessageEmbedThumbnail{
				URL: Data.Thumbnail,
			}
		}
	}

	// if len(optional) > 0 {
	// 	for _, option := range optional {
	// 		switch fmt.Sprintf("%T", option) {
	// 		case "*discordgo.MessageEmbedThumbnail":
	// 			if option.(*discordgo.MessageEmbedThumbnail).URL != "" {
	// 				embed.Thumbnail = option.(*discordgo.MessageEmbedThumbnail)
	// 			}
	// 		case "*discordgo.MessageEmbedImage":
	// 			if option.(*discordgo.MessageEmbedImage).URL != "" {
	// 				embed.Image = option.(*discordgo.MessageEmbedImage)
	// 			}
	// 		}
	// 	}
	// }

	return embed
}

func CreateSimpleEmbed(msg string) *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		Description: msg,
		Color:       0x0091f9,
	}
}

func BoldText(text string) string {
	return "**" + text + "**"
}
