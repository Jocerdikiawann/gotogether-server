package entity

type Destination struct {
	Id                string    `json:"_id" bson:"_id"`
	GoogleId          string    `json:"googleId" bson:"googleId"`
	DestinationLatLng float64   `json:"destinationLatLng" bson:"destinationLatLng,omitempty"`
	Polyline          []float64 `json:"polyline" bson:"polyline,omitempty"`
}

type Location struct {
	Id       string  `json:"_id" bson:"_id"`
	GoogleId string  `json:"googleId" bson:"googleId"`
	Location float64 `json:"location" bson:"location,omitempty"`
}
