package database

import (
	"blogos/src/config"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func Connect() (*gorm.DB, error) {
	db, err := gorm.Open(config.DB_DRIVER, config.DBURL)
	if err != nil {
		return nil, err
	}

	return db, nil
}
