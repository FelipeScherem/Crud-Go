// Package controllerProdutos projeto404/src/Api/Controllers/ControllerUsuarios/ControllerCriarProdutos.go
package controllerProdutos

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	produtoModel "projeto404/src/Api/Models/ModelProdutos"
	repositorysProdutos "projeto404/src/Api/Repositorys/RepositorysProdutos"
)

// CriarProduto Faz Bind, valida e chama Repository
func CriarProduto(c *gin.Context) {
	// Model e payload
	var produto produtoModel.ProdutoStruct
	var produtoRequest ProdutoRequest

	// Faz o bind do corpo da solicitação
	if err := c.BindJSON(&produtoRequest); err != nil {
		c.JSON(404, gin.H{"error": "Erro ao decodificar os dados JSON"})
		return
	}

	// ######### Verificações #########
	// Verifica se nome vazio
	if produtoRequest.Nome == "" {
		c.JSON(404, gin.H{"error": "O campo de nome não pode ser vazio"})
		return
	}
	// Verifica se preço é vazio
	if produtoRequest.Preco == 0.0 {
		c.JSON(404, gin.H{"error": "O campo de preço não pode ser vazio"})
		return
	}

	// Popula a struct produto com os dados de produtoRequest
	produto.Nome = produtoRequest.Nome
	produto.Descricao = produtoRequest.Descricao
	produto.Preco = produtoRequest.Preco
	produto.Estoque = produtoRequest.Estoque
	produto.Disponivel = produtoRequest.Disponivel
	produto.TipoDeDesconto = produtoRequest.TipoDeDesconto
	produto.Valor = produtoRequest.Valor

	// Converte arrays de strings para JSON
	categorias, err := json.Marshal(produtoRequest.Categorias)
	if err != nil {
		c.JSON(500, gin.H{"error": "Erro ao converter categorias para JSON"})
		return
	}
	imagens, err := json.Marshal(produtoRequest.Imagens)
	if err != nil {
		c.JSON(500, gin.H{"error": "Erro ao converter imagens para JSON"})
		return
	}
	tamanho, err := json.Marshal(produtoRequest.Tamanho)
	if err != nil {
		c.JSON(500, gin.H{"error": "Erro ao converter tamanho para JSON"})
		return
	}
	cor, err := json.Marshal(produtoRequest.Cor)
	if err != nil {
		c.JSON(500, gin.H{"error": "Erro ao converter cor para JSON"})
		return
	}

	produto.Categorias = string(categorias)
	produto.Imagens = string(imagens)
	produto.Tamanho = string(tamanho)
	produto.Cor = string(cor)

	// Chama o repository para fazer a inserção
	err = repositorysProdutos.CriarProduto(produto)
	if err != nil {
		c.JSON(404, gin.H{"error": "Houve um erro na inserção dos dados no banco de dados",
			"detalhes": err.Error()})
		return
	}

	// Se o produto foi criado corretamente, retorna um ok
	c.JSON(201, gin.H{
		"mensagem": "Produto criado",
	})
	return
}

// ProdutoRequest Struct de request
type ProdutoRequest struct {
	Nome           string   `json:"nome"`
	Descricao      string   `json:"descricao"`
	Preco          float64  `json:"preco"`
	TipoDeDesconto string   `json:"tipo_de_desconto"`
	Valor          float64  `json:"desconto"`
	Estoque        int      `json:"estoque"`
	Disponivel     bool     `json:"disponivel"`
	Categorias     []string `json:"categorias"`
	Imagens        []string `json:"imagens"`
	Tamanho        []string `json:"tamanho"`
	Cor            []string `json:"cor"`
}
