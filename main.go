package main

import (
	"BalancingServers/config"
	"BalancingServers/endpoints"
	"BalancingServers/repository"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"log"
)

func main() {

	var conf config.Config
	conf.ReadTomlConfig("config.toml")

	_, err := repository.NewPostgresDB(conf.Database)
	if err != nil {
		log.Fatal(err)
	}

	runWebServer(conf.Host, conf.Port)
}

// Run wev server with routing
func runWebServer(host string, port uint16) {
	router := gin.Default()

	router.POST("/post_task", endpoints.PostTaskHandler).GET("/post_task", endpoints.PostTaskGetMethodHandler)

	// Run server
	er := router.Run(fmt.Sprintf("%s:%d", host, port))
	if er != nil {
		log.Fatal(er)
	}
}
