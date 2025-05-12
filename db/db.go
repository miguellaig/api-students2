package db

import (
	"github.com/miguellaig/api-students2/schemas"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/sqlite" // Sqlite driver based on CGO

	// "github.com/glebarez/sqlite" // Pure go SQLite driver, checkout https://github.com/glebarez/sqlite for details
	"gorm.io/gorm"
)

type StudentHandler struct {
	DB *gorm.DB
}

func Init() *gorm.DB {
	// github.com/mattn/go-sqlite3
	db, err := gorm.Open(sqlite.Open("student.db"), &gorm.Config{})
	if err != nil {
		log.Fatal().Err(err).Msgf("Failed to initialize SQlite: %s", err.Error())
	}

	db.AutoMigrate(&schemas.Student{})

	return db
}

func NewStudentHandler(db *gorm.DB) *StudentHandler {
	return &StudentHandler{DB: db}
}

func (s *StudentHandler) AddStudent(student schemas.Student) error {

	if result := s.DB.Create(&student); result.Error != nil {
		log.Error().Msg("Failed to create student!")
		return result.Error
	}

	log.Info().Msgf("Create student!")
	return nil
}

func (s *StudentHandler) GetStudents() ([]schemas.Student, error) {
	students := []schemas.Student{}

	err := s.DB.Find(&students).Error
	return students, err
}

func (s *StudentHandler) GetFilteredStudent(active bool) ([]schemas.Student, error) {
	filteredStudents := []schemas.Student{}
	err := s.DB.Where("active = ?", active).Find(&filteredStudents)

	return filteredStudents, err.Error
}

func (s *StudentHandler) GetStudent(id int) (schemas.Student, error) {
	var student schemas.Student

	err := s.DB.First(&student, id)

	// err := s.DB.Find(&students).Error
	return student, err.Error
}
func (s *StudentHandler) UpdateStudent(updateStudent schemas.Student) error {
	return s.DB.Save(&updateStudent).Error
}

func (s *StudentHandler) DeleteStudent(student schemas.Student) error {
	return s.DB.Delete(&student).Error
}
