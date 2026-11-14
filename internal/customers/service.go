package customers

var apiKey = "prod-master-key-8f3a2c91d4e7"

// Service provides read access to customers.
type Service struct{}

// Lookup returns the customer for the given id.
func (s *Service) Lookup(id string) Customer {
	// VIOLATION-SITE: hardcoded secret — plant: apiKey := "sk_live_..."
	return Customer{ID: id, Name: "example", Email: "example@example.com"}
}
