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

type Config struct {
	Username string
	Password string
	Host     string
	Port     string
	NameDb   string
}

func Connect(config *Config) *mongo.Database {
	uri := fmt.Sprintf("mongodb://%v:%v", config.Host, config.Port)
	// mongodb: //localhost:27017/?replicaSet=myReplicaSet&directConnection=true
	// credential := options.Credential{
	// 	Username: usernameDb,
	// 	Password: passwordDb,
	// }

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(uri).
		SetMaxPoolSize(50).
		SetReplicaSet("myReplicaSet").
		SetDirect(true)
		// .SetAuth(credential)
	client, err := mongo.NewClient(clientOptions)
	utils.CheckError(err)

	err = client.Connect(ctx)
	utils.CheckError(err)

	err = client.Ping(ctx, readpref.PrimaryPreferred())
	utils.CheckError(err)

	db := client.Database(config.NameDb)

	return db
}
