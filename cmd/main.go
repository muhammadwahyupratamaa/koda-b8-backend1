package main

import (
	"log"

	"koda-b8-backend1/internal/di"
	"koda-b8-backend1/internal/lib"

	"github.com/gin-gonic/gin"
)

func main() {

	db, err := lib.Conn()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	container := di.NewContainer(db)

	r := gin.Default()

	r.POST("/register", container.UserHandler().Register)
	r.POST("/login", container.UserHandler().Login)
	r.GET("/users", container.UserHandler().GetUser)

	r.Run(":8080")
}