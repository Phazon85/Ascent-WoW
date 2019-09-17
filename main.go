package main

import (
	"errors"
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
)

//Config holds system environment variables
type Config struct {
	BotKeyword  string
	InfoMessage *discordgo.MessageEmbed
	Logger      *log.Logger
}

// CheckEnvironmentConfig loads BOT_TOKEN and BOT_KEYWORD secrets
func CheckEnvironmentConfig() (*Config, string, error) {

	token := os.Getenv("BOT_TOKEN")
	if token == "" {
		err := errors.New("BOT_TOKEN not defined in environment variables")
		return nil, "", err
	}

	keyword := os.Getenv("BOT_KEYWORD")
	if keyword == "" {
		err := errors.New("BOT_KEYWORD not defined in environment variables")
		return nil, "", err
	}

	return &Config{
		BotKeyword: keyword,
	}, token, nil
}

func setupHandlers(dg *discordgo.Session, config *Config) {

	dg.AddHandler(MessageCreate)
	dg.AddHandler(Ready)
	dg.AddHandler(MessageReactionAdd)

}

func main() {
	log.Println("Starting bot")

	//load environment variables
	config, token, err := CheckEnvironmentConfig()
	if err != nil {
		log.Printf("Error loading environment variables: %s", err.Error())
	}

	//Create new discordgo session
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Printf("Error creating new discord session: %s", err.Error())
	}

	//Intialize handlers for discordgo session
	setupHandlers(dg, config)

	//Starts discord event listener
	err = dg.Open()
	if err != nil {
		log.Printf("Error opening discordgo session: %s", err.Error())
	}
	defer dg.Close()

	//Keeps listeners alive
	<-make(chan struct{})

}

// type DiscordAPIReponse struct {
// 	Code    int    `json:"code"`
// 	Message string `json:"message"`
// }

// type discordError struct {
// 	HTTPResponseMessage string
// 	APIResponse         *DiscordAPIReponse
// 	CustomMessage       string
// }

// // (dErr *discordError) Error satisfies the error interface
// func (dErr *discordError) Error() string {
// 	return "error from Discord API: " + dErr.String()
// }

// // (dErr *discordError) String satisfies the fmt.Stringer interface
// func (dErr *discordError) String() string {
// 	buf := bytes.NewBuffer([]byte{})

// 	if dErr.CustomMessage != "" {
// 		buf.Write([]byte("CustomMessage: " + dErr.CustomMessage + ", "))
// 	}

// 	buf.Write([]byte("HTTPResponseMessage: " + dErr.HTTPResponseMessage + ", "))
// 	buf.Write([]byte("APIResponse.Code: " + strconv.Itoa(dErr.APIResponse.Code) + ", "))
// 	buf.Write([]byte("APIResponse.Message: " + dErr.APIResponse.Message))

// 	return buf.String()
// }
