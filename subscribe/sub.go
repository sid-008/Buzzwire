package subscribe

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type message struct {
	Topic string `json:"topic"`
	Data  string `json:"data"`
}

func ping(c *gin.Context) {
	fmt.Println("Hello from sub node!")
}

func display(c *gin.Context) {
	var data message

	if err := c.BindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	log.Println("This is subscriber speaking!", data.Data, "was recv. on topic ", data.Topic)
	c.JSON(http.StatusOK, "Ack")
}

func StartPubNode() {
	r := gin.Default()
	r.GET("/ping", ping)
	r.POST("/sub", display)

	log.Println("Subscribe is running on port 3001")
	err := r.Run(":3001")
	if err != nil {
		log.Fatal(err)
	}
}
