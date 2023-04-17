package infra

import (
	"context"
	"fmt"
	"go-clean-arch/config"
	"go.mongodb.org/mongo-driver/event"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

type MongoStorage struct {
	client *mongo.Client
}

func InitMongo(ctx context.Context, cfg *config.MongoConfig) (*mongo.Client, error) {
	client, err := initMongoConnection(ctx, cfg)
	return client, err
}

func initMongoConnection(ctx context.Context, cfg *config.MongoConfig) (*mongo.Client, error) {
	monitor := &event.PoolMonitor{
		Event: HandlePoolMonitor,
	}
	opts := options.Client().ApplyURI(cfg.MgoUri).
		SetMaxPoolSize(cfg.MaxPoolSize).
		SetPoolMonitor(monitor)

	mgoClient, err := mongo.Connect(ctx, opts)
	if err != nil {
		zap.S().Debugf("Mongo connect failed: %+v", err.Error())
		return nil, err
	}

	err = mgoClient.Ping(ctx, nil)
	if err != nil {
		zap.S().Debugf("Mongo ping failed: %+v", err.Error())

		return nil, err
	}
	zap.S().Info("Mongo connect succeeded")

	return mgoClient, nil
}

func HandlePoolMonitor(evt *event.PoolEvent) {
	switch evt.Type {
	case event.PoolClosedEvent:
		fmt.Println("DB connection closed.")
	}
}
