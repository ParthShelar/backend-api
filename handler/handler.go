package handler

import (
	"context"
	"time"

	pb "github.com/backend-team/backend-api/proto/backend-api"
	"github.com/backend-team/backend-api/repository"
	"github.com/prometheus/client_golang/prometheus"
)

type Handler struct {
	Repo *repository.Repository
	pb.UnimplementedBackendApiServiceServer
}

// Prometheus metrics
var (
	requestCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "handler_request_count",
			Help: "Total number of requests processed by each handler",
		},
		[]string{"method", "status"},
	)

	requestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "handler_request_duration_seconds",
			Help:    "Histogram of response times for each handler",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method"},
	)
)

func init() {
	// Register Prometheus metrics
	prometheus.MustRegister(requestCount)
	prometheus.MustRegister(requestDuration)
}

// Instrumented handler methods
func (h *Handler) RegisterEmployee(ctx context.Context, req *pb.RegisterEmployeeRequest) (*pb.RegisterEmployeeResponse, error) {
	startTime := time.Now()
	method := "RegisterEmployee"
	defer func() {
		duration := time.Since(startTime).Seconds()
		requestDuration.WithLabelValues(method).Observe(duration)
	}()

	employee, err := h.Repo.RegisterEmployee(ctx, req.FirstName, req.LastName, req.JoiningDate, req.Salary)
	if err != nil {
		requestCount.WithLabelValues(method, "error").Inc()
		return nil, err
	}

	requestCount.WithLabelValues(method, "success").Inc()
	return &pb.RegisterEmployeeResponse{Employee: &pb.Employee{
		Id:          employee.ID,
		FirstName:   employee.FirstName,
		LastName:    employee.LastName,
		JoiningDate: employee.JoiningDate.String(),
		Salary:      employee.Salary,
	}}, nil
}
