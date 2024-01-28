package main

import (
	"BalancingServers/endpoints"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

const port uint16 = 8081

func main() {
	runWebServer(port)
	//runGinWebServer(port)
}

func runWebServer(port uint16) {
	router := gin.Default()

	router.POST("/post_task", endpoints.PostTaskHandler).GET("/post_task", endpoints.PostTaskGetMethodHandler)

	// Run server
	er := router.Run(fmt.Sprintf(":%d", port))
	if er != nil {
		log.Fatal(er)
	}
}
