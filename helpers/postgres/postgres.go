package postgres

import (
	"database/sql"

	"github.com/phazon85/multisql"
)

//PSQLService implements the Actions interface
type PSQLService struct {
	DB *sql.DB
}

//NewDBObject returns a PSQLService struct for package to use
func NewDBObject(filename, drivername string) *PSQLService {
	sql := multisql.NewSQLDBObject(filename, drivername)

	return &PSQLService{
		DB: sql,
	}

}

//LoadAvailableRaidGroups gets all raid groups available to start with
func (p *PSQLService) LoadAvailableRaidGroups() ([]*RaidGroup, error) {
	groups := []*RaidGroup{}
	rows, err := p.DB.Query(getAvailableRaidGroups)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		newGroup := &RaidGroup{}
		err = rows.Scan(&newGroup.ID, &newGroup.Name)
		if err != nil {
			return nil, err
		}
		groups = append(groups, newGroup)
	}

	return groups, nil
}

//CreateRaidGroup inserts a new Raidgroup into the database
func (p *PSQLService) CreateRaidGroup(guildid, userid string) error {
	_, err := p.DB.Exec(insertNewRaidGroup, guildid, userid)
	return err
}

// //CreateRaid inserts a new raid and marks it active
// func (p *PSQLService) CreateRaid(raid *Raid) error {
// 	result := p.checkActiveRaid(raid)
// 	if result != nil {
// 		return result
// 	}

// 	_, err := p.DB.Exec(insertNewRaid, raid.ID, uuid.New(), time.Now(), true)
// 	return err
// }

//CheckActiveRaid takes in a RaidGroup struct and checks
func (p *PSQLService) CheckActiveRaid(raidgroupid string) ([]Raid, error) {
	results := []Raid{}
	//query DB for all raids attached to raid group and append to results empty slice
	rows, err := p.DB.Query(getRaids, raidgroupid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		raid := Raid{}
		err = rows.Scan(&raid.ID, &raid.RaidID, &raid.Members, &raid.StartTime, &raid.EndTime, &raid.ItemsAwarded, &raid.Active)
		if err != nil {
			return nil, err
		}
		results = append(results, raid)
	}
	return results, nil
	//parse results for any raids marked as active
	//If no active raid exists, append new raid to results and UPDATE DB record with new json

	//return results or error
}

// //JoinRaid checks to see if the discord user is currently a raider in the active raid's RaidGroup.
// //If they are not in the current RaidGroup, then add them. Otherwise,
// func (p *PSQLService) JoinRaid(raidGroup *RaidGroup, Raid *Raid) error {

// 	// Get raid
// 	raid := p.getRaid(guildid)
// 	//Check to see if raider has already joined raid group from active server raid

// 	//If not add raider to raid group and add to active status

// 	//If apart of raid group, set active

// 	//return error if it didn't work

// }

// func (p *PSQLService) getRaid(raidid string) *Raid {
// 	result := &Raid{}
// 	row := p.DB.QueryRow(getRaid, raidid)
// }
