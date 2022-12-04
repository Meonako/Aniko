package utils

import (
	"bytes"
	"encoding/base64"

	"github.com/Meonako/go-logger/v2"
)

var i = 0

func DecodeBase64ToImage(base64Data string) *bytes.Reader {
	logger.ToTerminal("Start Decode Base64: ", i)
	decoded, err := base64.StdEncoding.DecodeString(base64Data)
	if err != nil {
		logger.ToTerminalRed("Cannot decode base64: ", err)
		return nil
	}

	i++
	logger.ToTerminal("Finished Decode Base64: ", i)
	return bytes.NewReader(decoded)
}
