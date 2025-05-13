package controller

import (
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/philipepompeu/fake-pefin/service"
)

func DebtHandler(c *gin.Context) {
	username := c.GetString("username")

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Erro ao ler body"})
		return
	}

	id := service.SaveRawDebt(string(body))

	c.JSON(http.StatusOK, gin.H{
		"message":         "Debt armazenado com sucesso",
		"transactionalId": id,
	})

	c.JSON(200, gin.H{"message": "Token válido!", "username": username})
}
