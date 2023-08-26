package request

type UserRequest struct {
	GoogleId string `json:"googleId" bson:"googleId" validate:"required"`
	Email    string `validate:"required,email"`
	Name     string `validate:"required"`
}
