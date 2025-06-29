package config

import (
	"flag"
	"fmt"
	"os"
)

type Config struct {
	dbUser     string
	dbPassword string
	dbHost     string
	dbPort     string
	dbName     string
	testDBHost string
	testDBName string
	apiPort    string
	migrate    string
}

func Get() *Config {
	config := &Config{}

	flag.StringVar(&config.dbUser, "dbUser", os.Getenv("POSTGRES_USER"), "DB user name")
	flag.StringVar(&config.dbPassword, "dbPassword", os.Getenv("POSTGRES_PASSWORD"), "DB password")
	flag.StringVar(&config.dbHost, "dbHost", os.Getenv("POSTGRES_HOST"), "DB host")
	flag.StringVar(&config.dbPort, "dbPort", os.Getenv("POSTGRES_PORT"), "DB port")
	flag.StringVar(&config.dbName, "dbName", os.Getenv("POSTGRES_DB"), "DB name")
	flag.StringVar(&config.testDBHost, "testDBHost", os.Getenv("TEST_DB_HOST"), "test DB host")
	flag.StringVar(&config.testDBName, "testDBName", os.Getenv("TEST_DB_NAME"), "test DB name")
	flag.StringVar(&config.apiPort, "apiPort", os.Getenv("API_PORT"), "API Port")
	flag.StringVar(&config.migrate, "migrate", "up", "specify if we should be migrating DB 'up' or 'down'")
	flag.Parse()
	return config
}

func (c *Config) GetDBConnStr() string {
	return c.getDBConnStr(c.dbHost, c.dbName)
}

func (c *Config) GetTestDBConnStr() string {
	return c.getDBConnStr(c.testDBHost, c.testDBName)
}

func (c *Config) GetAPIPort() string {
	return ":" + c.apiPort
}

func (c *Config) GetMigration() string {
	return c.migrate
}

func (c *Config) getDBConnStr(dbHost, dbName string) string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		c.dbUser,
		c.dbPassword,
		dbHost,
		c.dbPort,
		dbName,
	)
}
