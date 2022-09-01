package data

import (
	"OctaneServer/config"

	"errors"

	"github.com/rs/zerolog/log"
)

type dataBase interface {
	Open() error
	Close() error
	Insert(table string, key string, data interface{}) error
	Update(table string, key string, data interface{}) error
	Get(table string, key string) ([]byte, error)
	Delete(table string, key string) error
	Drop(table string) error
}

var (
	db         dataBase
	ErrNoEntry = errors.New("no such entry")
)

const (
	MongoDB       = "mongo"
	BBoltDB       = "bbolt"
	TableNodes    = "nodes"
	TableRooms    = "rooms"
	TableStatuses = "statuses"
	TableContents = "contents"
)

func init() {
	switch config.CurrentConfig.Db.Type {
	case MongoDB:
		db = &Mongo{Ready: false}
		err := db.Open()
		if err != nil {
			log.Fatal().Err(err).Msg(config.Msg[config.CurrentConfig.Lang].Console.Error.DBOpen)
		}
	case BBoltDB:
		db = &Bbolt{Ready: false}
		//Open at RW time cuz multi-process open causes deadlock
	default:
		log.Fatal().Msg(config.Msg[config.CurrentConfig.Lang].Console.Error.DBNotSupported)
	}
	err := newSnowFlakeNode()
	if err != nil {
		log.Fatal().Err(err).Msg(config.Msg[config.CurrentConfig.Lang].Console.Error.SnowFlake)
	}
}

func Close() error {
	err := closeSnowFlakeNode()
	if err != nil {
		return err
	}
	return db.Close()
}
