package noti

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "capsmhoo/gen/proto"
)

type notiServer struct {
	// Implements the generated NotiServer interface
	pb.UnimplementedNotiServiceServer
	repo NotiRepository
}

// func (s *notiServer) mustEmbedUnimplementedNotiServiceServer() {}

func (s *notiServer) GetAllNotiByUserId(ctx context.Context, req *pb.GetAllNotiByUserIdRequest) (*pb.GetAllNotiByUserIdResponse, error) {
	log.Default().Println("Get All Noti By User ID")
	notis, err := s.repo.GetAllNotiByUserID(req.UserId)
	if err != nil {
		log.Fatal("Error while getting all noti by user id")
		return nil, err
	}

	notiRes := []*pb.Noti{}

	for _, noti := range notis {
		notiRes = append(notiRes, &pb.Noti{
			Title:  noti.Title,
			Body:   noti.Body,
			UserId: noti.UserID,
			IsRead: noti.IsRead,
		})
	}

	return &pb.GetAllNotiByUserIdResponse{
		Notis: notiRes,
	}, nil
}

func (s *notiServer) ReadNoti(ctx context.Context, req *pb.ReadNotiRequest) (*pb.ReadNotiResponse, error) {
	log.Default().Println("Read Notification")
	err := s.repo.DeleteAllNotiByUserID(req.UserId)
	if err != nil {
		log.Fatal("Error while deleting all noti by user id")
		return &pb.ReadNotiResponse{
			Success: false,
		}, err
	}

	return &pb.ReadNotiResponse{
		Success: true,
	}, nil
}

func StartgRPCServer(
	repo NotiRepository,
	grpc_host string,
	grpc_port string,
) {
	log.Println("gRPC server running ...")

	lis, err := net.Listen("tcp", grpc_host+":"+grpc_port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterNotiServiceServer(s, &notiServer{repo: repo})
	log.Printf("Server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve : %v", err)
	}
	log.Default().Println("Go gRPC server started")
}
