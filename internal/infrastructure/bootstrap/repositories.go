package bootstrap

import (
	eywa "github.com/wmulabs/eywa"
	"github.com/wmulabs/eywa-starter/internal/infrastructure/config"
	eywamongo "github.com/wmulabs/eywa/mongo"
	eywaredis "github.com/wmulabs/eywa/redis"
)

// Repositories groups the persistence ports the engine needs. Echo and Chronicle are kept as their
// concrete types because they implement both the write port (used by the engine) and the read/query
// port (used by the management API).
type Repositories struct {
	Spirit    eywa.SpiritRepository
	Memory    eywa.MemoryRepository
	Echo      *eywamongo.EchoRepository
	Chronicle *eywamongo.ChronicleRepository
	AppToken  eywa.AppTokenRepository
	Bond      eywa.Bond
}

func InitializeRepositories(cfg *config.Config, db *DatabaseConnections) *Repositories {
	mongoDB := db.Mongo.GetDatabase()
	return &Repositories{
		Spirit:    eywamongo.NewSpiritRepository(mongoDB),
		Memory:    eywaredis.NewMemoryRepository(db.Redis.GetClient(), cfg.App.ServiceName, cfg.App.Environment, 3600, nil),
		Echo:      eywamongo.NewEchoRepository(mongoDB),
		Chronicle: eywamongo.NewChronicleRepository(mongoDB),
		AppToken:  eywamongo.NewAppTokenRepository(mongoDB),
		Bond:      eywaredis.NewBondManager(db.Redis.GetClient()),
	}
}
