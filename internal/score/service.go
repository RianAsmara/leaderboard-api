package score

import "time"

type Service interface {
	GetLeaderboard() ([]Score, error)
	AddScore(userID uint, game string, score float64) error
	// DeleteScore(id int) error
}

func NewService(repo Repository, redisRepo RedisRepository) Service {
	return &service{repo: repo, redisRepo: redisRepo}
}

type service struct {
	repo      Repository
	redisRepo RedisRepository
}

func (s *service) AddScore(userID uint, game string, score float64) error {
	// Add score to PostgreSQL
	if err := s.repo.AddScore(&Score{UserID: userID, Game: game, Score: score, Timestamp: time.Now()}); err != nil {
		return err
	}
	// Add score to Redis
	return s.redisRepo.SetScore(userID, score)
}

func (s *service) GetLeaderboard() ([]Score, error) {
	// Get leaderboard from Redis
	scores, err := s.redisRepo.GetLeaderboard()
	if err != nil {
		return nil, err
	}
	return scores, nil
}

// func (s *service) DeleteScore() error {
// 	// Clear scores from PostgreSQL and Redis
// 	if err := s.repo.DeleteScore(); err != nil {
// 		return err
// 	}
// 	return s.redisRepo.DeleteScore()
// }
