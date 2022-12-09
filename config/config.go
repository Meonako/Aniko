package config

import (
	"os"

	"github.com/Meonako/go-logger/v2"

	"github.com/bwmarrin/discordgo"
	"github.com/goccy/go-json"
)

type settings struct {
	// These group value will be load once bot is ready and ready event is trigger

	OWNER        *discordgo.User               // Will based on OWNER_ID field
	EMBED_AUTHOR *discordgo.MessageEmbedAuthor // Will based on OWNER_ID field
	EMBED_FOOTER *discordgo.MessageEmbedFooter // Will based on OWNER_ID field

	OWNER_ID          string
	REGISTER_GUILD_ID string // Guild ID to register. Empty register globally

	EMBED_MAIN_COLOR  int // Main Embed Color.
	EMBED_ERROR_COLOR int // Embed Color when sending Error message.

	BASE_URL          string
	API_TXT2IMG_PATH  string
	API_PROGRESS_PATH string
	API_STYLES_PATH   string

	OPEN_IN_WINDOWS_TERMINAL bool
	IF_NOT_FOUND_BINARY      int8
}

var (
	configFileName = "config.json"

	Conf = settings{
		OWNER_ID:          "572446728759017472",
		REGISTER_GUILD_ID: "966011438466474004",

		EMBED_MAIN_COLOR:  0x0091f9,
		EMBED_ERROR_COLOR: 0xFF0000,

		BASE_URL:          "http://127.0.0.1:7860",
		API_TXT2IMG_PATH:  "/sdapi/v1/txt2img",
		API_PROGRESS_PATH: "/sdapi/v1/progress",
		API_STYLES_PATH:   "/sdapi/v1/prompt-styles",

		OPEN_IN_WINDOWS_TERMINAL: false,
		IF_NOT_FOUND_BINARY:      0,
	}
)

func init() {
	logger.ToTerminal("Initializing config...")
	bytes, err := os.ReadFile(configFileName)
	if err != nil {
		logger.ToTerminalRed(err)
		return
	}

	err = json.Unmarshal(bytes, &Conf)
	if err != nil {
		logger.ToTerminalRed(err)
		return
	}
	logger.ToTerminal("Initialized config")
}

func Config() *settings {
	return &Conf
}

func Save() {
	file, err := os.OpenFile(configFileName, os.O_CREATE|os.O_RDWR, 0700)
	if err != nil {
		logger.ToTerminalRed(err)
		return
	}

	bytes, err := json.MarshalIndent(Conf, "", "    ")
	if err != nil {
		logger.ToTerminalRed(err)
		return
	}

	_, err = file.Write(bytes)
	if err != nil {
		logger.ToTerminalRed(err)
		return
	}
	logger.ToTerminal(logger.Yellow("Successfully saved config file"))
}
