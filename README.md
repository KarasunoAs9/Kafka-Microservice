# Kafka-Microservice

## Description
The project is a set of microservices on Go that interact through Apache Kafka. The producer sends house data to Kafka, and the supervisor processes this data and finds a home with maximum area.

## Requirements
- [Go 1.20+](https://golang.org/doc/go1.20)
- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)

## Installation

1. **Set up the necessary Go Dependencies:**

    In each microservice (producer and conximer) run the command:

    ```bash
    go mod tidy
    ```

2. **Run Kafka and Zookeeper via Docker Compose:**

    Go to the project root and run:

    ```bash
    docker-compose up -d
    ```

## Start

1. **Producer Launch:**

    The producer sends house data to Kafka.

    Go to the directory `producer/cmd', and run:

    ```bash
    go run main.go
    ```

2. **Start of the consumer:**

    The Controller reads data from Kafka and processes it.

    Go to the directory `consumer/cmd', and run:

    ```bash
    go run main.go
    ```

## Use

1. **Add file with house:**

    The producer uses a file called house.json which should contain house data in the following format:

    ```json
    [
    {"street": "Heroes of Ukraine", "area": 2000},
    {"street": "Maplewood Avenue", "area": 1220},
    {"street": "Cedar Hill Lane", "area": 912},
    {"street": "Sunrise Boulevard", "area": 4420}
    ]
    ```

2. **Sending data to Kafka:**

    The producer sends house data to the top of the page.

3. **Data processing:**

    The Consumer processes data from the top of the table, finds the house with the maximum area and exposes him.

## Project Structure

- **/microservice/producer** - The producer’s microservice that sends data to Kafka.
  - «/cmd»- The point of entry in the producer application.
  - «/service»- Producer’s logic.
  - '/entity'' - The data structures used in the producer.
- **/microservice/consumer** - The microservice of a consumer who processes data from Kafka.
  - «/cmd»- The point of entry in the consur application.
  - «/service»- The consumer’s logic.
  - '/entity'' - The data structures used in the conximum.
- **/docker-compose.yml** - File to run Kafka and Zookeeper via Docker Compose.
- **/house.json** - Example file with house data.

## Completion

To stop the microservices, just stop them in the terminal and to stop Kafka, run the command:

```bash
docker-compose down
```
