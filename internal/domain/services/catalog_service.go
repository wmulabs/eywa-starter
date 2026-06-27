// Package services holds the domain service ports the application layer depends on. Adapters that
// implement them live on the driven side (internal/infrastructure/driven). This is the hexagonal
// boundary: an Action depends on this interface, never on a concrete client.
package services

import "context"

// CatalogItem is a domain entity returned by the catalog.
type CatalogItem struct {
	ID       string
	Name     string
	Category string
	Price    float64
}

// CatalogService looks up catalog items. Implement it with a real backend (HTTP API, DB) on the
// driven side; the example ships an in-memory adapter.
type CatalogService interface {
	Find(ctx context.Context, query string) ([]CatalogItem, error)
}
