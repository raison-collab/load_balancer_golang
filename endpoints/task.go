package endpoints

import (
	"BalancingServers/config"
	"BalancingServers/repository/postgres"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
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

	var identifyData *IdentifyTaskData = &IdentifyTaskData{}

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

	var pg postgres.PG
	var cfg config.Config

	cfg.SetupConfig()
	pg.New()

	var lastInsertId uint

	_ = pg.DB.QueryRow("INSERT INTO tasks_table (bash, ram, disk, cpu, priority) VALUES ($1, $2, $3, $4, $5) RETURNING id",
		postTask.Bash, postTask.Ram, postTask.Disk, postTask.CPU, postTask.Priority).Scan(&lastInsertId)

	identifyData.ID = lastInsertId

	dataJsonMessage.IdentifyData = *identifyData

	c.JSON(http.StatusOK, *dataJsonMessage)

}

func PostTaskGetMethodHandler(c *gin.Context) {
	var errorJsonMessage *ErrorJsonMessage = &ErrorJsonMessage{
		Message:    "error",
		StatusCode: http.StatusBadRequest,
		Error:      "You used GET method. POST needed",
	}
	c.JSON(http.StatusBadRequest, *errorJsonMessage)
}
