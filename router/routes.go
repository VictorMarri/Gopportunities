package router

import (
	"github.com/VictorMarri/Gopportunities.git/handler"
	"github.com/gin-gonic/gin"
)

// *gin.Engine como pointer para router do Gin
func initializeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{

		//Initialize handler
		handler.InitializeHandler()

		//http://localhost:8080/api/v1/opening

		//GET opening opportunities
		v1.GET("/opening", handler.ShowOpeningHandler)
		//POST new opportunities
		v1.POST("/opening", handler.CreateOpeningHandler)
		//DELETE an opportunity
		v1.DELETE("/opening", handler.DeleteOpeningHandler)
		//UPDATE an opportunity
		v1.PUT("/opening", handler.UpdateOpeningHandler)
		//GETs a list of available opportunities
		v1.GET("/openings", handler.ListOpeningHandler)
	}
}
