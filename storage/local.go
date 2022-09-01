package storage

import (
	"os"
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
