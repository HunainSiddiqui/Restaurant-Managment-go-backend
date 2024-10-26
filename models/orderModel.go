package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Order struct {
	ID         primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Order_id   string             `json:"order_id" binding:"required" bson:"order_id" validate:"required"`
	Table_id   string             `json:"table_id" binding:"required" bson:"table_id" validate:"required"`
	Order_Date time.Time          `json:"order_date" binding:"required" bson:"order_date" `
	CreatedAt  time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt  time.Time          `json:"updated_at" bson:"updated_at"`
}
