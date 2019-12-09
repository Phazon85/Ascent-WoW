package multisql

import (
	"database/sql"
	"log"

	//used for testing in pg_test.go
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"github.com/phazon85/multisql/config"
)

//NewSQL returns a sql DB object
func NewSQL(file string) *sql.DB {
	driverName, cfg := config.NewConfig(file)
	db, err := sql.Open(driverName, cfg)
	if err != nil {
		log.Printf("Error making config string or opening SQL DB: %s", err.Error())
	}
	err = db.Ping()
	if err != nil {
		log.Printf("Error pinging sql db connection: %s", err.Error())
	}
	log.Println("Successfully connected.")
	return db
}
