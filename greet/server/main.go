package main

import (
	pb "github.com/Durotimicodes/grpc-go/greet/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

//global address the server will be listening
var addr string = "0.0.0.0:50051"

//to listen @ port, create a server struct that has access to the Greet server API
type Server struct {
	pb.GreetServiceServer
}

func main() {

	Listen, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to Listen on : %v\n", err)
	}

	log.Printf("Listening on %s\n", addr)

	//create an instance of server
	s := grpc.NewServer()
	if err = s.Serve(Listen); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}

}
