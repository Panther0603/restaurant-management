package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrderdItem struct {
	Id            primitive.ObjectID `json:"is" bson:"id"`
	OrderItem_Uid string             `json:"orderitem_uid" bson:"orderitem_uid"`
	Food_id       string             `json:"order_id" bson:"order_id" validate:"required"`
	Quantity      uint               `json:"quantity" bson:"quantity" validate:"required"`
	Order_Id      string             `json:"table_id" bson:"table_id" validate:"required"`
	Order_At      time.Time          `json:"booked_at" bson:"booked_at"`
}
