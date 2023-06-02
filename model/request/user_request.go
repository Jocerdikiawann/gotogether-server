package request

type UserRequest struct {
	GoogleId string `validate:"required"`
	Email    string `validate:"required,email"`
	Name     string `validate:"required"`
}
