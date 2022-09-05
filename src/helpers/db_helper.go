package helpers

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DBHelper struct {
	DB *gorm.DB
}

func (dbHelper *DBHelper) Connect(constants Constants, loggerPrinter *log.Logger) {
	newLogger := logger.New(
		loggerPrinter, // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,       // Disable color
		},
	)

	var addOns = ""
	if constants.IsProduction() {
		addOns = "sslmode=require"
	} else if constants.IsLocal() {
		addOns = "sslmode=disable"
	} else {
		log.Panicf("Enviroment '%v' not defined, available: %v", constants.Env, constants.AvalaibleEnviroments())
	}

	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v TimeZone=%v "+addOns,
		constants.DatabaseCredentials.host,
		constants.DatabaseCredentials.user,
		constants.DatabaseCredentials.pass,
		constants.DatabaseCredentials.dbname,
		constants.DatabaseCredentials.port,
		constants.DatabaseCredentials.timezone,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})

	/*
		var db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{
			Logger: newLogger,
		})
	*/

	if err != nil {
		panic("connection database failed")
	}

	dbHelper.DB = db

}
