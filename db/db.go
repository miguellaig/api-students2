package db

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite" // Sqlite driver based on CGO
	// "github.com/glebarez/sqlite" // Pure go SQLite driver, checkout https://github.com/glebarez/sqlite for details
	"gorm.io/gorm"
)

type StudentHandler struct {
	DB *gorm.DB
}

type Student struct {
	gorm.Model
	Name   string `json:"name"`
	CPF    int    `json:"cpf"`
	Email  string `json:"email"`
	Age    int    `json:"age"`
	Active bool   `json:"registration"`
}

func Init() *gorm.DB {
	// github.com/mattn/go-sqlite3
	db, err := gorm.Open(sqlite.Open("student.db"), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&Student{})

	return db
}

func NewStudentHandler(db *gorm.DB) *StudentHandler {
	return &StudentHandler{DB: db}
}

func (s *StudentHandler) AddStudent(student Student) error {

	if result := s.DB.Create(&student); result.Error != nil {
		return result.Error
	}

	fmt.Println("Create student!")
	return nil
}

func (s *StudentHandler) GetStudents() ([]Student, error) {
	students := []Student{}

	err := s.DB.Find(&students).Error
	return students, err
}
