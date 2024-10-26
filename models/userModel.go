package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID            primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	First_name    string             `json:"first_name" binding:"required" bson:"first_name" validate:"required,min=2,max=100"`
	Last_name     string             `json:"last_name" binding:"required" bson:"last_name" validate:"required,min=2,max=100"`
	Password      string             `json:"pass_word" binding:"required" bson:"pass_word" validate:"required,min=6"`
	Email         string             `json:"email" binding:"required" bson:"email" validate:"required,email"`
	Avatar        string             `json:"avatar" binding:"required" bson:"avatar" `
	Phone         string             `json:"phone" binding:"required" bson:"phone" validate:"required,min=10,max=10"`
	Token         string             `json:"token" binding:"required" bson:"token" `
	Refresh_Token string             `json:"refresh_token" binding:"required" bson:"refresh_token" `
	User_id       string             `json:"user_id" binding:"required" bson:"user_id" `
	CreatedAt     time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt     time.Time          `json:"updated_at" bson:"updated_at"`
}
