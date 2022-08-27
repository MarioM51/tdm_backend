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

type _DBCredentials struct {
	host     string
	user     string
	pass     string
	dbname   string
	port     string
	timezone string
}

func (*DBHelper) getCredentials(env string) _DBCredentials {
	var dbCredentials _DBCredentials
	if env == "local" {
		dbCredentials = _DBCredentials{
			host:     "localhost",
			user:     "postgres",
			pass:     "postgres",
			dbname:   "carro_de_madera_db",
			port:     "5432",
			timezone: "America/Mexico_City",
		}
	} else if env == "prod" {
		dbCredentials = _DBCredentials{
			host:     "carrodemaderadb.postgres.database.azure.com",
			user:     "carrodemaderauserdb",
			pass:     "plumonrojO(88)",
			dbname:   "carro_de_madera_db",
			port:     "5432",
			timezone: "America/Mexico_City",
		}
	} else {
		panic(fmt.Sprintf("env '%v' not defined", env))
	}

	return dbCredentials
}

func (dbHelper *DBHelper) Connect(env string, loggerPrinter *log.Logger) {
	newLogger := logger.New(
		loggerPrinter, // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,       // Disable color
		},
	)

	credentials := dbHelper.getCredentials(env)

	var addOns = ""
	if env == "prod" {
		addOns = "sslmode=require"
	} else if env == "local" {
		addOns = "sslmode=disable"
	}

	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v TimeZone=%v "+addOns,
		credentials.host, credentials.user, credentials.pass, credentials.dbname, credentials.port, credentials.timezone,
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
