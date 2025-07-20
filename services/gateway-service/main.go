package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/royalsrivastavagithub/hello-microservices/pb/hello"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		conn, err := grpc.NewClient("localhost:50051",
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		)
		if err != nil {
			http.Error(w, "failed to connect to gRPC server", http.StatusInternalServerError)
			log.Printf("gRPC dial error: %v", err)
			return
		}
		defer conn.Close()

		client := pb.NewGreeterClient(conn)

		req := &pb.HelloRequest{} // Fill fields if needed
		res, err := client.SayHello(ctx, req)
		if err != nil {
			http.Error(w, "failed to call SayHello", http.StatusInternalServerError)
			log.Printf("SayHello error: %v", err)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(res.Message))
	})

	log.Println("gateway-service is running on port :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("failed to serve HTTP: %v", err)
	}
}
