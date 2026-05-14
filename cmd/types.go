package cmd

import "time"

// Entry represents the essential unit of data in the portfolio.
// Each entry represents a line in the CSV file.
type Entry struct {
	DateTime   time.Time
	Resource   string
	IsPositive bool
	Amount     float64
}
