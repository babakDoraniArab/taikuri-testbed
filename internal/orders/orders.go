// Package orders holds the order domain model. It imports customers, creating
// the internal import edge that future map-edge tests exercise.
package orders

import "github.com/babakDoraniArab/taikuri-testbed/internal/customers"

// Order ties a customer to a line total (in cents).
type Order struct {
	ID         string
	Customer   customers.Customer
	TotalCents int
}
