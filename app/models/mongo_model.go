package models

import (
	"context"
	appfacades "goravel/app/facades"
	"reflect"

	"github.com/goravel/framework/facades"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoModel[T ModelTypeInterface] struct {
	Client     *mongo.Client
	Database   string
	collection string
}

type ModelTypeInterface interface {
	CollectionName() string
}

func (m *MongoModel[T]) Init() {
	model := new(T)
	m.Client = appfacades.Mongo()
	m.collection = (*model).CollectionName()
	m.Database = facades.Config().GetString("database.connections.mongodb.database")
}

func (m *MongoModel[T]) Collection() *mongo.Collection {
	return m.Client.Database(m.Database).Collection(m.collection)
}

func (m *MongoModel[T]) Insert(ctx context.Context, data *T) (string, error) {
	if m.Client == nil {
		m.Init()
	}
	reflected := reflect.ValueOf(data).Elem()
	if reflected.Kind() == reflect.Struct {
		newId := primitive.NewObjectID()
		reflected.FieldByName("ID").Set(reflect.ValueOf(&newId))
	}
	result, err := m.Collection().InsertOne(ctx, data)
	if err != nil {
		return "", err
	}
	id := result.InsertedID.(primitive.ObjectID)
	return id.String(), nil
}

func (m *MongoModel[T]) All(ctx context.Context) ([]*T, error) {
	filter := bson.D{{}}
	return m.Get(ctx, filter)

}
func (m *MongoModel[T]) Get(ctx context.Context, filter bson.D) ([]*T, error) {
	if m.Client == nil {
		m.Init()
	}
	var items []*T
	cur, err := m.Collection().Find(ctx, filter)
	if err != nil {
		return items, err
	}
	for cur.Next(ctx) {
		var item T
		err := cur.Decode(&item)
		if err != nil {
			return items, err
		}
		items = append(items, &item)
	}
	if err := cur.Err(); err != nil {
		return items, err
	}
	// once exhausted, close the cursor
	cur.Close(ctx)
	if len(items) == 0 {
		return items, mongo.ErrNoDocuments
	}
	return items, nil
}

func (m *MongoModel[T]) DeleteFirst(ctx context.Context, filter bson.D) error {
	if m.Client == nil {
		m.Init()
	}
	res, err := m.Collection().DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	if res.DeletedCount == 0 {
		return mongo.ErrNoDocuments
	}
	return nil
}

func (m *MongoModel[T]) DeleteMany(ctx context.Context, filter bson.D) error {
	if m.Client == nil {
		m.Init()
	}
	res, err := m.Collection().DeleteMany(ctx, filter)
	if err != nil {
		return err
	}
	if res.DeletedCount == 0 {
		return mongo.ErrNoDocuments
	}
	return nil
}

func (m *MongoModel[T]) GetMeasurementKey(userId int64) string {
	if m.Client == nil {
		m.Init()
	}
	return ""
}
