package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "grpc-client/proto"

	"google.golang.org/grpc"
)

func main() {
	// Conectar al servidor gRPC
	conn, err := grpc.Dial("grpc-server:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("No se pudo conectar al servidor: %v", err)
	}
	defer conn.Close()

	// Crear un cliente gRPC
	client := pb.NewUserServiceClient(conn)

	// Enviar solicitudes en un bucle infinito
	for i := 1; ; i++ {
		name := fmt.Sprintf("Usuario%d", i)
		age := 20 + (i % 30) // Genera edades entre 20 y 50

		req := &pb.UserRequest{
			Name: name,
			Age:  int32(age),
		}

		resp, err := client.SendUserData(context.Background(), req)
		if err != nil {
			log.Fatalf("Error en la llamada al servicio: %v", err)
		}

		// Imprimir respuesta del servidor
		fmt.Println(resp.GetMessage())

		// Esperar 2 segundos antes de enviar la siguiente solicitud
		time.Sleep(2 * time.Second)
	}
}
