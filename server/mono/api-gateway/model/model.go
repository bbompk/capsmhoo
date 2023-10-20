package model

type Team struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Profile string `json:"profile"`
}

type TeamJoinRequest struct {
	ID        string `json:"id"`
	TeamID    string `json:"team_id"`
	StudentID string `json:"student_id"`
}

type Notification struct {
	Title  string `json:"title"`
	Body   string `json:"body"`
	UserID string `json:"user_id"`
	IsRead bool   `json:"is_read"`
}
