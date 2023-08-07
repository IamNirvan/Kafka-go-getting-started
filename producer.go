package main

import (
	"fmt"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

//44.216.224.162:9093

//Kafka port: 9093
//Zookeeper port: 2181 (kafka tool)

func main() {
	// fmt.Println("Command line arguments")
	// fmt.Println(os.Args)

	//* get the properties file that is specified when executing the program
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <config-file-path>\n",
			os.Args[0])
		os.Exit(1)
	}

	//* Grab the config file from the cmd arguments
	configFile := os.Args[1]

	//* Execute ReadConfig() and store the resulting map
	//* (contains configuration details as key value pairs) as a variable
	conf := ReadConfig(configFile)

	var topic string
	fmt.Printf("\nEnter a topic: ")
	fmt.Scanln(&topic)

	// ----------------------
	// DEMO CODE
	// ----------------------
	// topic := "purchases"
	// ----------------------
	p, err := kafka.NewProducer(&conf)

	if err != nil {
		fmt.Printf("Failed to create producer: %s", err)
		os.Exit(1)
	}

	// Go-routine to handle message delivery reports and
	// possibly other event types (errors, stats, etc)
	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Failed to deliver message: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Produced event to topic %s: key = %-10s value = %s\n",
						*ev.TopicPartition.Topic, string(ev.Key), string(ev.Value))
				}
			}
		}
	}()

	// -------------------------------------------------------------------------------------------------
	// DEMO CODE
	// -------------------------------------------------------------------------------------------------
	// users := [...]string{"eabara", "jsmith", "sgarcia", "jbernard", "htanaka", "awalther"}
	// items := [...]string{"book", "alarm clock", "t-shirts", "gift card", "batteries", "Something"}
	//
	// for n := 0; n < 10; n++ {
	// 	key := users[rand.Intn(len(users))]
	// 	data := items[rand.Intn(len(items))]
	// 	p.Produce(&kafka.Message{
	// 		TopicPartition: kafka.TopicPartition{
	// 			Topic:     &topic,
	// 			Partition: kafka.PartitionAny,
	// 		},
	// 		Key:   []byte(key),
	// 		Value: []byte(data),
	// 	}, nil)
	// }
	// -------------------------------------------------------------------------------------------------

	//* Run an infinite for loop that pusblishes events
	//* to Kafka under a certain topic
	for {
		p.Produce(
			&kafka.Message{
				TopicPartition: kafka.TopicPartition{
					Topic:     &topic,
					Partition: kafka.PartitionAny,
				},
				Key:   []byte("Key"),
				Value: []byte("Value"),
			}, nil,
		)
	}

	//* Wait for all messages to be delivered then close
	p.Flush(15 * 1000)
	p.Close()
}
