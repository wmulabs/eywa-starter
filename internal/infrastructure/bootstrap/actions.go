package bootstrap

import (
	"log"

	eywa "github.com/wmulabs/eywa"
	"github.com/wmulabs/eywa-starter/internal/application/actions"
)

// initializeActions builds the Action registry. Register your tools here; inject domain services for
// tools that reach the outside world (see catalog_lookup).
func initializeActions(svc *Services) eywa.ActionRegistry {
	registry := eywa.NewActionRegistry()

	if err := registry.Register(actions.NewTimeAction()); err != nil {
		log.Fatalf("register time action: %v", err)
	}
	if err := registry.Register(actions.NewCatalogLookupAction(svc.Catalog)); err != nil {
		log.Fatalf("register catalog_lookup action: %v", err)
	}

	return registry
}
