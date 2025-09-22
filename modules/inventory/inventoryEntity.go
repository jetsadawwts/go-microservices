package inventory

import "go.mongodb.org/mongo-driver/bson/primitive"

type (
	Inventory struct {
		Id        primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
		UserId    string `json:"user_id" bson:"user_id"`
		ProductId string `json:"product_id" bson:"product_id"`
	}
)
