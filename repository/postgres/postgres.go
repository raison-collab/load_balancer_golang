package postgres

import (
	"BalancingServers/config"
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
)

type PG struct {
	DB *sqlx.DB
}

func (p *PG) NewPostgresDB(cfg config.DatabaseConfig) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DatabaseName, cfg.Password, cfg.SSlMode,
	))
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	p.DB = db
}

func (p *PG) ClosePostgresDB() {
	if err := p.DB.Close(); err != nil {
		log.Fatal(err)
	}
}

//func (p *PG) InsertTask(pt *endpoints.PostTask) uint {
//	res, err := p.DB.Exec("INSERT INTO task(bash, ram, disk, cpu, priority) VALUES ($1, $2, $3, $4, $5)",
//		pt.Bash, pt.Ram, pt.Disk, pt.CPU, pt.Priority)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	lastId, err := res.LastInsertId()
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	return uint(lastId)
//}
