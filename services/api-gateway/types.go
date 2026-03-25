package main

import (
	pb "ride-sharing/shared/proto/trip"
	"ride-sharing/shared/types"
)

type (
	previewTripRequest struct {
		UserID      string           `json:"userID"`
		Pickup      types.Coordinate `json:"pickup"`
		Destination types.Coordinate `json:"destination"`
	}
	startTripRequest struct {
		RideFareID string `json:"rideFareID"`
		UserID     string `json:"userID"`
	}
)

func (p *previewTripRequest) ToProto() *pb.PreviewTripRequest {
	return &pb.PreviewTripRequest{
		UserID: p.UserID,
		StartLocation: &pb.Coordinate{
			Latitude:  p.Pickup.Latitude,
			Longitude: p.Pickup.Longitude,
		},
		EndLocation: &pb.Coordinate{
			Latitude:  p.Destination.Latitude,
			Longitude: p.Destination.Longitude,
		},
	}
}

func (s *startTripRequest) ToProto() *pb.CreateTripRequest {
	return &pb.CreateTripRequest{
		RideFareID: s.RideFareID,
		UserID:     s.UserID,
	}
}
