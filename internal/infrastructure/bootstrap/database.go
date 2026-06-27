package bootstrap

import (
	"context"
	"fmt"

	"github.com/wmulabs/eywa-starter/internal/infrastructure/config"
	eywamongo "github.com/wmulabs/eywa/mongo"
	eywaredis "github.com/wmulabs/eywa/redis"
)

type DatabaseConnections struct {
	Mongo *eywamongo.MongoConnection
	Redis *eywaredis.RedisRepository
}

func InitializeDatabases(ctx context.Context, cfg *config.Config) (*DatabaseConnections, error) {
	mongo, err := eywamongo.NewMongoConnection(ctx, cfg.Database.MongoURL, cfg.Database.MongoDatabase, cfg.App.ServiceName)
	if err != nil {
		return nil, fmt.Errorf("connect mongo: %w", err)
	}

	redis, err := eywaredis.NewRedisConnection(ctx, cfg.Cache.RedisURL, cfg.App.ServiceName)
	if err != nil {
		return nil, fmt.Errorf("connect redis: %w", err)
	}

	return &DatabaseConnections{Mongo: mongo, Redis: redis}, nil
}

func (d *DatabaseConnections) Close(ctx context.Context) {
	if d.Mongo != nil {
		d.Mongo.DisconnectMongoDB(ctx)
	}
	if d.Redis != nil {
		d.Redis.DisconnectRedisDB(ctx)
	}
}
