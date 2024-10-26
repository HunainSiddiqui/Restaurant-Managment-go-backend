package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)


type Menu struct {
	ID         primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name       string             `json:"name" binding:"required" bson:"name" validate:"required" minLength:"2" maxLength:"100"`
	Menu_id   string             `json:"menu_id" binding:"required" bson:"menu_id" `
	Catoganry  string 		   `json:"catoganry" binding:"required" bson:"catoganry" validate:"required"`
	CreatedAt  time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt  time.Time          `json:"updated_at" bson:"updated_at"`
	Start_Date  *time.Time          `json:"start_date" bson:"start_date"`
	End_Date  *time.Time          `json:"end_date" bson:"end_date"`
}