package publish

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

func ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Message": "Pong",
	})
}

func publish(c *gin.Context) {
	var req message

	if err := c.BindJSON(&req); err != nil {
		log.Println(err)
		return
	}

	log.Println("This is pub node sending:", req.Data, " to topic ", req.Topic)

	reqJson, _ := json.Marshal(req)
	resp := bytes.NewBuffer(reqJson)

	_, err := http.Post("http://localhost:3003/queue", "application/json", resp)

	if err != nil {
		log.Println(err)
	}
	c.JSON(http.StatusOK, string(reqJson))
}

func StartPubNode() {
	r := gin.Default()
	r.GET("/ping", ping)
	r.POST("/publish", publish)

	log.Println("Publish is running on port 3000")

	err := r.Run(":3000")
	if err != nil {
		log.Fatal(err)
	}
}
