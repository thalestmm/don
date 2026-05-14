package cmd

import "time"

// Registry represents the entire portfolio registry.
// It is the main structure for the stored JSON file.
type Registry struct {
	Currency string  `json:"currency,omitempty"`
	Entries  []Entry `json:"entries"`
}

// Entry represents the essential unit of data in the portfolio.
type Entry struct {
	DateTime   time.Time `json:"datetime"`
	Resource   string    `json:"resource"`
	IsPositive bool      `json:"isPositive"`
	Amount     float64   `json:"amount"`
}
