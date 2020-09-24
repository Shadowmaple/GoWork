package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"
	pb "google.golang.org/grpc/examples/helloworld/helloworld"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

var conn *grpc.ClientConn

func init() {
	// Set up a connection to the server.
	var err error
	conn, err = grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
}

func main() {
	c := &conn
	for {
		Greet()
		log.Println(&conn, c)
		if c != &conn {
			log.Println("diff")
			break
		}
		time.Sleep(time.Minute)
	}
}

func Greet() {
	c := pb.NewGreeterClient(conn)

	// Contact the server and print out its response.
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	fmt.Println(r)

	// log.Printf("Greeting: %s", r.Message)
	fmt.Println("OK")
}
