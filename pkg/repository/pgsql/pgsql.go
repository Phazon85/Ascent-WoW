package pgsql

import (
	"database/sql"
	"errors"

	"github.com/bwmarrin/discordgo"
)

var (
	errNoGuildID       = errors.New("guildid cannot be nil")
	errActiveRaid      = errors.New("Active raid currently exists for guild")
	errNoRaidGroup     = errors.New("Could not find an active raidgroup")
	errRaidGroupExists = errors.New("Raid Group already exists")
)

const (
	checkActiveRaidGroup = "SELECT id FROM raid_groups WHERE id=$1;"
	createRaidGroup      = "INSERT INTO raid_groups (id, author) VALUES ($1, $2);"
)

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

//Client ...
type Client struct {
	DB *sql.DB
}

//New returns a Client interface
func New(db *sql.DB) *Client {
	return &Client{
		DB: db,
	}

}

//InitRaidGroup ...
func (c *Client) InitRaidGroup(mc *discordgo.MessageCreate) error {
	err := c.checkIfAccountExists(mc.ChannelID)
	if err != nil {
		return err
	}

	//Creating new raidgroup with channelID as ID and message author as Name
	newGroup := newRaidGroup(mc.ChannelID, mc.Author.Username)
	_, err = c.DB.Exec(
		createRaidGroup,
		newGroup.ID,
		newGroup.Name,
	)
	return nil
}

func newRaidGroup(id, name string) *RaidGroup {
	return &RaidGroup{
		ID:   id,
		Name: name,
	}
}

func (c *Client) checkIfAccountExists(id string) error {
	result := &RaidGroup{}
	row := c.DB.QueryRow(checkActiveRaidGroup, id)
	if _ = row.Scan(&result.ID); result.ID == id {
		return errRaidGroupExists
	}
	return nil
}

// //GetRaidGroup ...
// func (c *Client) GetRaidGroup() {}

// //LoadAvailableRaidGroups gets all raid groups available to start with
// func (p *Client) LoadAvailableRaidGroups() ([]*RaidGroup, error) {
// 	groups := []*RaidGroup{}
// 	rows, err := p.DB.Query(getAvailableRaidGroups)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()
// 	for rows.Next() {
// 		newGroup := &RaidGroup{}
// 		err = rows.Scan(&newGroup.ID, &newGroup.Name)
// 		if err != nil {
// 			return nil, err
// 		}
// 		groups = append(groups, newGroup)
// 	}

// 	return groups, nil
// }

// //CreateRaidGroup inserts a new Raidgroup into the database
// func (p *Client) CreateRaidGroup(guildid, userid string) error {
// 	_, err := p.DB.Exec(insertNewRaidGroup, guildid, userid)
// 	return err
// }

// //CheckActiveRaid takes in a RaidGroup struct and checks
// func (p *Client) CheckActiveRaid(raidgroupid string) ([]Raid, error) {
// 	results := []Raid{}
// 	//query DB for all raids attached to raid group and append to results empty slice
// 	rows, err := p.DB.Query(getRaids, raidgroupid)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()
// 	for rows.Next() {
// 		raid := Raid{}
// 		err = rows.Scan(&raid.ID, &raid.RaidID, &raid.Members, &raid.StartTime, &raid.EndTime, &raid.ItemsAwarded, &raid.Active)
// 		if err != nil {
// 			return nil, err
// 		}
// 		results = append(results, raid)
// 	}
// 	return results, nil
// 	//parse results for any raids marked as active
// 	//If no active raid exists, append new raid to results and UPDATE DB record with new json

// 	//return results or error
// }
