package database

import (
	"fmt"
	"log"
	"myGram/models"
	"os"

	_ "github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// var (
// 	errLoad = godotenv.Load()
// )

var (
	host     = os.Getenv("PGHOST")
	port     = os.Getenv("PGPORT")
	user     = os.Getenv("PGUSER")
	password = os.Getenv("PGPASSWORD")
	dbname   = os.Getenv("PGDATABASE")
	db       *gorm.DB
	err      error
)

func StartDB() {
	// if errLoad != nil {
	// 	fmt.Println("Error loading .env file")
	// 	os.Exit(1)
	// }

	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, port)

	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}

	fmt.Println("Success connecting to database.")
	db.Debug().AutoMigrate(models.User{}, models.Socialmedia{}, models.Photo{}, models.Comment{})
}

func GetDB() *gorm.DB {
	return db
}
