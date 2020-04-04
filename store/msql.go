package store

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"github.com/rmukubvu/amakhosi/resource"
	"github.com/square/squalor"
)
import _ "github.com/go-sql-driver/mysql"

//var db *sqlx.DB
var db *squalor.DB

func init() {
	const driver string = "mysql"
	dataSourceName := resource.DataSourceName()
	sdb, err := sql.Open(driver, dataSourceName)
	if err != nil {
		panic(err)
	}
	db, _ = squalor.NewDB(sdb)
}

func Insert(query string, args ...interface{}) (int64, error) {
	result := db.MustExec(query, args...)
	return result.RowsAffected()
}

func Fetch(query string, args ...interface{}) (dest interface{}, err error) {
	err = db.Select(&dest, query, args...)
	return
}

func CloseDb() {
	if db != nil {
		db.Close()
	}
}
