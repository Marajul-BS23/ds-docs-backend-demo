package routes

import (
	"github.com/BrainStation-23/dsdoc-backend/controllers"
	"github.com/BrainStation-23/dsdoc-backend/middleware"
	"github.com/gin-gonic/gin"
)

func TaskRouter(router *gin.Engine){
	router.Use(middleware.Authenticate)
	router.POST("/task",controllers.CreateTaks)
	router.GET("/task",controllers.GetTask)
	router.PATCH("/task/:id",controllers.UpdateTask)
}