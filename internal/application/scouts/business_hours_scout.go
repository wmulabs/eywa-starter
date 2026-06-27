// Package scouts holds Scout implementations: they enrich a Pulse with context (into its Knowledge)
// before the Spirit reasons. Scouts run sequentially and fail open.
package scouts

import (
	"context"
	"time"

	eywa "github.com/wmulabs/eywa"
)

// BusinessHoursScout adds whether the request arrived during business hours to the Pulse's Knowledge,
// so the Spirit can adapt its tone/answer. A real Scout might call a CRM, a feature flag, etc.
type BusinessHoursScout struct{}

func NewBusinessHoursScout() eywa.Scout { return &BusinessHoursScout{} }

func (s *BusinessHoursScout) GetName() string { return "business_hours" }

// IsApplicable gates whether this Scout runs for a given Pulse. Return true to always enrich.
func (s *BusinessHoursScout) IsApplicable(_ *eywa.Pulse) bool { return true }

func (s *BusinessHoursScout) Harvest(_ context.Context, event *eywa.Pulse) error {
	hour := time.Now().UTC().Hour()
	event.AddKnowledge("is_business_hours", hour >= 9 && hour < 18)
	return nil
}
