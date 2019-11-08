package postgres

import "errors"

var (
	errNoGuildID  = errors.New("guildid cannot be nil")
	errActiveRaid = errors.New("Active raid currently exists for guild")
)

const (
	getAvailableRaidGroups = "SELECT id, author FROM raid_groups;"
	insertNewRaidGroup     = "INSERT INTO raid_groups (id, author) VALUES ($1, $2);"
	insertNewRaid          = "INSERT INTO raids(id, raidid, start_time, active) VALUES($1, $2, $3, $4);"
	checkActiveRaid        = "SELECT * FROM raids WHERE id = '$1' and active = true'"
	getRaids               = "SELECT * FROM raids WHERE id = '%s'"
	updateRaid             = `
		UPDATE raids
		SET members = '$1'
		WHERE id = '$2'
	`
)

//Actions provide ways to interact with the postgresDB
type Actions interface {
	LoadAvailableRaidGroups() ([]*RaidGroup, error)
	CreateRaidGroup(guildid, userid string) error
	// CreateRaid(guildid string) error
	// JoinRaid(guildid, memberid string) error
	CheckActiveRaid(raidgroupid string) ([]Raid, error)
}

//RaidGroup hold specific raid group's DKP
type RaidGroup struct {
	ID    string
	Name  string
	DKP   map[string]int
	Raids []Raid
}

//Raid hold individual raid data
type Raid struct {
	ID           string
	RaidID       string
	Members      map[string]bool
	StartTime    string
	EndTime      string
	ItemsAwarded map[int]string
	Active       bool
}
