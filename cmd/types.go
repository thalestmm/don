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

func (r *Ledger) AddEntry(entry Entry) {
	if entry.DateTime.IsZero() {
		entry.DateTime = time.Now()
	}
	r.Entries = append(r.Entries, entry)
}

func (r *Ledger) ListAllEntries() []Entry {
	return r.Entries
}

func (r *Ledger) ListEntriesByResource(resource string) []Entry {
	if !r.resourceExists(resource) {
		return nil
	}

	var entries []Entry
	for _, entry := range r.Entries {
		if entry.Resource == resource {
			entries = append(entries, entry)
		}
	}
	return entries
}

func (r *Ledger) Total() float64 {
	total := 0.0
	for _, entry := range r.Entries {
		if entry.IsPositive {
			total += entry.Amount
		} else {
			total -= entry.Amount
		}
	}
	return total
}

func (r *Ledger) Resources() []string {
	var resources []string
	for _, entry := range r.Entries {
		resources = append(resources, entry.Resource)
	}
	return resources
}

func (r *Ledger) resourceExists(resource string) bool {
	resources := r.Resources()

	return slices.Contains(resources, resource)
}

func (r *Ledger) TotalByResource(resource string) (float64, error) {
	total := 0.0

	// Early return an error if the resource does not exist
	if !r.resourceExists(resource) {
		return total, ErrResourceNotFound
	}

	// Resource exists, query entries registered
	for _, entry := range r.Entries {
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
func (r *Ledger) SetAmountForResource(resource string, amount float64) error {
	total, err := r.TotalByResource(resource)

	// If resource does not exist, create a new entry with the desired amount
	switch {
	case errors.Is(err, ErrResourceNotFound):
		r.AddEntry(Entry{
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

		r.AddEntry(Entry{
			Resource:   resource,
			IsPositive: delta > 0,
			Amount:     delta,
		})
		return nil
	}
}
