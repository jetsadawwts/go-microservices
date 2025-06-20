package utils

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ConvertToObjectId(id string) primitive.ObjectID {
	onjectId, _ := primitive.ObjectIDFromHex(id)
	return onjectId
}
