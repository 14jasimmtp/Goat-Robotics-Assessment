package middleware

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/14jasimmtp/Goat-Robotics-Assessment/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func AuthMiddleware(c *gin.Context) {
	tokenString := strings.TrimPrefix(c.GetHeader("Authorization"), "Bearer ")

	var secretKey = viper.GetString("ATokenSecret")
	fmt.Println(secretKey, tokenString)

	Claims, err := utils.IsValidAccessToken(secretKey, tokenString)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Not Authorised"})
		return
	}
	fmt.Println("claimID" , Claims.ID)
	c.Set("User_id", Claims.ID)

	log.Println("MW: User Authorized")
	c.Next()
}
