package noti

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	mongoClient *mongo.Client
}

type NotiRepository interface {
	GetAllNotis() ([]Notification, error)
	GetNotiByID(id string) (*Notification, error)
	CreateNoti(noti Notification) (*Notification, error)
	UpdateNotiByID(id string, noti Notification) (*Notification, error)
	DeleteNotiByID(id string) (*Notification, error)
	DeleteAllNotiByUserID(userID string) error
	GetAllNotiByUserID(userID string) ([]Notification, error)
}

func (r *Repository) GetAllNotis() ([]Notification, error) {
	coll := r.mongoClient.Database("capsmhoo").Collection("notifications")
	cur, err := coll.Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.Background())

	var notis []Notification
	for cur.Next(context.Background()) {
		var noti Notification
		err := cur.Decode(&noti)
		if err != nil {
			return nil, err
		}
		notis = append(notis, noti)
	}
	return notis, nil
}

func (r *Repository) GetNotiByID(id string) (*Notification, error) {
	coll := r.mongoClient.Database("capsmhoo").Collection("notifications")
	var noti Notification
	err := coll.FindOne(context.Background(), bson.M{"_id": id}).Decode(&noti)
	if err != nil {
		return nil, err
	}
	return &noti, nil
}

func (r *Repository) CreateNoti(noti Notification) (*Notification, error) {
	coll := r.mongoClient.Database("capsmhoo").Collection("notifications")
	_, err := coll.InsertOne(context.Background(), noti)
	if err != nil {
		return nil, err
	}
	return &noti, nil
}

func (r *Repository) UpdateNotiByID(id string, noti Notification) (*Notification, error) {
	coll := r.mongoClient.Database("capsmhoo").Collection("notifications")
	_, err := coll.UpdateOne(context.Background(), bson.M{"_id": id}, noti)
	if err != nil {
		return nil, err
	}
	return &noti, nil
}

func (r *Repository) DeleteNotiByID(id string) (*Notification, error) {
	coll := r.mongoClient.Database("capsmhoo").Collection("notifications")
	var noti Notification
	err := coll.FindOneAndDelete(context.Background(), bson.M{"_id": id}).Decode(&noti)
	if err != nil {
		return nil, err
	}
	return &noti, nil
}

func (r *Repository) DeleteAllNotiByUserID(userID string) error {
	coll := r.mongoClient.Database("capsmhoo").Collection("notifications")
	_, err := coll.DeleteMany(context.Background(), NotificationFilter{UserID: userID})
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) GetAllNotiByUserID(userID string) ([]Notification, error) {
	coll := r.mongoClient.Database("capsmhoo").Collection("notifications")
	cur, err := coll.Find(context.Background(), NotificationFilter{UserID: userID})
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.Background())

	var notis []Notification
	for cur.Next(context.Background()) {
		var noti Notification
		err := cur.Decode(&noti)
		if err != nil {
			return nil, err
		}
		notis = append(notis, noti)
	}
	return notis, nil
}

func ProvideNotiRepository(mongoClient *mongo.Client) *Repository {
	return &Repository{mongoClient}
}
