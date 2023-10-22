package project

import (
	"errors"

	"gorm.io/gorm"
)

// Define Dependencies
type StudentRepositoryStruct struct {
	db *gorm.DB
}

// Define what this will do
type StudentRepository interface {
	GetAllStudentByTeamID(teamID string) ([]Student, error)
}

func (r *StudentRepositoryStruct) GetAllStudentByTeamID(teamID string) ([]Student, error) {
	var students []Student
	if err := r.db.Table("students").Where("team_id = ?", teamID).Find(&students).Error; err != nil {
		return nil, errors.New("Student not found")
	}
	return students, nil
}

// Dependency Injection
func ProvideStudentRepository(db *gorm.DB) *StudentRepositoryStruct {
	return &StudentRepositoryStruct{db: db}
}
