package repository

import (
	"io/ioutil"
	"encoding/json"
	"playground/model"
)

func LoadTimeline(filePath string) (*model.TimelineHistory, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	history := &model.TimelineHistory{}

	err = json.Unmarshal(data, &history)
	if err != nil {
		return nil, err
	}

	return history, nil
}