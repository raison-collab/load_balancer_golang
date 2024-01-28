package repository

import (
	"BalancingServers/config"
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
)

var conf config.Config
var databaseConfig config.DatabaseConfig = conf.Database

func readConf() {
	conf.ReadTomlConfig("config.toml")
}

func NewPostgresDB(cfg config.DatabaseConfig) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DatabaseName, cfg.Password, cfg.SSlMode,
	))
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
