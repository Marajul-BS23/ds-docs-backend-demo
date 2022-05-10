package middleware

import (
	"fmt"
	"log"
	"net/http"

	helper "github.com/BrainStation-23/ds-docs-backend-demo/helpers"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Authenticate(c *gin.Context)  {
		
		clientToken,error := c.Cookie("token")
		
		if error != nil || clientToken=="" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("No Authorization header provided")})
			c.Abort()
			return
		}

		claims, err := helper.ValidateToken(clientToken)
		if err != "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			c.Abort()
			return
		}
		
		c.Set("email", claims.Email)
		c.Set("first_name", claims.First_name)
		c.Set("last_name", claims.Last_name)
		c.Set("uid", claims.Uid)
		c.Set("user_type", claims.User_type)
		c.Next()
	
}




func IsAdmin(c *gin.Context){
	clientToken, err := c.Cookie("token");
	_ = err
	
	claims := jwt.MapClaims{} 
	
	token , parseErr := jwt.ParseWithClaims(
		clientToken,
		&claims,
		func(token *jwt.Token)(interface{}, error){
			return []byte(helper.SECRET_KEY), nil
		},
	)
	_ = token
	if parseErr!=nil {
		log.Fatal("middleware token parse error")
	}

	if(claims["User_type"]!="ADMIN"){
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		c.Abort()
	}

}