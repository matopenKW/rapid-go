package main

import (
	"fmt"
	"net/http"

	"github.com/abyssparanoia/rapid-go/api/src/lib/log"

	"github.com/go-chi/chi"
)

var addr = ":3001"

func main() {
	// Dependency
	d := &Dependency{}
	d.Inject()

	// Routing
	r := chi.NewRouter()
	Routing(r, d)

	// Run

	fmt.Printf("[START] server. port: %s\n", addr)
	http.ListenAndServe(addr, log.Logger(r))
}
