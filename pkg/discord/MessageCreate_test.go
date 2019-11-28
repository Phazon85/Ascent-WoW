package discord

// import (
// 	"fmt"
// 	"testing"

// 	"go.uber.org/zap/zapcore"

// 	"github.com/bwmarrin/discordgo"
// 	"github.com/phazon85/Ascent-WoW/helpers/logging"

// 	"github.com/phazon85/Ascent-WoW/helpers/mock"
// )

// func TestConfig_MessageCreate(t *testing.T) {
// 	session, err := mock.Session()
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	log := logging.NewLogger()
// 	log.Level.SetLevel(zapcore.FatalLevel)

// 	config := &Config{
// 		Logger:     log,
// 		BotKeyword: "testKeyword",
// 	}

// 	// message from a bot
// 	sendBotMessage(session, config)

// 	tests := []string{
// 		"ixnay", // no keyword
// 		fmt.Sprintf("%s %s", config.BotKeyword, "ixnay"), // keyword, unrecognized command
// 		fmt.Sprintf("%s %s", config.BotKeyword, "info"),  // keyword, incomplete command
// 		fmt.Sprintf("%s %s", config.BotKeyword, "log_level debug"),
// 		fmt.Sprintf("%s %s", config.BotKeyword, "log_level info"),
// 		fmt.Sprintf("%s %s", config.BotKeyword, "log_level warn"),
// 		fmt.Sprintf("%s %s", config.BotKeyword, "log_level error"),
// 		fmt.Sprintf("%s %s", config.BotKeyword, "log_level fatal"),
// 		fmt.Sprintf("%s %s", config.BotKeyword, "log_level panic"),
// 	}

// 	for _, test := range tests {
// 		sendMessage(session, config, test)
// 	}
// }

// func sendBotMessage(s *discordgo.Session, config *Config) {
// 	config.MessageCreate(
// 		s,
// 		&discordgo.MessageCreate{
// 			Message: &discordgo.Message{
// 				Author: &discordgo.User{
// 					Username: "testUsername",
// 					Bot:      true,
// 				},
// 				GuildID:   "testGuild",
// 				ChannelID: "testChannel",
// 				Content:   "",
// 			},
// 		},
// 	)
// }

// func sendMessage(s *discordgo.Session, config *Config, message string) {
// 	config.MessageCreate(
// 		s,
// 		&discordgo.MessageCreate{
// 			Message: &discordgo.Message{
// 				Author: &discordgo.User{
// 					Username: "testUsername",
// 					Bot:      false,
// 				},
// 				GuildID:   "testGuild",
// 				ChannelID: "testChannel",
// 				Content:   message,
// 			},
// 		},
// 	)
// }
