package shortenurl

import (
	"net/http"

	"gorm.io/gorm"
)

type URL struct {
	ID        uint   `gorm:"primary_key"`
	Original  string `gorm:"not null"`
	Shortened string `gorm:"not null"`
}

func RedirectURL(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	// The code r.URL.Path[1:] takes a slice of the r.URL.Path string,
	// starting from the second character (index 1) to the end of the string.
	// This is done to extract the random string of the URL to be redirected, which is at the end of the URL.
	id := r.URL.Path[1:]
	var url URL
	shortened := "http://localhost:8080/" + id
	db.First(&url, "shortened = ?", shortened)
	http.Redirect(w, r, url.Original, http.StatusFound)
}
