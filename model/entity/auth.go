package entity

type Auth struct {
	Id       string `json:"_id" bson:"_id"`
	GoogleId string `json:"googleId" bson:"googleId"`
	Email    string `json:"email" bson:"email"`
	Name     string `json:"name" bson:"name"`
}
