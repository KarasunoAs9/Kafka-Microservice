package handler

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/IBM/sarama"
	entity_producer "github.com/KarasunoAs9/microservice-go/microservice/producer/service/entity"
)

func readHouseForFile(fileName string) ([]entity_producer.House, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	var houses []entity_producer.House
	err = json.Unmarshal(data, &houses)
	if err != nil {
		return nil, err
	}

	return houses, nil
}

func HandlerProducer() {
	houses, err := readHouseForFile("house.json")
	if err != nil {
		fmt.Printf("error with reading data in file: %v", err)
	}

	config := sarama.NewConfig()
	config.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer([]string{"localhost:9092"}, config)
	if err != nil {
		fmt.Printf("error with creating producer: %v", err)
	}

	defer producer.Close()

	for _, house := range houses {
		message, err := json.Marshal(house)
		if err != nil {
			fmt.Printf("error with marhaling house: %v", err)
		}

		msg := &sarama.ProducerMessage{
			Topic: "house_topic",
			Value: sarama.ByteEncoder(message),
		}

		partition, offset, err := producer.SendMessage(msg)
		if err != nil {
			fmt.Printf("error with sending message in producer: %v", err)
		}

		fmt.Printf("Message sucsessufully send in partition(%d) and offset(%d)", partition, offset)
	}

}
