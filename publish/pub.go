package publish

import (
	"bytes"
	"encoding/json"
	"log"
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
)

type message struct {
	Id   int    `json:"Id"`
	Data string `json:"data"`
}

func ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Message": "Pong",
	})
}

func publish(c *gin.Context) {
	var req message

	if err := c.BindJSON(&req); err != nil {
		return
	}

	log.Println("This is pub node sending:", req.Data)

	reqJson, _ := json.Marshal(req)
	resp := bytes.NewBuffer(reqJson)

	_, err := http.Post("http://localhost:3003/queue", "application/json", resp)

	if err != nil {
		log.Println(err)
	}
	c.JSON(http.StatusOK, string(reqJson))
}

func test(c *gin.Context) {
	var req message
	req.Id = rand.Int()
	req.Data = c.Param("data")
	reqJson, _ := json.Marshal(req)
	c.JSON(http.StatusOK, string(reqJson))
}

func StartPubNode() {
	r := gin.Default()
	r.GET("/ping", ping)
	r.POST("/publish", publish)
	r.GET("/publish/:data", test)

	log.Println("Publish is running on port 3000")

	err := r.Run(":3000")
	if err != nil {
		log.Fatal(err)
	}
}
