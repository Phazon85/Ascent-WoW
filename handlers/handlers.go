package handlers

import (
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
)

//Config implements the config struct
type Config interface {
	MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate)
	// MessageReactionAdd(s *discordgo.Session, mra *discordgo.MessageReactionAdd)
	StateReady(s *discordgo.Session, r *discordgo.Ready)
}

//Config implements the Config interface
type config struct {
	BotKeyword  string
	Log         *log.Logger
	InfoMessage *discordgo.MessageEmbed
	HelpMessage *discordgo.MessageEmbed
}

type mraEvent struct {
	Session      *discordgo.Session
	Guild        *discordgo.Guild
	GuildMember  *discordgo.Member
	GuildRoleMap map[string]*discordgo.Role
}

//NewHandlers takes in helper and info
func NewHandlers(
	hm *discordgo.MessageEmbed,
	im *discordgo.MessageEmbed,
	botkey string,
) Config {

	infoMessage := im
	helpMessage := hm
	return &config{
		BotKeyword:  botkey,
		InfoMessage: infoMessage,
		HelpMessage: helpMessage,
	}
}



//MessageCreate is the callback for the MessageCrate event from Discord
func (c *config) MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	//Checks if bot keyword is present in message
	if !strings.HasPrefix(m.Content, c.BotKeyword) {
		return
	}

	// [BOT_KEYWORD] [command] [options] :: "!asc" "log_level" "debug"
	contentTokens := strings.Split(strings.TrimSpace(m.Content))
	if len(contentTokens) < 2 {
		return
	}

	switch strings.ToLower(contentTokens[1]) {
	case "info":
		_, err := s.ChannelMessageSendEmbed(m.ChannelID, c.)
		if err != nil {
			log.Printf("Error sending info message: %s", err.Error())
			return
		}

	case "help":
		_, err := s.ChannelMessageSendEmbed(m.ChannelID, c.HelpMessage)
		if err != nil {
			log.Printf("Error sending help message: %s", err.Error())
			return
		}
	default:
		//Silently fail for any unregistered commands
	}
}

//StateReady is the callback for the Ready event from Discord
func (c *config) StateReady(s *discordgo.Session, r *discordgo.Ready) {
	idleSince := 0

	usd := discordgo.UpdateStatusData{
		IdleSince: &idleSince,
		Game: &discordgo.Game{
			Name: strings.TrimSpace(c.BotKeyword),
			Type: discordgo.GameTypeWatching,
		},
		AFK:    false,
		Status: "online",
	}

	err := s.UpdateStatusComplex(usd)
	if err != nil {
		log.Printf("Error setting status: %s", err.Error())
	}
}

//MessageReactionAdd is the callback for the MessageReactionAdd event for Discord
// func (c *config) MessageReactionAdd(s *discordgo.Session, mra *discordgo.MessageReactionAdd) {
// 	//Get the user
// 	user, err := s.User(mra.UserID)
// 	if err != nil {
// 		log.Printf("Error getting user in MessageReactionAdd: %s", err.Error())
// 		return
// 	}

// 	//Get the guild
// 	guild, err := s.Guild(mra.GuildID)
// 	if err != nil {
// 		log.Printf("Error getting guild in MessageReactionAdd: %s", err.Error())
// 		return
// 	}

// 	found := false
// 	var guildMember *discordgo.Member

// 	for _, member := range guild.Members {
// 		if member.User.ID == user.ID {
// 			found = true
// 			guildMember = member
// 			break
// 		}
// 	}

// 	if !found {
// 		log.Printf("User not found in guild in MessageReactionAdd: %s", err.Error())
// 		return
// 	}

// 	event := &mraEvent{}

// }
