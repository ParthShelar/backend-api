package main

import (
	"context"
	"expvar"
	"log"
	"net"
	"net/http"

	"cloud.google.com/go/spanner"
	"github.com/backend-team/backend-api/handler"
	config "github.com/backend-team/backend-api/config"
	pb "github.com/backend-team/backend-api/proto/backend-api"
	"github.com/backend-team/backend-api/repository"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	ctx := context.Background()
	dbPath := config.LoadConfig(ctx).GetDBPath()

	// Initialize Spanner client
	client, err := spanner.NewClient(ctx, dbPath)
	if err != nil {
		log.Fatalf("Failed to create Spanner client: %v", err)
	}
	defer client.Close()

	// Initialize repository and handler
	repo := &repository.Repository{Client: client}
	h := &handler.Handler{
		Repo: repo,
	}

	// Start gRPC server in a goroutine
	go func() {
		listener, err := net.Listen("tcp", ":1000")
		if err != nil {
			log.Fatalf("Failed to listen: %v", err)
		}

		grpcServer := grpc.NewServer()
		pb.RegisterBackendApiServiceServer(grpcServer, h)
		reflection.Register(grpcServer)

		log.Println("gRPC server is running on port :1000")
		if err := grpcServer.Serve(listener); err != nil {
			log.Fatalf("gRPC server failed: %v", err)
		}
	}()

	// Start HTTP server for Prometheus and expvar metrics
	go func() {
		http.Handle("/metrics", promhttp.Handler())  // Prometheus metrics
		http.Handle("/debug/vars", expvar.Handler()) // expvar metrics

		log.Println("Metrics server is running on port :2112")
		if err := http.ListenAndServe(":2112", nil); err != nil {
			log.Fatalf("Metrics server failed: %v", err)
		}
	}()

	// Block the main goroutine to keep servers running
	select {}
}
