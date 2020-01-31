package main 

import (
	"io/ioutil"
	"rabbitmq-http-publisher/app/infrastructure/ginlogrus"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
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

	r.POST("/payload", func(c *gin.Context) {
		body, err := ioutil.ReadAll(c.Request.Body);
		if err != nil {
			log.Error(err)
			c.Error(err);
		}
		log.Println(body)
		c.Status(202)
	})
	r.Run("0.0.0.0:9000") // listen and serve on 0.0.0.0:8080
}