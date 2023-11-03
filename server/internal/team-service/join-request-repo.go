package team

import (
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TeamJoinRequestRepository interface {
	GetJoinRequests() ([]TeamJoinRequest, error)
	GetJoinRequestByID(id string) (*TeamJoinRequest, error)
	GetJoinRequestByTeamID(teamID string) ([]TeamJoinRequest, error)
	CreateJoinRequest(request TeamJoinRequest) (*TeamJoinRequest, error)
	UpdateJoinRequestByID(id string, request TeamJoinRequest) (*TeamJoinRequest, error)
	DeleteJoinRequestByID(id string) (*TeamJoinRequest, error)
}

type JoinRequestRepository struct {
	db *gorm.DB
}

func (r *JoinRequestRepository) GetJoinRequests() ([]TeamJoinRequest, error) {
	var requests []TeamJoinRequest
	if err := r.db.Table("team_join_requests").Find(&requests).Error; err != nil {
		return nil, err
	}
	return requests, nil
}

func (r *JoinRequestRepository) GetJoinRequestByID(id string) (*TeamJoinRequest, error) {
	var request TeamJoinRequest
	if err := r.db.Table("team_join_requests").Where("id = ?", id).First(&request).Error; err != nil {
		return nil, errors.New("Join request not found.")
	}
	return &request, nil
}

func (r *JoinRequestRepository) GetJoinRequestByTeamID(teamID string) ([]TeamJoinRequest, error) {
	var requests []TeamJoinRequest
	if err := r.db.Table("team_join_requests").Where("team_id = ?", teamID).Find(&requests).Error; err != nil {
		return nil, errors.New("Join request not found.")
	}
	return requests, nil
}

func (r *JoinRequestRepository) CreateJoinRequest(request TeamJoinRequest) (*TeamJoinRequest, error) {
	id := uuid.New()
	request.ID = id.String()

	if err := r.db.Table("team_join_requests").Create(&request).Error; err != nil {
		return nil, err
	}
	return &request, nil
}

func (r *JoinRequestRepository) UpdateJoinRequestByID(id string, updatedRequest TeamJoinRequest) (*TeamJoinRequest, error) {
	if err := r.db.Table("team_join_requests").Where("id = ?", id).Updates(&updatedRequest).Error; err != nil {
		return nil, errors.New("Could not update the join request.")
	}
	return &updatedRequest, nil
}

func (r *JoinRequestRepository) DeleteJoinRequestByID(id string) (*TeamJoinRequest, error) {
	var team TeamJoinRequest
	if err := r.db.Table("team_join_requests").Where("id = ?", id).First(&team).Error; err != nil {
		return nil, errors.New("Join request not found.")
	}
	if errr := r.db.Table("team_join_requests").Where("id = ?", id).Delete(team).Error; errr != nil {
		return nil, errors.New("Could not delete the join request.")
	}
	return &team, nil
}

func ProvideJoinRequestRepository(db *gorm.DB) *JoinRequestRepository {
	return &JoinRequestRepository{db: db}
}
