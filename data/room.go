package data

import (
	"OctaneServer/config"
	"strconv"
	"time"

	"github.com/goccy/go-json"
	"github.com/rs/zerolog/log"
)

func NewRoom(name string) (int64, error) {
	var (
		room  Room
		id    int64
		strId string
	)
	for {
		id = SnowFlake()
		strId = strconv.FormatInt(id, 10)
		_, err := db.Get(TableRooms, strId)
		if err == ErrNoEntry {
			break
		} else if err != nil {
			return -1, err
		}
		log.Warn().Msg(config.Msg[config.CurrentConfig.Lang].Console.Error.DuplicatedID)
	}

	room = Room{
		Id:      id,
		Name:    name,
		Devices: []Device{},
	}
	err := db.Insert(TableRooms, strId, room)
	if err != nil {
		return -1, err
	}
	return id, nil
}

func GetRoom(id int64) (Room, error) {
	room := Room{}
	result, err := db.Get(TableRooms, strconv.FormatInt(id, 10))
	if err != nil {
		return room, err
	}
	log.Debug().Bytes("result", result).Msg("GetRoom: Returned result")
	err = json.Unmarshal(result, &room)
	if err != nil {
		return room, err
	}
	return room, nil
}

func DeleteRoom(id int64) error {
	return db.Delete(TableRooms, strconv.FormatInt(id, 10))
}

func ConnectRoom(name string, id int64) error {
	room, err := GetRoom(id)
	if err != nil {
		return err
	}
	room.Devices = append(room.Devices, Device{Name: name, Timestamp: time.Now().Unix()})
	err = db.Update(TableRooms, strconv.FormatInt(id, 10), room)
	if err != nil {
		return err
	}
	return nil
}
