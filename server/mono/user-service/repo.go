package user

// Define Dependencies
type Repository struct {
	// should have a database connection here
}

// Define what this will do
type UserRepository interface {
	GetUser() User
}

func (r *Repository) GetUser() User {
	return User{
		ID:       "1",
		Email:    "email@example.com",
		Password: "password",
	}
}

// Dependency Injection
func ProvideRepository() *Repository {
	return &Repository{}
}
