package db

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/jackc/pgx/v5"
)

const (
	postgresURL = "postgres://username:password@localhost:5432/leaderboard_db" // Update with your DB details.
)

// userNames is a sample list of usernames for seeding.
var userNames = []string{"Alice", "Bob", "Charlie", "David", "Eve", "Frank", "Grace", "Hank", "Ivy", "Jack"}

func main() {
	// Connect to PostgreSQL.
	conn, err := pgx.Connect(context.Background(), postgresURL)
	if err != nil {
		log.Fatalf("Unable to connect to PostgreSQL: %v", err)
	}
	defer conn.Close(context.Background())

	// Seed users and scores.
	seedUsersAndScores(conn)
}

// seedUsersAndScores creates random users and assigns random scores.
func seedUsersAndScores(conn *pgx.Conn) {
	rand.Seed(time.Now().UnixNano())

	// Create users.
	for _, name := range userNames {
		_, err := conn.Exec(context.Background(), "INSERT INTO users (username) VALUES ($1) ON CONFLICT (username) DO NOTHING", name)
		if err != nil {
			log.Printf("Failed to insert user %s: %v", name, err)
		}
	}

	// Assign random scores to each user.
	for i := 0; i < 50; i++ { // Generate 50 random scores.
		userID := rand.Intn(len(userNames)) + 1
		score := rand.Float64() * 100
		_, err := conn.Exec(context.Background(), "INSERT INTO scores (user_id, score) VALUES ($1, $2)", userID, score)
		if err != nil {
			log.Printf("Failed to insert score for user_id %d: %v", userID, err)
		} else {
			fmt.Printf("Inserted score %.2f for user_id %d\n", score, userID)
		}
	}
}
