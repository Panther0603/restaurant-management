package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Menu struct {
	Id         primitive.ObjectID `json:"id" bson:"_id"`
	Menu_Uid   string             `json:"menu_uid" bson:"menu_uid" validate:"required"`
	Name       string             `json:"name" bson:"name" validate:"required"`
	Category   string             `json:"category" validate:"required"`
	Start_Date time.Time          `json:"start_date" bson:"start_date"`
	End_Date   time.Time          `json:"end_date" bson:"end_date"`
	Created_at time.Time          `json:"created_at" bson:"created_at"`
	Updated_at time.Time          `json:"updated_date"  bson:"updated_date"`
}
