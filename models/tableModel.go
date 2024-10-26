package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Table struct {
	ID              primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Table_id        string             `json:"table_id" binding:"required" bson:"table_id" validate:"required"`
	Table_number    int                `json:"table_number" binding:"required" bson:"table_number" validate:"required"`
	Number_of_seats int                `json:"number_of_seats" binding:"required" bson:"number_of_seats" validate:"required"`
	Table_status    string             `json:"table_status" binding:"required" bson:"table_status" validate:"eq=available|eq=booked|eq=reserved"`
	CreatedAt       time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt       time.Time          `json:"updated_at" bson:"updated_at"`
}
