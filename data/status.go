package data

import (
	"strconv"

	"github.com/goccy/go-json"
)

func SetStatus(id int64, status Status) error {
	_, err := GetRoom(id)
	if err != nil {
		return err
	}
	_, err = db.Get(TableStatuses, strconv.FormatInt(id, 10))
	if err == ErrNoEntry {
		err = db.Insert(TableStatuses, strconv.FormatInt(id, 10), &status)
		if err != nil {
			return err
		}
		return nil
	} else if err != nil {
		return err
	}
	err = db.Update(TableStatuses, strconv.FormatInt(id, 10), &status)
	if err != nil {
		return err
	}

	return nil
}

func DeleteStatus(id int64) error {
	_, err := GetRoom(id)
	if err != nil {
		return err
	}
	err = db.Delete(TableStatuses, strconv.FormatInt(id, 10))
	if err != nil {
		return err
	}
	return nil
}

func GetStatus(id int64) (Status, error) {
	status := Status{}
	_, err := GetRoom(id)
	if err != nil {
		return status, err
	}
	result, err := db.Get(TableStatuses, strconv.FormatInt(id, 10))
	if err != nil {
		return status, err
	}
	err = json.Unmarshal(result, &status)
	if err != nil {
		return status, err
	}
	return status, nil
}
