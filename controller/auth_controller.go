package controller

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/philipepompeu/fake-pefin/service"
)

func AuthHandler(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if !strings.HasPrefix(authHeader, "Basic ") {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header inválido"})
		return
	}

	token, err := service.Authenticate(authHeader)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
