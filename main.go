package main

import (
	"log"

	"github.com/VolticFroogo/checks-api/db"
	"github.com/VolticFroogo/checks-api/operator"
	"github.com/gin-gonic/gin"
)

func main() {
	err := db.Init()
	if err != nil {
		log.Fatalf("Error initialising DB: %s", err)
	}

	r := gin.Default()

	r.GET("/operator/", operator.All)
	r.GET("/operator/:id", operator.Specific)
	r.POST("/operator/", operator.Insert)

	err = r.Run()
	if err != nil {
		log.Fatalf("Error running Gin router: %s", err)
	}
}
