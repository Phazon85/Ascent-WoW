package postgres

import (
	"github.com/phazon85/Ascent-WoW/helpers/postgres"
	"log"
	"testing"
)

func TestCheckActiveRaid(t *testing.T) {
	raidGroupID := "125446049853603841"
	raid, err := postgres.Actions.
	if err != nil {
		log.Printf("Error getting active raid")
	}
	if raid == nil {
		log.Printf("Error empty raid slice")
	}
}
