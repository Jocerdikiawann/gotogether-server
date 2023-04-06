package repository

import (
	"context"

	"github.com/Jocerdikiawann/server_share_trip/model/entity"
	"github.com/Jocerdikiawann/server_share_trip/model/request"
	"github.com/Jocerdikiawann/server_share_trip/repository/design"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type RouteRepositoryImpl struct {
	db *mongo.Database
}

func NewRouteRepository(db *mongo.Database) design.RouteRepository {
	return &RouteRepositoryImpl{
		db: db,
	}
}

func (repo *RouteRepositoryImpl) WatchLocation() (stream *mongo.ChangeStream, err error) {
	pipeline := mongo.Pipeline{}

	options := options.ChangeStream().SetFullDocument(options.UpdateLookup)
	stream, err = repo.db.Collection("location").Watch(context.Background(), pipeline, options)
	return
}

func (repo *RouteRepositoryImpl) GetDestinationAndPolyline(context context.Context, id string) (destination entity.Destination, err error) {
	filter := bson.M{
		"_id": id,
	}

	result := repo.db.Collection("destination").FindOne(context, filter)
	err = result.Decode(&destination)
	return
}

func (repo *RouteRepositoryImpl) SendLocation(context context.Context, request request.LocationRequest) (id string, err error) {
	result, err := repo.db.Collection("location").InsertOne(context, request)

	id = result.InsertedID.(primitive.ObjectID).Hex()
	return
}

func (repo *RouteRepositoryImpl) SendDestinationAndPolyline(context context.Context, request request.DestinationAndPolylineRequest) (id string, err error) {
	result, err := repo.db.Collection("destination").InsertOne(context, request)
	id = result.InsertedID.(primitive.ObjectID).Hex()
	return
}
