package repository

import "BalancingServers/repository/postgres"

type Repository struct {
	Postgres *postgres.PG
}
