package db

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// DB is the global database connection for PostgreSQL.
var DB *gorm.DB

// RedisClient is the global Redis client.
var RedisClient *redis.Client

// InitPostgres initializes the PostgreSQL database connection using GORM.
func InitPostgres() {
	// Get the PostgreSQL connection string from the environment variable.
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("DATABASE_URL environment variable is not set")
	}

	// Open a connection to PostgreSQL using GORM with custom configurations.
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // Set the logging level to Info.
	})
	if err != nil {
		log.Fatalf("Failed to connect to PostgreSQL: %v", err)
	}

	// Test the connection by pinging the database.
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatalf("Failed to get generic database object from GORM DB: %v", err)
	}

	// Set database connection pool configurations.
	sqlDB.SetMaxIdleConns(10)           // Maximum number of idle connections.
	sqlDB.SetMaxOpenConns(100)          // Maximum number of open connections.
	sqlDB.SetConnMaxLifetime(time.Hour) // Maximum amount of time a connection may be reused.

	// Ping the database to verify that the connection is alive.
	if err := sqlDB.Ping(); err != nil {
		log.Fatalf("Failed to ping PostgreSQL database: %v", err)
	}

	log.Println("Successfully connected to PostgreSQL")
}

// InitRedis initializes the Redis client connection.
func InitRedis() {
	// Get the Redis URL from the environment variable.
	redisURL := os.Getenv("REDIS_URL")
	if redisURL == "" {
		log.Fatal("REDIS_URL environment variable is not set")
	}

	// Parse the Redis URL to extract options including password.
	opt, err := redis.ParseURL(redisURL)
	if err != nil {
		log.Fatalf("Failed to parse Redis URL: %v", err)
	}

	// Initialize the Redis client with the parsed options.
	RedisClient = redis.NewClient(opt)

	// Set a context with a timeout to handle the ping.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Try to ping the Redis server to test the connection.
	_, err = RedisClient.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}

	log.Println("Successfully connected to Redis")
}

// Close closes the database connections gracefully.
func Close() {
	// Close the PostgreSQL connection pool.
	if DB != nil {
		sqlDB, err := DB.DB()
		if err != nil {
			log.Printf("Error retrieving database object from GORM DB: %v", err)
		} else {
			if err := sqlDB.Close(); err != nil {
				log.Printf("Error closing PostgreSQL connection: %v", err)
			}
		}
	}

	// Close the Redis connection.
	if RedisClient != nil {
		if err := RedisClient.Close(); err != nil {
			log.Printf("Error closing Redis connection: %v", err)
		}
	}
}
