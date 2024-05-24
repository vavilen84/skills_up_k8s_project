package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc/credentials/insecure"
	"net/http"

	pb "github.com/vavilen84/skills_up_k8s_project/protobuf"
	"google.golang.org/grpc"
)

func handler(w http.ResponseWriter, r *http.Request) {
	conn, err := grpc.NewClient("service2:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		http.Error(w, "Failed to connect to gRPC server", http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	c := pb.NewGreeterClient(conn)
	res, err := c.SayHello(context.Background(), &pb.HelloRequest{Name: "World"})
	if err != nil {
		http.Error(w, "Failed to call gRPC method", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, res.Message)
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", nil)
}
