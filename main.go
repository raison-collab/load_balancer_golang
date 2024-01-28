package main

import (
	"BalancingServers/config"
	"BalancingServers/endpoints"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {

	var conf config.Config
	conf.ReadConfig("config.toml")

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
