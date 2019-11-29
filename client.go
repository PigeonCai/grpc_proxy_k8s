package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"google.golang.org/grpc"
	pb "google.golang.org/grpc/examples/helloworld/helloworld"
)

const (
	address = "envoy:81"
	// address     = "localhost:8080"
	defaultName = "world"
)

type helloHandler struct{}

func (h *helloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	// create grpc client

	for i := 0; i < 100; i++ {
		conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		defer conn.Close()
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
		fmt.Println("Greeting: %s", r.GetMessage())
	}

	w.Write([]byte("Hello, world!"))
}

func main() {
	// start http server
	http.Handle("/", &helloHandler{})
	fmt.Printf("start")
	http.ListenAndServe(":8080", nil)
}

// func main() {
// 	// Set up a connection to the server.
// 	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
// 	if err != nil {
// 		log.Fatalf("did not connect: %v", err)
// 	}
// 	defer conn.Close()
// 	c := pb.NewGreeterClient(conn)

// 	// Contact the server and print out its response.
// 	name := defaultName
// 	if len(os.Args) > 1 {
// 		name = os.Args[1]
// 	}
// 	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
// 	defer cancel()
// 	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})
// 	if err != nil {
// 		log.Fatalf("could not greet: %v", err)
// 	}
// 	log.Printf("Greeting: %s", r.GetMessage())
// }
