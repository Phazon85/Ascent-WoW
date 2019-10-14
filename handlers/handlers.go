package handlers

import (
	"bytes"
	"strconv"

	"github.com/bwmarrin/discordgo"
	"go.uber.org/zap"
)

//Actions implements the different Discord callbacks
type Actions interface {
	MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate)
	MessageReactionAdd(s *discordgo.Session, mra *discordgo.MessageReactionAdd)
	StateReady(s *discordgo.Session, r *discordgo.Ready)
}

//Config implements the Config interface
type Config struct {
	Log        *zap.Logger
	BotKeyword string
}

//DiscordAPIResponse is a receiving struct for API JSON responses
type DiscordAPIResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type discordError struct {
	HTTPResponseMessage string
	APIResponse         *DiscordAPIResponse
	CustomMessage       string
}

func (dErr *discordError) Error() string {
	return "error from Discord API: " + dErr.String()
}

func (dErr *discordError) String() string {
	buf := bytes.NewBuffer([]byte{})

	if dErr.CustomMessage != "" {
		buf.Write([]byte("CustomMessage: " + dErr.CustomMessage + ", "))
	}

	buf.Write([]byte("HTTPResponseMessage: " + dErr.HTTPResponseMessage + ", "))
	buf.Write([]byte("APIResponse.Code: " + strconv.Itoa(dErr.APIResponse.Code) + ", "))
	buf.Write([]byte("APIResponse.Message: " + dErr.APIResponse.Message))

	return buf.String()
}

//NewHandlers takes in bot info and returns a Config
func NewHandlers(
	botkey string,
	log *zap.Logger,
) *Config {
	return &Config{
		Log:        log,
		BotKeyword: botkey,
	}
}
