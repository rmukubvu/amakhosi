package repository

import (
	"encoding/json"
	"github.com/rmukubvu/amakhosi/model"
	"github.com/rmukubvu/amakhosi/store"
)

const locationBucket string = "loc.bucket"

//AddLocation to bolt db
func AddLocation(p model.Pumps) error {
	value, err := json.Marshal(p)
	if err != nil {
		return err
	}
	return store.Insert(locationBucket, p.ID, value)
}

//LocationById get location by id
func LocationById(key int) ([]model.Pumps, error) {
	buf, err := store.SingleOrDefault(locationBucket, key)
	if err != nil {
		return nil, err
	}
	return modelFromByte(buf), nil
}

func modelFromByte(data []byte) []model.Pumps {
	var result []model.Pumps
	err := json.Unmarshal(data, &result)
	if err != nil {
		return []model.Pumps{}
	}
	return result
}
