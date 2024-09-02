package score

import "time"

// Score represents a score submission
type Score struct {
	ID        uint `gorm:"primaryKey"`
	UserID    uint
	Game      string
	Score     float64
	Timestamp time.Time
}
