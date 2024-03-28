package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type BalanceResponse struct {
	Total int    `json:"total"`
	Date  string `json:"balance"`
	Limit string `json:"limit"`
}

type TransactionResponse struct {
	Value       int
	Type        string
	Description string
	Date        string
}

type Response struct {
	Balance             BalanceResponse       `json:"balance"`
	TransactionResponse []TransactionResponse `json:"lastTransactions"`
}
type Balance struct {
	Limit   int    `json:"limit"`
	Balance string `json:"balance"`
}

type TransactionRequest struct {
	Value       int    `json:"value" binding:"required"`
	Type        string `json:"type" binding:"required"`
	Description string `json:"description" binding:"required"`
}

func main() {
	server := routers()
	server.Run(":8080")
}

func routers() *gin.Engine {
	ginServer := gin.Default()
	ginServer.GET("/clientes/:id/extrato", func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(http.StatusOK, id)
	})

	ginServer.POST("/clientes/:id/transacoes", func(c *gin.Context) {
		var request TransactionRequest
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusUnprocessableEntity, err.Error())
			return
		}

	})

	return ginServer
}
