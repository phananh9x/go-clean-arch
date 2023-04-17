package repository

import (
	"github.com/go-redis/redis/v8"
	"go-clean-arch/config"
	"go-clean-arch/service/repository/mgo"
	"go.mongodb.org/mongo-driver/mongo"
)

const database = "go-clean-arch"

type IRepo interface {
	NewCustomersRepository() mgo.CustomersRepository
}

type repo struct {
	mgo   *mongo.Client
	cfg   *config.AppConfig
	redis *redis.Client
}

func (r repo) NewCustomersRepository() mgo.CustomersRepository {
	return mgo.NewCustomersRepository(r.mgo.Database(database))
}

func NewRepo(mgo *mongo.Client, cfg *config.AppConfig, redis *redis.Client) IRepo {
	return &repo{
		mgo:   mgo,
		cfg:   cfg,
		redis: redis,
	}
}
