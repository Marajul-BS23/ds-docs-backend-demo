package routes

import (
	"github.com/BrainStation-23/ds-docs-backend-demo/controllers"
	"github.com/gin-gonic/gin"
)

func AuthRouter(router *gin.Engine){
	
	router.POST("/login",controllers.Login);
	router.POST("/signup",controllers.Signup);
}