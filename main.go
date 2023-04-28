// Package main is the entry point of the application
package main

import (
	"encoding/json"
	"fmt"
	"hexagonal-fiber/cmd"
	databsDomain "hexagonal-fiber/domain/database"
	secureDomain "hexagonal-fiber/domain/security"

	"hexagonal-fiber/infrastructure/repository/postgres"
	"hexagonal-fiber/infrastructure/repository/redis"

	"log"
	"os"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/spf13/viper"
	"github.com/valyala/fasthttp"

	"hexagonal-fiber/infrastructure/restapi/routes"
)

// main services
func main() {

	// initialize config
	router := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})
	router.Use(logger.New())
	router.Use(limiter.New())
	router.Use(cors.New(cors.ConfigDefault))

	// init databases
	databases := initDatabase()

	// commands handler
	cmd.Execute(databases)

	// getting key ssh
	err := secureDomain.GettingKeySSH()
	if err != nil {
		panic(fmt.Errorf("fatal error in getting key ssh: %s", err))
	}

	// root routes
	routes.ApplicationRootRouter(router, databases)

	// postgres routes
	routes.ApplicationV1Router(router, databases)

	// running config
	startServer(router)
}

// initial databases
func initDatabase() databsDomain.Database {
	// postgres connection
	postgresDB, err := postgres.NewGorm()
	if err != nil {
		panic(fmt.Errorf("fatal error in postgres: %s", err))
	}

	// redis connection
	redisDB, err := redis.InitRedis()
	if err != nil {
		panic(fmt.Errorf("fatal error in redis: %s", err))
	}

	return databsDomain.Database{
		Postgre: postgresDB,
		Redis:   redisDB,
	}
}

// start server config
func startServer(app *fiber.App) {
	viper.SetConfigFile("config.json")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("fatal error in config file: %s", err.Error())
	}

	// check environment
	environment := os.Getenv("ENV")
	if environment == "railway-production" {
		log.Fatal(app.Listen(":"))

	} else {
		serverPort := fmt.Sprintf(":%s", viper.GetString("ServerPort"))

		s := &fasthttp.Server{
			Handler:            app.Handler(),
			ReadTimeout:        18000 * time.Second,
			WriteTimeout:       18000 * time.Second,
			MaxRequestBodySize: 1000 << 20,
		}

		if err := s.ListenAndServe(serverPort); err != nil {
			log.Fatalf("fatal error description: %s", strings.ToLower(err.Error()))
		}
	}
}
