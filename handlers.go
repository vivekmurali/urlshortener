package main

import (
	bolt "go.etcd.io/bbolt"
	"net/http"
)

func getURLS(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Yay!!!"))
}

func newURL(w http.ResponseWriter, r *http.Request) {
	// w.Write([]byte("Yay!!!"))
	err := db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("widgets"))
		if err != nil {
			return err
		}
		err = b.Put([]byte("key"), []byte("value"))
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		w.Write([]byte(err.Error()))
	}
	w.Write([]byte("Added to DATABASE"))

}
