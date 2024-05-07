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

var MenuCollection *mongo.Collection = database.OpenCollection(database.Client, "menu")

var V *validator.Validate = validator.New()

func CreateMenu() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cacel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cacel()

		V = validator.New()

		var menu models.Menu
		err := c.BindJSON(&menu)

		if err != nil && ErrorAbort(c, http.StatusBadRequest, custom.CInValidBody) {
			return
		}

		menu.Id = primitive.NewObjectID()
		menu.Menu_Uid = uuid.NewString()
		menu.Created_at = helper.GetCurentTime()
		menu.Updated_at = helper.GetCurentTime()

		err = V.Struct(menu)
		if err != nil {
			ErrorAbort(c, http.StatusBadRequest, custom.CMissingReqField)
			return
		}

		_, err = MenuCollection.InsertOne(ctx, &menu)

		if err != nil && ErrorAbort(c, http.StatusInternalServerError, custom.CDataNotSaves+" menu") {
			return
		}
		c.JSON(http.StatusCreated, gin.H{"message": custom.DataSavedSucess})
	}
}

func GetMenus() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var allMenu []models.Menu

		pointCursor, err := MenuCollection.Find(ctx, bson.M{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": custom.CErrFindAllData + " menu"})
			c.Abort()
			return
		}

		err = pointCursor.All(ctx, &allMenu)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": custom.CErrListParsing + " of menu"})
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, &allMenu)
	}
}

func UpdateMenu() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := GetApiContext()
		defer cancel()

		var menu models.Menu
		err := c.BindJSON(&menu)
		if err != nil {
			ErrorAbort(c, http.StatusBadRequest, custom.CInValidBody)
			return
		}

		if menu.Id.IsZero() {
			ErrorAbort(c, http.StatusBadRequest, custom.CMissingReqField)
			return
		}

		filter := bson.M{"_id": menu.Id}

		update := bson.D{
			{Key: "$set", Value: bson.D{
				{Key: "start_date", Value: menu.Start_Date},
				{Key: "end_date", Value: menu.End_Date},
			}},
		}

		if menu.Name != "" {
			update[0].Value = append(update[0].Value.(bson.D), bson.E{Key: "name", Value: menu.Name})
		}

		if menu.Category != "" {
			update[0].Value = append(update[0].Value.(bson.D), bson.E{Key: "category", Value: menu.Category})
		}

		upsert := true
		opt := options.UpdateOptions{
			Upsert: &upsert,
		}

		log.Println("Now will update")

		updatedMenu, err := MenuCollection.UpdateOne(ctx, filter, update, &opt)
		if err != nil {
			ErrorAbort(c, http.StatusInternalServerError, "menu "+custom.CNotUpdated)
			return
		}
		log.Printf("update count %v", updatedMenu.ModifiedCount)

		c.JSON(http.StatusOK, gin.H{"message": custom.DataUpdatedSucess + " menu"})
	}
}

func GetMenuById() gin.HandlerFunc {
	return func(c *gin.Context) {

		ctx, cacel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cacel()
		var menu models.Menu
		menuId := c.Query("id")

		menutId, err := primitive.ObjectIDFromHex(menuId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": custom.CInvalidId})
			c.Abort()
			return
		}
		err = MenuCollection.FindOne(ctx, bson.M{"_id": menutId}).Decode(&menu)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": custom.CIdNotFoud})
			c.Abort()
			return
		}

		c.Header("Content-Type", "application/json")
		c.JSON(http.StatusOK, menu)

	}
}

func ErrorAbort(c *gin.Context, statusCode int, errMsg string) bool {
	if errMsg != "" {
		c.JSON(statusCode, gin.H{"error": errMsg})
		c.Abort()
		return true // Indicate that an error occurred
	}
	return false

}

func IsValidTimeStartEnd(start time.Time, end time.Time) bool {
	return start.After(time.Now()) && end.After(start)
}

func ParseMenuTime(menu *models.Menu, start, end string) error {
	layout := "2006-01-02T15:04:05Z07:00" // Time layout for JSON dates
	startTime, err := time.Parse(layout, start)
	if err != nil {
		return err
	}
	endTime, err := time.Parse(layout, end)
	if err != nil {
		return err
	}
	menu.Start_Date = startTime
	menu.End_Date = endTime
	return nil
}
