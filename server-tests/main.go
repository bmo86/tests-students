package main

import (
	"log"
	"net"
	"student-proto/database"
	"student-proto/server"
	"student-proto/testpb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	list, err := net.Listen("tcp", ":5070")

	if err != nil {
		log.Fatal(err)
	}

	repo, err := database.NewPortgressRepository("postgres://postgres:postgres@localhost:54321/postgres?sslmode=disable")

	server := server.NewTestServer(repo)

	if err != nil {
		log.Fatal(err)
	}

	t := grpc.NewServer()
	testpb.RegisterTestServiceServer(t, server)

	reflection.Register(t)

	if err := t.Serve(list); err != nil {
		log.Fatal(err)
	}
}
