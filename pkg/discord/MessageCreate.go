package discord

import (
	"strings"

	"go.uber.org/zap"

	"github.com/bwmarrin/discordgo"
)

const (
	errRaidGroupAlreadyExists = `Raid group already exists for this channel`
	errActiveRaid             = `Active raid currently exists for your channel. Stop current raid before attempting to start a new one.`
)

func messageCreate(dkp DKP, d *Discord) func(*discordgo.Session, *discordgo.MessageCreate) {
	return func(s *discordgo.Session, mc *discordgo.MessageCreate) {
		// Ignore all messages created by the bot itself
		if mc.Author.ID == s.State.User.ID {
			return
		}

		//Checks if bot keyword is present in message
		if !strings.HasPrefix(mc.Content, d.Prefix) {
			return
		}

		// [BOT_KEYWORD] [command] [options] :: "!asc" "log_level" "debug"
		contentTokens := strings.Split(strings.TrimSpace(mc.Content), " ")
		if len(contentTokens) < 2 {
			return
		}

		switch strings.ToLower(contentTokens[1]) {
		case "info":
			_, err := s.ChannelMessageSendEmbed(mc.ChannelID, infoMessage())
			if err != nil {
				d.Logger.Debug("Sending Info message", zap.String("mc: ", err.Error()))
				return
			}

		case "help":
			_, err := s.ChannelMessageSendEmbed(mc.ChannelID, helpMessage())
			if err != nil {
				d.Logger.Debug("Sending Help message", zap.String("mc: ", err.Error()))
				return
			}
		case "raid":
			if len(contentTokens) >= 3 {
				raidOpt := strings.ToLower(contentTokens[2])

				switch raidOpt {
				case "init":
					err := dkp.InitRaidGroup(mc)
					if err != nil {
						d.Logger.Debug("init raid request", zap.String("dkp: ", err.Error()))
					}
					_, err = s.ChannelMessageSend(mc.ChannelID, errRaidGroupAlreadyExists)
				case "start":
					err := dkp.StartRaid(mc)
					if err != nil {
						d.Logger.Debug("start raid request", zap.String("dkp: ", err.Error()))
						_, err = s.ChannelMessageSend(mc.ChannelID, errActiveRaid)
						return
					}
					_, err = s.ChannelMessageSend(mc.ChannelID, "Successfully started Raid")
				case "stop":
					err := dkp.StopRaid(mc)
					if err != nil {
						d.Logger.Debug("stop raid request", zap.String("dkp: ", err.Error()))
						_, err = s.ChannelMessageSend(mc.ChannelID, "Failed to stop active raid. Please try again")
						return
					}
					_, err = s.ChannelMessageSend(mc.ChannelID, "Successfully stopped Raid")
				case "join":
					err := dkp.
				}
			}
		default:
			//Silently fail for any unregistered commands
		}

	}
}

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
