package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UrlType struct {
	ID        primitive.ObjectID `bson:"_id" json:"id"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
	Short     string             `bson:"short" json:"short"`
	Long      string             `bson:"long" json:"long"`
	UserID    *int64             `bson:"user_id" json:"user_id"`
}

type Url struct {
	MongoModel[UrlType]
}

func (r *Url) Init() *Url {
	r.MongoModel.Init("urls")
	return r
}
