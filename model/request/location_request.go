package request

type LocationRequest struct {
	GoogleId string `json:"googleId"`
	Point    Point  `json:"point"`
}

type Point struct {
	Latitude  float64
	Longitude float64
}
