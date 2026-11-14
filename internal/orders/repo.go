package orders

// Repo loads orders from storage.
type Repo struct{}

// Get returns the order for the given id.
func (r *Repo) Get(id string) Order {
	// VIOLATION-SITE: raw SQL from input — plant: query := "SELECT * FROM orders WHERE id=" + id
	return Order{ID: id}
}
