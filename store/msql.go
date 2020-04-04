package store

import (
	"database/sql"
	"github.com/rmukubvu/amakhosi/resource"
	"github.com/square/squalor"
)
import _ "github.com/go-sql-driver/mysql"

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

func Insert(bind string, model interface{}) error {
	_, err := db.BindModel(bind, model)
	if err != nil {
		panicOnError(err)
	}
	err = db.Insert(model)
	return err
}

func Fetch(query string, args ...interface{}) (map[string]interface{}, error) {
	rows, err := db.Query(query, args...) // Note: Ignoring errors for brevity
	if err != nil {
		panicOnError(err)
	}
	cols, _ := rows.Columns()

	for rows.Next() {
		// Create a slice of interface{}'s to represent each column,
		// and a second slice to contain pointers to each item in the columns slice.
		columns := make([]interface{}, len(cols))
		columnPointers := make([]interface{}, len(cols))
		for i, _ := range columns {
			columnPointers[i] = &columns[i]
		}
		// Scan the result into the column pointers...
		if err := rows.Scan(columnPointers...); err != nil {
			return nil, err
		}
		// Create our map, and retrieve the value for each column from the pointers slice,
		// storing it in the map with the name of the column as the key.
		m := make(map[string]interface{})
		for i, colName := range cols {
			val := columnPointers[i].(*interface{})
			m[colName] = *val
		}
		return m, nil
	}
	return nil, nil
}

func panicOnError(err error) {
	if err != nil {
		closeDb()
		panic(err)
	}
}

func closeDb() {
	if db != nil {
		db.Close()
	}
}
