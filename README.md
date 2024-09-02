# Real-Time Leaderboard System

This project is a real-time leaderboard system built with Go, Fiber, GORM, PostgreSQL, and Redis. It ranks users based on their scores in various games or activities, providing real-time updates and efficient querying using Redis sorted sets.

## Features

- **User Authentication**: Register and login via Google OAuth.
- **Score Submission**: Users can submit scores for various games or activities.
- **Leaderboard Updates**: Display a global leaderboard showing top users.
- **User Rankings**: Users can view their rankings on the leaderboard.
- **Top Players Report**: Generate reports on the top players for a specific period.
- **Role-Based Access Control (RBAC)**: Manage user permissions and roles.

## Prerequisites

Before setting up the project, ensure you have the following installed:

- Go 1.20+
- Docker & Docker Compose
- Redis
- PostgreSQL

## Environment Variables

Create a `.env` file in the project root with the following content:

```env
# Application
PORT=3333

# PostgreSQL
DATABASE_URL=postgres://postgres:postgres@localhost:5432/leaderboard?sslmode=disable

# Redis
REDIS_URL=redis://:yourpassword@localhost:6379

# Google OAuth Credentials
GOOGLE_CLIENT_ID=your-google-client-id
GOOGLE_CLIENT_SECRET=your-google-client-secret
REDIRECT_URL=http://localhost:3333/auth/google/callback
Replace yourpassword, your-google-client-id, and your-google-client-secret with the appropriate values.

```

## Setup and Run

Step 1: Clone the Repository
Clone the project repository to your local machine.

```
git clone <https://github.com/yourusername/real-time-leaderboard.git>
cd real-time-leaderboard
```

Step 2: Install Dependencies
Ensure all Go dependencies are installed.

```
go mod tidy
```

Step 3: Build and Run with Docker Compose
Use Docker Compose to build and run the application. The application will only start after PostgreSQL and Redis are up and running, verified through health checks.

```

docker-compose up --build

```

This command will:

Build the Go application.
Set up PostgreSQL and Redis with health checks to ensure they are ready before the app starts.
Launch the application, which listens on the port defined in the .env file.
Step 4: Run Migrations and Seeder
After the application is up, run migrations and seed the database with random data to simulate the leaderboard.

```

docker exec -it app-container-name bash
go run seeder/seed.go

```

## Testing

Run unit, integration, and benchmark tests using the following commands:

```

# Run all tests

go test ./... -v

# Run benchmark tests

go test -bench=.

# Run coverage tests

go test -cover

```

## Health Check

The application includes automatic health checks for PostgreSQL and Redis. Upon starting, the application verifies the connection to these services and logs the results:

PostgreSQL: Checks connectivity before performing any database operations.
Redis: Pings Redis to ensure it's ready for handling real-time leaderboard updates.

## Contact

For any issues or inquiries, you can reach out through the following channels:

Telegram: [@rianasmaraputra](https://t.me/rianasmaraputra)

<!-- Email: <rputra711@gmail.com> -->
