package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Invoice struct {
	ID               primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Order_id         string             `json:"order_id" binding:"required" bson:"order_id" validate:"required"`
	Invoice_id       string             `json:"invoice_id" binding:"required" bson:"invoice_id" validate:"required"`
	Payment_Method   string             `json:"payment_method" binding:"required" bson:"payment_method" validate:"eq=credit_card|eq=debit_card|eq=paypal|eq=gopay|eq="`
	Payment_Status   string             `json:"payment_status" binding:"required" bson:"payment_status" validate:"eq=pending|eq=success|eq=failed|eq=expired"`
	Payment_Due_Date time.Time          `json:"payment_due_date" binding:"required" bson:"payment_due_date" `
	CreatedAt        time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt        time.Time          `json:"updated_at" bson:"updated_at"`
}
