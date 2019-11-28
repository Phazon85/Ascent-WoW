package discord

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

const (
	hordeEmoji      = "633097926402637824"
	guildMemberRole = "616443279051063313"
	priestEmoji     = "633096835241869343"
	priestRole      = "649346745607913524"
	warriorEmoji    = "633096699266596900"
	warriorRole     = "649346732697845811"
	shamanEmoji     = "633096699312734238"
	shamanRole      = "649346747696676886"
	paladinEmoji    = "633096699115601923"
	paladinRole     = "" // I do not currently have a paladin role
	mageEmoji       = "633096699308539914"
	mageRole        = "649347010180153344"
	warlockEmoji    = "633096699031846925"
	warlockRole     = "649347095685365800"
	rogueEmoji      = "633096666802552842"
	rogueRole       = "649346750473175049"
	hunterEmoji     = "633096699073658892"
	hunterRole      = "649346973585113090"
	druidEmoji      = "633096698968801291"
	druidRole       = "649346874725105665"
)

func messageReactionAdd() func(*discordgo.Session, *discordgo.MessageReactionAdd) {
	return func(s *discordgo.Session, mra *discordgo.MessageReactionAdd) {
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

		switch mra.Emoji.ID {
		case priestEmoji:
			checkGuildRole(event, priestRole)
		case warriorEmoji:
			checkGuildRole(event, warriorRole)
		case shamanEmoji:
			checkGuildRole(event, shamanRole)
		case paladinEmoji:
			checkGuildRole(event, paladinRole)
		case mageEmoji:
			checkGuildRole(event, mageRole)
		case warlockEmoji:
			checkGuildRole(event, warlockRole)
		case rogueEmoji:
			checkGuildRole(event, rogueRole)
		case hunterEmoji:
			checkGuildRole(event, hunterRole)
		case druidEmoji:
			checkGuildRole(event, druidRole)
		default:
			//Silently fail
		}

	}
}

func checkGuildRole(event *mraEvent, guildRoleID string) {
	for _, guildRole := range event.GuildRoleMap {
		if guildRole.ID == guildRoleID {
			grantGuildMemberRole(event, guildRole)
			return
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
