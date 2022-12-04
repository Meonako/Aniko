package model

import (
	"bytes"
	"strconv"

	"github.com/Meonako/Aniko/utils"

	"github.com/bwmarrin/discordgo"
)

type txt2ImgRespond struct {
	Images     []string   `json:"images"`     // This one is base64 encoded
	Parameters Txt2ImgAPI `json:"parameters"` // Same as the information sent
	Info       string     `json:"info"`       // Same as Parameters field but in long long string
	ERROR      string
}

// Return reader that is ready to sent to discord
func (res *txt2ImgRespond) DecodeImage(index int) *bytes.Reader {
	if index >= len(res.Images) {
		return nil
	}

	rawImage := res.Images[index]
	return utils.DecodeBase64ToImage(rawImage)
}

func (res *txt2ImgRespond) GenerateDiscordFile() []*discordgo.File {
	var files []*discordgo.File

	for index := range res.Images {
		files = append(files, &discordgo.File{
			Name:        "images (" + strconv.Itoa(index+1) + ").png",
			ContentType: "image/png",
			Reader:      res.DecodeImage(index),
		})
	}

	return files
}

func (res *txt2ImgRespond) IsEmpty() bool {
	return res.Images == nil || res.Info == ""
}
