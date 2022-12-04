package model

import (
	"bytes"
	"fmt"
	"math"

	"github.com/Meonako/Aniko/http/api"
	"github.com/Meonako/Aniko/utils"

	"github.com/Meonako/go-logger/v2"

	"github.com/goccy/go-json"
)

type progress struct {
	Progress     float64 `json:"progress"`
	ETA          float64 `json:"eta_relative"`
	State        state   `json:"state"`
	CurrentImage string  `json:"current_image"`
	Error        string  `json:"ERROR"`
}

type state struct {
	Skipped             bool   `json:"skipped"`
	Interrupted         bool   `json:"interrupted"`
	Job                 string `json:"job"`
	JobCount            int    `json:"job_count"`
	JobNo               int    `json:"job_no"`
	CurrentStep         int    `json:"sampling_step"`
	TargetSamplingSteps int    `json:"sampling_steps"`
}

func SendProgress() *progress {
	resp := progress{}
	rawData := api.Progress()
	if len(rawData) <= 0 {
		logger.ToTerminalRed("Progress Raw Data is empty")
		return &progress{}
	}

	logger.ToTerminalRedIfError(json.Unmarshal(rawData, &resp))
	if resp.Error != "" {
		logger.ToTerminalRed("ERROR: ", resp.Error)
		return &progress{}
	}

	return &resp
}

func (p *progress) GetProgress() string {
	return fmt.Sprintf("%.2f", p.Progress*100) + "%"
}

func (p *progress) GetETA() string {
	return fmt.Sprintf("%v Minutes %.2f Seconds", math.Floor(p.ETA/60), math.Mod(p.ETA, 60))
}

func (p *progress) GetCurrentImage() (*bytes.Reader, bool) {
	if p.CurrentImage == "" {
		return nil, false
	}

	return utils.DecodeBase64ToImage(p.CurrentImage), true
}
