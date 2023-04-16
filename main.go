package main

import (
	"myGram/database"
	"myGram/routers"
	"os"
)

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.
func main() {
	database.StartDB()
	var PORT = os.Getenv("PORT")
	routers.StartApp().Run(":" + PORT)
}
