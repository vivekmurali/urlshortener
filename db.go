package main

import (
	"log"

	bolt "go.etcd.io/bbolt"
)

func dbInit() *bolt.DB {
	db, err := bolt.Open("my.db", 0666, nil)
	if err != nil {
		log.Fatal(err)
		_ = db
	}
	return db
}
