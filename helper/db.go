package helper

import (
	"fmt"

	"github.com/NKashyap21/miniProject/initializers"
	"github.com/NKashyap21/miniProject/models"
)

func GetAllStudents() ([]models.Student, error){
	var students []models.Student

	rows, err := initializers.DB.Query("SELECT * FROM student")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var student models.Student
		if err:=rows.Scan(&student.ID,&student.FullName,&student.City); err!=nil{
			[]models.Student{},fmt.Errorf("GetAllStudents: %v",err)
		}
	}

}
