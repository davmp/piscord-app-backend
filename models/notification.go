package models

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type NotificationType string

const (
	NotificationTypeNewMessage          NotificationType = "NEW_MESSAGE"
	NotificationTypeUserJoined          NotificationType = "USER_JOINED"
	NotificationTypeUserLeft            NotificationType = "USER_LEFT"
	NotificationTypeFriendRequest       NotificationType = "FRIEND_REQUEST"
	NotificationTypeFriendRequestAccept NotificationType = "FRIEND_REQUEST_ACCEPTED"
	NotificationTypeRoomInvite          NotificationType = "ROOM_INVITE"
	NotificationTypeMention             NotificationType = "MENTION"
	NotificationTypeSystem              NotificationType = "SYSTEM"
)

type Notification struct {
	ID        bson.ObjectID    `json:"id" bson:"_id,omitempty"`
	UserID    bson.ObjectID    `json:"userId" bson:"userId"`
	Title     string           `json:"title" bson:"title"`
	Body      string           `json:"Body" bson:"body"`
	Link      string           `json:"link,omitempty" bson:"link,omitempty"`
	Picture   string           `json:"picture,omitempty" bson:"picture,omitempty"`
	Type      NotificationType `json:"type" bson:"type"`
	IsRead    bool             `json:"isRead" bson:"isRead"`
	CreatedAt time.Time        `json:"createdAt" bson:"createdAt"`
}
