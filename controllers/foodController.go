package controllers

import (
	"context"
	"net/http"
	"restraument-management/custom"
	"restraument-management/database"
	"restraument-management/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var FoodCollection *mongo.Collection = database.OpenCollection(database.Client, "foods")
var avgApitme time.Duration = 100 * time.Second

func CreateFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		var food models.Food
		ctx, cancel := GetApiContext()
		defer cancel()

		err := c.BindJSON(&food)
		if err != nil && ErrorAbort(c, http.StatusBadRequest, custom.CInValidBody) {
			return
		}

		food.Id = primitive.NewObjectID()
		food.Food_Uid = uuid.NewString()
		food.Created_at = time.Now()
		food.Updated_at = time.Now()
		err = v.Struct(food)

		if err != nil && ErrorAbort(c, http.StatusBadRequest, custom.CMissingReqField) {
			return
		}

		count, err := FoodCollection.CountDocuments(ctx, bson.M{"menu_id": food.Menu_id})
		if err != nil && ErrorAbort(c, http.StatusBadRequest, "Please provide valid menu id") {
			return
		}
		if count > 0 && ErrorAbort(c, http.StatusBadRequest, "Food is already present in the menu ") {
			return
		}

		_, err = FoodCollection.InsertOne(ctx, &food)
		if err != nil && ErrorAbort(c, http.StatusBadRequest, custom.CDataNotSaves) {
			return
		}
	}
}

func GetFoods() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), avgApitme)
		defer cancel()

		var allFood []models.Food
		foodCursor, err := FoodCollection.Find(ctx, bson.M{})

		if err != nil && ErrorAbort(c, http.StatusInternalServerError, custom.CErrFindAllData) {
			return
		}

		err = foodCursor.All(ctx, &allFood)
		if err != nil && ErrorAbort(c, http.StatusInternalServerError, custom.CErrListParsing) {
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": allFood})
	}
}

func UpdateFood() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func GetFoodById() gin.HandlerFunc {
	return func(c *gin.Context) {
		var food models.Food
		// used heleper function instead of the native func
		ctx, cancel := GetApiContext()
		defer cancel()

		foodId, changed := c.GetQuery("id")

		if !changed && ErrorAbort(c, http.StatusBadRequest, custom.CInvalidId) {
			return
		}
		foodTId, err := primitive.ObjectIDFromHex(foodId)
		if err != nil && ErrorAbort(c, http.StatusBadRequest, custom.CInvalidId) {
			return
		}
		err = FoodCollection.FindOne(ctx, bson.M{"_id": foodTId}).Decode(&food)
		if err != nil && ErrorAbort(c, http.StatusNotFound, custom.CIdNotFoud) {
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": food})
	}
}

// func round(num float64) int {

// }

// func toFixed(num float64, precision float64) float64 {

// }

func GetApiContext() (ctx context.Context, cancel context.CancelFunc) {
	ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	return ctx, cancel
}
