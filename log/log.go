package log

import "go.mongodb.org/mongo-driver/bson/primitive"

type Log struct {
	ID        primitive.ObjectID `bson:"_id"`
	Server    primitive.ObjectID `bson:"server"`
	Operator  primitive.ObjectID `bson:"operator"`
	Timestamp int64
	Updated   bool
	Restarted bool
	Comment   string
}
