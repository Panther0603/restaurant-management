package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrderStatus string

const (
	StatusBooking   OrderStatus = "BOOKED"
	StatusPreparing OrderStatus = "PREPARING"
	StatusServed    OrderStatus = "SERVED" // corrected from StatusServer to StatusServed
)

type Order struct {
	Id              primitive.ObjectID `json:"id" bson:"_id"`
	Order_Uid       string             `json:"order_uid" bson:"order_uid" validate:"required"`
	Customer_Name   string             `json:"customer_name" bson:"customer_name" validate:"required,min=3,max=100"`
	Table_id        string             `json:"table_id" bson:"table_id" validate:"required"`
	Booked_At       time.Time          `json:"booked_at" bson:"booked_at"`
	Order_Status    OrderStatus        `json:"order_status" bson:"order_status" validate:"required,oneof=BOOKED PREPARING SERVED"` // removed spaces around oneof values
	Order_UpdatedAt time.Time          `json:"order_updatedat" bson:"order_updatedat"`
}
