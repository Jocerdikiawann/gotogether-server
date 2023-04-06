package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type Destination struct {
	Id                string  `json:"_id" bson:"_id"`
	GoogleId          string  `json:"googleId" bson:"googleId"`
	DestinationLatLng Point   `json:"destinationLatLng" bson:"destinationLatLng,omitempty"`
	Polyline          []Point `json:"polyline" bson:"polyline,omitempty"`
}

type Location struct {
	Id       primitive.ObjectID `json:"_id" bson:"_id"`
	GoogleId string             `json:"googleId" bson:"googleId"`
	Location Point              `json:"location" bson:"location,omitempty"`
}

type Point struct {
	Latitude  float64
	Longitude float64
}
