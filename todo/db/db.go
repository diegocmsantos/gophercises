package db

import (
	"encoding/binary"
	"log"
	"time"

	"github.com/boltdb/bolt"
)

var db *bolt.DB

func Init(dbPath string, bucketName string) error {
	var err error
	db, err = bolt.Open(dbPath, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return err
	}

	return db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(bucketName))
		return err
	})
}

// Update store some data to the database
func Update(bucket, task string) (int, error) {

	var id int
	err := db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(bucket))
		id64, _ := bucket.NextSequence()
		id = int(id64)
		key := itob(id)
		return bucket.Put(key, []byte(task))

	})

	if err != nil {
		log.Fatal(err)
		return -1, err
	}

	return id, nil
}

// ReadAll returns all keys
func ReadAll(bucket string) (map[int]string, error) {
	values := make(map[int]string)
	err := db.View(func(tx *bolt.Tx) error {
		// Assume bucket exists and has keys
		b := tx.Bucket([]byte(bucket))
		cursor := b.Cursor()

		for k, v := cursor.First(); k != nil; k, v = cursor.Next() {
			values[btoi(k)] = string(v)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return values, nil
}

func Delete(bucket string, key int) error {
	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		return b.Delete(itob(key))
	})
}

func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

func btoi(b []byte) int {
	return int(binary.BigEndian.Uint64(b))
}
