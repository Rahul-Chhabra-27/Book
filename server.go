package main

import (
	"context"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"rahulchhabra.io/go/proto/Book"
)

type Config struct {
	Book.UnimplementedBookServiceServer
}

func (*Config) Create(context.Context, *Book.BookRequest) (*Book.BookResponse, error) {
	return &Book.BookResponse{
		Messsage: "Book is successfully created",
	}, nil
}
func main() {
	// Create a listener on TCP port
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	// Create a gRPC server object
	s := grpc.NewServer()
	// Attach the Greeter service to the server
	Book.RegisterBookServiceServer(s, &Config{})
	// Serve gRPC server
	log.Println("Serving gRPC on 0.0.0.0:8080")
	go func() {
		log.Fatalln(s.Serve(lis))
	}()

	// Create a client connection to the gRPC server we just started
	// This is where the gRPC-Gateway proxies the requests
	conn, err := grpc.DialContext(
		context.Background(),
		"0.0.0.0:8080",
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}

	gwmux := runtime.NewServeMux()
	// Register Greeter
	err = Book.RegisterBookServiceHandler(context.Background(), gwmux, conn)
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}

	gwServer := &http.Server{
		Addr:    ":8090",
		Handler: gwmux,
	}

	log.Println("Serving gRPC-Gateway on http://0.0.0.0:8090")
	log.Fatalln(gwServer.ListenAndServe())
}

// func main() {
// 	fmt.Println("Server Started....")
// 	ctx := context.Background()
// 	ctx, cancel := context.WithCancel(ctx)
// 	defer cancel()

// 	// ** Register gRPC server endpoint
// 	// ** Make sure the gRPC server is running properly and accessible
// 	mux := runtime.NewServeMux()
// 	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
// 	grpcServerAddress := "localhost:50051"
// 	proxyServer := "localhost:[8081]"
// 	err := Book.RegisterBookServiceHandlerFromEndpoint(ctx, mux, grpcServerAddress, opts)

// 	if err != nil {
// 		log.Fatalf("Couldn't start grpc server %s", err)
// 	}
// 	fmt.Printf("Proxy Server is starting on port %s\n", proxyServer)
// 	fmt.Printf("Grpc Server is starting on port %s\n", grpcServerAddress)

// 	// ** Start HTTP server (and proxy calls to gRPC server endpoint)
// 	if err = http.ListenAndServe(":8081", mux); err != nil {
// 		log.Fatalf("Starting a proxy server")
// 	}
// }
