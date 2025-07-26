package router //Definindo subpackage

import "github.com/gin-gonic/gin"

func Initialize() {

	//Initialize router using default configs from Gin package
	router := gin.Default()

	//Initialize routes from routes package
	initializeRoutes(router)

	//Rodando API por padrao na porta 8080
	router.Run(":8080") // listen and serve on 0.0.0.0:8080
}
