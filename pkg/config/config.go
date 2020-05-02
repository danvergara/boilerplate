package config

import (
	"flag"
	"fmt"
	"os"
)

// Config object that stores the DB config
type Config struct {
	dbUser     string
	dbPswd     string
	dbHost     string
	dbPort     string
	dbName     string
	testDBHost string
	testDBName string
}

// Get retuns an Config object with all the DB data
func Get() *Config {
	conf := &Config{}

	flag.StringVar(&conf.dbUser, "dbuser", os.Getenv("POSTGRES_USER"), "DB user name")
	flag.StringVar(&conf.dbPswd, "dbpswd", os.Getenv("POSTGRES_PASSWORD"), "DB pass")
	flag.StringVar(&conf.dbPort, "dbposrt", os.Getenv("POSTGRES_PORT"), "DB port")
	flag.StringVar(&conf.dbHost, "dbhost", os.Getenv("POSTGRES_HOST"), "DB host")
	flag.StringVar(&conf.dbName, "dbname", os.Getenv("POSTGRES_DB"), "DB name")
	flag.StringVar(&conf.testDBHost, "testdbhost", os.Getenv("TEST_DB_HOST"), "test database host")
	flag.StringVar(&conf.testDBName, "testdbname", os.Getenv("TEST_DB_NAME"), "test database name")

	flag.Parse()

	return conf
}

// GetDBConnStr build the str connection to the DB
func (conf *Config) GetDBConnStr() string {
	return conf.getDBConnStr(conf.dbHost, conf.dbName)
}

// GetTestDBConnStr build the str connection to the test DB
func (conf *Config) GetTestDBConnStr() string {
	return conf.getDBConnStr(conf.testDBHost, conf.testDBName)
}

func (conf *Config) getDBConnStr(dbhost, dbname string) string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		conf.dbUser,
		conf.dbPswd,
		dbhost,
		conf.dbPort,
		dbname,
	)
}
