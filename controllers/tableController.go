package controllers

import (
	"log"
	"net/http"
	"restraument-management/custom"
	"restraument-management/database"
	"restraument-management/helper"
	"restraument-management/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var TableCollection *mongo.Collection = database.OpenCollection(database.Client, "tables")

func CreateTable() gin.HandlerFunc {
	return func(c *gin.Context) {
		var table models.Table
		ctx, cancel := GetApiContext()
		defer cancel()

		err := c.BindJSON(&table)
		if err != nil {
			ErrorAbort(c, http.StatusBadRequest, custom.CInValidBody)
			return
		}

		log.Println("bind")
		table.Id = primitive.NewObjectID()
		table.Table_Uid = uuid.NewString()
		table.Created_At = helper.GetCurentTime()

		V = validator.New()
		err = V.Struct(table)

		if err != nil {
			ErrorAbort(c, http.StatusBadRequest, custom.CMissingReqField)
			return
		}

		count, err := TableCollection.CountDocuments(ctx, bson.D{{Key: "table_no", Value: table.Table_No}})
		if err != nil {
			ErrorAbort(c, http.StatusBadRequest, "Please provide valid Table no")
			return
		}
		if count > 0 {
			ErrorAbort(c, http.StatusBadRequest, "Table is already present in the table no ")
			return
		}

		_, err = TableCollection.InsertOne(ctx, &table)
		if err != nil {
			ErrorAbort(c, http.StatusBadRequest, custom.CDataNotSaves)
			return
		}

		c.JSON(http.StatusCreated, gin.H{"message": custom.DataSavedSucess + " table"})
	}
}

func GetTables() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := GetApiContext()
		defer cancel()

		var allTable []models.Table

		pointCursor, err := TableCollection.Find(ctx, bson.M{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": custom.CErrFindAllData + " table"})
			c.Abort()
			return
		}

		err = pointCursor.All(ctx, &allTable)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": custom.CErrListParsing + " of table"})
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, allTable)
	}
}

func UpdateTable() gin.HandlerFunc {
	return func(c *gin.Context) {

		ctx, cancel := GetApiContext()
		defer cancel()

		var table models.Table

		err := c.BindJSON(&table)

		if err != nil {
			ErrorAbort(c, http.StatusBadRequest, custom.CInValidBody)
			return
		}
		if table.Id.String() == "" {
			ErrorAbort(c, http.StatusBadRequest, custom.CMissingReqField)
			return
		}

		filter := bson.D{{Key: "_id", Value: table.Id}}
		update := bson.D{{
			Key:   "$set",
			Value: bson.D{{}},
		}}

		if table.Table_No != 0 {
			count, err := TableCollection.CountDocuments(ctx, bson.D{{Key: "table_no", Value: table.Table_No}})
			if err != nil {
				ErrorAbort(c, http.StatusBadRequest, "Please provide valid Table no")
				return
			}
			if count > 0 {
				ErrorAbort(c, http.StatusBadRequest, "Table is already present in the table no ")
				return
			}
			update[0].Value = append(update[0].Value.(bson.D), bson.E{Key: "table_no", Value: table.Table_No})
		}

		upsert := true

		opt := options.UpdateOptions{
			Upsert: &upsert,
		}

		updatedFood, err := TableCollection.UpdateOne(ctx, filter, update, &opt)
		if err != nil {
			ErrorAbort(c, http.StatusInternalServerError, "table"+custom.CNotUpdated)
		}
		c.JSON(http.StatusOK, updatedFood)

	}

}

func GetTableById() gin.HandlerFunc {
	return func(c *gin.Context) {
		var table models.Table
		// used heleper function instead of the native func
		ctx, cancel := GetApiContext()
		defer cancel()

		tableId, changed := c.GetQuery("id")

		if !changed && ErrorAbort(c, http.StatusBadRequest, custom.CInvalidId) {
			return
		}
		tableITd, err := primitive.ObjectIDFromHex(tableId)
		if err != nil && ErrorAbort(c, http.StatusBadRequest, custom.CInvalidId) {
			return
		}
		err = TableCollection.FindOne(ctx, bson.M{"_id": tableITd}).Decode(&table)
		if err != nil && ErrorAbort(c, http.StatusNotFound, custom.CIdNotFoud) {
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": table})
	}
}
