package main

import (
	"context"
	pb "github.com/erikperttu/shippy-consignment-service/proto/consignment"
	vesselProto "github.com/erikperttu/shippy-vessel-service/proto/vessel"
	"gopkg.in/mgo.v2"
	"log"
)

// Implement all methods from the protobuf def
type service struct {
	session      *mgo.Session
	vesselClient vesselProto.VesselServiceClient
}

func (s *service) GetRepo() Repository {
	return &ConsignmentRepository{s.session.Clone()}
}

// From consignment.pb.go
//CreateConsignment(context.Context, *Consignment, *Response) error
func (s *service) CreateConsignment(ctx context.Context, req *pb.Consignment, res *pb.Response) error {
	repo := s.GetRepo()
	defer repo.Close()

	// Call the client instance of the vessel service with the weight and capacity
	vesselResponse, err := s.vesselClient.FindAvailable(context.Background(), &vesselProto.Specification{
		MaxWeight: req.Weight,
		Capacity:  int32(len(req.Containers)),
	})
	if err != nil {
		return err
	}
	log.Printf("Found vessel: %s \n", vesselResponse.Vessel.Name)

	// Set the vessel id from the response
	req.VesselId = vesselResponse.Vessel.Id

	// Save the consignment
	err = repo.Create(req)
	if err != nil {
		return err
	}
	// Set response values
	res.Created = true
	res.Consignment = req
	return nil
}

// From consignment.pb.go
// GetConsignments(context.Context, *GetRequest, *Response) error
func (s *service) GetConsignments(ctx context.Context, req *pb.GetRequest, res *pb.Response) error {
	repo := s.GetRepo()
	defer repo.Close()

	consignments, err := repo.GetAll()
	if err != nil {
		return err
	}
	res.Consignments = consignments
	return nil
}
