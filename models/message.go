package models

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Message struct {
	ID        bson.ObjectID  `json:"id" bson:"_id,omitempty"`
	RoomID    bson.ObjectID  `json:"roomId" bson:"roomId"`
	UserID    bson.ObjectID  `json:"userId" bson:"userId"`
	Content   string         `json:"content" bson:"content"`
	Type      string         `json:"type" bson:"type"` // "text", "image", "file", "system"
	FileURL   string         `json:"fileUrl,omitempty" bson:"fileUrl,omitempty"`
	ReplyTo   *bson.ObjectID `json:"replyTo,omitempty" bson:"replyTo,omitempty"`
	IsEdited  bool           `json:"isEdited" bson:"isEdited"`
	IsDeleted bool           `json:"isDeleted" bson:"isDeleted"`
	CreatedAt time.Time      `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt" bson:"updatedAt"`
}

type SendMessageRequest struct {
	RoomID  string `json:"roomId" binding:"required"`
	Content string `json:"content" binding:"required"`
	Type    string `json:"type,omitempty"`
	ReplyTo string `json:"replyTo,omitempty"`
}

type MessageResponse struct {
	ID           bson.ObjectID           `json:"id"`
	RoomID       bson.ObjectID           `json:"roomId"`
	UserID       bson.ObjectID           `json:"userId"`
	Username     string                  `json:"username"`
	Picture      string                  `json:"picture,omitempty"`
	Content      string                  `json:"content"`
	Type         string                  `json:"type"`
	FileURL      string                  `json:"fileUrl,omitempty"`
	IsOwnMessage bool                    `json:"isOwnMessage"`
	ReplyTo      *MessagePreviewResponse `json:"replyTo,omitempty"`
	IsEdited     bool                    `json:"isEdited"`
	CreatedAt    time.Time               `json:"createdAt"`
	UpdatedAt    time.Time               `json:"updatedAt"`
}

type MessagePreviewResponse struct {
	ID        bson.ObjectID `json:"id"`
	Username  string        `json:"username"`
	Content   string        `json:"content"`
	Picture   string        `json:"picture,omitempty"`
	CreatedAt time.Time     `json:"createdAt"`
}

type WSMessage struct {
	Type    string `json:"type"`
	Payload any    `json:"payload,omitempty"`
}

type WSResponse struct {
	Type    string `json:"type"`
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
	Data    any    `json:"data,omitempty"`
}
