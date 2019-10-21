package handlers

import (
	"bytes"
	"strconv"

	"github.com/phazon85/Ascent-WoW/helpers/logging"

	"github.com/bwmarrin/discordgo"
)

//Actions implements the different Discord callbacks
type Actions interface {
	MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate)
	MessageReactionAdd(s *discordgo.Session, mra *discordgo.MessageReactionAdd)
	StateReady(s *discordgo.Session, r *discordgo.Ready)
}

//Config implements the Config interface
type Config struct {
	Logger     *logging.ZapLogger
	BotKeyword string
	Mongo      *mongo.Client
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
	log *logging.ZapLogger,
	mongo *mongo.MongoService,
) *Config {
	return &Config{
		Logger:     log,
		BotKeyword: botkey,
		Mongo:      mongo,
	}
}
