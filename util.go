package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func ReadConfig(configFile string) kafka.ConfigMap {

	//* Create a map to store configurations as key value pairs
	m := make(map[string]kafka.ConfigValue)

	file, err := os.Open(configFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to open file: %s", err)
		os.Exit(1)
	}
	defer file.Close()

	//* Read the contents of the properties file, line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		//* Condition is priarily regarding the comment and
		//* empty lines in the properties file
		if !strings.HasPrefix(line, "#") && len(line) != 0 {

			//* There is an '=' between the key and value pairs.
			//* Splitting the contents of the line will give the key and its value
			kv := strings.Split(line, "=")
			parameter := strings.TrimSpace(kv[0])
			value := strings.TrimSpace(kv[1])

			//* The key along with its value is added to the map
			m[parameter] = value
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Failed to read file: %s", err)
		os.Exit(1)
	}

	return m
}
