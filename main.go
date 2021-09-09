package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wildanadt/go-rest/controllers"
	"github.com/wildanadt/go-rest/storage"
)

func main() {

	db := storage.ConnectToDB()
	fmt.Println(db)

	router := gin.Default()

	router.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	public := router.Group("/api")

	public.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Welcome to Wildanadt API"})
	})

	public.POST("/register", controllers.Register)

	router.Run("0.0.0.0:8080")
}
