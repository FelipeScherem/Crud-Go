// Package controllerUsuarios projeto404/src/Api/Controllers/ControllerUsuarios/ControllerDeletarUsuarios.go
package controllerUsuarios

import (
	"github.com/gin-gonic/gin"
	"net/http"
	modelUsuario "projeto404/src/Api/Models/ModelUsers"
	repositoryUsuarios "projeto404/src/Api/Repositorys/RepositorysUsuarios"
	"strconv"
)

func DeletarUsuarios(c *gin.Context) {

	var usuarioRequest DeletarUsuarioRequest

	// Recupera o id do usario do entpoint
	idDoUsuarioStr := c.Param("id")
	idDoUsuario, err := strconv.Atoi(idDoUsuarioStr) // Converte para INT
	// Verifica se houve erro na conversão
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"mensagem": "ID inválido", "erro": err})
		return
	}

	// Bind JSON request body para UsuarioRequest
	if err := c.ShouldBindJSON(&usuarioRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"mensagem": "Erro ao processar requisição JSON", "erro": err})
		return
	}

	// Valida se a senha esta correta para fazer as atualizações
	validarSenha, err := repositoryUsuarios.ValidarSenha(idDoUsuario, usuarioRequest.Senha)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err})
		return
	} else if !validarSenha {
		c.JSON(http.StatusBadRequest, gin.H{"mensagem": "Senha não confere"})
		return
	}

	// Popula a  struct com a senha
	usuarioModel := modelUsuario.UsuarioStruct{
		Senha: usuarioRequest.Senha,
	}

	// Chama o repository para inserção
	mensagemDeErro, err := repositoryUsuarios.DeletarUsuario(usuarioModel, idDoUsuario)
	// Verifica se houve erro
	if err != nil && mensagemDeErro != "" {
		c.JSON(http.StatusBadRequest, gin.H{"mensagem": mensagemDeErro, "error": err})
		return
	}

	// Retorna um ok
	c.JSON(http.StatusOK, gin.H{"mensagem": "Usuário deletado com sucesso"})
}

// AtualizarUsuarioRequest Docs swagger
type DeletarUsuarioRequest struct {
	Senha string `json:"senha"`
}
