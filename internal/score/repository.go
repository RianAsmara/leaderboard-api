package score

import "gorm.io/gorm"

type Repository interface {
	GetLeaderboard() ([]Score, error)
	AddScore(score *Score) error
	// DeleteScore(id int) error
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}

type repository struct {
	db *gorm.DB
}

func (r *repository) AddScore(score *Score) error {
	return r.db.Create(score).Error
}

func (r *repository) GetLeaderboard() ([]Score, error) {
	var scores []Score
	err := r.db.Find(&scores).Error
	return scores, err
}
