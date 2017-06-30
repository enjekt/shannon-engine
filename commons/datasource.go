package commons

import (
	"fmt"
	"github.com/boltdb/bolt"
	"log"
)

type Database struct {
	DBName string
	DB     *bolt.DB
}

func (database *Database) Open() {

	db, err := bolt.Open(database.DBName, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	database.DB = db

}

func (database *Database) Close() {
	database.DB.Close()
}

func (database *Database) Upsert(bucket string, key string, value string) {
	database.Open()
	defer database.Close()
	fmt.Printf("Adding value %s for bucket %s and key %s \n", value, bucket, key)
	database.DB.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucket([]byte(bucket))
		fmt.Printf("Bucket: %q \n", bucket)

		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		keyArr := []byte(key)
		fmt.Printf("Key: %q \n", keyArr)
		valArr := []byte(value)
		fmt.Printf("Value: %q \n", valArr)
		err = bucket.Put(keyArr, valArr)
		fetched := bucket.Get(keyArr)
		fmt.Printf("Feteched byte[] from db: %q \n", fetched)

		if err != nil {
			return fmt.Errorf("Put bucket: %s", err)
		}
		return nil
	})

}

func (database *Database) Get(bucket string, key string) string {

	database.Open()
	defer database.Close()
	fmt.Printf("Retrieve value for bucket %s and key %s \n", bucket, key)
	var value string

	database.DB.View(func(tx *bolt.Tx) error {

		bucket := tx.Bucket([]byte(bucket))
		fmt.Printf("Bucket: %q \n", bucket)

		if bucket == nil {
			return fmt.Errorf("Bucket %q not found!", bucket)
		}
		fetched := bucket.Get([]byte(key))
		//fmt.Printf("Fetehed byte[] from db: %q \n",fetched)
		if fetched == nil {
			return fmt.Errorf("Value for token %q not found!", key)
		}

		value = string(fetched)
		//fmt.Printf("String returned from db Get: %s \n",value)
		return nil
	})

	return value
}
