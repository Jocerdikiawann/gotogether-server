package repository

import (
	"context"

	"github.com/Jocerdikiawann/server_share_trip/model/entity"
	"github.com/Jocerdikiawann/server_share_trip/model/request"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type RouteRepositoryImpl struct {
	db *mongo.Database
}

func NewRouteRepository(db *mongo.Database) *RouteRepositoryImpl {
	return &RouteRepositoryImpl{
		db: db,
	}
}

func (repo *RouteRepositoryImpl) GetDestinationAndPolyline(context context.Context, id string) (entity.Destination, error) {
	var destination entity.Destination
	objId, err := primitive.ObjectIDFromHex(id)
	filter := bson.M{
		"_id": objId,
	}

	if err != nil {
		return destination, err
	}

	errDb := repo.db.Collection("destination").FindOne(context, filter).Decode(&destination)
	return destination, errDb
}

func (repo *RouteRepositoryImpl) SendDestinationAndPolyline(context context.Context, request request.DestinationAndPolylineRequest) (id string, err error) {
	result, err := repo.db.Collection("destination").InsertOne(context, request)
	id = result.InsertedID.(primitive.ObjectID).Hex()
	return
}
