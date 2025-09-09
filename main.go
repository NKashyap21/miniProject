package main

import (
	"github.com/NKashyap21/miniProject/initializers"
	"github.com/gin-gonic/gin"
)

type Student struct {
	Id       int64
	FullName string
	City     string
}

func init() {
	initializers.LoadEnvVariables()
	initializers.SetupDatabase()
}

func main() {
	router := gin.Default()

	router.GET("/all",HandleAll)

}
