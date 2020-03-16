package customer

import "go.mongodb.org/mongo-driver/bson/primitive"

type Customer struct {
	ID   primitive.ObjectID
	Name string
}
