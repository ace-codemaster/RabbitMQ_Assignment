package main

import (
	"EasternEnterprise/helpers/logginghelper"
	"EasternEnterprise/model"
	"EasternEnterprise/services"
	"encoding/json"
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {

	//defer dbhelper.HotelDB.Close()

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			transmittedObject := &model.Transmission{}
			//log.Printf("Received a message: %s", d.Body)
			err := json.Unmarshal(d.Body, transmittedObject)
			if err != nil {
				fmt.Println("Error..!!", err)
			} else {
				logginghelper.Info("Received a message: ", transmittedObject)
				status := services.SaveHotelData(transmittedObject)
				if status {
					logginghelper.Info("Data Saved Successfully")
				} else {
					logginghelper.Error("Failed to save data")
				}

			}
		}
	}()

	logginghelper.Info(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
