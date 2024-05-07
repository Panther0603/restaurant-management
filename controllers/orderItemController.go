package controllers

import (
	"restraument-management/database"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var OrderItemCollection *mongo.Collection = database.OpenCollection(database.Client, "orderitems")

func CreateOrderItem() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func GetOrderItems() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func UpdateOrderItem() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func GetOrderItemById() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

// func ItemsByOrder(id string) (orderdItem []primitive.M, err error) {

// }
