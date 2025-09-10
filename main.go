package main

import (
	"github.com/NKashyap21/miniProject/handlers"
	"github.com/NKashyap21/miniProject/initializers"
	"github.com/gin-gonic/gin"
)


func init() {
	initializers.LoadEnvVariables()
	initializers.SetupDatabase()
}

func main() {
	router := gin.Default()

	router.GET("/student/all", handlers.HandleAll)
	router.GET("/student/:id", handlers.HandleStudentById)
	router.POST("/student/add", handlers.HandleAddStudent)
	router.DELETE("/student/delete/:id",handlers.HandleDeleteByID)
	router.Run()

}
