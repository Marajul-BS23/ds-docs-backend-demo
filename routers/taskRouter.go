package routes

import (
	"github.com/BrainStation-23/ds-docs-backend-demo/controllers"
	"github.com/BrainStation-23/ds-docs-backend-demo/middleware"
	"github.com/gin-gonic/gin"
)

func TaskRouter(router *gin.Engine){
	router.Use(middleware.Authenticate)
	router.POST("/task",controllers.CreateTask)
	router.GET("/task",controllers.GetTask)
	router.PUT("/task/:id",controllers.UpdateTask)
	router.DELETE("/task/:id",controllers.DeleteTask)


	//  need to use middleware for admin 
	router.Use(middleware.IsAdmin)
	router.GET("/task/all", controllers.GetAllTask)
	router.GET("/task/:id", controllers.GetTaskByUID)

}