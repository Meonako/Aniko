package utils

import (
	"bytes"
	"encoding/base64"

	"github.com/Meonako/go-logger/v2"
)

func DecodeBase64ToImage(base64Data string) *bytes.Reader {
	decoded, err := base64.StdEncoding.DecodeString(base64Data)
	if err != nil {
		logger.ToTerminalRed("Cannot decode base64: ", err)
		return nil
	}
	return bytes.NewReader(decoded)
}
