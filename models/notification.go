package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type NotificationType string

const (
	NotificationTypeNewMessage          NotificationType = "NEW_MESSAGE"
	NotificationTypeMessageReaction     NotificationType = "MESSAGE_REACTION"
	NotificationTypeFriendRequest       NotificationType = "FRIEND_REQUEST"
	NotificationTypeFriendRequestAccept NotificationType = "FRIEND_REQUEST_ACCEPTED"
	NotificationTypeRoomInvite          NotificationType = "ROOM_INVITE"
	NotificationTypeMention             NotificationType = "MENTION"
	NotificationTypeSystem              NotificationType = "SYSTEM"
	NotificationTypeOther               NotificationType = "OTHER"
)

type Notification struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserID    primitive.ObjectID `bson:"user_id" json:"user_id"`
	Content   string             `bson:"content" json:"content"`
	Type      NotificationType   `bson:"type" json:"type"`
	ReadAt    *time.Time         `bson:"read_at,omitempty" json:"read_at,omitempty"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
}

type NotificationResponse struct {
	ID        primitive.ObjectID `json:"id"`
	Content   string             `json:"content"`
	Type      NotificationType   `json:"type"`
	ReadAt    *time.Time         `json:"read_at,omitempty"`
	CreatedAt time.Time          `json:"created_at"`

	Title   string `json:"title,omitempty"`
	Link    string `json:"link,omitempty"`
	Picture string `json:"picture,omitempty"`
}
