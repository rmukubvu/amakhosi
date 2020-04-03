package store

import (
	"encoding/binary"
	"github.com/boltdb/bolt"
	"log"
)

//Db handler for bolt
var db *bolt.DB

//InitDB open db
func init() {
	// Open the pumplocations.db data file in your current directory.
	// It will be created if it doesn't exist.
	b, err := bolt.Open("pumplocations.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	db = b
}

//Itob returns an 8-byte big endian representation of v.
func Itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

/*
func GetNextID(bucketName string) (int, error) {
	// Open a writable transaction
	tx, err := db.Begin(true)
	if err != nil {
		return -1, err
	}
	defer tx.Rollback()

	b, _ := tx.CreateBucketIfNotExists([]byte(bucketName))
	seq, _ := b.NextSequence()

	// Commit changes
	if err := tx.Commit(); err != nil {
		return -1, err
	}
	return int(seq), nil
}

/Delete delete record from bucket
func Delete(bucketName string, key int) (err error) {

	db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		err = b.Delete(Itob(key))
		return err
	})
	return
}
*/

//SingleOrDefault get unique record
func SingleOrDefault(bucketName string, key int) ([]byte, error) {
	var buf []byte

	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		// SingleOrDefault only one key
		buf = b.Get(Itob(key))
		return nil
	})
	return buf, nil
}

//Insert add a record to a bucket
func Insert(bucketName string, key int, value []byte) error {
	// Open a writable transaction
	tx, err := db.Begin(true)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// SingleOrDefault bucket
	b, _ := tx.CreateBucketIfNotExists([]byte(bucketName))
	// Persist bytes to bucket.
	err = b.Put(Itob(key), value)
	if err != nil {
		return err
	}
	// Commit changes
	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}

/*
func InsertMany(bucketName string,key int,value []byte) error {
	// Start the transaction.
	tx, err := db.Begin(true)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// Retrieve the root bucket for the account.
	// Assume this has already been created when the account was set up.
	root := tx.Bucket([]byte(bucketName))

	// Setup the users bucket.
	bkt, err := root.CreateBucketIfNotExists([]byte("locations"))
	if err != nil {
		return err
	}

	// Generate an ID for the new user.
	userID, err := bkt.NextSequence()
	if err != nil {
		return err
	}
	u.ID = userID

	// Marshal and save the encoded user.
	if buf, err := json.Marshal(u); err != nil {
		return err
	} else if err := bkt.Put([]byte(strconv.FormatUint(u.ID, 10)), buf); err != nil {
		return err
	}

	// Commit the transaction.
	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}*/

func CloseDb() error {
	return db.Close()
}
