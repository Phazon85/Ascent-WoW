package handlers

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

type mraEvent struct {
	Session      *discordgo.Session
	Guild        *discordgo.Guild
	GuildMember  *discordgo.Member
	GuildRoleMap map[string]*discordgo.Role
}

//MessageReactionAdd is the callback for the MessageReactionAdd event for Discord
func (c *Config) MessageReactionAdd(s *discordgo.Session, mra *discordgo.MessageReactionAdd) {
	//Get the user
	user, err := s.User(mra.UserID)
	if err != nil {
		log.Printf("Error getting user in MessageReactionAdd: %s", err.Error())
		return
	}

	//Get the guild
	guild, err := s.Guild(mra.GuildID)
	if err != nil {
		log.Printf("Error getting guild in MessageReactionAdd: %s", err.Error())
		return
	}

	// temp variables for error checking
	found := false
	var guildMember *discordgo.Member

	for _, member := range guild.Members {
		if member.User.ID == user.ID {
			found = true
			guildMember = member
			break
		}
	}

	if !found {
		log.Printf("User not found in guild in MessageReactionAdd: %s", err.Error())
		return
	}

	//Store event info into struct
	event := &mraEvent{
		Session:      s,
		Guild:        guild,
		GuildMember:  guildMember,
		GuildRoleMap: make(map[string]*discordgo.Role),
	}

	for _, role := range event.Guild.Roles {
		event.GuildRoleMap[role.ID] = role
	}

	log.Printf("Channel ID: " + mra.ChannelID)
	log.Printf("Emoji ID: " + mra.Emoji.ID)
	log.Printf("Emoji Name: " + mra.Emoji.Name)
	log.Printf("User ID: " + mra.UserID)

	//Checking to see if user already has the Guild Member role
	for _, memberRoleID := range event.GuildMember.Roles {
		if event.GuildRoleMap[memberRoleID].ID == "616443279051063313" {
			return
		}
	}

	for _, guildRole := range event.GuildRoleMap {
		if guildRole.ID == "616443279051063313" {
			grantGuildMemberRole(event, guildRole)
		}
	}

}

func grantGuildMemberRole(event *mraEvent, guildRole *discordgo.Role) {
	err := event.Session.GuildMemberRoleAdd(event.Guild.ID, event.GuildMember.User.ID, guildRole.ID)
	if err != nil {
		log.Printf("Failed to add role in granteGuildMemberRole: %s", err.Error())
		return
	}
}
