package main 

import (
	"rabbitmq-http-publisher/app/infrastructure/ginlogrus"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type Payload struct {
	ExchangeName string   `json:"exchangeName"`
	Topic string   `json:"topic"`
	Data map[string]interface{} `json:"data"`
}

func main() {
	log.SetFormatter(&log.JSONFormatter{})
	logg := log.New()
	r := gin.New()
	r.Use(ginlogrus.Logger(logg))
	r.Use(gin.Recovery())
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/payload", func(c *gin.Context) {
		var json Payload
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		
		log.Println(json)
		
		c.Status(202)
	})
	r.Run("0.0.0.0:9000") // listen and serve on 0.0.0.0:8080
}