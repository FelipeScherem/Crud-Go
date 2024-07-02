// Package controllerProdutos projeto404/src/Api/Controllers/ControllerProdutos/ControllerCriarProdutos.go
package controllerProdutos

import (
	"github.com/gin-gonic/gin"
	produtoModel "projeto404/src/Api/Models/ModelProdutos"
	repositorysProdutos "projeto404/src/Api/Repositorys/RepositorysProdutos"
)

// CriarProduto Faz Bind, valida e chama Repository
func CriarProduto(c *gin.Context) {
	// Model e payload
	var produto produtoModel.ProdutoStruct
	var produtoRequest ProdutoStruct

	// Faz o bind do corpo da solicitação
	if err := c.BindJSON(&produtoRequest); err != nil {
		c.JSON(404, gin.H{"error": "Erro ao decodificar os dados JSON"})
		return
	}

	// ######### Verificações #########
	// Se nome vazio
	if produtoRequest.Nome == "" {
		c.JSON(404, gin.H{"error": "O campo de nome não pode ser vazio"})
	}
	// Se preço vazio

	// Popula a struct produto com os dados de produtoRequest
	produto.Nome = produtoRequest.Nome
	produto.Descricao = produtoRequest.Descricao
	produto.Preco = produtoRequest.Preco
	produto.Estoque = produtoRequest.Estoque
	produto.Disponivel = produtoRequest.Disponivel
	produto.Categorias = produtoRequest.Categorias
	produto.Imagens = produtoRequest.Imagens
	produto.Tamanho = produtoRequest.Tamanho
	produto.Cor = produtoRequest.Cor
	produto.DataCriacao = produtoRequest.DataCriacao
	produto.DataAtualizacao = produtoRequest.DataAtualizacao
	// Popula os descontos (caso existam)
	produto.Desconto = produtoRequest.Desconto

	// Chama o repository para fazer a inserção
	err := repositorysProdutos.CriarProduto(produto)
	if err != nil {
		c.JSON(404, gin.H{"error": "Houve um erro na inserção dos dados no banco de dados",
			"detalhes": err.Error})
		return
	}

	// Se o produto foi criado corretamente, retorna um ok
	c.JSON(201, gin.H{
		"mensagem": "Produto criado",
	})
	return
}

// ProdutoStruct Struct de request
type ProdutoStruct struct {
	Nome            string     `json:"nome"`
	Descricao       string     `json:"descricao"`
	Preco           float64    `json:"preco"`
	Desconto        []struct{} `json:"Desconto"`
	Estoque         int        `json:"estoque"`
	Disponivel      bool       `json:"disponivel"`
	Categorias      []string   `json:"categorias"`
	Imagens         []string   `json:"imagens"`
	Tamanho         []string   `json:"tamanho"`
	Cor             []string   `json:"cor"`
	DataCriacao     string     `json:"data_criacao"`
	DataAtualizacao string     `json:"data_atualizacao"`
}

// Desconto Struct para representar uma variante do produto (combinação de tamanho e cor)
type Desconto struct {
	TipoDeDesconto string  `json:"tipo_de_desconto"`
	Desconto       float64 `json:"desconto"`
}
