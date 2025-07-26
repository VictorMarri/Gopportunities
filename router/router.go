package router //Definindo subpackage

import "github.com/gin-gonic/gin"

func Initialize() {
	//Incializa router usando configs Default do pacote Gin
	router := gin.Default()
	//Define uma rota
	router.GET("/ping", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "pong",
		})
	})
	//Rodando API por padrao na porta 8080
	router.Run(":8080") // listen and serve on 0.0.0.0:8080
}
