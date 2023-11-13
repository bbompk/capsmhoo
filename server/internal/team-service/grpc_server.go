package team

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "capsmhoo/gen/proto"
	joinRequestPb "capsmhoo/gen/team-join-request-pb"
)

type teamServer struct {
	// Implements the generated TeamServer interface
	pb.UnimplementedTeamServiceServer
	repo         TeamRepository
	student_repo StudentRepository
}

func (s *teamServer) mustEmbedUnimplementedTeamServiceServer() {}

func (s *teamServer) GetAllTeams(ctx context.Context, empty *pb.Empty) (*pb.TeamList, error) {
	fmt.Println("Get Teams")
	teams := s.repo.GetTeams()
	teamRes := []*pb.Team{}

	for _, team := range teams {
		teamRes = append(teamRes, &pb.Team{
			Id:      team.ID,
			Name:    team.Name,
			Profile: team.Profile,
		})
	}

	return &pb.TeamList{
		Teams: teamRes,
	}, nil
}

func (s *teamServer) GetTeamById(ctx context.Context, teamId *pb.TeamId) (*pb.Team, error) {
	fmt.Println("Get Team By ID")
	team, err := s.repo.GetTeamByID(teamId.Id)
	if err != nil {
		return nil, errors.New("team not found")
	}

	teamRes := pb.Team{
		Id:      team.ID,
		Name:    team.Name,
		Profile: team.Profile,
	}

	return &teamRes, nil
}

func (s *teamServer) GetTeamByUserId(ctx context.Context, user_id *pb.UserId) (*pb.Team, error) {
	fmt.Println("Get Team By UserID")
	team, err := s.repo.GetTeamByUserID(user_id.Id)
	if err != nil {
		return nil, err
	}

	teamRes := pb.Team{
		Id:      team.ID,
		Name:    team.Name,
		Profile: team.Profile,
	}

	return &teamRes, nil
}

func (s *teamServer) CreateTeam(ctx context.Context, team *pb.Team) (*pb.Team, error) {
	fmt.Println("Create Team")

	createdTeam, err := s.repo.CreateTeam(Team{
		Name:    team.Name,
		Profile: team.Profile,
	})
	if err != nil {
		return nil, err
	}

	createdTeamRes := pb.Team{
		Id:      createdTeam.ID,
		Name:    createdTeam.Name,
		Profile: createdTeam.Profile,
	}

	return &createdTeamRes, nil
}

func (s *teamServer) UpdateTeam(ctx context.Context, team *pb.Team) (*pb.Team, error) {
	fmt.Println("Update Team")

	updatedTeam, err := s.repo.UpdateTeamByID(team.Id, Team{
		ID:      team.Id,
		Name:    team.Name,
		Profile: team.Profile,
	})
	if err != nil {
		return nil, err
	}

	updatedTeamRes := pb.Team{
		Id:      updatedTeam.ID,
		Name:    updatedTeam.Name,
		Profile: updatedTeam.Profile,
	}

	return &updatedTeamRes, nil
}

func (s *teamServer) DeleteTeam(ctx context.Context, teamId *pb.TeamId) (*pb.Team, error) {
	fmt.Println("Delete Team")

	team, err := s.repo.DeleteTeamByID(teamId.Id)
	if err != nil {
		return nil, err
	}

	deletedTeamRes := pb.Team{
		Id:      team.ID,
		Name:    team.Name,
		Profile: team.Profile,
	}

	return &deletedTeamRes, nil
}

func (s *teamServer) AddStudentToTeam(ctx context.Context, teamAndStudentID *pb.TeamAndStudentID) (*pb.Student, error) {
	fmt.Println("Add Student To Team")

	updatedStudent, err := s.student_repo.UpdateStudentTeam(teamAndStudentID.StudentId, teamAndStudentID.TeamId)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	student := pb.Student{
		Id:     updatedStudent.ID,
		Name:   updatedStudent.Name,
		UserId: updatedStudent.UserID,
		TeamId: *updatedStudent.TeamID,
	}

	return &student, nil
}

