package repository

import (
	"github.com/rmukubvu/amakhosi/model"
	"github.com/rmukubvu/amakhosi/store"
)

//AddLocation to bolt db
func AddLocation(p model.Pumps) error {
	return store.Insert("pumps", &p)
}

//LocationById get location by id
func LocationsById(pumpReference string) ([]model.Pumps, error) {
	return store.FetchPumps("select * from pumps where pump_reference = ?", pumpReference)
}
