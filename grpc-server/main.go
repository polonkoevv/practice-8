package main

import (
	pb "grpc-server/grpc-server/grpc"
	"grpc-server/service"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()
	profileService := service.NewProfileService()

	// Добавим несколько тестовых профилей
	profileService.AddProfile(1, "John Doe", "john@example.com")
	profileService.AddProfile(2, "Jane Smith", "jane@example.com")

	pb.RegisterProfileServiceServer(server, profileService)

	log.Printf("Server is listening on :50051")
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
