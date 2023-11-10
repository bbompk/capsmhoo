package team

import (
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

type TeamRepository interface {
	GetTeams() []Team
	GetTeamByID(id string) (*Team, error)
	GetTeamByUserID(user_id string) (*Team, error)
	CreateTeam(team Team) (*Team, error)
	UpdateTeamByID(id string, team Team) (*Team, error)
	DeleteTeamByID(id string) (*Team, error)
	DeleteAll() error
}

func (r *Repository) GetTeams() []Team {
	var teams []Team
	r.db.Table("teams").Find(&teams)
	return teams
}

func (r *Repository) GetTeamByID(id string) (*Team, error) {
	var team Team
	if err := r.db.Table("teams").Where("id = ?", id).First(&team).Error; err != nil {
		return nil, errors.New("Team not found.")
	}
	return &team, nil
}

func (r *Repository) GetTeamByUserID(user_id string) (*Team, error) {
	var student Student

	if err := r.db.Table("students").Where("user_id = ?", user_id).First(&student).Error; err != nil {
		return nil, errors.New("User not found in student database.")
	}

	var team Team
	if err := r.db.Table("teams").Where("id = ?", student.TeamID).First(&team).Error; err != nil {
		return nil, errors.New("This user is not in a team.")
	}
	return &team, nil
}

func (r *Repository) CreateTeam(team Team) (*Team, error) {
	// Assign a unique ID to the new user (you may use a UUID generator)
	id := uuid.New()
	team.ID = "TEAM-" + id.String() // Replace with your logic

	if err := r.db.Table("teams").Create(&team).Error; err != nil {
		return nil, err
	}
	return &team, nil
}

func (r *Repository) UpdateTeamByID(id string, team Team) (*Team, error) {
	team.ID = id
	if err := r.db.Table("teams").Where("id = ?", id).Updates(&team).Error; err != nil {
		return nil, errors.New("Team not found.")
	}
	return &team, nil
}

func (r *Repository) DeleteTeamByID(id string) (*Team, error) {
	var team Team
	if err := r.db.Table("teams").Where("id = ?", id).First(&team).Error; err != nil {
		return nil, errors.New("Team not found.")
	}
	if err := r.db.Table("teams").Where("id = ?", id).Delete(Team{}).Error; err != nil {
		return nil, errors.New("Team not found.")
	}
	return &team, nil
}
func (r *Repository) DeleteAll() error {

	if err := r.db.Table("teams").Where("id > ''").Delete(&Team{}).Error; err != nil {
		// Handle database error.
		return err
	}
	// User was deleted successfully.
	return nil
}

// Dependency Injection
func ProvideRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}
