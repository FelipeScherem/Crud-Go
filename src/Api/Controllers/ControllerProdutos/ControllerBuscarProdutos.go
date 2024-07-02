// Package controllerProdutos projeto404/src/Api/Controllers/ControllerUsuarios/ControllerBuscarProdutos.go
package controllerProdutos

import (
	"github.com/gin-gonic/gin"
)

func BuscarProduto(c *gin.Context) {
	c.JSON(400, gin.H{
		"mensagem": "Lista de Produtos",
	})
}
