package api

import (
	"bytes"
	"net/http"

	"github.com/Meonako/Aniko/config"
)

func TXT2Img(data []byte) string {
	API_URL := config.Conf.BASE_URL + config.Conf.API_TXT2IMG_PATH
	resp, err := http.Post(API_URL, "application/json", bytes.NewBuffer(data))
	if err != nil {
		return "ERROR: " + err.Error()
	}
	defer resp.Body.Close()

	return readString(resp)
}
