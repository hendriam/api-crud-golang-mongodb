package lib

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database struct {
	*mongo.Client
}

const defaultTimeout = 2 * time.Second

func LoadDatabase() (Database, error) {
	config := LoadConfig()
	logging := LoadLogging()

	clientOptions := options.Client().ApplyURI(config.Database.MongoDB.Dsn)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		logging.Error().Msgf("[Database] error connected => %s", err)
		panic(err.Error())
	}

	ctx, cancel := defaultContext()
	defer cancel()

	err = client.Ping(ctx, nil)
	if err != nil {
		logging.Error().Msgf("[Database] error ping => %s", err)
		panic(err.Error())
	}

	logging.Info().Msgf("[Database] success connected")
	return Database{client}, nil
}

func defaultContext() (context.Context, context.CancelFunc) {
	return (context.WithTimeout(context.Background(), defaultTimeout))
}
