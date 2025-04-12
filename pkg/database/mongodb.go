// pkg/database/mongodb.go
package database

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoConfig MongoDB配置
type MongoConfig struct {
	URI         string
	Database    string
	MaxPoolSize uint64
	Timeout     time.Duration
}

// NewMongoClient 创建MongoDB客户端
func NewMongoClient(cfg MongoConfig) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), cfg.Timeout)
	defer cancel()

	opts := options.Client().
		ApplyURI(cfg.URI).
		SetMaxPoolSize(cfg.MaxPoolSize)

	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		return nil, err
	}

	return client, client.Ping(ctx, nil)
}
