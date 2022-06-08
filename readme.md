# Golang Base

### Go version
`golang-base` requires Go version 1.18 or later.


### Kafka
`producer` example
```go
package main

import (
    "github.com/Jungle20m/golang-base/utils"
    "github.com/segmentio/kafka-go"
)

func main() {
    kafkaInstance := utils.NewKafkaInstance("localhost:9092")

    message := kafka.Message{
        Topic: "topic",
        Key:   []byte("hello"),
        Value: []byte("world"),
    }

    kafkaWriter, err := kafkafInstance.NewWriter()
    if err != nil {
        fmt.Println(err)
    }
    defer kafkaWriter.Close()

    kafkaWriter.WriteMessages(context.Background(), message)
}




```