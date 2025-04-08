package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "grpc-server/proto"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedUserServiceServer
}

func (s *server) SendUserData(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {
	// Imprimir en consola cada solicitud recibida
	log.Printf("Recibido en Servidor gRCP: Nombre=%s, Edad=%d\n", req.Name, req.Age)
	
	// Responder al cliente
	message := fmt.Sprintf("Datos recibidos correctamente: Nombre=%s, Edad=%d", req.Name, req.Age)
	return &pb.UserResponse{Message: message}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, &server{})

	log.Println("Servidor gRPC corriendo en el puerto 50051...")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Error al ejecutar el servidor: %v", err)
	}
}
