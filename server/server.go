package server

import "go.mongodb.org/mongo-driver/bson/primitive"

type Server struct {
	ID       primitive.ObjectID
	Customer primitive.ObjectID
	Name     string
}
