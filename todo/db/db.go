package db

import (
	"fmt"
	"log"

	"github.com/mitchellh/go-homedir"

	"github.com/boltdb/bolt"
)

// Open opens the database and return a instance of it
func Open() (*bolt.DB, error) {

	homeDirectory, err := homedir.Dir()
	if err != nil {
		return nil, err
	}

	db, err := bolt.Open(homeDirectory+"/todo.db", 0644, nil)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return db, nil

}

// Update store some data to the database
func Update(db *bolt.DB, bucket, key, value string) error {
	// store some data
	err := db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte(bucket))
		if err != nil {
			return err
		}

		err = bucket.Put([]byte(key), []byte(value))
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}

// Read retrieve data based on the passed key
func Read(db *bolt.DB, bucket, key []byte) error {
	// retrieve the data
	err := db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(bucket)
		if bucket == nil {
			return fmt.Errorf("bucket %q not found", bucket)
		}

		val := bucket.Get(key)
		fmt.Println(string(val))

		return nil
	})

	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}

// ReadAll returns all keys
func ReadAll(db *bolt.DB, bucket string) (map[string]string, error) {
	values := make(map[string]string)
	db.View(func(tx *bolt.Tx) error {
		// Assume bucket exists and has keys
		b := tx.Bucket([]byte(bucket))

		if b == nil {
			return nil
		}

		b.ForEach(func(k, v []byte) error {
			values[string(k)] = string(v)
			return nil
		})
		return nil
	})

	return values, nil
}
