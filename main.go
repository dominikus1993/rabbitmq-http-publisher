package main 

import (
	"encoding/json"
	"rabbitmq-http-publisher/app/application/dto"
	"rabbitmq-http-publisher/app/infrastructure/ginlogrus"
	"rabbitmq-http-publisher/app/infrastructure/rabbitmq"
	"runtime"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

func logReq(ch *amqp.Channel, messagesChannel chan *dto.Payload) {
	for d := range messagesChannel {
		b, err := json.Marshal(d.Data)
		if err != nil {
			log.Error(err);
		} else {
			rabbitmq.PublishMessage(ch, d.ExchangeName, d.Topic, &b)
		}
	}
}

func produceConsumers(q int, ch *amqp.Channel, messagesChannel chan *dto.Payload) {
	for index := 0; index < q; index++ {
		go logReq(ch, messagesChannel)
	}
}


func main() {
	log.SetFormatter(&log.JSONFormatter{})
	ch := make(chan *dto.Payload)

	rabbitMqConnection := rabbitmq.ConnectToRabbitMq(rabbitmq.GetAmpqConnection())
	rabbitMqChannel := rabbitmq.CreateChannel(rabbitMqConnection)
	defer rabbitmq.CloseRabbit(rabbitMqConnection, rabbitMqChannel)

	produceConsumers(runtime.NumCPU(), rabbitMqChannel, ch)
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