package controllerUsuarios

import (
	"github.com/gin-gonic/gin"
	"net/http"
	modelUsuario "projeto404/src/Api/Models/ModelUsers"
	repositoryUsuarios "projeto404/src/Api/Repositorys/RepositorysUsuarios"
	"strconv"
	"strings"
)

// AtualizarUsuarios Valida e atualiza dados dos usuarios
func AtualizarUsuarios(c *gin.Context) {
	var usuarioRequest AtualizarUsuarioRequest

	// Recupera o id do usario do entpoint
	idDoUsuarioStr := c.Param("id")
	idDoUsuario, err := strconv.Atoi(idDoUsuarioStr) // Converte para INT
	// Verifica se houve erro na conversão
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"mensagem": "ID inválido", "erro": err})
		return
	}

	// Bind JSON request body para UsuarioRequest
	if err = c.ShouldBindJSON(&usuarioRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"mensagem": "Erro ao processar requisição JSON", "erro": err})
		return
	}

	// ############## VALIDAÇÕES ##############

	// Valida se a senha esta correta para fazer as atualizações
	validarSenha, err := repositoryUsuarios.ValidarSenha(idDoUsuario, usuarioRequest.Senha)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err})
		return
	}
	if !validarSenha {
		c.JSON(http.StatusBadRequest, gin.H{"mensagem": "Senha não confere"})
		return
	}

	// Se uma nova senha informada, entra na validação
	if usuarioRequest.NovaSenha != "" {
		usuarioRequest.Senha = usuarioRequest.NovaSenha
	}

	// Popula os dados da struct de inserção
	usuarioModel := modelUsuario.UsuarioStruct{
		Nome:                     usuarioRequest.Nome,
		Email:                    usuarioRequest.Email,
		Telefone:                 usuarioRequest.Telefone,
		Senha:                    usuarioRequest.Senha,
		Foto:                     usuarioRequest.Foto,
		CidadeDeAtuacao:          usuarioRequest.CidadeDeAtuacao,
		ServicosPrestados:        strings.Join(usuarioRequest.ServicosPrestados, ", "),
		RazaoSocial:              usuarioRequest.RazaoSocial,
		NomeFantasia:             usuarioRequest.NomeFantasia,
		QuantidadeDeFuncionarios: usuarioRequest.QuantidadeDeFuncionarios,
	}

	// Chama o repository para inserção
	mensagemDeErro, err := repositoryUsuarios.AtualizarUsuario(usuarioModel, idDoUsuario)
	// Verifica se houve erro
	if err != nil && mensagemDeErro != "" {
		c.JSON(http.StatusBadRequest, gin.H{"mensagem": mensagemDeErro, "error": err})
		return
	}

	// Retorna um ok
	c.JSON(http.StatusOK, gin.H{"mensagem": "Usuário atualizado com sucesso"})
}

// AtualizarUsuarioRequest Docs swagger
type AtualizarUsuarioRequest struct {
	Nome                     string   `json:"nome"`
	Email                    string   `json:"email"`
	Telefone                 int      `json:"telefone"`
	Senha                    string   `json:"senha"`
	NovaSenha                string   `json:"novaSenha"`
	Foto                     string   `json:"foto"`
	CidadeDeAtuacao          string   `json:"cidadeDeAtuacao"`
	ServicosPrestados        []string `json:"tipoDeServico"`
	RazaoSocial              string   `json:"razaoSocial"`
	NomeFantasia             string   `json:"nomeFantasia"`
	QuantidadeDeFuncionarios string   `json:"quantidadeDeFuncionarios"`
}
