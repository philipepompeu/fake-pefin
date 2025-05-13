package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/philipepompeu/fake-pefin/controller"
)

func main() {
	r := gin.Default()

	r.POST("/security/iam/v1/client-identities/login", controller.AuthHandler)

	protected := r.Group("/")
	protected.Use(JWTAuthMiddleware())
	{
		protected.POST("/collection/debt/", controller.DebtHandler)
	}

	port := ":3005"

	fmt.Println("Servidor iniciado em : localhost" + port)
	r.Run(port)
}
