package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type Destination struct {
	Id                primitive.ObjectID `json:"id" bson:"_id"`
	GoogleId          string             `json:"googleId" bson:"googleId"`
	DestinationLatLng Point              `json:"destination" bson:"destination,omitempty"`
	InitialLocation   Point              `json:"initialLocation" bson:"initialLocation,omitempty"`
	EncodedRoute      string             `json:"encodedRoute" bson:"encodedRoute,omitempty"`
}

type Location struct {
	Id       primitive.ObjectID `json:"id" bson:"_id"`
	GoogleId string             `json:"googleId" bson:"googleId"`
	Location Point              `json:"location" bson:"location,omitempty"`
}

type Point struct {
	Latitude  float64
	Longitude float64
}
