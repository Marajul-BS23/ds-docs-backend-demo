package main

import (
	"log"
	"os"

	routes "github.com/BrainStation-23/dsdoc-backend/routers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)
func main() {

	err := godotenv.Load(".env")
	if err!=nil {
		log.Fatal("Error loding Env File")
	}


	port := os.Getenv("PORT")
	if port == "" {
		port = "5000";
	}

	router := gin.New()
	router.Use(gin.Logger())

	routes.AuthRouter(router)
	routes.TaskRouter(router)


	router.Run(":" + port)

}