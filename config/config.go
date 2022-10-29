package config

import (
	"fmt"
	"github.com/kelseyhightower/envconfig"
	"github.com/phuslu/log"
	"github.com/rostikts/fintech_test_project/infrastructure/datastore"
	"path/filepath"
	"runtime"
)

type config struct {
	DBConfig datastore.Config
	FilePath string
}

var Config *config

func init() {
	var cfg config
	err := envconfig.Process("", &cfg)
	if err != nil {
		log.DefaultLogger.Fatal().Err(err).Msg("Error occurred during config init")
	}
	cfg.FilePath = getBaseFilePath()
	Config = &cfg
}

func getBaseFilePath() string {
	_, b, _, _ := runtime.Caller(0)

	root := filepath.Join(filepath.Dir(b), "../")
	return fmt.Sprintf("%v/tmp", root)
}
