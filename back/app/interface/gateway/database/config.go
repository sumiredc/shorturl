package database

import "os"

type Config struct {
	User     string
	Password string
	Host     string
	Port     string
	Database string
}

func newConfig() *Config {
	return &Config{
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_PORT"),
		os.Getenv("MYSQL_DATABASE"),
	}
}
