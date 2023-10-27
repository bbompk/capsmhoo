package team

import (
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Define Dependencies
type StudentRepositoryStruct struct {
	db *gorm.DB
}

// Define what this will do
type StudentRepository interface {
	// GetStudent() Student
	GetStudents() []Student
	GetStudentByID(id string) (*Student, error)
	CreateStudent(student Student) (*Student, error)
	UpdateStudentByID(id string, student Student) (*Student, error)
	DeleteStudentByID(id string) error
	UpdateStudentTeam(stid string, tid string) (*Student, error)
}

func (r *StudentRepositoryStruct) GetStudents() []Student {
	var students []Student
	r.db.Table("students").Find(&students)
	return students
}

func (r *StudentRepositoryStruct) GetStudentByID(id string) (*Student, error) {
	var student Student
	if err := r.db.Table("students").Where("id = ?", id).First(&student).Error; err != nil {
		return nil, errors.New("Student not found")
	}
	return &student, nil
}

func (r *StudentRepositoryStruct) CreateStudent(student Student) (*Student, error) {
	// Assign a unique ID to the new student (you may use a UUID generator)
	id := uuid.New()
	student.ID = id.String() // Replace with your logic
	if err := r.db.Table("students").Create(&student).Error; err != nil {
		return nil, err
	}
	return &student, nil
}

func (r *StudentRepositoryStruct) UpdateStudentByID(id string, student Student) (*Student, error) {
	if err := r.db.Table("students").Where("id = ?", id).Updates(&student).Error; err != nil {
		return nil, errors.New("Student cannot be update")
	}
	var updatedstudent Student
	if err := r.db.Table("students").Where("id = ?", id).First(&updatedstudent).Error; err != nil {
		return nil, errors.New("Student not found")
	}
	return &updatedstudent, nil
}

func (r *StudentRepositoryStruct) DeleteStudentByID(id string) error {

	if err := r.db.Table("students").Where("id = ?", id).Delete(Student{}).Error; err != nil {
		return errors.New("Student not found")
	}
	return nil
}

func (r *StudentRepositoryStruct) UpdateStudentTeam(stid string, tid string) (*Student, error) {
	var student Student
	if err := r.db.Table("students").Where("id = ?", stid).First(&student).Error; err != nil {
		return nil, errors.New("Student not found")
	}
	student.TeamID = &tid
	if err := r.db.Table("students").Where("id = ?", stid).Updates(&student).Error; err != nil {
		return nil, errors.New("Student cannot be update")
	}
	var updatedstudent Student
	if err := r.db.Table("students").Where("id = ?", stid).First(&updatedstudent).Error; err != nil {
		return nil, errors.New("Student not found")
	}
	return &updatedstudent, nil
}

// Dependency Injection
func ProvideStudentRepository(db *gorm.DB) *StudentRepositoryStruct {
	return &StudentRepositoryStruct{db: db}
}
