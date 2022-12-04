package model

import "strings"

const (
	PROMPT          = "prompt"
	NEGATIVE_PROMPT = "negative-prompt"
	SAMPLING_STEPS  = "sampling-steps"
	SAMPLING_METHOD = "sampling-method"
	WIDTH           = "width"
	HEIGHT          = "height"
	CFG_SCALE       = "cfg-scale"
	SEED            = "seed"
	COUNT           = "count"
)

func ToReadable(text string) string {
	return strings.ToUpper(strings.ReplaceAll(text, "-", " "))
}
