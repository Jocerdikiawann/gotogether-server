package config

import (
	"context"
	"fmt"
	"time"

	"github.com/Jocerdikiawann/server_share_trip/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func Connect(usernameDb, passwordDb, nameDb, hostDb, portDb string) *mongo.Database {
	uri := fmt.Sprintf("mongodb://%s:%v", hostDb, portDb)
	credential := options.Credential{
		Username: usernameDb,
		Password: passwordDb,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(uri).SetAuth(credential).SetMaxPoolSize(50)
	client, err := mongo.NewClient(clientOptions)
	utils.CheckError(err)

	err = client.Connect(ctx)
	utils.CheckError(err)

	err = client.Ping(ctx, readpref.Primary())
	utils.CheckError(err)

	db := client.Database(nameDb)

	return db
}
