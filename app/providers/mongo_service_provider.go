package providers

import (
	"context"

	"github.com/goravel/framework/contracts/foundation"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoServiceProvider struct {
}

const BindingMongo = "goravel.mongo"

func (receiver *MongoServiceProvider) Register(app foundation.Application) {

	app.Singleton(BindingMongo, func(app foundation.Application) (any, error) {
		mongoUrl := app.MakeConfig().GetString("database.connections.mongodb.url")
		serverAPI := options.ServerAPI(options.ServerAPIVersion1)
		opts := options.Client().ApplyURI(mongoUrl).SetServerAPIOptions(serverAPI)

		client, err := mongo.Connect(context.Background(), opts)
		if err != nil {
			panic(err)
		}
		return client, nil
	})

}

func (receiver *MongoServiceProvider) Boot(app foundation.Application) {

}
