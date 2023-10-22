package project

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "capsmhoo/gen/projectpb"
)

type projectServer struct {
	// Implements the generated ProjectServer interface
	pb.UnimplementedProjectServiceServer
	repo ProjectRepository
}

func (s *projectServer) mustEmbedUnimplementedProjectServiceServer() {}

func (s *projectServer) GetAllProjects(ctx context.Context, empty *pb.Empty) (*pb.ProjectList, error) {
	fmt.Println("Get All Projects")
	projects := s.repo.GetProjects()
	projectRes := []*pb.Project{}

	for _, project := range projects {
		projectRes = append(projectRes, convertProjectRes(&project))
	}

	return &pb.ProjectList{
		Projects: projectRes,
	}, nil
}

func (s *projectServer) GetProjectById(ctx context.Context, projectId *pb.ProjectId) (*pb.Project, error) {
	fmt.Println("Get Project By ID")
	project, err := s.repo.GetProjectByID(projectId.ProjectId)
	if err != nil {
		return nil, errors.New("project not found")
	}

	projectRes := convertProjectRes(project)

	return projectRes, nil
}

func (s *projectServer) CreateProject(ctx context.Context, project *pb.Project) (*pb.Project, error) {
	fmt.Println("Create Project")

	createdProject, err := s.repo.CreateProject(Project{
		ProfessorID: project.ProfessorId,
		Name:        project.Name,
		Description: project.Description,
	})
	if err != nil {
		return nil, err
	}

	createdProjectRes := convertProjectRes(createdProject)

	return createdProjectRes, nil
}

func (s *projectServer) UpdateProject(ctx context.Context, project *pb.Project) (*pb.Project, error) {
	fmt.Println("Update Project")

	updatedProject, err := s.repo.UpdateProjectByID(project.ProjectId, Project{
		ProjectID:   project.ProjectId,
		TeamID:      sql.NullString{String: project.TeamId},
		ProfessorID: project.ProfessorId,
		Name:        project.Name,
		Description: project.Description,
	})
	if err != nil {
		return nil, err
	}

	updatedProjectRes := convertProjectRes(updatedProject)

	return updatedProjectRes, nil
}

func (s *projectServer) DeleteProject(ctx context.Context, projectId *pb.ProjectId) (*pb.Project, error) {
	fmt.Println("Delete Project")

	project, err := s.repo.DeleteProjectByID(projectId.ProjectId)
	if err != nil {
		return nil, err
	}

	deletedProjectRes := convertProjectRes(project)

	return deletedProjectRes, nil
}

// func (s *projectServer) AddStudentToTeam(ctx context.Context, teamAndStudentID *pb.TeamAndStudentID) (*pb.Empty, error) {
// 	return nil, nil
// }

// func (s *projectServer) RemoveStudentFromTeam(ctx context.Context, teamAndStudentID *pb.TeamAndStudentID) (*pb.Empty, error) {
// 	return nil, nil
// }

func (s *projectServer) CreateProjectRequest(ctx context.Context, projectRequest *pb.ProjectRequest) (*pb.ProjectRequest, error) {
	fmt.Println("Create Project Request")

	createdProjectRequest, err := s.repo.CreateProjectRequest(ProjectRequest{
		ProjectID: projectRequest.ProjectId,
		TeamID:    projectRequest.TeamId,
		Message:   projectRequest.Message,
		Status:    projectRequest.Status,
	})
	if err != nil {
		return nil, err
	}

	createdProjectRequestRes := &pb.ProjectRequest{
		ProjectRequestId: createdProjectRequest.ProjectRequestID,
		ProjectId:        createdProjectRequest.ProjectID,
		TeamId:           createdProjectRequest.TeamID,
		Message:          createdProjectRequest.Message,
		Status:           createdProjectRequest.Status,
	}

	return createdProjectRequestRes, nil
}

func (s *projectServer) AcceptProjectRequest(ctx context.Context, projectRequest *pb.ProjectRequest) (*pb.SuccessResponse, error) {
	fmt.Println("Accept Project Request")

	projReq, err := s.repo.GetProjectRequestByID(projectRequest.ProjectRequestId)
	if err != nil {
		return nil, err
	}

	err = s.repo.AddTeamToProject(projReq.TeamID, projReq.ProjectID)
	if err != nil {
		return nil, err
	}

	err = s.repo.AcceptProjectRequest(projectRequest.ProjectRequestId)
	if err != nil {
		return nil, err
	}

	return &pb.SuccessResponse{
		Success: true,
	}, nil
}

func (s *projectServer) RejectProjectRequest(ctx context.Context, projectRequest *pb.ProjectRequest) (*pb.SuccessResponse, error) {
	fmt.Println("Reject Project Request")

	err := s.repo.RejectProjectRequest(projectRequest.ProjectRequestId)
	if err != nil {
		return nil, err
	}

	return &pb.SuccessResponse{
		Success: true,
	}, nil
}

func convertProjectRes(project *Project) *pb.Project {
	projectRes := pb.Project{
		ProjectId:   project.ProjectID,
		ProfessorId: project.ProfessorID,
		Name:        project.Name,
		Description: project.Description,
	}
	if project.TeamID.Valid {
		projectRes.TeamId = project.TeamID.String
	}
	return &projectRes
}

func StartgRPCServer(
	repo ProjectRepository,
	grpc_host string,
	grpc_port string,
) {
	log.Println("gRPC server running ...")

	lis, err := net.Listen("tcp", grpc_host+":"+grpc_port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterProjectServiceServer(s, &projectServer{repo: repo})
	log.Printf("Server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve : %v", err)
	}
	fmt.Println("Go gRPC server started")
}
