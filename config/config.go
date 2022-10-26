package config

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/phuslu/log"
	"github.com/rostikts/fintech_test_project/infrastructure/datastore"
)

type config struct {
	DBConfig datastore.Config
}

var Config *config

func init() {
	var cfg config
	err := envconfig.Process("", &cfg)
	if err != nil {
		log.DefaultLogger.Fatal().Err(err).Msg("Error occurred during config init")
	}
	Config = &cfg
}
