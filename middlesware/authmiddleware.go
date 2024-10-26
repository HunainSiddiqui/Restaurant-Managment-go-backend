package authention

import (
	"net/http"
	help "restaurant-golang/helper"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the value of the Authorization header
		clientToken := c.GetHeader("token")
		if clientToken == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "No Authorization header provided"})
			c.Abort()
			return
		}

		// // Split the token into "Bearer" and the token value
		// extractedToken := strings.Split(clientToken, "Bearer ")
		// if len(extractedToken) == 2 {
		// 	clientToken = extractedToken[1]
		// } else {
		// 	c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token provided"})
		// 	c.Abort()
		// 	return
		// }

		// Validate the token
		claims, err := help.ValidateToken(clientToken)

		if err != "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err})
			c.Abort()
			return
		}

		c.Set("email", claims.Email)
		c.Set("first_name", claims.FirstName)
		c.Set("last_name", claims.LastName)
		c.Set("uid", claims.Uid)
		c.Next()

	}
}
