package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Invoice struct {
	Id               primitive.ObjectID `json:"id" bson:"_id"`
	Invoice_uid      string             `json:"invoice_uid" bson:"invoice_uid"`
	Order_Id         string             `json:"order_id" bson:"order_id" validate:"required"`
	Paymemt_Method   *string            `json:"payment_method" bson:"payment_method" validate:"eq=CARD|eq=CASH|eq=UPI|eq="`
	Paymneent_Status *string            `json:"payment_status" bson:"payment_status" validate:"eq=PENDING|eq=PAID"`
	Payment_Due_date time.Time          `json:"payment_due_date" bson:"payment_due_date"`
	Created_at       time.Time          `json:"created_at" bson:"created_at"`
	Upadated_At      time.Time          `json:"updated_at" bson:"updated_at"`
}
