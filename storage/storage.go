package storage

import (
	"OctaneServer/config"

	"github.com/rs/zerolog/log"
)

const (
	S3Storage    = "s3"
	LocalStorage = "local"
)

var (
	Storage storage
)

type storage interface {
	Open() error
	Save(key string, data []byte) error
	Delete(key string) error
}

func init() {
	switch config.CurrentConfig.Storage.Type {
	case S3Storage:
		Storage = &S3{}
	case LocalStorage:
		Storage = &Local{}
	default:
		log.Fatal().Msg(config.Msg[config.CurrentConfig.Lang].Console.Error.StorageNotSupported)
	}
	err := Storage.Open()
	if err != nil {
		log.Fatal().Err(err).Msg(config.Msg[config.CurrentConfig.Lang].Console.Error.StorageOpen)
	}
}

func Uri(key string) string {
	if config.CurrentConfig.Storage.Type == LocalStorage {
		return config.CurrentConfig.Root + "uploads/" + key
	}
	return config.CurrentConfig.Storage.Endpoint + config.CurrentConfig.Storage.Name + "/" + key
}
