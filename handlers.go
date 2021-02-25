package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"

	bolt "go.etcd.io/bbolt"
)

func getURLS(w http.ResponseWriter, r *http.Request) {
	// w.Write([]byte("Yay!!!"))
	URL := r.URL.Query()["url"][0]
	db, err := bolt.Open("mydb", 0666, nil)
	if err != nil {
		w.Write([]byte(err.Error()))
	}
	err = db.View(func(tx *bolt.Tx) error {
		v := tx.Bucket([]byte("urlshortener")).Get([]byte(URL))
		w.Write(v)

		return nil
	})
	if err != nil {
		w.Write([]byte(err.Error()))
	}

	err = db.Close()
	if err != nil {
		log.Fatal(err)
	}

}

type Body struct {
	Url   string `json:"url"`
	Short string `json:"short"`
}

func newURL(w http.ResponseWriter, r *http.Request) {

	var (
		url,
		shortURL string
	)
	if r.Header.Get("content-type") == "application/json" {
		var data Body
		gg, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.Write([]byte(err.Error()))
		}
		err = json.Unmarshal([]byte(gg), &data)
		if err != nil {
			w.Write([]byte(err.Error()))
		}

		url = data.Url
		shortURL = data.Short

	} else {

		err := r.ParseForm()
		if err != nil {
			w.Write([]byte(err.Error()))
		}
		url = r.FormValue("url")
		shortURL = r.FormValue("short")
		// fmt.Println(r.Form)

	}
	db, err := bolt.Open("mydb", 0666, nil)
	if err != nil {
		w.Write([]byte(err.Error()))
	}

	err = db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("urlshortener"))
		if err != nil {

			return err
		}
		v := b.Get([]byte(url))
		if v != nil {
			return errors.New("URL already exists")
		}
		err = b.Put([]byte(url), []byte(shortURL))
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		w.Write([]byte(err.Error()))
		db.Close()
		return
	}
	err = db.Close()
	if err != nil {
		// w.Write([]byte(err.Error()))
		log.Fatal(err.Error())
	}
	w.Write([]byte("Added to DATABASE"))

}
