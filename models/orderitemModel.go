package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrderItem struct {
	ID            primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Order_id      string             `json:"order_id" binding:"required" bson:"order_id" validate:"required"`
	Food_id       string             `json:"food_id" binding:"required" bson:"food_id" validate:"required"`
	Order_item_id string             `json:"order_item_id" binding:"required" bson:"order_item_id" validate:"required"`
	Quantity      int                `json:"quantity" binding:"required" bson:"quantity" validate:"required"`
	Unit_price    float64            `json:"unit_price" binding:"required" bson:"unit_price" validate:"required" min:"1"`
	CreatedAt     time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt     time.Time          `json:"updated_at" bson:"updated_at"`
}
