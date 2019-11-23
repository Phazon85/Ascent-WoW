package discord

// import (
// 	"log"
// 	"strings"

// 	"github.com/bwmarrin/discordgo"
// )

// //StateReady is the callback for the Ready event from Discord
// func (c *Config) StateReady(s *discordgo.Session, r *discordgo.Ready) {
// 	idleSince := 0

// 	usd := discordgo.UpdateStatusData{
// 		IdleSince: &idleSince,
// 		Game: &discordgo.Game{
// 			Name: strings.TrimSpace(c.BotKeyword),
// 			Type: discordgo.GameTypeWatching,
// 		},
// 		AFK:    false,
// 		Status: "online",
// 	}

// 	err := s.UpdateStatusComplex(usd)
// 	if err != nil {
// 		log.Printf("Error setting status: %s", err.Error())
// 	}
// }
