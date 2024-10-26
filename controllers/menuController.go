package controllers

import (
	"context"
	"fmt"
	"net/http"
	"restaurant-golang/database"
	"restaurant-golang/models"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var menucollection *mongo.Collection = database.OpenCollection(database.Client, "menu")

func inTimeSpan(start, end, check time.Time) bool {
	return check.After(start) && check.Before(end)
}

func GetMenus() gin.HandlerFunc {
	return func(c *gin.Context) {
		result, err := menucollection.Find(context.TODO(), bson.M{})
		if err != nil {
			c.JSON(500, gin.H{
				"message": "Error",
			})
		}
		var menus []bson.M
		if err = result.All(context.TODO(), &menus); err != nil {
			c.JSON(500, gin.H{
				"message": "Error",
			})
		}
		c.JSON(200, menus)
	}
}

func GetMenu() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancle = context.WithTimeout(context.Background(), 100*time.Second)
		MenuID := c.Param("id")
		var menu models.Menu
		err := menucollection.FindOne(ctx, bson.M{"menu_id": MenuID}).Decode(&menu)
		defer cancle()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while feacthing menu"})

		}
		c.JSON(http.StatusOK, menu)
	}
}

func CreateMenu() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancle = context.WithTimeout(context.Background(), 100*time.Second)
		var menu models.Menu
		defer cancle()
		//Parsing the Request Body (JSON) into Menu Struct
		if err := c.BindJSON(&menu); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		validationErr := validate.Struct(menu)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}
		menu.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		menu.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		menu.ID = primitive.NewObjectID()
		menu.Menu_id = menu.ID.Hex()

		result, insertErr := menucollection.InsertOne(ctx, menu)
		if insertErr != nil {
			msg := fmt.Sprintf("Menu Item was not created")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}

		defer cancle()
		c.JSON(http.StatusOK, result)
	}
}

func UpdateMenu() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var menu models.Menu
		if err := c.BindJSON(&menu); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Extract the menu ID from URL parameters and convert it to ObjectID
		menuId := c.Param("id")
		objectId, err := primitive.ObjectIDFromHex(menuId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
			return
		}

		// Create the filter to find the correct menu by its ObjectID
		filter := bson.M{"_id": objectId}

		// Prepare the update object
		var updateObj primitive.D

		if menu.Name != "" {
			updateObj = append(updateObj, bson.E{Key: "name", Value: menu.Name})
		}
		if menu.Catoganry != "" {
			updateObj = append(updateObj, bson.E{Key: "catoganry", Value: menu.Catoganry})
		}
		if menu.Start_Date != nil && menu.End_Date != nil {
			if menu.Start_Date.After(*menu.End_Date) {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Start Date should be before End Date"})
				return
			}
			updateObj = append(updateObj, bson.E{Key: "start_date", Value: menu.Start_Date})
			updateObj = append(updateObj, bson.E{Key: "end_date", Value: menu.End_Date})
		}

		// Set the updated_at field to the current time
		menu.UpdatedAt = time.Now()
		updateObj = append(updateObj, bson.E{Key: "updated_at", Value: menu.UpdatedAt})

		// Perform the update with upsert option set to false
		opt := options.UpdateOptions{Upsert: new(bool)}
		result, err := menucollection.UpdateOne(ctx, filter, bson.D{{Key: "$set", Value: updateObj}}, &opt)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while updating Menu"})
			return
		}

		if result.MatchedCount == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "Menu not found"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Menu updated successfully", "result": result})
	}
}
