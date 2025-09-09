package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var db *sql.DB

type Student struct {
	Id       int64
	FullName string
	City     string
}

func main() {
	godotenv.Load()

	cfg := mysql.NewConfig()
	cfg.User = os.Getenv("DBUSER")
	cfg.Passwd = os.Getenv("DBPASS")
	cfg.Net = "tcp"
	cfg.DBName = "college"

	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	fmt.Println("Connected")
	
	newStudent := Student{
		FullName:"Krishna",
		City:"HYD",
	}
	id,err:=addStudent(newStudent)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("The new student add at id: %v",id)

	student,err:=studentById(id)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(student)

	studentArr,err := studentsByCity("HYD")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(studentArr)
}

func studentsByCity(cityCode string) ([]Student,error) {
	var students []Student

	rows,err := db.Query("SELECT * FROM Student WHERE City = ?",cityCode)
	if err != nil{
		return nil, fmt.Errorf("studentsByCity: %v",err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var student Student
		if err := rows.Scan(&student.Id,&student.FullName,&student.City);err != nil {
			return nil, fmt.Errorf("studentsByCity: %v",err.Error())
		}
		students = append(students, student)
	}

	if err := rows.Err(); err!=nil {
		return nil, fmt.Errorf("studentsByCity: %v",err.Error())
	}

	return students,nil

}

func studentById(id int64) (Student,error) {
	var student Student

	row := db.QueryRow("SELECT * FROM Student WHERE id = ?",id)
	if err:= row.Scan(&student.Id,&student.FullName,&student.City); err != nil {
		if err == sql.ErrNoRows {
			return student, fmt.Errorf("studentByID: %v",err)
		}
		return student, fmt.Errorf("studentByID: %v",err)
	}
	return student, nil
}

func addStudent(student Student) (int64,error) {
	result,err:=db.Exec("INSERT INTO Student (FullName,City) VALUES (?,?)",student.FullName,student.City)
	if err != nil {
		return 0,fmt.Errorf("addStudent: %v",err.Error())
	}
	id,err:=result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("addStudent: %v",err.Error())
	}
	return id,nil
}