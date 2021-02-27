package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"

	bolt "go.etcd.io/bbolt"
)

type ErrorJson struct {
	Error string `json:"error"`
}

type URLJson struct {
	Url string `json:"url"`
}

type ResJson struct {
	Response string `json:"response"`
}

func getURLS(w http.ResponseWriter, r *http.Request) {
	// w.Write([]byte("Yay!!!"))
	shortURL := r.URL.Query()["short"][0]
	db, err := bolt.Open("mydb", 0666, nil)
	if err != nil {
		e := ErrorJson{
			Error: err.Error(),
		}
		b, _ := json.Marshal(e)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(b)
		return
	}
	err = db.View(func(tx *bolt.Tx) error {
		v := tx.Bucket([]byte("urlshortener")).Get([]byte(shortURL))
		e := URLJson{
			Url: string(v),
		}
		b, _ := json.Marshal(e)
		w.Write(b)

		return nil
	})
	if err != nil {
		e := ErrorJson{
			Error: err.Error(),
		}
		b, _ := json.Marshal(e)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(b)
		return
	}

	err = db.Close()
	if err != nil {
		log.Fatal(err)
		return
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
			e := ErrorJson{
				Error: err.Error(),
			}
			b, _ := json.Marshal(e)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(b)
			return
		}

		err = json.Unmarshal([]byte(gg), &data)
		if err != nil {
			e := ErrorJson{
				Error: err.Error(),
			}
			b, _ := json.Marshal(e)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(b)
			return
		}

		url = data.Url
		shortURL = data.Short

	} else {

		err := r.ParseForm()
		if err != nil {
			e := ErrorJson{
				Error: err.Error(),
			}
			b, _ := json.Marshal(e)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(b)
			return
		}
		url = r.FormValue("url")
		shortURL = r.FormValue("short")
		// fmt.Println(r.Form)

	}
	db, err := bolt.Open("mydb", 0666, nil)
	if err != nil {
		e := ErrorJson{
			Error: err.Error(),
		}
		b, _ := json.Marshal(e)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(b)
		return
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
		err = b.Put([]byte(shortURL), []byte(url))
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		e := ErrorJson{
			Error: err.Error(),
		}
		b, _ := json.Marshal(e)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(b)
		db.Close()
		return
	}
	err = db.Close()
	if err != nil {
		// w.Write([]byte(err.Error()))
		log.Fatal(err.Error())
	}
	e := ResJson{
		Response: "Added to database",
	}
	b, _ := json.Marshal(e)
	w.Write(b)

}
