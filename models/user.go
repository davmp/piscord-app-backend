package models

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type User struct {
	ID        bson.ObjectID `json:"id" bson:"_id,omitempty"`
	Username  string        `json:"username" bson:"username" binding:"required,min=3,max=30"`
	Password  string        `json:"password" bson:"password" binding:"required,min=4"`
	Picture   string        `json:"picture" bson:"picture"`
	Bio       string        `json:"bio" bson:"bio"`
	CreatedAt time.Time     `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time     `json:"updatedAt" bson:"updatedAt"`
}

type UserResponse struct {
	ID        bson.ObjectID `json:"id"`
	Username  string        `json:"username"`
	Picture   string        `json:"picture,omitempty"`
	Bio       string        `json:"bio,omitempty"`
	CreatedAt time.Time     `json:"createdAt"`
}

type ProfileResponse struct {
	UserResponse
	IsOnline     bool           `json:"isOnline"`
	DirectChatID *bson.ObjectID `json:"directChatId,omitempty"`
}

type UpdateProfileRequest struct {
	Username string `json:"username,omitempty" binding:"omitempty,min=3,max=30"`
	Password string `json:"password,omitempty" binding:"omitempty,min=6"`
	Picture  string `json:"picture,omitempty"`
	Bio      string `json:"bio,omitempty"`
}

type UserRegisterRequest struct {
	Username string `json:"username" binding:"required,min=3,max=30"`
	Password string `json:"password" binding:"required,min=6"`
	Picture  string `json:"picture,omitempty"`
}

type AuthRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type AuthResponse struct {
	Token string        `json:"token"`
	User  *UserResponse `json:"user"`
}
