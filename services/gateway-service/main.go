package main

import (
	"context"
	"log"
	"net/http"

	"google.golang.org/grpc"

	pb "github.com/royalsrivastavagithub/hello-microservices/pb/hello"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()
		conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		defer conn.Close()
		client := pb.NewGreeterClient(conn)

		req := &pb.HelloRequest{}
		res, err := client.SayHello(ctx, req)
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		w.Write([]byte(res.Message))
	})

	log.Println("gateway-service is running on port :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

