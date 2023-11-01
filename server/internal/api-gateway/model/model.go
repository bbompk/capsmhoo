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

type User struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Student struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	TeamID string `json:"team_id"`
	UserID string `json:"user_id"`
}

type Professor struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Profile string `json:"profile"`
	UserID  string `json:"user_id"`
}

type Project struct {
	ID          string `json:"id"`
	TeamID      string `json:"team_id"`
	ProfessorID string `json:"professor_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      string `json:"status"`
	Label       string `json:"label"`
}

type ProjectRequest struct {
	ProjectRequestID string `json:"project_request_id"`
	ProjectID        string `json:"project_id"`
	TeamID           string `json:"team_id"`
	Message          string `json:"message"`
	Status           string `json:"status"`
}

type SuccessResponse struct {
	Success bool `json:"success"`
}

////////////////////////////////////// Bad Coding but i dont care now , i just want it to be done TT //////////////////////////////////////

type StudentRequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

type ProfessorRequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Profile  string `json:"profile"`
}

type UserRequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type StudentResponseBody struct {
	Code  string  `json:"code"`
	Data  Student `json:"data"`
	Error string  `json:"error"`
}

type ProfessorResponseBody struct {
	Code  string    `json:"code"`
	Data  Professor `json:"data"`
	Error string    `json:"error"`
}

type UserResponseBody struct {
	Code  string `json:"code"`
	Data  User   `json:"data"`
	Error string `json:"error"`
}

type UserListResponseBody struct {
	Code  string `json:"code"`
	Data  []User `json:"data"`
	Error string `json:"error"`
}

type StudentListResponseBody struct {
	Code  string    `json:"code"`
	Data  []Student `json:"data"`
	Error string    `json:"error"`
}

type ProfessorListResponseBody struct {
	Code  string      `json:"code"`
	Data  []Professor `json:"data"`
	Error string      `json:"error"`
}
