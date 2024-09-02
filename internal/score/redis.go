package score

import (
	"context"
	"strconv"

	"github.com/go-redis/redis/v8"
)

// Redis interface for leaderboard operations
type RedisRepository interface {
	SetScore(userID uint, score float64) error
	GetLeaderboard() ([]Score, error)
	GetUserRank(userID uint) (int, error)
	DeleteScore() error
}

// Redis implementation
type redisRepository struct {
	client *redis.Client
	ctx    context.Context
}

// NewRedisRepository creates a new Redis repository
func NewRedisRepository(client *redis.Client, ctx context.Context) RedisRepository {
	return &redisRepository{client: client, ctx: ctx}
}

func (r *redisRepository) SetScore(userID uint, score float64) error {
	key := "leaderboard"
	_, err := r.client.ZAdd(r.ctx, key, &redis.Z{
		Score:  score,
		Member: strconv.Itoa(int(userID)),
	}).Result()
	return err
}

func (r *redisRepository) GetLeaderboard() ([]Score, error) {
	key := "leaderboard"
	z := &redis.ZRangeBy{
		Min: "-inf",
		Max: "+inf",
	}

	results, err := r.client.ZRangeByScoreWithScores(r.ctx, key, z).Result()
	if err != nil {
		return nil, err
	}

	var scores []Score
	for _, result := range results {
		userID, _ := strconv.Atoi(result.Member.(string))
		scores = append(scores, Score{
			UserID: uint(userID),
			Score:  result.Score,
		})
	}

	return scores, nil
}

func (r *redisRepository) GetUserRank(userID uint) (int, error) {
	key := "leaderboard"
	rank, err := r.client.ZRevRank(r.ctx, key, strconv.Itoa(int(userID))).Result()
	if err != nil {
		return 0, err
	}
	return int(rank) + 1, nil // +1 because rank is 0-based
}

func (r *redisRepository) DeleteScore() error {
	key := "leaderboard"
	_, err := r.client.Del(r.ctx, key).Result()
	return err
}
