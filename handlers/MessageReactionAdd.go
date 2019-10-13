package handlers

// //MessageReactionAdd is the callback for the MessageReactionAdd event for Discord
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

// 	// temp variables for error checking
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

// 	//Store event info into struct
// 	event := &mraEvent{
// 		Session: s,
// 		Guild: guild,
// 		GuildMember: guildMember,
// 		GuildRoleMap: make(map[string]*discordgo.Role),
// 	}

// 	for _, role := range event.Guild.Role {
// 		event.GuildRoleMap[role.ID] = role
// 	}

// }
