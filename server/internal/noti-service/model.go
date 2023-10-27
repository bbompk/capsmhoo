package noti

type Notification struct {
	Title  string `json:"title"`
	Body   string `json:"body"`
	UserID string `json:"user_id"`
	IsRead bool   `json:"is_read"`
}

type NotificationFilter struct {
	UserID string `bson:"user_id" json:"user_id"`
}

type CreateNotificationMessage struct {
	Title      string `json:"title"`
	Body       string `json:"body"`
	ReceiverID string `json:"receiver_id"`
}
