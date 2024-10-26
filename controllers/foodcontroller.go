package controllers

import (
	"context"
	"fmt"
	"math"
	"net/http"
	"restaurant-golang/database"
	"restaurant-golang/models"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var foodcollection *mongo.Collection = database.OpenCollection(database.Client, "food")

var validate = validator.New()

func toFixed(value float64, precision int) float64 {
	power := math.Pow(10, float64(precision))
	return math.Round(value*power) / power
}

func GetFoods() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get page and limit from query parameters, with default values
		pageParam := c.DefaultQuery("page", "1")
		limitParam := c.DefaultQuery("limit", "10")

		// Convert page and limit to integers
		page, err := strconv.Atoi(pageParam)
		if err != nil || page < 1 {
			page = 1 // Default to page 1 if invalid
		}

		limit, err := strconv.Atoi(limitParam)
		if err != nil || limit < 1 {
			limit = 10 // Default to 10 if invalid
		}

		// Calculate the number of documents to skip
		skip := (page - 1) * limit

		// MongoDB aggregation pipeline for pagination
		pipeline := mongo.Pipeline{
			{{Key: "$skip", Value: skip}},
			{{Key: "$limit", Value: limit}},
		}

		// Execute the aggregation pipeline
		cursor, err := foodcollection.Aggregate(context.TODO(), pipeline)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching foods"})
			return
		}

		var foods []bson.M
		if err = cursor.All(context.TODO(), &foods); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error decoding foods"})
			return
		}

		// Send the paginated result as JSON
		c.JSON(http.StatusOK, gin.H{
			"page":  page,
			"limit": limit,
			"data":  foods,
		})
	}
}

func GetFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancle = context.WithTimeout(context.Background(), 100*time.Second)
		foodID := c.Param("id")
		var food models.Food
		err := foodcollection.FindOne(ctx, bson.M{"food_id": foodID}).Decode(&food)
		defer cancle()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while feacthing food"})

		}
		c.JSON(http.StatusOK, food)

	}
}

func CreateFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancle = context.WithTimeout(context.Background(), 100*time.Second)
		var food models.Food
		var menu models.Menu
		defer cancle()
		//Parsing the Request Body (JSON) into food Struct
		if err := c.BindJSON(&food); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		validationErr := validate.Struct(food)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}

		menucollection := database.OpenCollection(database.Client, "menu")
		defer cancle()
		err := menucollection.FindOne(ctx, bson.M{"menu_id": food.Menu_id}).Decode(&menu)

		if err != nil {
			msg := fmt.Sprintf("Menu not found")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}
		food.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		food.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		food.ID = primitive.NewObjectID()
		food.Food_id = food.ID.Hex()
		var num = toFixed(*food.Price, 2)
		food.Price = &num

		result, insertErr := foodcollection.InsertOne(ctx, food)
		if insertErr != nil {
			msg := fmt.Sprintf("Food Item was not created")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}

		defer cancle()
		c.JSON(http.StatusOK, result)

	}
}

func UpdateFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Update Food",
		})
	}
}
