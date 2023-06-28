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

type Db struct {
	Username string
	Password string
	Host     string
	Port     string
	NameDb   string
}

func Connect(dbConfig *Db) *mongo.Database {
	uri := fmt.Sprintf("mongodb://%v:%v", dbConfig.Host, dbConfig.Port)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(uri).
		SetMaxPoolSize(50).
		SetReplicaSet("myReplicaSet").
		SetDirect(true)
	client, err := mongo.NewClient(clientOptions)
	utils.CheckError(err)

	err = client.Connect(ctx)
	utils.CheckError(err)

	err = client.Ping(ctx, readpref.PrimaryPreferred())
	utils.CheckError(err)

	db := client.Database(dbConfig.NameDb)

	return db
}
