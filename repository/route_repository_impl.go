package repository

import (
	"context"
	"fmt"

	"github.com/Jocerdikiawann/server_share_trip/model/entity"
	"github.com/Jocerdikiawann/server_share_trip/model/request"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type RouteRepositoryImpl struct {
	db *mongo.Database
}

func NewRouteRepository(db *mongo.Database) *RouteRepositoryImpl {
	return &RouteRepositoryImpl{
		db: db,
	}
}

func (repo *RouteRepositoryImpl) WatchLocation(id string) (stream *mongo.ChangeStream, err error) {
	pipeline := mongo.Pipeline{
		bson.D{
			{
				Key: "$match", Value: bson.D{
					{Key: "operationType", Value: "insert"},
					{Key: "fullDocument.googleid", Value: bson.D{
						{Key: "$eq", Value: id},
					}},
				},
			},
		},
	}

	options := options.ChangeStream().SetFullDocument(options.Default)
	stream, err = repo.db.Collection("location").Watch(context.Background(), pipeline, options)
	return
}

func (repo *RouteRepositoryImpl) GetDestinationAndPolyline(context context.Context, id string) (destination entity.Destination, err error) {
	objId, err := primitive.ObjectIDFromHex(id)
	filter := bson.M{
		"_id": objId,
	}

	err = repo.db.Collection("destination").FindOne(context, filter).Decode(&destination)
	fmt.Println(destination)
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
