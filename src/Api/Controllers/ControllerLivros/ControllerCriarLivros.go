// proteto400/src/Api/Controllers/ControllerLivros/ControllerCriarLivros.go
package controllerLivros

import (
	"github.com/gin-gonic/gin"
)

// DOCS DO SWAGGER
func CriarLivro(c *gin.Context) {
	// Variavel do payload
	var livro LivroStruct

	// Faz o bind do corpo da solicitação para a variável livro
	if err := c.BindJSON(&livro); err != nil {
		c.JSON(400, gin.H{"error": "Erro ao decodificar os dados JSON"})
		return
	}

	// Verifica se o nome não esta vazio
	if livro.Nome == "" {
		c.JSON(400, gin.H{"error": "O campo de nome não pode ser vazio"})
	}

    // Chama o repository para fazer a inserção 
	

    // Se o livro foi criado corretamente, retorna um ok
	c.JSON(200, gin.H{
		"mensagem": "Livro criado",
		"Livro":    livro.Nome,
	})
}

// Struct de request
type LivroStruct struct {
	Nome       string `json:"nome"`
	Autor      string `json:"autor"`
	Ano        string `json:"ano"`
	Lancamento string `json:"lancamento"`
}
