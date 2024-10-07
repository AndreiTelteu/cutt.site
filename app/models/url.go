package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Url struct {
	ID        *primitive.ObjectID `bson:"_id" json:"id"`
	CreatedAt time.Time           `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time           `bson:"updated_at" json:"updated_at"`
	Short     string              `bson:"short" json:"short"`
	Long      string              `bson:"long" json:"long"`
	UserID    *int64              `bson:"user_id" json:"user_id"`
}

func (r Url) CollectionName() string {
	return "urls"
}

func (r Url) Model() *MongoModel[Url] {
	return new(MongoModel[Url])
}
