package request

type LocationRequest struct {
	GoogleId   string `json:"googleId" validate:"required"`
	Point      Point  `json:"point" validate:"required"`
	IsFinish bool   `json:"isFinish" validate:"required"`
}

type Point struct {
	Latitude  float64
	Longitude float64
}
