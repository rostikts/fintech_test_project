package test_utils

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/phuslu/log"
	"github.com/rostikts/fintech_test_project/infrastructure/datastore"
)

func NewTestDB(config datastore.Config) *sqlx.DB {
	config.Host = "localhost"
	config.Port = 4011
	dns := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Host,
		config.Port,
		config.User,
		config.Password,
		config.Name,
	)
	db, err := sqlx.Connect("postgres", dns)
	if err != nil {
		log.DefaultLogger.Fatal().Err(err).Msg("Error occurred during db init")
	}
	return db
}
