package db

import (
	"database/sql"
	"fmt"

	"github.com/kelseyhightower/envconfig"
	_ "github.com/lib/pq"
)

var Db *sql.DB

var config = struct {
	PGDatabase string `default:"cliff"`
	PGHost     string `default:"127.0.0.1"`
	PGPort     int    `default:"5432"`
	PGUser     string `default:"cliff"`
	PGPassword string
	PGSSLMode  string `default:"disable"`
}{}

func init() {
	err := envconfig.Process("master", &config)
	if err != nil {
		panic(err)
	}

	db, err := sql.Open(
		"postgres",
		fmt.Sprintf(
			"user=%s password=%s dbname=%s host=%s sslmode=%s",
			config.PGUser, config.PGPassword, config.PGDatabase,
			config.PGHost, config.PGSSLMode,
		),
	)

	if err != nil {
		panic(err)
	}

}
