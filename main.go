package main 

import (
	"rabbitmq-http-publisher/app/application/dto"
	"rabbitmq-http-publisher/app/infrastructure/ginlogrus"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func logReq(q int, messagesChannel chan *dto.Payload) {
	for d := range messagesChannel {
		log.Errorln(d)
		log.Println(q)
	}
}

func produceConsumers(q int, messagesChannel chan *dto.Payload) {
	for index := 0; index < q; index++ {
		go logReq(index, messagesChannel)
	}
}


func main() {
	ch := make(chan *dto.Payload)
	produceConsumers(10, ch)
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
		var json dto.Payload
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		
		ch <- &json
		
		c.Status(202)
	})
	r.Run("0.0.0.0:9000") // listen and serve on 0.0.0.0:8080
}