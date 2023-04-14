package main

import (
	"myGram/database"
	"myGram/routers"
	"os"
)

func main() {
	database.StartDB()
	var PORT = os.Getenv("PORT")
	routers.StartApp().Run(":" + PORT)
}
