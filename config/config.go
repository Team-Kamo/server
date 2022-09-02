package config

import (
	"errors"
	"os"

	"github.com/rs/zerolog/log"

	"github.com/goccy/go-yaml"
)

type Config struct {
	Port      string
	Debug     bool
	Lang      string
	LogFormat string
	Root      string
	Db        DbConfig
	Storage   StorageConfig
	Token     []string
}

type DbConfig struct {
	Type string
	Uri  string
	Name string
}

type StorageConfig struct {
	Type      string
	Uri       string
	Name      string
	Endpoint  string
	AccessKey string
	SecretKey string
}

var CurrentConfig Config
var ErrPortUndefined = errors.New("port is not defined")

func init() {
	loadMsg()
	bytes, err := os.ReadFile("./config.yaml")
	if err != nil {
		log.Fatal().Err(err)
	}
	err = yaml.Unmarshal(bytes, &CurrentConfig)
	if err != nil {
		log.Fatal().Err(err).Msg(yaml.FormatError(err, true, true))
	}
	log.Debug().Interface("CurrentConfig", CurrentConfig)
	if CurrentConfig.Port == "" {
		log.Fatal().Err(ErrPortUndefined)
	}
}
