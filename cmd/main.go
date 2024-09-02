package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/RianAsmara/leaderboard-api/pkg/db"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	fiberSwagger "github.com/swaggo/fiber-swagger"
	_ "github.com/swaggo/fiber-swagger/example/docs"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io
// @BasePath /v2
func main() {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Initialize PostgreSQL and Redis connections.
	// db.InitPostgres()
	db.InitRedis()

	// Wait until both Redis and PostgreSQL are ready.
	// waitForPostgres()
	waitForRedis()

	app := fiber.New()

	// Routes
	app.Get("/", HealthCheck)

	app.Get("/swagger/*", fiberSwagger.WrapHandler)

	// Retrieve port from environment variable or default to 8080
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	// Start the server on the specified port
	log.Printf("Server is running on port %s", port)
	if err := app.Listen(":" + port); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}

}

// waitForPostgres repeatedly pings the PostgreSQL database until it's ready.
// func waitForPostgres() {
// 	for {
// 		sqlDB, err := db.DB.DB()
// 		if err == nil {
// 			if err := sqlDB.Ping(); err == nil {
// 				log.Println("PostgreSQL is ready.")
// 				return
// 			}
// 		}
// 		log.Println("Waiting for PostgreSQL to be ready...")
// 		time.Sleep(2 * time.Second)
// 	}
// }

// waitForRedis repeatedly pings the Redis server until it's ready.
func waitForRedis() {
	for {
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()

		_, err := db.RedisClient.Ping(ctx).Result()
		if err == nil {
			log.Println("Redis is ready.")
			return
		}
		log.Println("Waiting for Redis to be ready...")
		time.Sleep(2 * time.Second)
	}
}

// HealthCheck godoc
// @Summary Show the status of server.
// @Description get the status of server.
// @Tags root
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router / [get]
func HealthCheck(c *fiber.Ctx) error {
	res := map[string]interface{}{
		"data": "Server is up and running",
	}

	if err := c.JSON(res); err != nil {
		return err
	}

	return nil
}
