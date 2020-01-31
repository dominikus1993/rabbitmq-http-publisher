package main 

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"rabbitmq-http-publisher/app/infrastructure/ginlogrus"
)


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
	r.Run("0.0.0.0:9000") // listen and serve on 0.0.0.0:8080
}