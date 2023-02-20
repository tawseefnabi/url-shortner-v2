package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/tawseefnabi/url-shortner-v2/shortenurl"
)

func main() {
	db, err := shortenurl.Connect()
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("db", db, &shortenurl.URL{}, err)
	db.AutoMigrate(&shortenurl.URL{})
	http.HandleFunc("/shorten", func(w http.ResponseWriter, r *http.Request) {
		original := r.FormValue("url")
		shortened := shortenurl.ShortenUrl(original)
		// fmt.Printf(shortened)
		jsonBody, err := json.Marshal(shortened)

		db.Create(&shortenurl.URL{Original: original, Shortened: shortened})
		if err != nil {
			log.Println("error: ", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"status":"failed","message": "failed to unmarshal"}`))
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(jsonBody)
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		shortenurl.RedirectURL(db, w, r)
	})
	http.ListenAndServe(":8080", nil)
}
