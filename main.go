package main

import (
	"fmt"

	"github.com/tawseefnabi/url-shortner-v2/shortenurl"
)

func main() {
	db, err := shortenurl.Connnect()
	fmt.Println("db", db, err)
}
