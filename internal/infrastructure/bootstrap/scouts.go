package bootstrap

import (
	"log"

	eywa "github.com/wmulabs/eywa"
	"github.com/wmulabs/eywa-starter/internal/application/scouts"
)

// initializeScouts builds the Scout registry. Register context-enrichment Scouts here; a Link opts
// into them by name via WithScouts (see events.go).
func initializeScouts() eywa.ScoutRegistry {
	registry := eywa.NewScoutRegistry()

	if err := registry.Register(scouts.NewBusinessHoursScout()); err != nil {
		log.Fatalf("register business_hours scout: %v", err)
	}

	return registry
}
