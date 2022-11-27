package routes

import (
	"fmt"
	"net/http"

	"github.com/ShivanshVerma-coder/link-tracking/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.Use(CORSMiddleware())

	router.GET("/", func(c *gin.Context) {
		// helpers.PrettyPrint(c.RemoteIP())

		fmt.Printf("%v", c.ClientIP())
		c.JSON(http.StatusOK, "Server is running")
	})

	linkStoreController := controllers.NewLinksController()

	router.POST("/generate", linkStoreController.Generate)
	router.GET("/:id", linkStoreController.GetTargetLink)

	return router

}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
