package user

import (
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Define Dependencies
type ProfessorRepositoryStruct struct {
	db *gorm.DB
}

// Define what this will do
type ProfessorRepository interface {
	// GetProfessor() Professor
	GetProfessors() []Professor
	GetProfessorByID(id string) (*Professor, error)
	CreateProfessor(professor Professor) (*Professor, error)
	UpdateProfessorByID(id string, professor Professor) (*Professor, error)
	DeleteProfessorByID(id string) error
}

func (r *ProfessorRepositoryStruct) GetProfessors() []Professor {
	var professors []Professor
	r.db.Find(&professors)
	return professors
}

func (r *ProfessorRepositoryStruct) GetProfessorByID(id string) (*Professor, error) {
	var professor Professor
	if err := r.db.Where("id = ?", id).First(&professor).Error; err != nil {
		return nil, errors.New("Professor not found")
	}
	return &professor, nil
}

func (r *ProfessorRepositoryStruct) CreateProfessor(professor Professor) (*Professor, error) {
	// Assign a unique ID to the new professor (you may use a UUID generator)
	id := uuid.New()
	professor.ID = id.String() // Replace with your logic

	if err := r.db.Create(&professor).Error; err != nil {
		return nil, err
	}
	return &professor, nil
}

func (r *ProfessorRepositoryStruct) UpdateProfessorByID(id string, professor Professor) (*Professor, error) {
	if err := r.db.Where("id = ?", id).Updates(&professor).Error; err != nil {
		return nil, errors.New("Professor not found")
	}
	var updatedprofessor Professor
	if err := r.db.Where("id = ?", id).First(&updatedprofessor).Error; err != nil {
		return nil, errors.New("Professor not found")
	}
	return &updatedprofessor, nil
}

func (r *ProfessorRepositoryStruct) DeleteProfessorByID(id string) error {

	if err := r.db.Where("id = ?", id).Delete(Professor{}).Error; err != nil {
		return errors.New("Professor not found")
	}
	return nil
}

// Dependency Injection
func ProvideProfessorRepository(db *gorm.DB) *ProfessorRepositoryStruct {
	return &ProfessorRepositoryStruct{db: db}
}
