package discord

import (
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
)

//MessageCreate is the callback for the MessageCrate event from Discord
func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	//Checks if bot keyword is present in message
	if !strings.HasPrefix(m.Content, "!asc") {
		return
	}

	// [BOT_KEYWORD] [command] [options] :: "!asc" "log_level" "debug"
	contentTokens := strings.Split(strings.TrimSpace(m.Content), " ")
	if len(contentTokens) < 2 {
		return
	}

	switch strings.ToLower(contentTokens[1]) {
	case "info":
		_, err := s.ChannelMessageSendEmbed(m.ChannelID, infoMessage())
		if err != nil {
			log.Printf("Error sending info message: %s", err.Error())
			return
		}

	case "help":
		_, err := s.ChannelMessageSendEmbed(m.ChannelID, helpMessage())
		if err != nil {
			log.Printf("Error sending help message: %s", err.Error())
			return
		}
		// 		// case "log_level":
		// 		// 	if len(contentTokens) >= 3 {
		// 		// 		levelOpt := strings.ToLower(contentTokens[2])

		// 		// 		switch levelOpt {
		// 		// 		case "debug":
		// 		// 			c.Logger.Level.SetLevel(zapcore.DebugLevel)
		// 		// 			c.Logger.Log.Debug("Log level changed to Debug")
		// 		// 		case "info":
		// 		// 			c.Logger.Level.SetLevel(zapcore.InfoLevel)
		// 		// 			c.Logger.Log.Info("Log level changed to Info")
		// 		// 		case "warn":
		// 		// 			c.Logger.Level.SetLevel(zapcore.WarnLevel)
		// 		// 			c.Logger.Log.Warn("Log level changed to Warn")
		// 		// 		case "error":
		// 		// 			c.Logger.Level.SetLevel(zapcore.ErrorLevel)
		// 		// 			c.Logger.Log.Error("Log level changed to Error")
		// 		// 		case "fatal":
		// 		// 			c.Logger.Level.SetLevel(zapcore.FatalLevel)
		// 		// 		case "panic":
		// 		// 			c.Logger.Level.SetLevel(zapcore.PanicLevel)
		// 		// 		}
		// 		// 	}

		// 		// case "raid":
		// 		// 	if len(contentTokens) >= 3 {
		// 		// 		raidOpt := strings.ToLower(contentTokens[2])

		// 		// 		switch raidOpt {
		// 		// 		case "init":
		// 		// 			c.Logger.Log.Debug("Starting Raid")
		// 		// 			groups, err := c.SQL.LoadAvailableRaidGroups()
		// 		// 			if err != nil {
		// 		// 				c.Logger.Log.Debug("Failed to retrieve all Raid Groups")
		// 		// 			}

		// 		// 			for _, raidgroup := range groups {
		// 		// 				if raidgroup.ID == m.GuildID {
		// 		// 					s.ChannelMessageSend(m.ChannelID, "Raid group already exists")
		// 		// 					return
		// 		// 				}
		// 		// 			}

		// 		// 			err = c.SQL.CreateRaidGroup(m.GuildID, m.Author.ID)
		// 		// 			if err != nil {
		// 		// 				c.Logger.Log.Info("Failed to create raid group", zap.String("Error: ", err.Error()))
		// 		// 				return
		// 		// 			}

		// 		// 			s.ChannelMessageSend(m.ChannelID, "New Raid group created")
		// 		// 			return
		// 		// 		case "start":
		// 		// 			raid := &postgres.Raid{
		// 		// 				ID: m.GuildID,
		// 		// 			}
		// 		// 			c.Logger.Log.Debug("Starting raid group")
		// 		// 			err := c.SQL.CreateRaid(raid)
		// 		// 			if err != nil {
		// 		// 				s.ChannelMessageSend(m.ChannelID, "Active Raid: There is currently an raid marked active for your guild.")
		// 		// 				return
		// 		// 			}
		// 		// 			s.ChannelMessageSend(m.ChannelID, "New Raid started")
		// 		// 			return
		// 		// 		}

	}
}

// 	default:
// 		//Silently fail for any unregistered commands
// 	}
// }

func infoMessage() *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		URL:    "https://github.com/Phazon85/Ascent-WoW",
		Title:  "Ascent-WoW Discord Bot",
		Color:  0x031cfc,
		Footer: &discordgo.MessageEmbedFooter{Text: "Made using the discordgo library"},
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "About",
				Value:  "Ascent-WoW is a bot... To be added",
				Inline: false,
			},
			{
				Name:   "Author",
				Value:  "Ascent-WoW bot was created by Phazon85/Tourian",
				Inline: false,
			},
			{
				Name:   "Library",
				Value:  "Ascent-WoW uses the discordgo library by bwmarrin",
				Inline: false,
			},
		},
	}
}

func helpMessage() *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		URL:   "https://github.com/Phazon85/Ascent-WoW",
		Title: "Help Section",
		Color: 0x031cfc,
		Fields: []*discordgo.MessageEmbedField{
			{
				Name: "Help Section",
				Value: `
			help - Displays this! :D
			info - Displays information about the bot
			`,
				Inline: false,
			},
		},
	}
}
