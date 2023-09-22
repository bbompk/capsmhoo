package team

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "capsmhoo/gen/proto"
)

type teamServer struct {
	// Implements the generated TeamServer interface
	pb.UnimplementedTeamServiceServer
	repo TeamRepository
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

func (s *teamServer) AddStudentToTeam(ctx context.Context, teamAndStudentID *pb.TeamAndStudentID) (*pb.Empty, error) {
	return nil, nil
}

func (s *teamServer) RemoveStudentFromTeam(ctx context.Context, teamAndStudentID *pb.TeamAndStudentID) (*pb.Empty, error) {
	return nil, nil
}

func StartgRPCServer(
	repo TeamRepository,
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
	log.Printf("Server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve : %v", err)
	}
	fmt.Println("Go gRPC server started")
}
