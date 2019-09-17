package services

import (
	"encoding/json"
	"log"
	"regexp"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var (
	infoMessage = &discordgo.MessageEmbed{
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

	helpMessage = &discordgo.MessageEmbed{
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
)

//CallBacks contains methods for interacting with the discordgo api
type CallBacks interface {
	MessageCreate(*discordgo.Session, *discordgo.MessageCreate)
	Ready(*discordgo.Session, *discordgo.Ready)
	MessageReactionAdd(s *discordgo.Session, mr *discordgo.MessageReactionAdd)
}

//Config implements the CallBacks interface and holds discord information
type Config struct {
	BotKeyword  string
	InfoMessage *discordgo.MessageEmbed
	Logger      *log.Logger
}

//MessageCreate is the callback for the MessageCrate event from Discord
func (e *Config) MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	//Checks if bot keyword is present in message
	if !strings.HasPrefix(m.Content, e.BotKeyword) {
		return
	}

	// [BOT_KEYWORD] [command] [options] :: "!asc" "log_level" "debug"
	contentTokens := strings.Split(strings.TrimSpace(m.Content), " ")
	if len(contentTokens) < 2 {
		return
	}

	switch strings.ToLower(contentTokens[1]) {
	case "info":
		_, err := s.ChannelMessageSendEmbed(m.ChannelID, infoMessage)
		if err != nil {
			log.Printf("Error sending info message: %s", err.Error())
			return
		}

	case "help":
		_, err := s.ChannelMessageSendEmbed(m.ChannelID, helpMessage)
		if err != nil {
			log.Printf("Error sending help message: %s", err.Error())
			return
		}
	default:
		//Silently fail for any unregistered commands
	}
}

//Ready is the callback for the Ready event from Discord
func (e *Config) Ready(s *discordgo.Session, r *discordgo.Ready) {
	idleSince := 0

	usd := discordgo.UpdateStatusData{
		IdleSince: &idleSince,
		Game: &discordgo.Game{
			Name: strings.TrimSpace(e.BotKeyword),
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

//MessageReactionAdd is the call back for the MessageReactionAdd event from Discord
func (e *Config) MessageReactionAdd(s *discordgo.Session, mr *discordgo.MessageReactionAdd) {

	//Get the User ID
	user, err := s.User(mr.UserID)
	if err != nil {
		log.Printf("Error getting user ID for MessageReactionAdd: %s", err.Error())
	}

	//Get the Guild ID
	guild, err := s.Guild(mr.GuildID)
	if err != nil {
		log.Printf("Error getting guild ID for MessageReactionAdd: %s", err.Error())
		return
	}

	//Get the Guild's roles
	guildRoles, dErr := guildRoles(s, mr.GuildID)
	if dErr != nil {
		log.Printf("Error getting guild roles for MessageReactionAdd: %s", err.Error())
		return
	}

	// Get the guild member's roles
	memberRoles, err := guildMemberRoles(s, user, guild)
	if err != nil {
		log.Printf("Error getting guild member's roles for MessageReactionAdd: %s", err.Error())
		return
	}

	//Check to see if the role already exists in the guild
	for _, role := range guildRoles {
		if role.Name != "Guild Member" {
			continue
		}

		//Check to see if the member already has the role
		for _, mRole := range memberRoles {
			if mRole.ID == role.ID {
				return // User already has role, no change
			}
		}

		//Add role to member
		grantRole(s, user, guild, role)
	}
}

func guildRoles(s *discordgo.Session, guildID string) (roles []*discordgo.Role, dErr *discordError) {
	roles, err := s.GuildRoles(guildID)
	if err != nil {
		//Find the JSON with Regular Expression
		rx := regexp.MustCompile("{.*}")
		errHTTPString := rx.ReplaceAllString(err.Error(), "")
		errJSONString := rx.FindString(err.Error())

		dAPIResp := &DiscordAPIReponse{}

		dErr = &discordError{
			HTTPResponseMessage: errHTTPString,
			APIResponse:         dAPIResp,
			CustomMessage:       "",
		}

		unmarshallErr := json.Unmarshal([]byte(errJSONString), dAPIResp)
		if unmarshallErr != nil {
			dAPIResp.Code = -1
			dAPIResp.Message = "Unable to unmarshall API JSON Response: " + errJSONString
			return
		}

		//Add Custom Messages as appropriate
		switch dErr.APIResponse.Code {
		case 50013: // Code 50013: "Missing Permissions"
			dErr.CustomMessage = "Insufficient role permission to query guild roles"
		}
	}
	return
}

func guildMemberRoles(s *discordgo.Session, user *discordgo.User, guild *discordgo.Guild) ([]*discordgo.Role, error) {
	//Get Guild member
	guildMember, err := s.GuildMember(guild.ID, user.ID)
	if err != nil {
		return make([]*discordgo.Role, 0),
			err
	}

	//Get guild roles
	guildRoles, dErr := guildRoles(s, guild.ID)
	if dErr != nil {
		return make([]*discordgo.Role, 0), dErr
	}

	//map our member roles
	memberRoleIDs := make(map[string]bool)
	for _, roleID := range guildMember.Roles {
		memberRoleIDs[roleID] = true
	}

	memberRoles := make([]*discordgo.Role, 0)

	for _, role := range guildRoles {
		if memberRoleIDs[role.ID] {
			memberRoles = append(memberRoles, role)
		}
	}

	return memberRoles, nil
}

func grantRole(s *discordgo.Session, user *discordgo.User, guild *discordgo.Guild, role *discordgo.Role) {
	err := s.GuildMemberRoleAdd(guild.ID, user.ID, role.ID)
	if err != nil {
		log.Printf("Error adding role to member: %s", err.Error())
		return
	}
}
