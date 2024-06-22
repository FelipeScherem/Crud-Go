// Package controllerProdutos proteto400/src/Api/Controllers/ControllerProdutos/ControllerDeletarProdutos.go
package controllerProdutos

import (
	"fmt"
	"github.com/gin-gonic/gin"
	repositorysProdutos "projeto404/src/Api/Repositorys/RepositorysProdutos"
	util "projeto404/src/Api/Uteis"
)

func DeletarProduto(c *gin.Context) {

	// Declara as variavel de request
	var idsParaDeletar []int8
	var idParaDeletarRequest IdParaDeletarRequest

	// Faz bind
	if err := c.BindJSON(&idParaDeletarRequest); err != nil {
		c.JSON(404, gin.H{"error": "Houve um erro ao receber os dados", "detalhes": err.Error})
		return
	}

	// Mapeia os dados
	idsParaDeletar = idParaDeletarRequest.ids

	// Valida ids se int8
	for _, id := range idsParaDeletar {
		if !util.ValidaInt8(id) {
			c.JSON(404, gin.H{"error": fmt.Sprintf("ID inválido: %v não é do tipo int8", id)})
			return
		}
	}
	// Chama o repository que deleta
	err := repositorysProdutos.DeletarProduto(idsParaDeletar)
	if err != nil {
		c.JSON(404, gin.H{"error": "Houve um erro ao excluir os produtos", "detalhes": err.Error})
	}

	// Retorna
	c.JSON(204, gin.H{
		"mensagem": "Produtos Deletados",
	})
}

// IdParaDeletarRequest Recebe 1 ou mais Ids para deletar os produtos
type IdParaDeletarRequest struct {
	ids []int8 `json:"id"`
}
