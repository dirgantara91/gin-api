/*
Package main provides a simple API for managing regents using the Gin framework.
*/
package main

import (
	// controller package provides the RegentController for handling regent-related requests.
	controller "gin-api/controllers"
	_ "gin-api/docs"

	// gin is the Gin framework for building web applications.
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

/*
main is the entry point of the application.
*/
// @title Geography API
// @description API for managing geography data
// @version 1.0
// @host localhost:5000
// @BasePath /
func main() {
	// Create a new Gin router
	// r is the Gin router instance.
	r := gin.Default()

	// Define the endpoints
	// rc is an instance of the RegentController.
	rc := &controller.RegentController{}
	// GET /regent returns a list of regents.
	r.GET("/regent", rc.GetRegents)
	// GET /regent/:id returns a regent by ID.
	r.GET("/regent/:id", rc.GetRegent)

	// Register Swagger middleware
	// Register Swagger middleware
	//url := ginSwagger.URL("http://localhost:5000/swagger/doc.json") // url
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Run the API
	// Start the Gin server and listen on 0.0.0.0:5000.
	r.Run(":5000") // listen and serve on 0.0.0.0:5000
}
