package teamjoinrequest

import (
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TeamJoinRequestRepository interface {
	GetJoinRequests() ([]TeamJoinRequest, error)
	GetJoinRequestByID(id string) (*TeamJoinRequest, error)
	CreateJoinRequest(request TeamJoinRequest) (*TeamJoinRequest, error)
	UpdateJoinRequestByID(id string, request TeamJoinRequest) (*TeamJoinRequest, error)
	DeleteJoinRequestByID(id string) error
}

type Repository struct {
	db *gorm.DB
}

func (r *Repository) GetJoinRequests() ([]TeamJoinRequest, error) {
	var requests []TeamJoinRequest
	if err := r.db.Table("team_join_requests").Find(&requests).Error; err != nil {
		return nil, err
	}
	return requests, nil
}

func (r *Repository) GetJoinRequestByID(id string) (*TeamJoinRequest, error) {
	var request TeamJoinRequest
	if err := r.db.Table("team_join_requests").Where("id = ?", id).First(&request).Error; err != nil {
		return nil, errors.New("Join request not found.")
	}
	return &request, nil
}

func (r *Repository) CreateJoinRequest(request TeamJoinRequest) (*TeamJoinRequest, error) {
	id := uuid.New()
	request.ID = id.String()

	if err := r.db.Table("team_join_requests").Create(&request).Error; err != nil {
		return nil, err
	}
	return &request, nil
}

func (r *Repository) UpdateJoinRequestByID(id string, updatedRequest TeamJoinRequest) (*TeamJoinRequest, error) {
	if err := r.db.Table("team_join_requests").Where("id = ?", id).Updates(&updatedRequest).Error; err != nil {
		return nil, errors.New("Could not update the join request.")
	}
	return &updatedRequest, nil
}

func (r *Repository) DeleteJoinRequestByID(id string) error {
	if err := r.db.Table("team_join_requests").Where("id = ?", id).Delete(&TeamJoinRequest{}).Error; err != nil {
		return errors.New("Could not delete the join request.")
	}
	return nil
}

func ProvideRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}
