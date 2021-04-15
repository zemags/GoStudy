package main

import (
	"context"
	"log"

	"github.com/zemags/GoStudy/repeat_study/managment_platform/service_consignment/pb"

	"github.com/micro/go-micro/v2"
)

type repository interface {
	Create(*pb.Consignment) (*pb.Consignment, error)
	GetAll() []*pb.Consignment
}

// Repository simulates the use of a datastore
type Repository struct {
	consignments []*pb.Consignment
}

// Create a new consignment
func (repo *Repository) Create(consignment *pb.Consignment) (*pb.Consignment, error) {
	updated := append(repo.consignments, consignment)
	repo.consignments = updated
	return consignment, nil
}

func (repo *Repository) GetAll() []*pb.Consignment {
	return repo.consignments
}

type consignmentService struct {
	repo repository
}

// CreateConsignment - create method takes a context and a
// request as an argument
func (s *consignmentService) CreateConsignment(ctx context.Context, req *pb.Consignment, res *pb.Response) error {
	// save consignment
	consignment, err := s.repo.Create(req)
	if err != nil {
		return err
	}
	res.Created = true
	res.Consignment = consignment
	return nil
}

func (s *consignmentService) GetConsignments(ctx context.Context, req *pb.GetRequest, res *pb.Response) error {
	consignments := s.repo.GetAll()
	res.Consignments = consignments
	return nil
}

func main() {
	repo := &Repository{}
	service := micro.NewService(micro.Name("service.consignment"))
	// parse command line flags
	service.Init()

	if err := pb.RegisterShippingServiceHandler(service.Server(), &consignmentService{repo}); err != nil {
		log.Panic(err)
	}

	log.Println("Running server")
	if err := service.Run(); err != nil {
		log.Panic(err)
	}
}
