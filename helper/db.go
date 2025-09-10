package helper

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/NKashyap21/miniProject/initializers"
	"github.com/NKashyap21/miniProject/models"
)

func GetAllStudents() ([]models.Student, error) {
	var students []models.Student

	rows, err := initializers.DB.Query("SELECT * FROM student;")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var student models.Student
		if err := rows.Scan(&student.ID, &student.FullName, &student.City); err != nil {
			return []models.Student{}, fmt.Errorf("GetAllStudents: %v", err)
		}
		students = append(students, student)
	}

	if err := rows.Err(); err != nil {
		return []models.Student{}, fmt.Errorf("GetAllStudents: %v", err)
	}

	return students, nil
}

func GetStudentById(id int) (models.Student, error) {
	var student models.Student

	row := initializers.DB.QueryRow("SELECT * FROM Student WHERE id=?", id)
	if err := row.Scan(&student.ID, &student.FullName, &student.City); err != nil {
		if err == sql.ErrNoRows {
			return models.Student{}, fmt.Errorf("GetStudentById: %v", err)
		} else {
			return models.Student{}, fmt.Errorf("GetStudentById: %v", err)
		}
	}
	return student, nil
}

func AddStudent(student models.Student) (id int64, err error) {
	fullName := student.FullName
	city := student.City

	result, err := initializers.DB.Exec("INSERT INTO Student (FullName,City) VALUES (?,?)", fullName, city)
	if err != nil {
		return 0, err
	}
	id, err = result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, err
}

func DeleteStudentByID(id int) error {
	result, err := initializers.DB.Exec("DELETE FROM Student WHERE id = ?", id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no student found with id %d", id)
	}

	return nil
}
