package domain

import (
	pb "ride-sharing/shared/proto/trip"

	tripTypes "ride-sharing/services/trip-service/pkg/types"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RideFareModel struct {
	ID                primitive.ObjectID         `bson:"_id,omitempty"`
	UserID            string                     `bson:"userID"`
	PackageSlug       string                     `bson:"packageSlug"` // ex: van, luxury, sedan
	TotalPriceInCents float64                    `bson:"totalPriceInCents"`
	Route             *tripTypes.OsrmApiResponse `bson:"route"`
}

func (rf *RideFareModel) ToProto() *pb.RideFare {
	return &pb.RideFare{
		Id:                rf.ID.Hex(),
		UserID:            rf.UserID,
		PackageSlug:       rf.PackageSlug,
		TotalPriceInCents: rf.TotalPriceInCents,
	}
}

func ToRideFaresProto(fares []*RideFareModel) []*pb.RideFare {
	rideFares := make([]*pb.RideFare, len(fares))

	for i, fare := range fares {
		rideFares[i] = fare.ToProto()
	}

	return rideFares
}
