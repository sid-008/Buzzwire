package broker

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type message struct {
	Id   int    `json:"Id"`
	Data string `json:"data"`
}

func handlePublish(c *gin.Context) {
	var data message

	if err := c.BindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, ":(")
		return
	}
	log.Println("This is broker speaking.", data.Data, "was recv.")
	// storeEnqueue(data)
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
