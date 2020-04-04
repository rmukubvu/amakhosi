package repository

import (
	"encoding/json"
	"github.com/mitchellh/mapstructure"
	"github.com/rmukubvu/amakhosi/model"
	"github.com/rmukubvu/amakhosi/store"
)

//AddLocation to bolt db
func AddLocation(p model.Pumps) error {
	value, err := json.Marshal(p)
	if err != nil {
		return err
	}
	return store.Insert("pumps", value)
}

//LocationById get location by id
func LocationsById(pumpReference string) ([]model.Pumps, error) {
	rs, err := store.Fetch("select * from pumps where pump_reference = ?", pumpReference)
	if err != nil {
		return nil, err
	}
	var result []model.Pumps
	err = mapstructure.Decode(rs, &result)
	return result, nil
}
