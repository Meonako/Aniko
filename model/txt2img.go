package model

import (
	"strings"

	"github.com/Meonako/Aniko/http/api"
	"github.com/Meonako/Aniko/utils"

	"github.com/Meonako/go-logger/v2"

	"github.com/goccy/go-json"
)

var (
	DefaultValue = map[string]any{
		"sampling-steps-Min":     10,
		"sampling-steps-Default": 28,
		"sampling-steps-Max":     35,
		"width-Min":              192,
		"width-Default":          512,
		"width-Max":              1088,
		"height-Min":             192,
		"height-Default":         512,
		"height-Max":             1088,
		"cfg-scale-Min":          5.0,
		"cfg-scale-Default":      12.0,
		"cfg-scale-Max":          20.0,
		"count-Min":              1,
		"count-Default":          1,
		"count-Max":              4,

		SAMPLING_METHOD + "-Default": "Euler",
		NEGATIVE_PROMPT + "-Default": "lowres, bad anatomy, bad hands, text, error, missing fingers, extra digit, fewer digits, cropped, worst quality, low quality, normal quality, jpeg artifacts,signature, watermark, username, blurry, artist name",
	}
)

type Txt2ImgAPI struct {
	Prompt         string  `json:"prompt,omitempty"`
	NegativePrompt string  `json:"negative_prompt,omitempty"`
	SamplingSteps  int     `json:"steps,omitempty"`
	SamplingMethod string  `json:"sampler_index,omitempty"`
	Width          int     `json:"width,omitempty"`
	Height         int     `json:"height,omitempty"`
	CFGScale       float64 `json:"cfg_scale,omitempty"`
	Seed           int     `json:"seed,omitempty"`
	Count          int     `json:"n_iter,omitempty"`

	Styles []string `json:"styles,omitempty"`
}

func GetDefault[T string | int | float64](key string, currentValue any) T {
	var value T
	if currentValue != nil {
		value = currentValue.(T)
	} else {
		value = utils.GetNonNil[T](currentValue)
	}

	if def, ok := DefaultValue[key+"-Default"]; ok && value == utils.GetZeroValue[T]() {
		return def.(T)
	} else if min, ok := DefaultValue[key+"-Min"]; ok && value < min.(T) {
		return min.(T)
	} else if max, ok := DefaultValue[key+"-Max"]; ok && value > max.(T) {
		return max.(T)
	}

	return value
}

func (api *Txt2ImgAPI) SendTXT2IMG() *txt2ImgRespond {
	resp := txt2ImgRespond{}
	rawData := api.PostInfo()
	if rawData == "" {
		logger.ToTerminalRed("TXT2IMG Raw data is empty.")
		return &resp
	} else if strings.HasPrefix(rawData, "ERROR") {
		logger.ToTerminalRed(rawData)
		resp.ERROR = rawData
		return &resp
	}

	logger.ToTerminalRedIfError(json.Unmarshal([]byte(rawData), &resp))
	return &resp
}

func (ins *Txt2ImgAPI) PostInfo() string {
	Json, err := json.Marshal(ins)
	logger.ToTerminalRedFormatIfError(err, "Cannot marshal data to JSON: %v")
	return api.TXT2Img(Json)
}
