package pgsql

import (
	"database/sql"
	"errors"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/google/uuid"
)

var (
	errNoGuildID         = errors.New("guildid cannot be nil")
	errActiveRaid        = errors.New("Active raid currently exists for guild")
	errNoRaidGroup       = errors.New("Could not find an active raidgroup")
	errRaidGroupExists   = errors.New("Raid Group already exists")
	errNoRaidGroupExists = errors.New("Raid Group does not exists")
)

const (
	checkActiveRaidGroup = "SELECT id FROM raid_groups WHERE id=$1;"
	createRaidGroup      = "INSERT INTO raid_groups (id, author) VALUES ($1, $2);"
	createRaid           = `INSET INTO raids (id, raidid, start_time, active) 
					VALUES ($1, $2, $3, $4);`
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
	RaidID       uuid.UUID
	Members      map[string]bool
	StartTime    time.Time
	EndTime      time.Time
	ItemsAwarded map[string]int
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
	err := c.checkIfRaidGroupExists(mc.ChannelID)
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

func (c *Client) checkIfRaidGroupExists(id string) error {
	result := &RaidGroup{}
	row := c.DB.QueryRow(checkActiveRaidGroup, id)
	if _ = row.Scan(&result.ID); result.ID == id {
		return errRaidGroupExists
	}
	return nil
}

//StartRaid ...
func (c *Client) StartRaid(mc *discordgo.MessageCreate) error {
	err := c.checkIfRaidGroupExists(mc.ChannelID)
	if err == nil {
		return errNoRaidGroupExists
	}
	raid := newRaid(mc)

	_, err = c.DB.Exec(createRaid, raid.ID, raid.RaidID, raid.StartTime, raid.Active)
	if err != nil {
		return err
	}
	return nil
}

func newRaid(mc *discordgo.MessageCreate) *Raid {
	return &Raid{
		ID:        mc.ChannelID,
		RaidID:    uuid.New(),
		StartTime: time.Now(),
		Active:    true,
	}
}
