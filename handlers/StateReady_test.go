package handlers

import (
	"testing"

	"go.uber.org/zap/zapcore"

	"github.com/bwmarrin/discordgo"
	"github.com/phazon85/Ascent-WoW/helpers/logging"
	"github.com/phazon85/Ascent-WoW/helpers/mock"
)

func TestConfig_StateReady(t *testing.T) {
	session, err := mock.Session()
	if err != nil {
		t.Fatal(err)
	}

	log := logging.NewLogger()
	log.Level.SetLevel(zapcore.FatalLevel)

	config := &Config{
		Logger:     log,
		BotKeyword: "testKeyword",
	}

	config.StateReady(
		session,
		&discordgo.Ready{
			Guilds: make([]*discordgo.Guild, 0),
		},
	)
}
