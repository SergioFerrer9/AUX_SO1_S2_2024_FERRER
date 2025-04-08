package main

import (
    "context"
    "fmt"
    "github.com/segmentio/kafka-go"
    "os"
    "time"
)

func main() {
    topic := "my-topic"
    broker := "my-cluster-kafka-bootstrap.kafka:9092"

    writer := &kafka.Writer{
        Addr:     kafka.TCP(broker),
        Topic:    topic,
        Balancer: &kafka.LeastBytes{},
    }

    defer writer.Close()

    fmt.Println("Iniciando producer continuo...")

    for {
        now := time.Now().Format("2006-01-02 15:04:05")
        msg := fmt.Sprintf("Mensaje enviado a las %s", now)

        err := writer.WriteMessages(context.Background(),
            kafka.Message{Value: []byte(msg)},
        )

        if err != nil {
            fmt.Fprintf(os.Stderr, "Error al enviar mensaje: %v\n", err)
        } else {
            fmt.Printf("%s\n", msg)
        }

        time.Sleep(5 * time.Second)
    }
}
