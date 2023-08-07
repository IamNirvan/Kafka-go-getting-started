
# Golang, getting started with Kafka 

This project contains the codebase for a simple producer and consumer that connects to Confluent Cloud as of now for testing purposes. All files including the properties file that manages the connection properties is also included.

## Technologies
- Go 1.19.12
- Kafka

## Execution

Firstly, build the producer.
```bash
go build -o out/producer util.go producer.go
```

Then build the consumer
```bash
go build -o out/consumer util.go consumer.go
```

Lastly, use the below commands to execute the producer and consumer respectively
```bash
./out/producer getting-started.properties
```
```bash
./out/consumer getting-started.properties 
```

