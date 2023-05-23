package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/imnmania/go_crud/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.SyncDatabase()
}

func main() {
	fmt.Println("Welcome to CRUD API with Golang!!!")
	fmt.Println("----------------------------------")

	r := gin.Default()

	r.GET("/home", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Welcome to CRUD API with Golang!!!",
		})
	})

	r.Run()
}
