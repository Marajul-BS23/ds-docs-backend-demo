package routes

import (
	"github.com/BrainStation-23/ds-docs-backend-demo/controllers"
	"github.com/BrainStation-23/ds-docs-backend-demo/middleware"
	"github.com/gin-gonic/gin"
)

func TaskRouter(router *gin.Engine){
	router.Use(middleware.Authenticate)
	router.POST("/task",controllers.CreateTaks)
	router.GET("/task",controllers.GetTask)
	router.PUT("/task/:id",controllers.UpdateTask)
	router.DELETE("/task/:id",controllers.DeleteTask)
}