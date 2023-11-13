package client_grpc

import (
	"context"

	pb "capsmhoo/gen/projectpb"
	"capsmhoo/internal/api-gateway/model"
)

type ProjectClient struct {
	client *pb.ProjectServiceClient
}

type ProjectgRPCClient interface {
	GetAllProjects(ctx context.Context) ([]*model.Project, error)
	GetProjectByID(ctx context.Context, id string) (*model.Project, error)
	GetProjectByTeamID(ctx context.Context, teamId string) (*model.Project, error)
	GetProjectByProfessorID(ctx context.Context, professorId string) ([]*model.Project, error)
	CreateProject(ctx context.Context, project *model.Project) (*model.Project, error)
	UpdateProjectByID(ctx context.Context, id string, project *model.Project) (*model.Project, error)
	DeleteProjectByID(ctx context.Context, id string) (*model.Project, error)
	// AddStudentToProject(ctx context.Context, projectID string, studentID string) error
	// RemoveStudentFromProject(ctx context.Context, projectID string, studentID string) error
	GetProjectRequestByProjectID(ctx context.Context, projectID string) ([]*model.ProjectRequest, error)
	CreateProjectRequest(ctx context.Context, projectRequest *model.ProjectRequest) (*model.ProjectRequest, error)
	AcceptProjectRequest(ctx context.Context, id string) (*model.SuccessResponse, error)
	RejectProjectRequest(ctx context.Context, id string) (*model.SuccessResponse, error)
}

func (p *ProjectClient) GetAllProjects(ctx context.Context) ([]*model.Project, error) {
	res, err := (*p.client).GetAllProjects(ctx, &pb.Empty{})
	if err != nil {
		return nil, err
	}
	var projects []*model.Project
	for _, project := range res.Projects {
		projects = append(projects, &model.Project{
			ID:          project.ProjectId,
			TeamID:      project.TeamId,
			ProfessorID: project.ProfessorId,
			Name:        project.Name,
			Description: project.Description,
			Status:      project.Status,
			Label:       project.Label,
		})
	}
	return projects, nil
}

func (p *ProjectClient) GetProjectByID(ctx context.Context, id string) (*model.Project, error) {
	projectRes, err := (*p.client).GetProjectById(ctx, &pb.ProjectId{ProjectId: id})
	if err != nil {
		return nil, err
	}
	project := &model.Project{
		ID:          projectRes.ProjectId,
		TeamID:      projectRes.TeamId,
		ProfessorID: projectRes.ProfessorId,
		Name:        projectRes.Name,
		Description: projectRes.Description,
		Status:      projectRes.Status,
		Label:       projectRes.Label,
	}
	return project, nil
}

func (p *ProjectClient) GetProjectByTeamID(ctx context.Context, teamId string) (*model.Project, error) {
	projectRes, err := (*p.client).GetProjectByTeamId(ctx, &pb.TeamId{TeamId: teamId})
	if err != nil {
		return nil, err
	}
	project := &model.Project{
		ID:          projectRes.ProjectId,
		TeamID:      projectRes.TeamId,
		ProfessorID: projectRes.ProfessorId,
		Name:        projectRes.Name,
		Description: projectRes.Description,
		Status:      projectRes.Status,
		Label:       projectRes.Label,
	}
	return project, nil
}

func (p *ProjectClient) GetProjectByProfessorID(ctx context.Context, professorId string) ([]*model.Project, error) {
	res, err := (*p.client).GetProjectByProfessorId(ctx, &pb.ProfessorId{ProfessorId: professorId})
	if err != nil {
		return nil, err
	}
	var projects []*model.Project
	for _, project := range res.Projects {
		projects = append(projects, &model.Project{
			ID:          project.ProjectId,
			TeamID:      project.TeamId,
			ProfessorID: project.ProfessorId,
			Name:        project.Name,
			Description: project.Description,
			Status:      project.Status,
			Label:       project.Label,
		})
	}
	if projects == nil {
		return []*model.Project{}, nil
	}
	return projects, nil
}

func (p *ProjectClient) CreateProject(ctx context.Context, project *model.Project) (*model.Project, error) {
	projectRes, err := (*p.client).CreateProject(ctx, &pb.Project{ProfessorId: project.ProfessorID, Name: project.Name, Description: project.Description, Label: project.Label})
	if err != nil {
		return nil, err
	}
	project = &model.Project{
		ID:          projectRes.ProjectId,
		TeamID:      projectRes.TeamId,
		ProfessorID: projectRes.ProfessorId,
		Name:        projectRes.Name,
		Description: projectRes.Description,
		Status:      projectRes.Status,
		Label:       projectRes.Label,
	}
	return project, nil
}

