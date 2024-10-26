package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Food struct {
	ID         primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name       string             `json:"name" binding:"required" bson:"name" validate:"required" minLength:"2" maxLength:"100"`
	Price      *float64            `json:"price" binding:"required" bson:"price" validate:"required" min:"1"`
	Food_image string             `json:"food_image" binding:"required" bson:"food_image" validate:"required"`
	Food_id    string             `json:"food_id" binding:"required" bson:"food_id" `
	Menu_id    string             `json:"menu_id" binding:"required" bson:"menu_id" validate:"required"`
	CreatedAt  time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt  time.Time          `json:"updated_at" bson:"updated_at"`
}
