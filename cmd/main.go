package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/RianAsmara/leaderboard-api/pkg/db"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

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
