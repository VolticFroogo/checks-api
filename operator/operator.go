package operator

import (
	"context"
	"fmt"
	"net/http"

	"github.com/VolticFroogo/checks-api/db"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Operator struct {
	ID   primitive.ObjectID `bson:"_id"  json:"id"`
	Name string             `bson:"name" json:"name"`
}

func All(c *gin.Context) {
	ctx := context.Background()

	cursor, err := db.Operator.Find(ctx, bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("finding all operators: %s", err.Error()),
		})
		return
	}

	var operators []Operator
	err = cursor.All(ctx, &operators)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("decoding operators from cursor: %s", err.Error()),
		})
		return
	}

	c.JSON(http.StatusOK, operators)
}

func Specific(c *gin.Context) {
	ctx := context.Background()

	var operator Operator

	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("decoding operator id: %s", err.Error()),
		})
		return
	}

	err = db.Operator.FindOne(ctx, bson.M{
		"_id": id,
	}).Decode(&operator)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.Status(http.StatusNotFound)
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("finding operator: %s", err.Error()),
		})
		return
	}

	c.JSON(http.StatusOK, operator)
}

func Insert(c *gin.Context) {
	ctx := context.Background()

	var req struct {
		Name string `json:"name"`
	}

	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("parsing json body: %s", err.Error()),
		})
		return
	}

	operator := Operator{
		ID:   primitive.NewObjectID(),
		Name: req.Name,
	}

	_, err = db.Operator.InsertOne(ctx, operator)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("inserting operator: %s", err.Error()),
		})
		return
	}

	c.JSON(http.StatusCreated, operator)
}
