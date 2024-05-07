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

var FoodCollection *mongo.Collection = database.OpenCollection(database.Client, "foods")
var avgApitme time.Duration = 100 * time.Second

func CreateFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		var food models.Food
		ctx, cancel := GetApiContext()
		defer cancel()

		err := c.BindJSON(&food)
		if err != nil {
			ErrorAbort(c, http.StatusBadRequest, custom.CInValidBody)
			return
		}

		log.Println("bind")
		food.Id = primitive.NewObjectID()
		food.Food_Uid = uuid.NewString()
		food.Created_at = helper.GetCurentTime()
		food.Updated_at = helper.GetCurentTime()

		log.Println("mapping done")

		V = validator.New()
		err = V.Struct(food)

		log.Println("valoidating")
		if err != nil {
			ErrorAbort(c, http.StatusBadRequest, custom.CMissingReqField)
			return
		}

		count, err := FoodCollection.CountDocuments(ctx, bson.D{{Key: "menu_id", Value: food.Menu_id}, {Key: "name", Value: food.Name}})
		if err != nil {
			ErrorAbort(c, http.StatusBadRequest, "Please provide valid menu id")
			return
		}
		if count > 0 {
			ErrorAbort(c, http.StatusBadRequest, "Food is already present in the menu ")
			return
		}

		if food.Menu_id != "" {
			menuId, err := primitive.ObjectIDFromHex(food.Menu_id)
			if err != nil {
				ErrorAbort(c, http.StatusBadRequest, custom.CInvalidParentId)
				return
			}
			count, err := MenuCollection.CountDocuments(ctx, bson.D{{Key: "_id", Value: menuId}})
			if err != nil || count < 0 {
				ErrorAbort(c, http.StatusBadRequest, custom.CParentDataNoFound)
				return
			}
		}
		_, err = FoodCollection.InsertOne(ctx, &food)
		if err != nil {
			ErrorAbort(c, http.StatusBadRequest, custom.CDataNotSaves)
			return
		}
		c.JSON(http.StatusCreated, gin.H{"message": custom.DataSavedSucess + " food"})
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

		ctx, cancel := GetApiContext()
		defer cancel()

		var food models.Food

		err := c.BindJSON(&food)

		if err != nil {
			ErrorAbort(c, http.StatusBadRequest, custom.CInValidBody)
		}
		if food.Id.String() == "" {
			ErrorAbort(c, http.StatusBadRequest, custom.CMissingReqField)
		}

		filter := bson.D{{Key: "_id", Value: food.Id}}
		update := bson.D{{}}

		if food.Menu_id != "" {
			menuId, err := primitive.ObjectIDFromHex(food.Menu_id)
			if err != nil {
				ErrorAbort(c, http.StatusBadRequest, custom.CInvalidParentId)
			}
			count, err := MenuCollection.CountDocuments(ctx, bson.D{{Key: "_id", Value: menuId}})
			if err != nil || count < 0 {
				ErrorAbort(c, http.StatusBadRequest, custom.CParentDataNoFound)
			}
			update[0].Value = append(update[0].Value.(bson.D), bson.E{Key: "menu_id", Value: food.Menu_id})
		}

		if food.Name != "" {
			update[0].Value = append(update[0].Value.(bson.D), bson.E{Key: "name", Value: food.Name})
		}
		if food.Food_Image != "" {
			update[0].Value = append(update[0].Value.(bson.D), bson.E{Key: "food_image", Value: food.Food_Image})
		}
		if food.Price != nil {
			update[0].Value = append(update[0].Value.(bson.D), bson.E{Key: "price", Value: food.Price})
		}

		if err != nil {
			log.Println("time parse error")
		}
		update[0].Value = append(update[0].Value.(bson.D), bson.E{Key: "updated_at", Value: helper.GetCurentTime()})

		upsert := true

		opt := options.UpdateOptions{
			Upsert: &upsert,
		}

		updatedFood, err := FoodCollection.UpdateOne(ctx, filter, update, &opt)
		if err != nil {
			ErrorAbort(c, http.StatusInternalServerError, "food"+custom.CNotUpdated)
		}
		c.JSON(http.StatusOK, updatedFood)

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

func GetFoodsBYPerPage() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), avgApitme)
		defer cancel()

		var allFood []models.Food

		pageSize, pageNo := helper.GetPagenationPoint(c)

		matchStage := bson.D{{Key: "$match", Value: bson.D{{}}}}
		groupStage := bson.D{{Key: "$group", Value: bson.D{{Key: "_id", Value: "$id"}, {Key: "$sum", Value: 1}, {Key: "$push", Value: "$$ROOT"}}}}
		pageNationStage := bson.D{{Key: "$skip", Value: (pageNo * pageSize)}, {Key: "$limit", Value: pageSize}}

		pipeline := mongo.Pipeline{matchStage, groupStage, pageNationStage}

		foodCursor, err := FoodCollection.Aggregate(ctx, pipeline)

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

// func round(num float64) int {

// }

// func toFixed(num float64, precision float64) float64 {

// }

func GetApiContext() (ctx context.Context, cancel context.CancelFunc) {
	ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	return ctx, cancel
}
