package bootstrap

import (
	"context"
	"fmt"

	eywa "github.com/wmulabs/eywa"
	"github.com/wmulabs/eywa-starter/internal/infrastructure/config"
	eywagemini "github.com/wmulabs/eywa/providers/gemini"
)

// EngineComponents holds the assembled Weave.
type EngineComponents struct {
	Weave *eywa.Weave
}

// InitializeEngine builds the Weave, seeds the default Spirit(s), and registers Scouts, converters,
// and event routing.
func InitializeEngine(ctx context.Context, cfg *config.Config, repos *Repositories, svc *Services) (*EngineComponents, error) {
	oracle, err := eywagemini.NewOracle(cfg.LLM.GeminiKey)
	if err != nil {
		return nil, fmt.Errorf("gemini oracle: %w", err)
	}

	weave, err := eywa.NewWeaveBuilder(ctx).
		WithRepositories(repos.Spirit, repos.Memory, repos.Echo, repos.Chronicle).
		WithBond(repos.Bond).
		WithActionRegistry(initializeActions(svc)).
		WithScoutRegistry(initializeScouts()).
		AddOracle(oracle).
		WithMaxIterations(cfg.Engine.MaxIterations).
		Build()
	if err != nil {
		return nil, err
	}

	if err := seedSpirits(ctx, repos.Spirit); err != nil {
		return nil, err
	}
	registerInboundConverters(weave)
	registerEvents(weave)

	return &EngineComponents{Weave: weave}, nil
}
