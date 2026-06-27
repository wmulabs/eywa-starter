package actions

import (
	"context"
	"fmt"
	"strings"

	eywa "github.com/wmulabs/eywa"
	"github.com/wmulabs/eywa-starter/internal/domain/services"
)

// CatalogLookupAction is a tool that depends on a DOMAIN port (services.CatalogService), not a
// concrete client — the hexagonal flow: Action (application) -> port (domain) -> adapter (driven).
// Copy this shape whenever a tool needs to reach the outside world.
type CatalogLookupAction struct {
	catalog services.CatalogService
}

func NewCatalogLookupAction(catalog services.CatalogService) eywa.Action {
	return &CatalogLookupAction{catalog: catalog}
}

func (a *CatalogLookupAction) GetName() string { return "catalog_lookup" }
func (a *CatalogLookupAction) GetDescription() string {
	return "Search the product catalog by name or category."
}
func (a *CatalogLookupAction) GetParameters() map[string]any {
	return map[string]any{
		"type": "object",
		"properties": map[string]any{
			"query": map[string]any{"type": "string", "description": "Name or category to search for."},
		},
		"required": []string{"query"},
	}
}
func (a *CatalogLookupAction) IsCritical() bool                 { return false }
func (a *CatalogLookupAction) GetCategory() eywa.ActionCategory { return eywa.ActionRetrieval }
func (a *CatalogLookupAction) Validate(map[string]any) error    { return nil }

func (a *CatalogLookupAction) Execute(ctx context.Context, args map[string]any) (string, error) {
	query, _ := args["query"].(string)
	if strings.TrimSpace(query) == "" {
		// Business error: surfaced to the model so it can ask for a query. (Infra errors are hidden.)
		return "", eywa.NewBusinessError("query is required")
	}

	items, err := a.catalog.Find(ctx, query)
	if err != nil {
		return "", eywa.NewInfrastructureError("catalog unavailable", err)
	}
	if len(items) == 0 {
		return fmt.Sprintf("No catalog items match %q.", query), nil
	}

	var b strings.Builder
	for _, it := range items {
		fmt.Fprintf(&b, "- %s (%s) $%.0f\n", it.Name, it.Category, it.Price)
	}
	return b.String(), nil
}
