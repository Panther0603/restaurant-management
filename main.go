package main

import (
	"fmt"
	"os"
	"restraument-management/middleware"
	"restraument-management/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var envError error = godotenv.Load()

func main() {

	if envError != nil {
		fmt.Print(envError)
	}
	port := os.Getenv("PORT")

	if port != "" {
		port = "6003"
	}

	router := gin.New()
	router.Use(gin.Logger())
	routes.UserRoutes(router)
	router.Use(middleware.Authenication())

	routes.FoodRoutes(router)
	routes.InvoiceRoutes(router)
	routes.MenuRoutes(router)
	routes.NoteRoutes(router)
	routes.OrderItemRoutes(router)
	routes.OrderRoutes(router)
	routes.TableRoutes(router)

	router.Run(":" + port)

}
