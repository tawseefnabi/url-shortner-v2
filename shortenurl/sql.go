package shortenurl

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Connnect() (*gorm.Db, error) {
	db, err := gorm.Open(sqlite.Open("urls.db", &gorm.Config{}))
	if err != nil {

		return nil, err
	}
	return db, nil
}
