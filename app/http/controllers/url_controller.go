package controllers

import (
	"goravel/app/models"

	"github.com/goravel/framework/contracts/http"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UrlController struct {
	//Dependent services
}

func NewUrlController() *UrlController {
	return &UrlController{
		//Inject services
	}
}

func (r *UrlController) Create(ctx http.Context) http.Response {
	res := http.Json{
		"Hello": "Goravel",
	}

	url := models.Url{
		Short: "dadjb213",
		Long:  "https://telteu.ro/long-url-something",
	}
	url.Model().Insert(ctx.Context(), &url)
	res["new"] = url

	items, _ := new(models.Url).Model().All(ctx.Context())
	res["items"] = items

	find, _ := new(models.Url).Model().Get(ctx.Context(), bson.D{
		primitive.E{Key: "short", Value: "da"},
	})
	res["where-find"] = find

	deleteErr := new(models.Url).Model().DeleteFirst(ctx.Context(), bson.D{
		primitive.E{Key: "short", Value: "dadjb213"},
	})
	res["delete-where"] = deleteErr

	return ctx.Response().Success().Json(res)
}
