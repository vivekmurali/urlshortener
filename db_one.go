package main

import (
	"fmt"
	"log"

	bolt "go.etcd.io/bbolt"
)

func insert(db *bolt.DB) {
	db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("widgets"))
		if err != nil {
			return err
		}
		err = b.Put([]byte("key"), []byte("value"))
		// err = b.Put([]byte("hello"), []byte("test"))
		if err != nil {
			return err
		}
		return nil
	})

}
func view(db *bolt.DB) {
	log.Println("fdksal;")
	db.View(func(tx *bolt.Tx) error {
		v := tx.Bucket([]byte("widgets")).Get([]byte("ky"))
		fmt.Printf("v = %s\n", v)
		return nil
	})
}

// func main() {
// 	db, err := bolt.Open("my.db", 0666, nil)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	// insert(db)
// 	view(db)

// }
