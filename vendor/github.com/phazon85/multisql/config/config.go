package config

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

const (
	psqlConnectionString  = "host=%s port=%d user=%s password=%s dbname=%s sslmode=disable"
	mysqlConnectionString = "%s:%s@%s/%s"
)

//SQLConnectionInfo holds connection info for SQL implementations
type SQLConnectionInfo struct {
	Host       string `yaml:"host"`
	Port       int64  `yaml:"port"`
	User       string `yaml:"user"`
	Password   string `yaml:"password"`
	DriverName string `yaml:"drivername"`
	DBName     string `yaml:"dbname"`
}

//NewConfig takes in a file path, opens it, and returns a connection string appropriate to the DB
func NewConfig(file string) (driverName string, connection string) {
	cfg := &SQLConnectionInfo{}
	var con string
	if err := load(cfg, file); err != nil {
		return "", ""
	}

	switch cfg.DriverName {
	case "mysql":
		con = fmt.Sprintf(mysqlConnectionString, cfg.User, cfg.Password, cfg.Host, cfg.DBName)
	case "postgres":
		con = fmt.Sprintf(psqlConnectionString, cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName)
	}

	return cfg.DriverName, con
}

func load(config interface{}, fname string) error {
	data, err := ioutil.ReadFile(fname)
	if err != nil {
		return err
	}

	if err := yaml.Unmarshal(data, config); err != nil {
		return err
	}
	return nil
}
