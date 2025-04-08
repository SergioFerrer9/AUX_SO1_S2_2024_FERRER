package main

import (
    "fmt"
    "github.com/segmentio/kafka-go"
    "context"
)

func main() {
    topic := "my-topic"
    broker := "my-cluster-kafka-bootstrap.kafka:9092"

    r := kafka.NewReader(kafka.ReaderConfig{
        Brokers: []string{broker},
        Topic:   topic,
        GroupID: "go-consumer-group",
    })

    for {
        m, err := r.ReadMessage(context.Background())
        if err != nil {
            fmt.Println("Error leyendo mensaje:", err)
            break
        }
        fmt.Printf("Mensaje recibido: %s\n", string(m.Value))
    }

    r.Close()
}
