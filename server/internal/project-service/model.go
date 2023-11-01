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
	Status      string         `json:"status"`
	Label       string         `json:"label"`
}

type ProjectRequest struct {
	ProjectRequestID string `json:"project_request_id"`
	ProjectID        string `json:"project_id"`
	TeamID           string `json:"team_id"`
	Message          string `json:"message"`
	Status           string `json:"status"`
}

type Notification struct {
	Title  string `json:"title"`
	Body   string `json:"body"`
	UserID string `json:"user_id"`
	IsRead bool   `json:"is_read"`
}

type Student struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	TeamID *string `json:"team_id,omitempty"`
	UserID string  `json:"user_id"`
}