func (s *teamServer) RemoveStudentFromTeam(ctx context.Context, teamAndStudentID *pb.TeamAndStudentID) (*pb.Empty, error) {
	return nil, nil
}

type teamJoinRequestServer struct {
	joinRequestPb.UnimplementedTeamJoinRequestServiceServer
	repo         TeamJoinRequestRepository
	student_repo StudentRepository
}

func (s *teamJoinRequestServer) GetAllJoinRequests(ctx context.Context, empty *joinRequestPb.TeamJoinReqeustEmpty) (*joinRequestPb.TeamJoinRequestList, error) {
	fmt.Println("Get All Join Requests")
	requests, err := s.repo.GetJoinRequests()
	if err != nil {
		return nil, err
	}

	requestRes := []*joinRequestPb.TeamJoinRequest{}
	for _, req := range requests {
		requestRes = append(requestRes, &joinRequestPb.TeamJoinRequest{
			Id:        req.ID,
			TeamId:    req.TeamID,
			StudentId: req.StudentID,
		})
	}

	return &joinRequestPb.TeamJoinRequestList{
		JoinRequests: requestRes,
	}, nil
}

func (s *teamJoinRequestServer) GetJoinRequestById(ctx context.Context, reqID *joinRequestPb.TeamJoinRequestId) (*joinRequestPb.TeamJoinRequest, error) {
	fmt.Println("Get Join Request By ID")

	req, err := s.repo.GetJoinRequestByID(reqID.Id)
	if err != nil {
		return nil, err
	}

	return &joinRequestPb.TeamJoinRequest{
		Id:        req.ID,
		TeamId:    req.TeamID,
		StudentId: req.StudentID,
	}, nil
}

func (s *teamJoinRequestServer) GetJoinRequestByTeamId(ctx context.Context, teamID *joinRequestPb.TeamJoinRequestTeamId) (*joinRequestPb.TeamJoinRequestList, error) {
	fmt.Println("Get Join Requests By Team ID")

	requests, err := s.repo.GetJoinRequestByTeamID(teamID.TeamId)
	if err != nil {
		return nil, err
	}

	requestRes := []*joinRequestPb.TeamJoinRequest{}
	for _, req := range requests {
		requestRes = append(requestRes, &joinRequestPb.TeamJoinRequest{
			Id:        req.ID,
			TeamId:    req.TeamID,
			StudentId: req.StudentID,
		})
	}

	return &joinRequestPb.TeamJoinRequestList{
		JoinRequests: requestRes,
	}, nil
}

func (s *teamJoinRequestServer) GetJoinRequestByStudentId(ctx context.Context, studentID *joinRequestPb.TeamJoinRequestStudentId) (*joinRequestPb.TeamJoinRequestList, error) {
	fmt.Println("Get Join Requests By Student ID")

	requests, err := s.repo.GetJoinRequestByStudentID(studentID.StudentId)
	if err != nil {
		return nil, err
	}

	requestRes := []*joinRequestPb.TeamJoinRequest{}
	for _, req := range requests {
		requestRes = append(requestRes, &joinRequestPb.TeamJoinRequest{
			Id:        req.ID,
			TeamId:    req.TeamID,
			StudentId: req.StudentID,
		})
	}

	return &joinRequestPb.TeamJoinRequestList{
		JoinRequests: requestRes,
	}, nil
}

func (s *teamJoinRequestServer) CreateJoinRequest(ctx context.Context, teamJoinRequest *joinRequestPb.TeamJoinRequest) (*joinRequestPb.TeamJoinRequest, error) {
	fmt.Println("Create Join Request")

	createdRequest, err := s.repo.CreateJoinRequest(TeamJoinRequest{
		ID:        teamJoinRequest.Id,
		TeamID:    teamJoinRequest.TeamId,
		StudentID: teamJoinRequest.StudentId,
	})
	if err != nil {
		return nil, err
	}

	return &joinRequestPb.TeamJoinRequest{
		Id:        createdRequest.ID,
		TeamId:    createdRequest.TeamID,
		StudentId: createdRequest.StudentID,
	}, nil
}

