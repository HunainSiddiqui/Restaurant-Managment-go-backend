package controllers

import (
	"context"
	"net/http"
	"restaurant-golang/database"
	helper "restaurant-golang/helper"
	"restaurant-golang/models"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var usercollection *mongo.Collection = database.OpenCollection(database.Client, "user")

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}
func GetUsers() gin.HandlerFunc {
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
		cursor, err := usercollection.Aggregate(context.TODO(), pipeline)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching foods"})
			return
		}

		var users []bson.M
		if err = cursor.All(context.TODO(), &users); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error decoding foods"})
			return
		}

		// Send the paginated result as JSON
		c.JSON(http.StatusOK, gin.H{
			"page":  page,
			"limit": limit,
			"data":  users,
		})

	}
}

func GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancle = context.WithTimeout(context.Background(), 100*time.Second)
		userID := c.Param("id")
		var user models.User
		err := usercollection.FindOne(ctx, bson.M{"user_id": userID}).Decode(&user)
		defer cancle()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while feacthing user"})

		}
		c.JSON(http.StatusOK, user)
	}
}

func CreateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancle = context.WithTimeout(context.Background(), 100*time.Second)
		var user models.User
		defer cancle()
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
			return
		}
		validate := validator.New()
		validationErr := validate.Struct(user)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error in Data Validation": validationErr.Error()})
			return
		}

		// If email or phone number already exists
		email := user.Email
		phone := user.Phone
		emailErr := usercollection.FindOne(ctx, bson.M{"email": email}).Decode(&user)
		phoneErr := usercollection.FindOne(ctx, bson.M{"phone": phone}).Decode(&user)
		if emailErr == nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Email already exists"})
			return
		}
		if phoneErr == nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Phone number already exists"})
			return
		}

		//Hasing the Password

		hashedPassword, err := HashPassword(user.Password)
		user.Password = hashedPassword
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occured while hashing password"})
			return
		}

		token, refreshtoken, _ := helper.GenerateAllTokens(user.First_name, user.Last_name, user.Email, user.User_id)

		user.Token = token
		user.Refresh_Token = refreshtoken

		user.ID = primitive.NewObjectID()
		user.User_id = user.ID.Hex()
		user.CreatedAt = time.Now()
		user.UpdatedAt = time.Now()
		result, insertErr := usercollection.InsertOne(ctx, user)
		if insertErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occured while inserting user"})
			return
		}
		defer cancle()
		c.JSON(http.StatusOK, result)
	}
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

func LoginUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancle := context.WithTimeout(context.Background(), 100*time.Second)
		var user LoginRequest
		defer cancle()
		var foundUser models.User
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
			return
		}
		defer cancle()
		email := user.Email
		err := usercollection.FindOne(ctx, bson.M{"email": email}).Decode(&foundUser)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "User not found"})
			return
		}
		err = bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(user.Password))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid Password"})
			return
		}
		token, refreshtoken, _ := helper.GenerateAllTokens(foundUser.First_name, foundUser.Last_name, foundUser.Email, foundUser.User_id)
		foundUser.Token = token
		foundUser.Refresh_Token = refreshtoken
		c.JSON(http.StatusOK, foundUser)

	}
}

func UpdateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "PUT User",
		})
	}
}

func DeleteUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "DELETE User",
		})
	}
}
