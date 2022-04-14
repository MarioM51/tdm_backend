package helpers

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DBHelper struct {
	DB *gorm.DB
}

func (dbHelper *DBHelper) Connect() {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,       // Disable color
		},
	)

	var db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		panic("connection database failed")
	}

	dbHelper.DB = db

}
