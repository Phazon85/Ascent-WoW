package postgres

import (
	"database/sql"

	"github.com/phazon85/multisql"
)

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
