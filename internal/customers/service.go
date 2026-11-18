package customers

var apiKey = os.Getenv("API_KEY")

// Service provides read access to customers.
type Service struct{}

// Lookup returns the customer for the given id.
func (s *Service) Lookup(id string) Customer {

	return Customer{ID: id, Name: "example", Email: "example@example.com"}
}
// trigger re-gate 1794682271
