// Package catalog is a driven adapter implementing services.CatalogService. Swap this in-memory stub
// for a real HTTP/DB client without touching the Action that consumes it — that is the point of the
// domain port.
package catalog

import (
	"context"
	"strings"

	"github.com/wmulabs/eywa-starter/internal/domain/services"
)

type MemoryCatalog struct {
	items []services.CatalogItem
}

func NewMemoryCatalog() services.CatalogService {
	return &MemoryCatalog{items: []services.CatalogItem{
		{ID: "svc-1", Name: "Helios Analytics", Category: "analytics", Price: 49},
		{ID: "svc-2", Name: "Nimbus Storage", Category: "storage", Price: 19},
		{ID: "svc-3", Name: "Comet Metrics", Category: "analytics", Price: 35},
	}}
}

func (c *MemoryCatalog) Find(_ context.Context, query string) ([]services.CatalogItem, error) {
	terms := strings.Fields(strings.ToLower(query))
	out := make([]services.CatalogItem, 0)
	for _, it := range c.items {
		haystack := strings.ToLower(it.Name + " " + it.Category)
		if len(terms) == 0 || matchesAnyTerm(haystack, terms) {
			out = append(out, it)
		}
	}
	return out, nil
}

func matchesAnyTerm(haystack string, terms []string) bool {
	for _, t := range terms {
		if strings.Contains(haystack, t) {
			return true
		}
	}
	return false
}
