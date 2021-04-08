package main

import (
	"context"
	"log"
	"net"
	"sync"

	"github.com/zemags/GoStudy/repeat_study/managment_platform/service_consignment/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

type repository interface {
	Create(*pb.Consignment) (*pb.Consignment, error)
	GetAll() []*pb.Consignment
}

// Repository simulates the use of a datastore
type Repository struct {
	mu           sync.RWMutex
	consignments []*pb.Consignment
}

// Create a new consignment
func (repo *Repository) Create(consignment *pb.Consignment) (*pb.Consignment, error) {
	repo.mu.Lock()
	updated := append(repo.consignments, consignment)
	repo.consignments = updated
	repo.mu.Unlock()
	return consignment, nil
}

func (repo *Repository) GetAll() []*pb.Consignment {
	return repo.consignments
}

type service struct {
	repo repository
}

// CreateConsignment - create method takes a context and a
// request as an argument
func (s *service) CreateConsignment(ctx context.Context, req *pb.Consignment) (*pb.Response, error) {
	// save consignment
	consignment, err := s.repo.Create(req)
	if err != nil {
		return nil, err
	}
	return &pb.Response{
		Created: true, Consignment: consignment,
	}, nil
}

func (s *service) GetConsignments(ctx context.Context, req *pb.GetRequest) (*pb.Response, error) {
	consignments := s.repo.GetAll()
	return &pb.Response{Consignments: consignments}, nil
}

func main() {
	repo := &Repository{}
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listed %v", err)
	}
	server := grpc.NewServer()
	pb.RegisterShippingServiceServer(server, &service{repo})

	reflection.Register(server)
	log.Println("Running on port:", port)
	if err := server.Serve(lis); err != nil {
		log.Fatalf("Failed to server listener %v", err)
	}
}
