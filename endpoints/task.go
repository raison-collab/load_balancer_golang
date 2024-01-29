package endpoints

import (
	"BalancingServers/repository"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"log"
	"net/http"
)

func PostTaskHandler(c *gin.Context) {
	var postTask PostTask

	var errorJsonMessage *ErrorJsonMessage = &ErrorJsonMessage{
		Message:    "error",
		Error:      "unexpected error",
		StatusCode: http.StatusBadRequest,
	}

	var dataJsonMessage *DataJsonMessage = &DataJsonMessage{
		Message:    "OK",
		StatusCode: http.StatusOK,
	}

	ct := c.Request.Header.Get("Content-Type")

	if ct != "application/json" {
		errorJsonMessage.Error = "Your header Content-Type must be application/json"

		c.JSON(http.StatusBadRequest, *errorJsonMessage)
		return
	}

	if er := c.ShouldBindWith(&postTask, binding.JSON); er != nil {
		c.JSON(http.StatusBadRequest, errorJsonMessage)
		return
	}

	dataJsonMessage.Data = postTask

	c.JSON(http.StatusOK, *dataJsonMessage)

	var repo repository.Repository

	res, err := repo.Postgres.DB.Exec("INSERT INTO task(bash, ram, disk, cpu, priority) VALUES ($1, $2, $3, $4, $5)",
		postTask.Bash, postTask.Ram, postTask.Disk, postTask.CPU, postTask.Priority)
	if err != nil {
		log.Fatal(err)
	}

	lastInsertId, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(lastInsertId)

}

func PostTaskGetMethodHandler(c *gin.Context) {
	var errorJsonMessage *ErrorJsonMessage = &ErrorJsonMessage{
		Message:    "error",
		StatusCode: http.StatusBadRequest,
		Error:      "You used GET method. POST needed",
	}
	c.JSON(http.StatusBadRequest, *errorJsonMessage)
}
