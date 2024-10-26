package help

import (
	"context"
	"os"
	"restaurant-golang/database"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type SignedDetails struct {
	Email     string
	FirstName string
	LastName  string
	Uid       string
	jwt.RegisteredClaims
}

var usercollection *mongo.Collection = database.OpenCollection(database.Client, "user")

var SECRATE_KEY = os.Getenv("SECRATE_KEY")

func GenerateAllTokens(First_name string, Last_name string, Email string, User_id string) (signedToken string, signedRefreshToken string, err error) {
	claims := &SignedDetails{
		Email:     Email,
		FirstName: First_name,
		LastName:  Last_name,
		Uid:       User_id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 15)),
		},
	}

	refreshclaims := jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 15)),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshclaims)

	signedToken, err = token.SignedString([]byte(SECRATE_KEY))
	if err != nil {
		return
	}

	signedRefreshToken, err = refreshToken.SignedString([]byte(SECRATE_KEY))
	if err != nil {
		return
	}

	return signedToken, signedRefreshToken, nil

}

func UpdateToken(signedtoken string, refreshsignedtoken, user_id string) {
	var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var updatedObj primitive.D
	updatedObj = append(updatedObj, bson.E{Key: "token", Value: signedtoken})
	updatedObj = append(updatedObj, bson.E{Key: "refresh_token", Value: refreshsignedtoken})
	Update_at, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	updatedObj = append(updatedObj, bson.E{Key: "updated_at", Value: Update_at})
	upsert := true
	filter := bson.M{"user_id": user_id}
	opt := options.UpdateOptions{
		Upsert: &upsert,
	}
	_, err := usercollection.UpdateOne(ctx, filter, bson.D{{Key: "$set", Value: updatedObj}}, &opt)
	if err != nil {
		return
	}
}

func ValidateToken(signedtoken string) (claiams *SignedDetails, msg string) {

	tokenClaims, err := jwt.ParseWithClaims(signedtoken, &SignedDetails{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SECRATE_KEY), nil
	})
	if err != nil {
		msg = err.Error()
		return
	}
	claiams, ok := tokenClaims.Claims.(*SignedDetails)
	if !ok {
	
		msg = "The token is invalid"
		return
	}
	if claiams.ExpiresAt.Before(time.Now()) {
		msg = "Token has Expired"
		return
	}
	return claiams, msg

}
