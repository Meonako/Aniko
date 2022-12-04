package model

import (
	"github.com/Meonako/Aniko/http/api"

	"github.com/Meonako/go-logger/v2"

	"github.com/goccy/go-json"
)

type styles struct {
	Name           string `json:"name"`
	Prompt         string `json:"prompt"`
	NegativePrompt string `json:"negative_prompt"`
	Error          string `json:"ERROR"`
}

func SendStyles() ([]styles, string) {
	resp := []styles{}
	rawData := api.Styles()
	if len(rawData) <= 0 {
		logger.ToTerminalRed("Styles Raw Data is empty")
		return resp, "Styles Raw Data is empty"
	}

	err := json.Unmarshal(rawData, &resp)
	if err != nil {
		logger.ToTerminalRed(err)
		logger.ToTerminal(string(rawData))
	}

	for _, res := range resp {
		if res.Error != "" {
			logger.ToTerminalRed("ERROR: ", res.Error)
			return resp, res.Error
		}
	}

	return resp[1:], ""
}
