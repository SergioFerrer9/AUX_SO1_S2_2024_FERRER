package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"time"

	pb "grpc-cliente/userpb"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("grpc-server:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("No se pudo conectar: %v", err)
	}
	defer conn.Close()

	client := pb.NewUserServiceClient(conn)

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt)

	go func() {
		for {

			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			user := &pb.User{Name: "Juan", Age: 30}

			response, err := client.SendUserInfo(ctx, user)
			if err != nil {
				log.Printf("Error al enviar información: %v", err)
			} else {
				log.Printf("Respuesta del servidor: %s", response.Message)
			}

			cancel()

			time.Sleep(3 * time.Second)
		}
	}()

	<-sigs
	log.Println("Interrupción recibida, cerrando cliente...")
}
