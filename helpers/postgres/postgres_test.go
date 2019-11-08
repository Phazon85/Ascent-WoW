package postgres_test

import (
	"log"
	"testing"
)

func TestLoadAvailableRaidGroups(t *testing.T) {
	rgs, err := LoadAvailableRaidGroups()
	if err != nil {
		log.Printf("error getting all raids")
	}

	if rgs != nil {
		log.Printf("error getting raid groups")
	}
}
