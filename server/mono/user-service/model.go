package user

type User struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Student struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	TeamID *string `json:"team_id,omitempty"`
	UserID string  `json:"user_id"`
}

type Professor struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Profile string `json:"profile"`
	UserID  string `json:"user_id"`
}
