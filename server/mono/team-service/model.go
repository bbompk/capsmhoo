package team

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

type Student struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	TeamID *string `json:"team_id,omitempty"`
	UserID string  `json:"user_id"`
}
