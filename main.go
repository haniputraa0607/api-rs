package main

import (
	"api-rs/configs"
	"api-rs/database"
	"api-rs/middlewares"
	"api-rs/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var db *gorm.DB = configs.SetupDatabaseConnection()

func main() {
	defer configs.CloseDatabaseConnection(db)

	argsWithProg := os.Args

	if len(argsWithProg) == 1 {
		argsWithProg = append(argsWithProg, "")
	}

	switch argsWithProg[1] {
	case "migrate":
		if len(argsWithProg) == 2 {
			log.Fatal("Invalid argument")
			os.Exit(0)
		}

		e := database.Migrate(db, argsWithProg[2:])
		if e != nil {
			log.Fatal(e)
		}
		log.Print("Success Migrate")
	case "seed":
		if len(argsWithProg) == 2 {
			log.Fatal("Invalid argument")
			os.Exit(0)
		}

		e := database.Seed(db, argsWithProg[2:])
		if e != nil {
			log.Fatal(e)
		}
		log.Print("Success Seed")
	case "drop":
		if len(argsWithProg) == 2 {
			log.Fatal("Invalid argument")
			os.Exit(0)
		}

		e := database.Drop(db, argsWithProg[2:])
		if e != nil {
			log.Fatal(e)
		}
		log.Print("Success Drop")
	default:
		r := gin.Default()
		r.Use(middlewares.CorsMiddleware())

		routes.SetupRouter(r, db)

		port := configs.GetPort()
		r.Run(port)
	}
}
