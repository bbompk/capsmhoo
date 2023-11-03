package project

import (
	"database/sql"
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

type ProjectRepository interface {
	GetProjects() []Project
	GetProjectByID(id string) (*Project, error)
	// GetProjectByTeamID(id string) (*Project, error)
	// GetProjectByProfessorID(id string) []Project
	CreateProject(project Project) (*Project, error)
	UpdateProjectByID(id string, project Project) (*Project, error)
	DeleteProjectByID(id string) (*Project, error)
	DeleteAll() error
	GetProjectRequests() []ProjectRequest
	GetProjectRequestByID(id string) (*ProjectRequest, error)
	GetProjectRequestsByProjectId(projectId string) ([]ProjectRequest, error)
	CreateProjectRequest(projectRequest ProjectRequest) (*ProjectRequest, error)
	UpdateProjectRequestByID(id string, projectRequest ProjectRequest) (*ProjectRequest, error)
	DeleteProjectRequestByID(id string) (*ProjectRequest, error)
	DeleteProjectRequestByProjectID(id string) (*ProjectRequest, error)
	AcceptProjectRequest(id string) error
	RejectProjectRequest(id string) error
	AddTeamToProject(teamID string, projectID string) error
}

func (r *Repository) GetProjects() []Project {
	var projects []Project
	r.db.Table("projects").Find(&projects)
	return projects
}

func (r *Repository) GetProjectByID(id string) (*Project, error) {
	var project Project
	if err := r.db.Table("projects").Where("project_id = ?", id).First(&project).Error; err != nil {
		return nil, errors.New("Project not found")
	}
	return &project, nil
}

func (r *Repository) CreateProject(project Project) (*Project, error) {
	// Assign a unique ID to the new project (you may use a UUID generator)
	id := uuid.New()
	project.ProjectID = id.String() // Replace with your logic

	project.TeamID = sql.NullString{}

	project.Status = "open"

	if err := r.db.Table("projects").Create(&project).Error; err != nil {
		return nil, err
	}
	return &project, nil
}

func (r *Repository) UpdateProjectByID(id string, project Project) (*Project, error) {
	project.ProjectID = id

	if project.TeamID.String == "" {
		project.TeamID = sql.NullString{}
	} else {
		project.TeamID = sql.NullString{String: project.TeamID.String, Valid: true}
	}

	if err := r.db.Table("projects").Where("project_id = ?", id).Updates(&project).Error; err != nil {
		return nil, errors.New("Project not found")
	}
	return &project, nil
}

func (r *Repository) DeleteProjectByID(id string) (*Project, error) {
	var project Project
	if err := r.db.Table("projects").Where("project_id = ?", id).First(&project).Error; err != nil {
		return nil, errors.New("Project not found")
	}
	if err := r.db.Table("projects").Where("project_id = ?", id).Delete(Project{}).Error; err != nil {
		return nil, errors.New("Project not found")
	}
	return &project, nil
}
func (r *Repository) DeleteAll() error {

	if err := r.db.Table("projects").Where("project_id > ''").Delete(&Project{}).Error; err != nil {
		// Handle database error.
		return err
	}
	// Projects were deleted successfully.
	return nil
}

func (r *Repository) GetProjectRequests() []ProjectRequest {
	var projectRequests []ProjectRequest
	r.db.Table("project_requests").Find(&projectRequests)
	return projectRequests
}

func (r *Repository) GetProjectRequestByID(id string) (*ProjectRequest, error) {
	var projectRequest ProjectRequest
	if err := r.db.Table("project_requests").Where("project_request_id = ?", id).First(&projectRequest).Error; err != nil {
		return nil, errors.New("ProjectRequest not found")
	}
	return &projectRequest, nil
}

func (r *Repository) GetProjectRequestsByProjectId(projectId string) ([]ProjectRequest, error) {
	var projectRequests []ProjectRequest
	if err := r.db.Table("project_requests").Where("project_id = ?", projectId).Find(&projectRequests).Error; err != nil {
		return nil, errors.New("ProjectRequest not found")
	}
	return projectRequests, nil
}

func (r *Repository) CreateProjectRequest(projectRequest ProjectRequest) (*ProjectRequest, error) {
	// Assign a unique ID to the new projectRequest (you may use a UUID generator)
	id := uuid.New()
	projectRequest.ProjectRequestID = id.String() // Replace with your logic

	if err := r.db.Table("project_requests").Create(&projectRequest).Error; err != nil {
		return nil, err
	}
	return &projectRequest, nil
}

func (r *Repository) UpdateProjectRequestByID(id string, projectRequest ProjectRequest) (*ProjectRequest, error) {
	projectRequest.ProjectRequestID = id
	if err := r.db.Table("project_requests").Where("project_request_id = ?", id).Updates(&projectRequest).Error; err != nil {
		return nil, errors.New("ProjectRequest not found")
	}
	return &projectRequest, nil
}

func (r *Repository) DeleteProjectRequestByID(id string) (*ProjectRequest, error) {
	var projectRequest ProjectRequest
	if err := r.db.Table("project_requests").Where("project_request_id = ?", id).First(&projectRequest).Error; err != nil {
		return nil, errors.New("ProjectRequest not found")
	}
	if err := r.db.Table("project_requests").Where("project_request_id = ?", id).Delete(ProjectRequest{}).Error; err != nil {
		return nil, errors.New("ProjectRequest not found")
	}
	return &projectRequest, nil
}

func (r *Repository) DeleteProjectRequestByProjectID(id string) (*ProjectRequest, error) {
	var projectRequest ProjectRequest
	if err := r.db.Table("project_requests").Where("project_id = ?", id).First(&projectRequest).Error; err != nil {
		return nil, errors.New("ProjectRequest not found")
	}
	if err := r.db.Table("project_requests").Where("project_id = ?", id).Delete(ProjectRequest{}).Error; err != nil {
		return nil, errors.New("ProjectRequest not found")
	}
	return &projectRequest, nil
}

func (r *Repository) AcceptProjectRequest(id string) error {
	if err := r.db.Table("project_requests").Where("project_request_id = ?", id).Update("status", "accepted").Error; err != nil {
		return errors.New("ProjectRequest not found")
	}
	return nil
}

func (r *Repository) RejectProjectRequest(id string) error {
	if err := r.db.Table("project_requests").Where("project_request_id = ?", id).Update("status", "rejected").Error; err != nil {
		return errors.New("ProjectRequest not found")
	}
	return nil
}

func (r *Repository) AddTeamToProject(teamID string, projectID string) error {
	if err := r.db.Table("projects").Where("project_id = ?", projectID).Update("team_id", teamID).Error; err != nil {
		return errors.New("Project not found")
	}

	if err := r.db.Table("projects").Where("project_id = ?", projectID).Update("status", "closed").Error; err != nil {
		return errors.New("Project not found")
	}

	return nil
}

// Dependency Injection
func ProvideRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}
