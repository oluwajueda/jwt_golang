package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/oluwajueda/jwt-with-golang/database"
	helper "github.com/oluwajueda/jwt-with-golang/helpers"
	"github.com/oluwajueda/jwt-with-golang/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")
var validate = validate.New()


func HashPassword()

func VerifyPassword()

func Signup()gin.HandlerFunc{
	return func(c *gin.Context){
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var user models.User

		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}
}

func Login()

func GetUsers()

func GetUser() gin.HandlerFunc{
         return func(c *gin.Context){
			userId := c.Param("user_id")

			if err := helper.MatchUserTypeToUid(c, userId); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
          var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
          var user models.User
		  err := userCollection.FindOne(ctx, bson.M{"user_id":userId}).Decode(&user)
		  defer cancel()

		  if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		  }
		  		c.JSON(http.StatusOK, user)
		 }
		}