package controllers

import (
	"context"
	"fmt"
	appfacades "goravel/app/facades"

	"github.com/goravel/framework/contracts/http"
)

type HealthController struct {
	//Dependent services
}

func NewHealthController() *HealthController {
	return &HealthController{
		//Inject services
	}
}

func (r *HealthController) Index(ctx http.Context) http.Response {
	client := appfacades.Mongo()
	// defer appfacades.Mongo().Disconnect(ctx)

	err := client.Ping(ctx, nil)
	if err != nil {
		return ctx.Response().String(http.StatusBadGateway, err.Error())
	} else {
		health := "OK"

		stats := client.Database("admin").RunCommand(
			context.Background(),
			map[string]interface{}{
				"connPoolStats": 1,
			},
		)
		var connPoolStats map[string]interface{}
		err := stats.Decode(&connPoolStats)
		if err != nil {
			return ctx.Response().String(http.StatusBadGateway, err.Error())
		}
		health += "\nNumber of client connections: " + fmt.Sprint(connPoolStats["numClientConnections"])
		health += "\nNumber of scoped connections: " + fmt.Sprint(connPoolStats["numAScopedConnections"])
		health += "\nTotal connections in use: " + fmt.Sprint(connPoolStats["totalInUse"])
		health += "\nTotal available connections: " + fmt.Sprint(connPoolStats["totalAvailable"])
		health += "\nTotal connections created: " + fmt.Sprint(connPoolStats["totalCreated"])

		return ctx.Response().Success().String(health)
	}
}
