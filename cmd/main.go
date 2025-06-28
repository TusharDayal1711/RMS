package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"rmssystem_1/config"
	"rmssystem_1/routes"

	"rmssystem_1/database"
)

func main() {
	config.LoadEnv()
	dbConnectionString := config.GetDatabaseString()
	database.Init(dbConnectionString)
	//handler.CreateSuperAdmin()
	defer database.DB.Close()
	r := routes.GetRoutes()
	fmt.Println("Starting server on port " + os.Getenv("SERVER_PORT"))
	fmt.Println("Server is running on http://localhost:8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal("Server failed:", err)
	}
}