func (s *teamJoinRequestServer) UpdateJoinRequest(ctx context.Context, updatedRequest *joinRequestPb.TeamJoinRequest) (*joinRequestPb.TeamJoinRequest, error) {
	fmt.Println("Update Join Request")

	req, err := s.repo.UpdateJoinRequestByID(updatedRequest.Id, TeamJoinRequest{
		ID:        updatedRequest.Id,
		TeamID:    updatedRequest.TeamId,
		StudentID: updatedRequest.StudentId,
	})
	if err != nil {
		return nil, err
	}

	return &joinRequestPb.TeamJoinRequest{
		Id:        req.ID,
		TeamId:    req.TeamID,
		StudentId: req.StudentID,
	}, nil
}

func (s *teamJoinRequestServer) DeleteJoinRequest(ctx context.Context, reqID *joinRequestPb.TeamJoinRequestId) (*joinRequestPb.TeamJoinRequest, error) {
	fmt.Println("Delete Join Request")

	req, err := s.repo.DeleteJoinRequestByID(reqID.Id)
	// _, err := s.repo.DeleteJoinRequestByID(reqID.Id)
	if err != nil {
		return nil, err
	}

	return &joinRequestPb.TeamJoinRequest{
		Id:        req.ID,
		TeamId:    req.TeamID,
		StudentId: req.StudentID,
	}, nil
	// return &joinRequestPb.TeamJoinReqeustEmpty{}, nil
}

func (s *teamJoinRequestServer) ApproveJoinRequest(ctx context.Context, reqID *joinRequestPb.TeamJoinRequestId) (*joinRequestPb.TeamJoinRequest, error) {
	fmt.Println("Approve Join Request")

	var req *TeamJoinRequest
	req, err := s.repo.GetJoinRequestByID(reqID.Id)
	if err != nil {
		return nil, err
	}

	std, err := s.student_repo.GetStudentByID(req.StudentID)
	if err != nil {
		return nil, err
	}

	updatedTeam := false

	if std.TeamID == nil {
		// Update student team
		_, err = s.student_repo.UpdateStudentTeam(req.StudentID, req.TeamID)
		if err != nil {
			return nil, err
		}
		updatedTeam = true
	}

	// Delete join request
	req, err = s.repo.DeleteJoinRequestByID(reqID.Id)
	if err != nil {
		return nil, err
	}

	if !updatedTeam {
		return &joinRequestPb.TeamJoinRequest{
			Id:        req.ID,
			TeamId:    req.TeamID,
			StudentId: req.StudentID,
		}, errors.New("student already in a team")
	}

	return &joinRequestPb.TeamJoinRequest{
		Id:        req.ID,
		TeamId:    req.TeamID,
		StudentId: req.StudentID,
	}, nil

}

func (s *teamJoinRequestServer) DeclineJoinRequest(ctx context.Context, reqID *joinRequestPb.TeamJoinRequestId) (*joinRequestPb.TeamJoinRequest, error) {
	fmt.Println("Decline Join Request")

	req, err := s.repo.DeleteJoinRequestByID(reqID.Id)
	if err != nil {
		return nil, err
	}

	return &joinRequestPb.TeamJoinRequest{
		Id:        req.ID,
		TeamId:    req.TeamID,
		StudentId: req.StudentID,
	}, nil
}

func StartgRPCServer(
	repo TeamRepository,
	repoo TeamJoinRequestRepository,
	student_repo StudentRepository,
	grpc_host string,
	grpc_port string,
) {
	log.Println("gRPC server running ...")

	lis, err := net.Listen("tcp", grpc_host+":"+grpc_port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterTeamServiceServer(s, &teamServer{repo: repo})
	joinRequestPb.RegisterTeamJoinRequestServiceServer(s, &teamJoinRequestServer{repo: repoo, student_repo: student_repo})
	log.Printf("Server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve : %v", err)
	}
	fmt.Println("Go gRPC server started")
}