func (p *ProjectClient) UpdateProjectByID(ctx context.Context, id string, project *model.Project) (*model.Project, error) {
	projectRes, err := (*p.client).UpdateProject(ctx, &pb.Project{ProjectId: id, ProfessorId: project.ProfessorID, TeamId: project.TeamID, Name: project.Name, Description: project.Description, Status: project.Status, Label: project.Label})
	if err != nil {
		return nil, err
	}
	project = &model.Project{
		ID:          projectRes.ProjectId,
		TeamID:      projectRes.TeamId,
		ProfessorID: projectRes.ProfessorId,
		Name:        projectRes.Name,
		Description: projectRes.Description,
		Status:      projectRes.Status,
		Label:       projectRes.Label,
	}
	return project, nil
}

func (p *ProjectClient) DeleteProjectByID(ctx context.Context, id string) (*model.Project, error) {
	projectRes, err := (*p.client).DeleteProject(ctx, &pb.ProjectId{ProjectId: id})
	if err != nil {
		return nil, err
	}
	project := &model.Project{
		ID:          projectRes.ProjectId,
		TeamID:      projectRes.TeamId,
		ProfessorID: projectRes.ProfessorId,
		Name:        projectRes.Name,
		Description: projectRes.Description,
		Status:      projectRes.Status,
		Label:       projectRes.Label,
	}
	return project, nil
}

// func (t *TeamClient) AddStudentToTeam(ctx context.Context, teamID string, studentID string) error {
// 	return nil
// }

// func (t *TeamClient) RemoveStudentFromTeam(ctx context.Context, teamID string, studentID string) error {
// 	return nil
// }

func (p *ProjectClient) GetProjectRequestByProjectID(ctx context.Context, projectID string) ([]*model.ProjectRequest, error) {
	res, err := (*p.client).GetProjectRequestsByProjectId(ctx, &pb.ProjectId{ProjectId: projectID})
	if err != nil {
		return nil, err
	}
	var projectRequests []*model.ProjectRequest
	for _, projectRequest := range res.ProjectRequests {
		projectRequests = append(projectRequests, &model.ProjectRequest{
			ProjectRequestID: projectRequest.ProjectRequestId,
			ProjectID:        projectRequest.ProjectId,
			TeamID:           projectRequest.TeamId,
			Message:          projectRequest.Message,
			Status:           projectRequest.Status,
		})
	}
	return projectRequests, nil
}

func (p *ProjectClient) CreateProjectRequest(ctx context.Context, projectRequest *model.ProjectRequest) (*model.ProjectRequest, error) {
	projectRequestRes, err := (*p.client).CreateProjectRequest(ctx, &pb.ProjectRequest{ProjectId: projectRequest.ProjectID, TeamId: projectRequest.TeamID, Message: projectRequest.Message, Status: projectRequest.Status})
	if err != nil {
		return nil, err
	}
	projectRequest = &model.ProjectRequest{
		ProjectRequestID: projectRequestRes.ProjectRequestId,
		ProjectID:        projectRequestRes.ProjectId,
		TeamID:           projectRequestRes.TeamId,
		Message:          projectRequestRes.Message,
		Status:           projectRequestRes.Status,
	}
	return projectRequest, nil
}

func (p *ProjectClient) AcceptProjectRequest(ctx context.Context, id string) (*model.SuccessResponse, error) {
	projectRequestRes, err := (*p.client).AcceptProjectRequest(ctx, &pb.ProjectRequest{ProjectRequestId: id})
	if err != nil {
		return nil, err
	}
	res := &model.SuccessResponse{
		Success: projectRequestRes.Success,
	}
	return res, nil
}

func (p *ProjectClient) RejectProjectRequest(ctx context.Context, id string) (*model.SuccessResponse, error) {
	projectRequestRes, err := (*p.client).RejectProjectRequest(ctx, &pb.ProjectRequest{ProjectRequestId: id})
	if err != nil {
		return nil, err
	}
	res := &model.SuccessResponse{
		Success: projectRequestRes.Success,
	}
	return res, nil
}

func ProvideProjectClient(client *pb.ProjectServiceClient) *ProjectClient {
	return &ProjectClient{client: client}
}
