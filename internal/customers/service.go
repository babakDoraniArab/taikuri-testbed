package customers

import "os"

var apiKey = os.Getenv("ACME_API_KEY")

// Service provides read access to customers.
type Service struct{}

// Lookup returns the customer for the given id.
func (s *Service) Lookup(id string) Customer {
	// VIOLATION-SITE: hardcoded secret — plant: apiKey := "sk_live_..."
	return Customer{ID: id, Name: "example", Email: "example@example.com"}
}
