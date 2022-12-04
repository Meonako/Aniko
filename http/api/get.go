package api

import (
	"net/http"

	"github.com/Meonako/Aniko/config"
)

func Progress() []byte {
	API_URL := config.Conf.BASE_URL + config.Conf.API_PROGRESS_PATH
	resp, err := http.Get(API_URL)
	if err != nil {
		return []byte(`{ "ERROR": "` + err.Error() + `" }`)
	}
	defer resp.Body.Close()

	return readByte(resp)
}

func Styles() []byte {
	API_URL := config.Conf.BASE_URL + config.Conf.API_STYLES_PATH
	resp, err := http.Get(API_URL)
	if err != nil {
		return []byte(`
		[
			{
				"ERROR": "` + err.Error() + `"
			}
		]`)
	}
	defer resp.Body.Close()

	return readByte(resp)
}
