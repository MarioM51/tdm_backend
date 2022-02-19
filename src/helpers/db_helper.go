package helpers

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DBHelper struct {
	DB *gorm.DB
}

func (dbHelper *DBHelper) Connect() {
	var db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		panic("connection database failed")
	}

	dbHelper.DB = db

}
