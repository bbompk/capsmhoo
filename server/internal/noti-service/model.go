package noti

type Notification struct {
	Title  string `bson:"title" json:"title"`
	Body   string `bson:"body" json:"body"`
	UserID string `bson:"user_id" json:"user_id"`
	IsRead bool   `bson:"is_read" json:"is_read"`
}

type NotificationFilter struct {
	UserID string `bson:"user_id" json:"user_id"`
}

type CreateNotificationMessage struct {
	Title      string `json:"title"`
	Body       string `json:"body"`
	ReceiverID string `json:"receiver_id"`
}
