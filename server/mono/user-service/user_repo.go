package user

import (
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Define Dependencies
type Repository struct {
	// should have a database connection here
	// users []User // Simulating a user database
	db *gorm.DB
}

// Define what this will do
type UserRepository interface {
	// GetUser() User
	GetUsers() []User
	GetUserByID(id string) (*User, error)
	CreateUser(user User) (*User, error)
	UpdateUserByID(id string, user User) (*User, error)
	DeleteUserByID(id string) error
	DeleteAll() error
}

func (r *Repository) GetUsers() []User {
	// return r.users

	var users []User
	r.db.Table("users").Find(&users)
	return users
}

func (r *Repository) GetUserByID(id string) (*User, error) {
	// for _, user := range r.users {
	// 	if user.ID == id {
	// 		return &user, nil
	// 	}
	// }
	// return nil, errors.New("User not found")
	var user User
	if err := r.db.Table("users").Where("id = ?", id).First(&user).Error; err != nil {
		return nil, errors.New("User not found")
	}
	return &user, nil
}

func (r *Repository) CreateUser(user User) (*User, error) {
	// Assign a unique ID to the new user (you may use a UUID generator)
	id := uuid.New()
	user.ID = id.String() // Replace with your logic

	if err := r.db.Table("users").Create(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
	// r.users = append(r.users, user)
	// return &user, nil
}

func (r *Repository) UpdateUserByID(id string, user User) (*User, error) {
	user.ID = id
	if err := r.db.Table("users").Where("id = ?", id).Updates(&user).Error; err != nil {
		return nil, errors.New("User not found")
	}
	return &user, nil
	// for i, u := range r.users {
	// 	if u.ID == id {
	// 		user.ID = id
	// 		r.users[i] = user
	// 		return &user, nil
	// 	}
	// }
	// return nil, errors.New("User not found")
}

func (r *Repository) DeleteUserByID(id string) error {

	if err := r.db.Table("users").Where("id = ?", id).Delete(User{}).Error; err != nil {
		return errors.New("User not found")
	}
	return nil
	// for i, user := range r.users {
	// 	if user.ID == id {
	// 		// Remove the user from the slice
	// 		r.users = append(r.users[:i], r.users[i+1:]...)
	// 		return nil
	// 	}
	// }
	// return errors.New("User not found")
}
func (r *Repository) DeleteAll() error {

	if err := r.db.Table("users").Where("id > ''").Delete(&User{}).Error; err != nil {
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
