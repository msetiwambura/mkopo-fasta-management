package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"usrmanagement/utils"
)

/*
	func AuthMiddleware() gin.HandlerFunc {
		return func(c *gin.Context) {
			tokenString := c.GetHeader("Authorization")
			if tokenString == "" {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing token"})
				c.Abort()
				return
			}
			tokenString = strings.TrimPrefix(tokenString, "Bearer ")
			claims, err := utils.ValidateToken(tokenString)
			if err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Token"})
				c.Abort()
				return
			}

			// Check if user has access to the requested page
			var user models.User
			if err := configs.DB.Preload("Role.Pages").Where("email = ?", claims.Username).First(&user).Error; err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
				c.Abort()
				return
			}

			// Check if user has access to the requested URL
			requestedURL := c.Request.URL.Path
			hasAccess := false
			fmt.Println("Protected Pages : ", len(user.Role.Pages))
			for _, page := range user.Role.Pages {
				fmt.Println("Protected Pages : ", page.URL)
				if page.URL == requestedURL {
					hasAccess = true
					break
				}
			}
			if !hasAccess {
				c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
				c.Abort()
				return
			}
			c.Set("username", claims.Username)
			c.Set("role", user.Role.Name)
			c.Next()
		}
	}
*/
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			c.Abort()
			return
		}
		tokenStr := strings.Split(authHeader, "Bearer ")[1]
		claims, err := utils.ValidateToken(tokenStr)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		c.Set("claims", claims)
		c.Next()
	}
}
