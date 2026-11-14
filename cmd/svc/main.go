// Command svc is a tiny HTTP service used as the taikuri test fixture.
package main

import (
	"fmt"
	"net/http"

	// CROSS-PACKAGE-IMPORT SITE: cmd/svc imports internal/orders, which in turn
	// imports internal/customers. This import chain is the fixture for future
	// map-edge conformance tests (e.g. "frontend must not import a backend-internal
	// package", "service A must not import service B"). Do NOT plant a forbidden
	// edge yet — this is where one would go.
	"github.com/babakDoraniArab/taikuri-testbed/internal/orders"
)

func main() {
	repo := &orders.Repo{}

	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprintln(w, "ok")
	})
	mux.HandleFunc("/orders/", func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Path[len("/orders/"):]
		o := repo.Get(id)
		fmt.Fprintf(w, "order %s for customer %s\n", o.ID, o.Customer.ID)
	})

	addr := ":8081"
	fmt.Println("taikuri-testbed svc listening on", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		fmt.Println("server stopped:", err)
	}
}
// m1 milestone demo 1794657438
