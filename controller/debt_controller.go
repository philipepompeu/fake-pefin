package controller

import (
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/philipepompeu/fake-pefin/service"
)

func DebtHandler(c *gin.Context) {
	body, err := io.ReadAll(c.Request.Body)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Erro ao ler body"})
		return
	}

	id := service.SaveRawDebt(string(body))

	c.JSON(http.StatusOK, gin.H{
		"message":         "Debt armazenado com sucesso",
		"transactionalId": id,
	})
}
