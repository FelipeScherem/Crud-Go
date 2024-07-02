// Package controllerProdutosusuario projeto404/src/Api/Controllers/ControllerProdutos/ControllerAtualizarProdutos.go
package controllerProdutos

import (
	"github.com/gin-gonic/gin"
)

func AtualizarProduto(c *gin.Context) {
	c.JSON(400, gin.H{
		"mensagem": "Produto Atualizado",
	})
}
