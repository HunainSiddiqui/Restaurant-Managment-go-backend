package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Note struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title     string             `json:"title" binding:"required" bson:"title" validate:"required" minLength:"2" maxLength:"100"`
	Content   string             `json:"content" binding:"required" bson:"content" validate:"required" minLength:"2" maxLength:"1000"`
	Note_id   string             `json:"note_id" binding:"required" bson:"note_id" `
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
}
