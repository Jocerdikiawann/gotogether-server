package repository

import (
	"context"

	"github.com/Jocerdikiawann/server_share_trip/model/entity"
	"github.com/Jocerdikiawann/server_share_trip/model/request"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthRepositoryImpl struct {
	db *mongo.Database
}

func NewUserRepository(db *mongo.Database) *AuthRepositoryImpl {
	return &AuthRepositoryImpl{
		db: db,
	}
}

func (repo *AuthRepositoryImpl) Authentication(context context.Context, request request.UserRequest) (data entity.Auth, err error) {
	result, err := repo.db.Collection("user").InsertOne(context, request)

	data = entity.Auth{
		Id:       result.InsertedID.(primitive.ObjectID).Hex(),
		GoogleId: request.GoogleId,
		Email:    request.Email,
		Name:     request.Name,
	}
	return
}
