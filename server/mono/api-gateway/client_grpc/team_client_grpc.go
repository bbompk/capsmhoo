package client_grpc

import (
	"context"

	pb "capsmhoo/gen/proto"
	joinRequestPb "capsmhoo/gen/team-join-request-pb"
	"capsmhoo/mono/api-gateway/model"
)

type TeamClient struct {
	client *pb.TeamServiceClient
}

type TeamJoinRequestClient struct {
	client *joinRequestPb.TeamJoinRequestServiceClient
}
type TeamgRPCClient interface {
	GetAllTeams(ctx context.Context) ([]*model.Team, error)
	GetTeamByID(ctx context.Context, id string) (*model.Team, error)
	CreateTeam(ctx context.Context, team *model.Team) (*model.Team, error)
	UpdateTeamByID(ctx context.Context, id string, team *model.Team) (*model.Team, error)
	DeleteTeamByID(ctx context.Context, id string) (*model.Team, error)
	AddStudentToTeam(ctx context.Context, teamID string, studentID string) error
	RemoveStudentFromTeam(ctx context.Context, teamID string, studentID string) error
}
type TeamJoinRequestgRPCClient interface {
	GetAllJoinRequests(ctx context.Context) ([]*model.TeamJoinRequest, error)
	GetJoinRequestByID(ctx context.Context, id string) (*model.TeamJoinRequest, error)
	CreateJoinRequest(ctx context.Context, request *model.TeamJoinRequest) (*model.TeamJoinRequest, error)
	UpdateJoinRequest(ctx context.Context, id string, request *model.TeamJoinRequest) (*model.TeamJoinRequest, error)
	DeleteJoinRequest(ctx context.Context, id string) (*joinRequestPb.TeamJoinRequest, error)
}

func (t *TeamClient) GetAllTeams(ctx context.Context) ([]*model.Team, error) {
	res, err := (*t.client).GetAllTeams(ctx, &pb.Empty{})
	if err != nil {
		return nil, err
	}
	var teams []*model.Team
	for _, team := range res.Teams {
		teams = append(teams, &model.Team{
			ID:      team.Id,
			Name:    team.Name,
			Profile: team.Profile,
		})
	}
	return teams, nil
}

func (t *TeamClient) GetTeamByID(ctx context.Context, id string) (*model.Team, error) {
	teamRes, err := (*t.client).GetTeamById(ctx, &pb.TeamId{Id: id})
	if err != nil {
		return nil, err
	}
	team := &model.Team{
		ID:      teamRes.Id,
		Name:    teamRes.Name,
		Profile: teamRes.Profile,
	}
	return team, nil
}

func (t *TeamClient) CreateTeam(ctx context.Context, team *model.Team) (*model.Team, error) {
	teamRes, err := (*t.client).CreateTeam(ctx, &pb.Team{Name: team.Name, Profile: team.Profile})
	if err != nil {
		return nil, err
	}
	team = &model.Team{
		ID:      teamRes.Id,
		Name:    teamRes.Name,
		Profile: teamRes.Profile,
	}
	return team, nil
}

func (t *TeamClient) UpdateTeamByID(ctx context.Context, id string, team *model.Team) (*model.Team, error) {
	teamRes, err := (*t.client).UpdateTeam(ctx, &pb.Team{Id: id, Name: team.Name, Profile: team.Profile})
	if err != nil {
		return nil, err
	}
	team = &model.Team{
		ID:      teamRes.Id,
		Name:    teamRes.Name,
		Profile: teamRes.Profile,
	}
	return team, nil
}

func (t *TeamClient) DeleteTeamByID(ctx context.Context, id string) (*model.Team, error) {
	teamRes, err := (*t.client).DeleteTeam(ctx, &pb.TeamId{Id: id})
	if err != nil {
		return nil, err
	}
	team := &model.Team{
		ID:      teamRes.Id,
		Name:    teamRes.Name,
		Profile: teamRes.Profile,
	}
	return team, nil
}

func (t *TeamClient) AddStudentToTeam(ctx context.Context, teamID string, studentID string) error {
	return nil
}

func (t *TeamClient) RemoveStudentFromTeam(ctx context.Context, teamID string, studentID string) error {
	return nil
}
func (t *TeamJoinRequestClient) GetAllJoinRequests(ctx context.Context) ([]*model.TeamJoinRequest, error) {
	res, err := (*t.client).GetAllJoinRequests(ctx, &joinRequestPb.TeamJoinReqeustEmpty{})
	if err != nil {
		return nil, err
	}
	var requests []*model.TeamJoinRequest
	for _, req := range res.JoinRequests {
		requests = append(requests, &model.TeamJoinRequest{
			ID:        req.Id,
			TeamID:    req.TeamId,
			StudentID: req.StudentId,
		})
	}
	return requests, nil
}

func (t *TeamJoinRequestClient) GetJoinRequestByID(ctx context.Context, id string) (*model.TeamJoinRequest, error) {
	res, err := (*t.client).GetJoinRequestById(ctx, &joinRequestPb.TeamJoinRequestId{Id: id})
	if err != nil {
		return nil, err
	}
	return &model.TeamJoinRequest{
		ID:        res.Id,
		TeamID:    res.TeamId,
		StudentID: res.StudentId,
	}, nil
}

func (t *TeamJoinRequestClient) CreateJoinRequest(ctx context.Context, request *model.TeamJoinRequest) (*model.TeamJoinRequest, error) {
	res, err := (*t.client).CreateJoinRequest(ctx, &joinRequestPb.TeamJoinRequest{
		Id:        request.ID,
		TeamId:    request.TeamID,
		StudentId: request.StudentID,
	})
	if err != nil {
		return nil, err
	}
	return &model.TeamJoinRequest{
		ID:        res.Id,
		TeamID:    res.TeamId,
		StudentID: res.StudentId,
	}, nil
}

func (t *TeamJoinRequestClient) UpdateJoinRequest(ctx context.Context, id string, request *model.TeamJoinRequest) (*model.TeamJoinRequest, error) {
	res, err := (*t.client).UpdateJoinRequest(ctx, &joinRequestPb.TeamJoinRequest{
		Id:        id,
		TeamId:    request.TeamID,
		StudentId: request.StudentID,
	})
	if err != nil {
		return nil, err
	}
	return &model.TeamJoinRequest{
		ID:        res.Id,
		TeamID:    res.TeamId,
		StudentID: res.StudentId,
	}, nil
}

func (t *TeamJoinRequestClient) DeleteJoinRequest(ctx context.Context, id string) (*joinRequestPb.TeamJoinRequest, error) {
	return (*t.client).DeleteJoinRequest(ctx, &joinRequestPb.TeamJoinRequestId{Id: id})
}

func ProvideTeamClient(client *pb.TeamServiceClient) *TeamClient {
	return &TeamClient{client: client}
}

func ProvideTeamJoinRequestClient(client *joinRequestPb.TeamJoinRequestServiceClient) *TeamJoinRequestClient {
	return &TeamJoinRequestClient{client: client}
}
