package database

import (
	"os"

	"gorm.io/gorm"
)

var (
	host     = os.Getenv("PGHOST")
	port     = os.Getenv("PGPORT")
	user     = os.Getenv("PGUSER")
	password = os.Getenv("PGPASSWORD")
	dbname   = os.Getenv("PGDBNAME")
	db       *gorm.DB
	err      error
)

func StartDB() {

}
