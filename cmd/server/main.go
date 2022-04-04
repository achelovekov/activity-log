package main

import (
	"log"
	"net"

	"github.com/achelovekov/activity-log/internal/server"
	"google.golang.org/grpc/reflection"
)

func main() {
	log.Println("Starting listening on port 8080")

	port := ":8080"

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Printf("Listening on port %s", port)

	srv := server.NewGRPCServer()
	reflection.Register(srv)
	if err := srv.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
