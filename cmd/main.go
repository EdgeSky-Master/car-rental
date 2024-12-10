package main

import (
	"car-rental/internal/routes"
	"car-rental/pkg/database"
)

func main() {
	//r := gin.Default()
	database.ConnectToMySQL()
	// if db != nil {
	// 	fmt.Println("Database connected")
	// }
	r := routes.SetupRouter()
	r.Run(":8081")
}