package postgres

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
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

//CreateRaid inserts a new raid and marks it active
func (p *PSQLService) CreateRaid(raid *Raid) error {
	result := p.checkActiveRaid(raid)
	if result != nil {
		return result
	}

	_, err := p.DB.Exec(insertNewRaid, raid.ID, uuid.New(), time.Now(), true)
	return err
}

func (p *PSQLService) checkActiveRaid(raid *Raid) error {
	result := &Raid{}
	row := p.DB.QueryRow(getRaid, raid.ID)
	_ = row.Scan(&result.ID, &result.RaidID, &result.Members, &result.StartTime, &result.EndTime, &result.ItemsAwarded, &result.Active)
	if result.ID == raid.ID {
		return errActiveRaid
	}
	return nil
}

// func (p *PSQLService) JoinRaid() error {
// 	// Get raid
// 	raid := getRaid(guildid)
// 	//Check to see if member has already joined from raid we just got

// 	//If not add member to raid

// 	//return error if it didn't work

// }

// func (p *PSQLService) getRaid(raidid string) *Raid {
// 	result := &Raid{}
// 	row := p.DB.QueryRow()
// }
