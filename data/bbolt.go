package data

import (
	"OctaneServer/config"

	"github.com/goccy/go-json"
	"github.com/rs/zerolog/log"
	"go.etcd.io/bbolt"
)

type Bbolt struct {
	Type  string
	Ready bool
	db    *bbolt.DB
}

func (db *Bbolt) Open() error {
	tmpdb, err := bbolt.Open(config.CurrentConfig.Db.Uri, 0666, nil)
	if err != nil {
		return err
	}
	db.db = tmpdb
	db.Ready = true
	db.Type = BBoltDB
	return nil
}

func (db *Bbolt) Close() error {
	return db.db.Close()
}

func (db *Bbolt) Drop(table string) error {
	err := db.Open()
	if err != nil {
		log.Fatal().Err(err).Msg(config.Msg[config.CurrentConfig.Lang].Console.Error.DBOpen)
		return err
	}
	defer func() {
		err = db.Close()
		if err != nil {
			log.Fatal().Err(err).Msg(config.Msg[config.CurrentConfig.Lang].Console.Error.DBOpen)
		}
	}()
	return db.db.Update(func(tx *bbolt.Tx) error {
		err := tx.DeleteBucket([]byte(table))
		if err != nil {
			return err
		}
		return nil
	})
}

func (db *Bbolt) Delete(table string, key string) error {
	bKey := []byte(key)
	err := db.Open()
	if err != nil {
		log.Fatal().Err(err).Msg(config.Msg[config.CurrentConfig.Lang].Console.Error.DBOpen)
		return err
	}
	defer func() {
		err = db.Close()
		if err != nil {
			log.Fatal().Err(err).Msg(config.Msg[config.CurrentConfig.Lang].Console.Error.DBOpen)
		}
	}()

	return db.db.Update(func(tx *bbolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte(table))
		if err != nil {
			return err
		}
		if bucket.Get(bKey) == nil {
			return ErrNoEntry
		}
		err = bucket.Delete(bKey)
		if err != nil {
			return err
		}
		return nil
	})
}

func (db *Bbolt) Get(table string, key string) ([]byte, error) {
	bKey := []byte(key)
	err := db.Open()
	if err != nil {
		log.Fatal().Err(err).Msg(config.Msg[config.CurrentConfig.Lang].Console.Error.DBOpen)
		return nil, err
	}
	defer func() {
		err = db.Close()
		if err != nil {
			log.Fatal().Err(err).Msg(config.Msg[config.CurrentConfig.Lang].Console.Error.DBOpen)
		}
	}()

	var result []byte
	err = db.db.View(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket([]byte(table))
		if bucket == nil {
			return ErrNoEntry
		}
		result = bucket.Get(bKey)
		if result == nil {
			return ErrNoEntry
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return result, err
}

func (db *Bbolt) Insert(table string, key string, data interface{}) error {
	bKey := []byte(key)
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	err = db.Open()
	if err != nil {
		log.Fatal().Err(err).Msg(config.Msg[config.CurrentConfig.Lang].Console.Error.DBOpen)
		return err
	}
	defer func() {
		err = db.Close()
		if err != nil {
			log.Fatal().Err(err).Msg(config.Msg[config.CurrentConfig.Lang].Console.Error.DBOpen)
		}
	}()
	return db.db.Update(func(tx *bbolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte(table))
		if err != nil {
			return err
		}
		err = bucket.Put(bKey, jsonData)
		if err != nil {
			return err
		}
		return nil
	})
}

func (db *Bbolt) Update(table string, key string, data interface{}) error {
	return db.Insert(table, key, &data)
}
