package bootstrap

import (
	"github.com/wmulabs/eywa-starter/internal/domain/services"
	"github.com/wmulabs/eywa-starter/internal/infrastructure/driven/catalog"
)

// Services groups the domain service ports, each backed by a driven adapter. Actions depend on these
// ports (not the adapters), so you can swap implementations without touching the application layer.
type Services struct {
	Catalog services.CatalogService
}

func InitializeServices() *Services {
	return &Services{
		Catalog: catalog.NewMemoryCatalog(),
	}
}
