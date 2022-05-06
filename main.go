package main

import (
	routes "github.com/BrainStation-23/ds-docs-backend-demo/routers"
	"github.com/gin-gonic/gin"
)
func main() {
	//  load env files to os in dev
	// err := godotenv.Load(".env")
	// if err!=nil {
	// 	log.Fatal("Error loding Env File")
	// }


	// port := os.Getenv("PORT")
	// if port == "" {
	// 	port = "5000";
	// }

	router := gin.New()
	router.Use(gin.Logger())

	routes.AuthRouter(router)
	routes.TaskRouter(router)

	// router.Run(":" + port)

	router.Run()
}