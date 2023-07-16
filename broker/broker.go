package broker

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

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
	log.Println(data.Data)
	c.JSON(http.StatusOK, "It worked")

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
