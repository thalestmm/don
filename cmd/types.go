package cmd

import (
	"errors"
	"slices"
	"time"
)

// Sentinel error definitions
var (
	ErrResourceNotFound = errors.New("resource not found")
)

// Ledger represents the entire portfolio registry.
// It is the main structure for the stored JSON file.
type Ledger struct {
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

func NewLedger(currency string) *Ledger {
	return &Ledger{
		Currency: currency,
		Entries:  []Entry{},
	}
}

func (l *Ledger) AddEntry(entry Entry) {
	if entry.DateTime.IsZero() {
		entry.DateTime = time.Now()
	}
	l.Entries = append(l.Entries, entry)
}

func (l Ledger) ListAllEntries() []Entry {
	return l.Entries
}

func (l Ledger) ListEntriesByResource(resource string) []Entry {
	if !l.resourceExists(resource) {
		return nil
	}

	var entries []Entry
	for _, entry := range l.Entries {
		if entry.Resource == resource {
			entries = append(entries, entry)
		}
	}
	return entries
}

func (l Ledger) Total() float64 {
	total := 0.0
	for _, entry := range l.Entries {
		if entry.IsPositive {
			total += entry.Amount
		} else {
			total -= entry.Amount
		}
	}
	return total
}

func (l Ledger) Resources() []string {
	var resources []string
	for _, entry := range l.Entries {
		resources = append(resources, entry.Resource)
	}
	return resources
}

func (l Ledger) resourceExists(resource string) bool {
	resources := l.Resources()

	return slices.Contains(resources, resource)
}

func (l Ledger) TotalByResource(resource string) (float64, error) {
	total := 0.0

	// Early return an error if the resource does not exist
	if !l.resourceExists(resource) {
		return total, ErrResourceNotFound
	}

	// Resource exists, query entries registered
	for _, entry := range l.Entries {
		if entry.Resource == resource {
			if entry.IsPositive {
				total += entry.Amount
			} else {
				total -= entry.Amount
			}
		}
	}

	return total, nil
}

// SetAmountForResource creates a new entry with the amount necessary to total
// the entire resource up to the desired amount.
func (l *Ledger) SetAmountForResource(resource string, amount float64) error {
	total, err := l.TotalByResource(resource)

	// If resource does not exist, create a new entry with the desired amount
	switch {
	case errors.Is(err, ErrResourceNotFound):
		l.AddEntry(Entry{
			Resource:   resource,
			IsPositive: true,
			Amount:     amount,
		})
		return nil
	case err != nil:
		return err
	default:
		// Resource exists, calculate the delta and create a new entry
		delta := amount - total

		l.AddEntry(Entry{
			Resource:   resource,
			IsPositive: delta > 0,
			Amount:     delta,
		})
		return nil
	}
}
