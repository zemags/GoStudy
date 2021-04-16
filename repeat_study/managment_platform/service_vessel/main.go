package main

import (
	"context"
	"errors"
	"log"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/v2"
	"github.com/zemags/GoStudy/repeat_study/managment_platform/service_vessel/pb"
)

type Repository interface {
	FindAvailable(*pb.Specification) (*pb.Vessel, error)
}

type VesselRepository struct {
	vessels []*pb.Vessel
}

// grpc service handler
type vesselService struct {
	repo Repository
}

// FindAvailable - checks a specification
func (repo *VesselRepository) FindAvailable(spec *pb.Specification) (*pb.Vessel, error) {
	for _, vessel := range repo.vessels {
		if spec.Capacity <= vessel.Capacity && spec.MaxWeight <= vessel.MaxWeight {
			return vessel, nil
		}
	}
	return nil, errors.New("Cant find vessel for specification")
}

func (s *vesselService) FindAvailable(ctx context.Context, req *pb.Specification, res *pb.Response) error {
	// find next availabel vessel
	vessel, err := s.repo.FindAvailable(req)
	if err != nil {
		return err
	}
	res.Vessel = vessel
	return nil
}

func main() {
	vessels := []*pb.Vessel{
		&pb.Vessel{
			Id:        "vessel001",
			Name:      "Boat",
			MaxWeight: 200999,
			Capacity:  500,
		},
	}
	repo := &VesselRepository{vessels}

	service := micro.NewService(micro.Name("service.vessel"))

	service.Init()

	if err := pb.RegisterShippingServiceHandler(service.Server(), &vesselService{repo}); err != nil {
		log.Panic(err)
	}
	if err != service.Run(); err != nil {
		log.Panic(err)
	}
}
