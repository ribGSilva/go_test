package structs

import "go.mongodb.org/mongo-driver/bson/primitive"

type CreateUserRequest struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type GetUserResponse struct {
	ID   primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name string             `json:"name"`
	Age  int                `json:"age"`
}
