package handler

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/IBM/sarama"
	entity_consumer "github.com/KarasunoAs9/microservice-go/microservice/consumer/service/entity"
)

func findMaxArea(partitionConsumer sarama.PartitionConsumer, done chan bool) entity_consumer.House {
	var MaxHouse entity_consumer.House
	var firstRecievedMessage bool

	for {
		select {
		case msg := <-partitionConsumer.Messages():
			if !firstRecievedMessage {
				firstRecievedMessage = true

				go func() {
					time.Sleep(time.Second * 4)
					done <- true
				}()
			}
			var house entity_consumer.House
			err := json.Unmarshal(msg.Value, &house)
			if err != nil {
				fmt.Printf("error with unmarshaling message: %v", err)
				continue
			}
			if house.Area > MaxHouse.Area {
				MaxHouse = house
			}

			fmt.Println("Recieved message: ", string(msg.Value))

		case <-done:
			fmt.Println("Stopping message queue")
			return MaxHouse
		}
	}
}

func HandlerConsumer() {
	config := sarama.NewConfig()
	config.Producer.Return.Errors = true

	consumer, err := sarama.NewConsumer([]string{"localhost:9092"}, config)
	if err != nil {
		fmt.Printf("error with creating consumer: %v", err)
	}

	defer consumer.Close()

	partitionConsumer, err := consumer.ConsumePartition("house_topic", 0, sarama.OffsetNewest)
	if err != nil {
		fmt.Printf("error with consumer: %v", err)
	}
	
	defer partitionConsumer.Close()

	done := make(chan bool)

	maxArea := findMaxArea(partitionConsumer, done)
	fmt.Println("House with max area: ", maxArea)
}
