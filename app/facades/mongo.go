package appfacades

import (
	"log"

	"github.com/goravel/framework/foundation"
	"go.mongodb.org/mongo-driver/mongo"
)

const BindingMongo = "goravel.mongo"

func Mongo() *mongo.Client {
	app := foundation.NewApplication()
	instance, err := app.Make(BindingMongo)
	if err != nil {
		log.Fatalln(err)
		return nil
	}
	return instance.(*mongo.Client)
}
