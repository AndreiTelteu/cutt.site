package bootstrap

import (
	"github.com/goravel/framework/foundation"

	"goravel/config"
)

func Boot() {
	app := foundation.NewApplication()

	//Bootstrap the application
	app.Boot()

	//Bootstrap the config.
	config.Boot()

	// interval
	// go func() {
	// 	ticker := time.NewTicker(5 * time.Second)
	// 	defer ticker.Stop()
	// 	client := appfacades.Mongo()
	// 	defer client.Disconnect(context.Background())
	// 	for {
	// 		<-ticker.C
	// 		fmt.Println("Tick at", time.Now())
	// 		stats := client.Database("admin").RunCommand(
	// 			context.Background(),
	// 			map[string]interface{}{
	// 				"connPoolStats": 1,
	// 			},
	// 		)
	// 		var connPoolStats map[string]interface{}
	// 		err := stats.Decode(&connPoolStats)
	// 		if err != nil {
	// 			log.Fatal(err)
	// 		}
	// 		fmt.Println(helpers.PrettyJson(map[string]any{
	// 			"Number of client connections": connPoolStats["numClientConnections"],
	// 			"Number of scoped connections": connPoolStats["numAScopedConnections"],
	// 			"Total connections in use":     connPoolStats["totalInUse"],
	// 			"Total available connections":  connPoolStats["totalAvailable"],
	// 			"Total connections created":    connPoolStats["totalCreated"],
	// 		}))
	// 	}
	// }()
}
