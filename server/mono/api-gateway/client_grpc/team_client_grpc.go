package client_grpc

import (
	"context"

	pb "capsmhoo/gen/proto"
	"capsmhoo/mono/api-gateway/model"
)

type TeamClient struct {
	client *pb.TeamServiceClient
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

func ProvideTeamClient(client *pb.TeamServiceClient) *TeamClient {
	return &TeamClient{client: client}
}
