package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Table struct {
	Id         primitive.ObjectID `json:"id" bson:"_id"`
	Table_No   uint               `json:"table_no" bson:"table_no" validate:"required"`
	Table_Uid  string             `json:"table_uid" bson:"table_uid" validate:"required"`
	Created_At time.Time          `json:"created_at" bson:"created_at"`
}
