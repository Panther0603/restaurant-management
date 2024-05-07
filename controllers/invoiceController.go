package controllers

import (
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
)

var InvoiceCollection *mongo.Collection = database.OpenCollection(database.Client, "Invoice")

func CreateInvoice() gin.HandlerFunc {
	return func(c *gin.Context) {
		var invoice models.Invoice
		ctx, cancel := GetApiContext()
		defer cancel()

		err := c.BindJSON(&invoice)
		if err != nil && ErrorAbort(c, http.StatusBadRequest, custom.CInValidBody) {
			return
		}

		invoice.Id = primitive.NewObjectID()
		invoice.Invoice_uid = uuid.NewString()
		invoice.Created_at = helper.GetCurentTime()
		invoice.Upadated_At = helper.GetCurentTime()
		V = validator.New()
		err = V.Struct(invoice)

		if err != nil && ErrorAbort(c, http.StatusBadRequest, custom.CMissingReqField) {
			return
		}

		count, err := InvoiceCollection.CountDocuments(ctx, bson.M{"order_id": invoice.Order_Id})
		if err != nil && ErrorAbort(c, http.StatusBadRequest, "Please provide valid order id") {
			return
		}
		if count > 0 && ErrorAbort(c, http.StatusBadRequest, "Invoice  is already created can't create again") {
			return
		}

		_, err = InvoiceCollection.InsertOne(ctx, &invoice)
		if err != nil && ErrorAbort(c, http.StatusBadRequest, custom.CDataNotSaves) {
			return
		}
	}
}

func GetInvoices() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := GetApiContext()
		defer cancel()

		var allInvoice []models.Invoice
		invoiceCursor, err := FoodCollection.Find(ctx, bson.M{})

		if err != nil && ErrorAbort(c, http.StatusInternalServerError, custom.CErrFindAllData) {
			return
		}

		err = invoiceCursor.All(ctx, &allInvoice)
		if err != nil && ErrorAbort(c, http.StatusInternalServerError, custom.CErrListParsing) {
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": allInvoice})
	}
}

func UpdateInvoice() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func GetInvoiceById() gin.HandlerFunc {
	return func(c *gin.Context) {

		ctx, cacel := GetApiContext()
		defer cacel()
		var invoice models.Menu
		invoiceId := c.Query("id")

		invoiceTId, err := primitive.ObjectIDFromHex(invoiceId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": custom.CInvalidId})
			c.Abort()
			return
		}
		err = InvoiceCollection.FindOne(ctx, bson.M{"_id": invoiceTId}).Decode(&invoice)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": custom.CIdNotFoud})
			c.Abort()
			return
		}

		c.Header("Content-Type", "application/json")
		c.JSON(http.StatusOK, invoice)
	}
}
