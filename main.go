package main

import (
	"flag"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/randijulio13/gogin/config"
	"github.com/randijulio13/gogin/database"
	"github.com/randijulio13/gogin/migrate"
	"github.com/randijulio13/gogin/route"
)

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("failed load .env")
	}
}

func connectDatabase() {
	database.ConnectDB()
	migrate.Run()
}

func run() {
	appEnv := config.Config("APP_ENV", "debug")
	gin.SetMode(appEnv)

	route.Route = gin.Default()
	route.Setup()

	port := config.Config("APP_PORT", "3000")
	route.Route.Run(fmt.Sprintf(":%s", port))
}

func checkFlag() {
	seedFlag := flag.Bool("seed", false, "run database seed")
	flag.Parse()
	if *seedFlag {
		fmt.Println("seeding database...")
		migrate.Seed()
	}
}

func main() {
	loadEnv()
	connectDatabase()
	checkFlag()
	run()
}
