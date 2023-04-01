package middleware

import (
	"fmt"
	"jwt-gin/initializers"
	"jwt-gin/models"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func RequireAuth() gin.HandlerFunc {
	return func(c *gin.Context) {

		// get the cookie off req

		tokenString, err := c.Cookie("Authorization")
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		// decode it

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// validate the alg is what we expect
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(os.Getenv("SECRET")), nil
		})

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			// check the exp
			if float64(time.Now().Unix()) > claims["exp"].(float64) {
				c.AbortWithStatus(http.StatusUnauthorized)
			}

			// find the user with token sub
			var user models.User
			initializers.DB.First(&user, claims["sub"])

			if user.ID == 0 {
				c.AbortWithStatus(http.StatusUnauthorized)
			}

			//attach to req
			c.Set("user", user)

			// continue
			c.Next()
			fmt.Println(claims["sub"], claims["exp"])
		} else {
			fmt.Println(err)
		}
	}

}
