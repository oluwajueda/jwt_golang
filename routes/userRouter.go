package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/oluwajueda/jwt-with-golang/controllers"
	middleware "github.com/oluwajueda/jwt-with-golang/middleware"
)

func UserRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/users", controller.GetUsers())
	incomingRoutes.GET("/users/:user_id", controller.GetUser())
}