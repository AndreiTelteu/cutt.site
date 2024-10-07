package routes

import (
	"goravel/app/http/controllers"

	"github.com/goravel/framework/contracts/route"
	"github.com/goravel/framework/facades"
)

func Api() {
	url := controllers.NewUrlController()
	health := controllers.NewHealthController()

	facades.Route().Prefix("/api").Group(func(router route.Router) {
		router.Get("/health", health.Index)

		router.Prefix("/v1").Group(func(api route.Router) {
			api.Get("/create", url.Create)
		})

	})
}
