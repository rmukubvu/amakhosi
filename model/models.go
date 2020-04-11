package model

import "fmt"

//Pumps holds the pump locations
type Pumps struct {
	ID            int     `json:"id" db:"id"`
	LocationName  string  `json:"locationName" db:"location_name"`
	Latitude      float64 `json:"lat" db:"lat"`
	Longitude     float64 `json:"lon" db:"lon"`
	PumpReference string  `json:"pumpReference" db:"pump_reference"`
}

type InternalError struct {
	Message string `json:"message"`
}

type DatabaseConfig struct {
	User   string
	Pwd    string
	Host   string
	Port   int
	DbName string
}

type Account struct {
	ID            int    `db:"ID"`
	AccountNumber string `db:"account_number"`
	Identifier    string `db:"personal_identifier"`
	IsActive      bool   `db:"is_active"`
	CreatedDate   string `db:"created_date"`
}

func (dc *DatabaseConfig) String() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", dc.User, dc.Pwd, dc.Host, dc.Port, dc.DbName)
}
