package models

import (
	"fmt"
	appfacades "goravel/app/facades"

	"github.com/goravel/framework/facades"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoModel[T any] struct {
	Client     *mongo.Client
	Database   string
	collection string
}

func (m *MongoModel[T]) Init(collection string) *MongoModel[T] {
	m.Client = appfacades.Mongo()
	m.collection = collection
	m.Database = facades.Config().GetString("database.connections.mongodb.database")
	return m
}

func (m *MongoModel[T]) Collection() *mongo.Collection {
	return m.Client.Database(m.Database).Collection(m.collection)
}

func (m *MongoModel[T]) GetMeasurementKey(userId int64) string {
	return fmt.Sprintf("user-%d", userId)
}
