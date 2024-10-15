package main

import (
	"context"
	"log"
	"net"

	pb "grpc-server/userpb" // Reemplaza con la ruta correcta a tu paquete generado

	"google.golang.org/grpc"
)

// server es un tipo que implementa el servicio UserService
type server struct {
	pb.UnimplementedUserServiceServer
}

// SendUserInfo maneja la solicitud del cliente
func (s *server) SendUserInfo(ctx context.Context, user *pb.User) (*pb.UserResponse, error) {
	log.Printf("Recibido: %s, Edad: %d", user.Name, user.Age)
	response := &pb.UserResponse{
		Message: "Información recibida con éxito",
	}
	return response, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Error al escuchar: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, &server{})

	log.Println("Servidor gRPC escuchando en :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Error al servir: %v", err)
	}
}
