package discord

import (
	"strings"

	"github.com/bwmarrin/discordgo"
	"go.uber.org/zap"
)

func stateReady(prefix string, logger *zap.Logger) func(*discordgo.Session, *discordgo.Ready) {
	return func(s *discordgo.Session, r *discordgo.Ready) {
		idleSince := 0

		usd := discordgo.UpdateStatusData{
			IdleSince: &idleSince,
			Game: &discordgo.Game{
				Name: strings.TrimSpace(prefix),
				Type: discordgo.GameTypeWatching,
			},
			AFK:    false,
			Status: "online",
		}

		err := s.UpdateStatusComplex(usd)
		if err != nil {
			logger.Debug("Setting Status to online", zap.String("status: ", err.Error()))
		}
	}
}
