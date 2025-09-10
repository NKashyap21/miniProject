package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/NKashyap21/miniProject/helper"
	"github.com/NKashyap21/miniProject/models"
	"github.com/gin-gonic/gin"
)

func HandleAll(ctx *gin.Context) {
	students, err := helper.GetAllStudents()
	if err != nil {
		fmt.Println(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Errorf("%v", err.Error()),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": students,
	})
}

func HandleStudentById(ctx *gin.Context) {
	idStr := ctx.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Unable to convert string to int.",
		})
		return
	}

	student, err := helper.GetStudentById(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": student,
	})
}

func HandleAddStudent(ctx *gin.Context) {
	var student models.Student

	if err := ctx.ShouldBindBodyWithJSON(&student); err != nil {
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	id, err := helper.AddStudent(student)
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": id,
	})
}

func HandleDeleteByID(ctx *gin.Context) {
	idStr := ctx.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = helper.DeleteStudentByID(id)
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(200, gin.H{
		"message": "Student successfully deleted",
	})
}
