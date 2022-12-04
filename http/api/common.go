package api

import (
	"io"
	"net/http"

	"github.com/Meonako/go-logger/v2"
)

func readString(resp *http.Response) string {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.ToTerminalRed("Cannot read response body: ", err)
		return "ERROR: " + err.Error()
	}

	if resp.StatusCode != 200 {
		return "ERROR: " + string(body)
	}
	resp.Body.Close()

	return string(body)
}

func readByte(resp *http.Response) []byte {
	body, err := io.ReadAll(resp.Body)
	logger.ToTerminalRedFormatIfError(err, "Cannot read response body: %v")
	if resp.StatusCode != 200 {
		return []byte(`{ "ERROR": "` + string(body) + `" }`)
	}
	resp.Body.Close()

	return body
}
