package token

import (
	"fmt"
	"strings"

	signupin "golesson/controller/signUp-In"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the Bearer Token from the Authorization header
		bearerToken := c.GetHeader("Authorization")
		fmt.Println(bearerToken)
		tokenString := strings.TrimPrefix(bearerToken, "Bearer ")

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Don't forget to validate the alg is what you expect:
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
			return signupin.SecretKey, nil
		})

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			fmt.Println(claims["email"], claims["timeStamp"])
		} else {
			fmt.Println(err)
		}

		// Verify that the Bearer Token is valid
		// token, err := jwt.ParseWithClaims(bearerToken, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		// 	// Verify that the signing method is HS256
		// 	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		// 		return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		// 	}

		// 	// Return the secret key used to sign the token
		// 	return []byte("my-secret-key"), nil
		// })

		// if err != nil {
		// 	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		// 	return
		// }

		// if !token.Valid {
		// 	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token is not valid"})
		// 	return
		// }

		// // Set the user ID in the context
		// claims := token.Claims.(*jwt.StandardClaims)
		// c.Set("userID", claims.Subject)

		c.Next()
	}
}
