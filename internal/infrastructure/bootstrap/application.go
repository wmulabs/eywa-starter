// Package bootstrap wires the application together, one concern per file: databases, repositories,
// the engine, actions, events. Initialize composes them; Shutdown tears them down.
package bootstrap

import (
	"context"
	"fmt"

	"github.com/wmulabs/eywa-starter/internal/infrastructure/config"
)

// Application is the dependency container assembled at startup.
type Application struct {
	Config       *config.Config
	Database     *DatabaseConnections
	Repositories *Repositories
	Services     *Services
	Engine       *EngineComponents
}

// Initialize bootstraps the whole application in dependency order.
func Initialize(ctx context.Context) (*Application, error) {
	cfg := config.Load()

	db, err := InitializeDatabases(ctx, cfg)
	if err != nil {
		return nil, fmt.Errorf("databases: %w", err)
	}

	repos := InitializeRepositories(cfg, db)
	svc := InitializeServices()

	engine, err := InitializeEngine(ctx, cfg, repos, svc)
	if err != nil {
		return nil, fmt.Errorf("engine: %w", err)
	}

	return &Application{
		Config:       cfg,
		Database:     db,
		Repositories: repos,
		Services:     svc,
		Engine:       engine,
	}, nil
}

// Shutdown releases external connections.
func (a *Application) Shutdown(ctx context.Context) {
	if a.Database != nil {
		a.Database.Close(ctx)
	}
}
