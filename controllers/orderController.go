package controllers

import (
	"context"
	"log"
	"net/http"
	"restraument-management/custom"
	"restraument-management/database"
	"restraument-management/helper"
	"restraument-management/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cacel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cacel()

		V = validator.New()

		var order models.Order
		err := c.BindJSON(&order)

		if err != nil && ErrorAbort(c, http.StatusBadRequest, custom.CInValidBody) {
			return
		}

		order.Id = primitive.NewObjectID()
		order.Order_Uid = uuid.NewString()
		order.Booked_At = helper.GetCurentTime()
		order.Order_UpdatedAt = helper.GetCurentTime()

		err = V.Struct(order)
		if err != nil {
			log.Print(err.Error())
			ErrorAbort(c, http.StatusBadRequest, custom.CMissingReqField)
			return
		}

		if order.Table_id != "" {
			menuId, err := primitive.ObjectIDFromHex(order.Table_id)
			if err != nil {
				ErrorAbort(c, http.StatusBadRequest, custom.CInvalidParentId)
			}
			count, err := MenuCollection.CountDocuments(ctx, bson.D{{Key: "_id", Value: menuId}})
			if err != nil || count < 0 {
				ErrorAbort(c, http.StatusBadRequest, custom.CParentDataNoFound)
			}
		}

		_, err = OrderCollection.InsertOne(ctx, order)
		if err != nil {
			ErrorAbort(c, http.StatusInternalServerError, custom.CDataNotSaves)
			return
		}
		c.JSON(http.StatusCreated, gin.H{"message": custom.DataSavedSucess})

	}
}

var OrderCollection *mongo.Collection = database.OpenCollection(database.Client, "orders")

func GetOrders() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := GetApiContext()
		defer cancel()

		var allOrder []models.Order

		pointCursor, err := OrderCollection.Find(ctx, bson.M{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": custom.CErrFindAllData + " order"})
			c.Abort()
			return
		}

		err = pointCursor.All(ctx, &allOrder)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": custom.CErrListParsing + " of order"})
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, &allOrder)
	}
}

func UpdateOrder() gin.HandlerFunc {
	return func(c *gin.Context) {

		ctx, cancel := GetApiContext()
		defer cancel()

		var order models.Order

		err := c.BindJSON(&order)

		if err != nil {
			ErrorAbort(c, http.StatusBadRequest, custom.CInValidBody)
		}
		if order.Id.String() == "" {
			ErrorAbort(c, http.StatusBadRequest, custom.CMissingReqField)
		}

		filter := bson.D{{Key: "_id", Value: order.Id}}
		update := bson.D{{
			Key:   "$set",
			Value: bson.D{},
		}}

		if order.Table_id != "" {
			menuId, err := primitive.ObjectIDFromHex(order.Table_id)
			if err != nil {
				ErrorAbort(c, http.StatusBadRequest, custom.CInvalidParentId)
			}
			count, err := TableCollection.CountDocuments(ctx, bson.D{{Key: "_id", Value: menuId}})
			if err != nil || count < 0 {
				ErrorAbort(c, http.StatusBadRequest, custom.CParentDataNoFound)
			}
			update[0].Value = append(update[0].Value.(bson.D), bson.E{Key: "table_id", Value: order.Table_id})
		}

		if order.Customer_Name != "" {
			update[0].Value = append(update[0].Value.(bson.D), bson.E{Key: "customer_name", Value: order.Customer_Name})
		}
		if order.Order_Status != "" {
			update[0].Value = append(update[0].Value.(bson.D), bson.E{Key: "order_status", Value: order.Order_Status})
		}

		if err != nil {
			log.Println("time parse error")
		}
		update[0].Value = append(update[0].Value.(bson.D), bson.E{Key: "order_updatedat", Value: helper.GetCurentTime()})

		upsert := true

		opt := options.UpdateOptions{
			Upsert: &upsert,
		}

		updatedFood, err := OrderCollection.UpdateOne(ctx, filter, update, &opt)
		if err != nil {
			ErrorAbort(c, http.StatusInternalServerError, "order"+custom.CNotUpdated)
		}
		c.JSON(http.StatusOK, updatedFood)

	}
}

func GetOrderById() gin.HandlerFunc {
	return func(c *gin.Context) {
		var order models.Order
		// used heleper function instead of the native func
		ctx, cancel := GetApiContext()
		defer cancel()

		frderId, changed := c.GetQuery("id")

		if !changed && ErrorAbort(c, http.StatusBadRequest, custom.CInvalidId) {
			return
		}
		frderTId, err := primitive.ObjectIDFromHex(frderId)
		if err != nil && ErrorAbort(c, http.StatusBadRequest, custom.CInvalidId) {
			return
		}
		err = OrderCollection.FindOne(ctx, bson.M{"_id": frderTId}).Decode(&order)
		if err != nil && ErrorAbort(c, http.StatusNotFound, custom.CIdNotFoud) {
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": order})
	}
}
