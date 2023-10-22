package project

import (
	"database/sql"
)

type Project struct {
	ProjectID   string         `json:"project_id"`
	TeamID      sql.NullString `json:"team_id"`
	ProfessorID string         `json:"professor_id"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
}

type ProjectRequest struct {
	ProjectRequestID string `json:"project_request_id"`
	ProjectID        string `json:"project_id"`
	TeamID           string `json:"team_id"`
	Message          string `json:"message"`
	Status           string `json:"status"`
}
