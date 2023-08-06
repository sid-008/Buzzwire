package broker

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type message struct {
	Topic string `json:"topic"`
	Data  string `json:"data"`
}

func handlePublish(c *gin.Context) {
	var data message

	if err := c.BindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, ":(")
		return
	}
	log.Println("This is broker speaking.", data.Data, "was recv.")
	// storeEnqueue(data)

	reqJson, _ := json.Marshal(data)
	resp := bytes.NewBuffer(reqJson)

	_, err := http.Post("http://localhost:3001/sub", "application/json", resp)

	if err != nil {
		log.Println(err)
	}
	c.JSON(http.StatusOK, "It worked")
}

func ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func StartBroker() {
	b := gin.Default()
	b.GET("/ping", ping)
	b.POST("/queue", handlePublish)

	log.Println("Broker is running on port 3003")

	err := b.Run(":3003")
	if err != nil {
		log.Fatal(err)
	}
}
