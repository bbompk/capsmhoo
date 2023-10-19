package teamjoinrequest

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "capsmhoo/gen/team-join-request-pb"
)

type teamJoinRequestServer struct {
	pb.UnimplementedTeamJoinRequestServiceServer
	repo TeamJoinRequestRepository
}

func (s *teamJoinRequestServer) GetJoinRequests(ctx context.Context, empty *pb.Empty) (*pb.TeamJoinRequestList, error) {
	fmt.Println("Get All Join Requests")
	requests, err := s.repo.GetJoinRequests()
	if err != nil {
		return nil, err
	}

	requestRes := []*pb.TeamJoinRequest{}
	for _, req := range requests {
		requestRes = append(requestRes, &pb.TeamJoinRequest{
			Id:        req.ID,
			TeamId:    req.TeamID,
			StudentId: req.StudentID,
		})
	}

	return &pb.TeamJoinRequestList{
		JoinRequests: requestRes,
	}, nil
}

func (s *teamJoinRequestServer) GetJoinRequestById(ctx context.Context, reqID *pb.TeamJoinRequestId) (*pb.TeamJoinRequest, error) {
	fmt.Println("Get Join Request By ID")

	req, err := s.repo.GetJoinRequestByID(reqID.Id)
	if err != nil {
		return nil, err
	}

	return &pb.TeamJoinRequest{
		Id:        req.ID,
		TeamId:    req.TeamID,
		StudentId: req.StudentID,
	}, nil
}

func (s *teamJoinRequestServer) CreateJoinRequest(ctx context.Context, teamJoinRequest *pb.TeamJoinRequest) (*pb.TeamJoinRequest, error) {
	fmt.Println("Create Join Request")

	createdRequest, err := s.repo.CreateJoinRequest(TeamJoinRequest{
		ID:        teamJoinRequest.Id,
		TeamID:    teamJoinRequest.TeamId,
		StudentID: teamJoinRequest.StudentId,
	})
	if err != nil {
		return nil, err
	}

	return &pb.TeamJoinRequest{
		Id:        createdRequest.ID,
		TeamId:    createdRequest.TeamID,
		StudentId: createdRequest.StudentID,
	}, nil
}

func (s *teamJoinRequestServer) UpdateJoinRequestByID(ctx context.Context, updatedRequest *pb.TeamJoinRequest) (*pb.TeamJoinRequest, error) {
	fmt.Println("Update Join Request")

	req, err := s.repo.UpdateJoinRequestByID(updatedRequest.Id, TeamJoinRequest{
		ID:        updatedRequest.Id,
		TeamID:    updatedRequest.TeamId,
		StudentID: updatedRequest.StudentId,
	})
	if err != nil {
		return nil, err
	}

	return &pb.TeamJoinRequest{
		Id:        req.ID,
		TeamId:    req.TeamID,
		StudentId: req.StudentID,
	}, nil
}

func (s *teamJoinRequestServer) DeleteJoinRequestByID(ctx context.Context, reqID *pb.TeamJoinRequestId) (*pb.Empty, error) {
	fmt.Println("Delete Join Request")

	if err := s.repo.DeleteJoinRequestByID(reqID.Id); err != nil {
		return nil, err
	}

	return &pb.Empty{}, nil
}

// func (s *teamJoinRequestServer) ApproveJoinRequest(ctx context.Context, reqID *pb.TeamJoinRequestId) (*pb.TeamJoinRequest, error) {
// 	fmt.Println("Approve Join Request")
// 	// Conversion and logic here...
// 	//...
// }

// func (s *teamJoinRequestServer) DeclineJoinRequest(ctx context.Context, reqID *pb.TeamJoinRequestId) (*pb.TeamJoinRequest, error) {
// 	fmt.Println("Decline Join Request")
// 	// Conversion and logic here...
// 	//...
// }

func StartgRPCServer(
	repo TeamJoinRequestRepository,
	grpc_host string,
	grpc_port string,
) {
	log.Println("Team Join Request gRPC server running ...")

	lis, err := net.Listen("tcp", grpc_host+":"+grpc_port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterTeamJoinRequestServiceServer(s, &teamJoinRequestServer{repo: repo})
	log.Printf("Server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	fmt.Println("Team Join Request gRPC server started")
}
