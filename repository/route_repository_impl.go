package repository

import (
	"context"

	"github.com/Jocerdikiawann/server_share_trip/model/entity"
	"github.com/Jocerdikiawann/server_share_trip/model/request"
	"github.com/Jocerdikiawann/server_share_trip/repository/design"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type RouteRepositoryImpl struct {
	db *mongo.Database
}

func NewRouteRepository(db *mongo.Database) design.RouteRepository {
	return &RouteRepositoryImpl{
		db: db,
	}
}

func (repo *RouteRepositoryImpl) GetLocation(context context.Context, id string) (location entity.Location, err error) {
	filter := bson.M{
		"_id": id,
	}

	result := repo.db.Collection("route").FindOne(context, filter)

	err = result.Decode(&location)
	return
}

func (repo *RouteRepositoryImpl) GetDestinationAndPolyline(context context.Context, id string) (entity.Destination, error) {
	return entity.Destination{}, nil
}

func (repo *RouteRepositoryImpl) SendLocation(context context.Context, request request.LocationRequest) error {
	return nil
}

func (repo *RouteRepositoryImpl) SendDestinationAndPolyline(context context.Context, request request.DestinationAndPolylineRequest) error {
	return nil
}
