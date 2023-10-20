package client_grpc

import (
	"context"

	"capsmhoo/mono/api-gateway/model"

	pb "capsmhoo/gen/proto"
)

type NotiClient struct {
	client *pb.NotiServiceClient
}

type NotigRPCClient interface {
	GetAllNotiByUserId(ctx context.Context, userID string) ([]*model.Notification, error)
	ReadNoti(ctx context.Context, userID string) error
}

func (n *NotiClient) GetAllNotiByUserId(ctx context.Context, userID string) ([]*model.Notification, error) {
	res, err := (*n.client).GetAllNotiByUserId(ctx, &pb.GetAllNotiByUserIdRequest{UserId: userID})
	if err != nil {
		return nil, err
	}
	var notis []*model.Notification
	for _, noti := range res.Notis {
		notis = append(notis, &model.Notification{
			Title:  noti.Title,
			Body:   noti.Body,
			UserID: noti.UserId,
			IsRead: noti.IsRead,
		})
	}

	return notis, nil
}

func (n *NotiClient) ReadNoti(ctx context.Context, userID string) error {
	_, err := (*n.client).ReadNoti(ctx, &pb.ReadNotiRequest{UserId: userID})
	if err != nil {
		return err
	}
	return nil
}

func ProvideNotiClient(client *pb.NotiServiceClient) NotigRPCClient {
	return &NotiClient{client: client}
}
