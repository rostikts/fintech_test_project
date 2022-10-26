package datastore

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/phuslu/log"
)

type Config struct {
	User     string `envconfig:"POSTGRES_USER"`
	Password string `envconfig:"POSTGRES_PASSWORD"`
	Name     string `envconfig:"POSTGRES_DB"`
	Host     string `envconfig:"POSTGRES_HOST"`
	Port     int    `envconfig:"POSTGRES_PORT"`
}

func NewDB(config Config) *sqlx.DB {
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
