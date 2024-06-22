// Package controllerProdutos proteto400/src/Api/Controllers/ControllerUsuarios/ControllerListarProdutos.go
package controllerProdutos

import (
	"github.com/gin-gonic/gin"
)

func ListarProduto(c *gin.Context) {
	c.JSON(400, gin.H{
		"mensagem": "Lista de Produtos",
	})
}
