package models

import (
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Food struct {
	Id         primitive.ObjectID `json:"id" bson:"_id"`
	Name       *string            `json:"name" ,bson:"name" validate:"required, min=3,max=100"`
	Price      *float64           `json:"price" bson:"price" validate:"required"`
	Food_Image string             `json:"food_image" bson:"food_image" validate:"required"`
	Food_Uid   string             `json:"food_uid" bson:"food_uid" validate:"required,unique"`
	Created_at time.Time          `json:"created_at" bson:"created_at"`
	Updated_at time.Time          `json:"updated_at" bson:"updated_at"`
	Menu_id    string             `json:"menu_id" bson:"menu_id" validate:"required"`
	Is_Active  bool               `json:"is_active" bson:"is_active"`
}

func NewFood() Food {
	return Food{
		Is_Active:  true,
		Created_at: time.Now(),
		Updated_at: time.Now(),
		Food_Uid:   uuid.NewString(),
	}
}
