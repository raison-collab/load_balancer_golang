package endpoints

import (
	"BalancingServers/token_generator"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"log"
	"net/http"
)

type PostTask struct {
	Bash string `json:"bash"`
	Ram  uint16 `json:"ram"`
}

type IdentifyTaskData struct {
	Token string `json:"token"`
	Hash  string `json:"hash"`
}

type ErrorJsonMessage struct {
	Message    string `json:"message"`
	Error      string `json:"error"`
	StatusCode uint16 `json:"status_code"`
}

type DataJsonMessage struct {
	Message      string           `json:"message"`
	StatusCode   uint16           `json:"status_code"`
	Data         interface{}      `json:"data"`
	IdentifyData IdentifyTaskData `json:"identify_data"`
}

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

	postTaskBytes, err := json.Marshal(postTask)
	if err != nil {
		log.Fatal(err)
	}

	var identifyData *IdentifyTaskData = &IdentifyTaskData{}

	var hasher token_generator.Hash
	hasher.GenerateHash(string(postTaskBytes))

	var token token_generator.Token
	token.GenerateToken(hasher.Hash)

	identifyData.Hash = hasher.Hash
	identifyData.Token = token.Token

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
