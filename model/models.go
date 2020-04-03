package model

//Pumps holds the pump locations
type Pumps struct {
	ID           int     `json:"id"`
	LocationName string  `json:"location_name"`
	Latitude     float64 `json:"lat"`
	Longitude    float64 `json:"lon"`
}
