package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// *gin.Engine como pointer para router do Gin
func initializeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		//Show opening opportunities
		//http://localhost:8080/api/v1/opening
		v1.GET("/opening", func(ginContext *gin.Context) {
			ginContext.JSON(http.StatusOK, gin.H{
				"msg": "GET Opening",
			})
		})

		v1.POST("/opening", func(ginContext *gin.Context) {
			ginContext.JSON(http.StatusOK, gin.H{
				"msg": "POST Opening",
			})
		})

		v1.DELETE("/opening", func(ginContext *gin.Context) {
			ginContext.JSON(http.StatusOK, gin.H{
				"msg": "DELETE Opening",
			})
		})

		v1.PUT("/opening", func(ginContext *gin.Context) {
			ginContext.JSON(http.StatusOK, gin.H{
				"msg": "PUT Opening",
			})
		})

		v1.GET("/openings", func(ginContext *gin.Context) {
			ginContext.JSON(http.StatusOK, gin.H{
				"msg": "GET Openings",
			})
		})
	}
}
