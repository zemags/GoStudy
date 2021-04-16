package main

import (
	"errors"

	"github.com/zemags/GoStudy/repeat_study/managment_platform/service_vessel/pb"
)

type Repository interface {
	FindAvailable(*pb.Specification) (*pb.Vessel, error)
}

type VesselRepository struct {
	vessels []*pb.Vessel
}

// grpc service handler
type vesselService strcut {
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

func main() {

}
