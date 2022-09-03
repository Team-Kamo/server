package storage

import (
	"OctaneServer/config"
	"os"
	"strings"

	"github.com/rs/zerolog/log"
)

type Local struct {
}

func (storage *Local) Open() error {
	return nil
}

func (storage *Local) Save(key string, data []byte) error {
	err := os.Mkdir("uploads", 0777)
	if err != nil {
		return err
	}
	return os.WriteFile("uploads/"+key, data, 0660)
}

func (storage *Local) Delete(key string) error {
	return os.Remove("uploads/" + key)
}

func (storage *Local) HouseKeeping() {
	dir, err := os.ReadDir("uploads/")
	if err != nil {
		log.Warn().Err(err).Msg(config.Msg[config.CurrentConfig.Lang].Console.Error.HouseKeeping)
		return
	}
	for _, v := range dir {
		delete := false
		for _, w := range dir {
			if strings.HasPrefix(w.Name(), strings.Split(v.Name(), "_")[0]) && w.Name() != v.Name() {
				vInfo, err := v.Info()
				if err != nil {
					log.Warn().Err(err).Msg(config.Msg[config.CurrentConfig.Lang].Console.Error.HouseKeeping)
					return
				}
				wInfo, err := w.Info()
				if err != nil {
					log.Warn().Err(err).Msg(config.Msg[config.CurrentConfig.Lang].Console.Error.HouseKeeping)
					return
				}
				if vInfo.ModTime().Before(wInfo.ModTime()) {
					delete = true
				}
			}
		}
		if delete {
			err = storage.Delete(v.Name())
			if err != nil {
				log.Warn().Err(err).Msg(config.Msg[config.CurrentConfig.Lang].Console.Error.HouseKeeping)
				return
			}
		}
	}
}
