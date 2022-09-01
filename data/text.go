package data

import (
	"fmt"
	"strconv"

	"github.com/goccy/go-json"
)

type textContainer struct {
	Text string
}

func SetTextContent(id int64, text string) error {
	_, err := GetRoom(id)
	if err != nil {
		return err
	}
	_, err = db.Get(TableContents, strconv.FormatInt(id, 10))
	if err == ErrNoEntry {
		err = db.Insert(TableContents, strconv.FormatInt(id, 10), textContainer{Text: text})
		if err != nil {
			return err
		}
		return nil
	} else if err != nil {
		return err
	}
	err = db.Update(TableContents, strconv.FormatInt(id, 10), textContainer{Text: text})
	if err != nil {
		return err
	}

	return nil
}

func DeleteTextContent(id int64) error {
	_, err := GetRoom(id)
	if err != nil {
		return err
	}
	err = db.Delete(TableContents, strconv.FormatInt(id, 10))
	if err != nil {
		return err
	}
	return nil
}

func GetTextContent(id int64) (string, error) {
	content := textContainer{}
	_, err := GetRoom(id)
	if err != nil {
		return content.Text, err
	}
	result, err := db.Get(TableContents, strconv.FormatInt(id, 10))
	if err != nil {
		return content.Text, err
	}
	fmt.Print(string(result))
	err = json.Unmarshal(result, &content)
	if err != nil {
		return content.Text, err
	}
	return content.Text, nil
}
