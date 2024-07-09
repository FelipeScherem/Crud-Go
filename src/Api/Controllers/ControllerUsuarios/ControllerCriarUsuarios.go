package controllerUsuarios

import (
	"net/http"
	modelUsuario "projeto404/src/Api/Models/ModelUsers"
	repositoryUsuarios "projeto404/src/Api/Repositorys/RepositorysUsuarios"
	"reflect"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// CriarUsuarios Valida e insere dados dos usuarios
func CriarUsuarios(c *gin.Context) {
	var usuarioRequest UsuarioRequest

	// Bind JSON request body para UsuarioRequest
	if err := c.ShouldBindJSON(&usuarioRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"mensagem": "Erro ao processar requisição JSON", "erro": err})
		return
	}

	// Chama a função para verificar se há algum dos campos vazio, pois para criar o usuario, nehum deve estar
	if validaSeCamposVazio(c, usuarioRequest) {
		return // A resposta JSON de erro já foi enviada pela função
	}

	// camposDuplicados é uma variavel que agrupa os dados repetidos
	camposDuplicados, err := repositoryUsuarios.VerificaSeDadosExistem(usuarioRequest.CPF, usuarioRequest.CNPJ, usuarioRequest.Email, usuarioRequest.Telefone)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"mensagem": "Houve um erro ao criar usuário", "erro": err.Error()})
		return
	} else if len(camposDuplicados) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"mensagem": "Campos duplicados encontrados", "campos": camposDuplicados})
		return
	}

	// Popula os dados da struct de inserção
	usuarioModel := modelUsuario.UsuarioStruct{
		Nome:                     usuarioRequest.Nome,
		Email:                    usuarioRequest.Email,
		Telefone:                 usuarioRequest.Telefone,
		Senha:                    usuarioRequest.Senha,
		Foto:                     usuarioRequest.Foto,
		CidadeDeAtuacao:          usuarioRequest.CidadeDeAtuacao,
		ServicosPrestados:        usuarioRequest.ServicosPrestados,
		CPF:                      usuarioRequest.CPF,
		CNPJ:                     usuarioRequest.CNPJ,
		NomeFantasia:             usuarioRequest.NomeFantasia,
		QuantidadeDeFuncionarios: usuarioRequest.QuantidadeDeFuncionarios,
		DataDeNascimento:         usuarioRequest.DataDeNascimento,
	}

	// Chama o repository para inserção
	err = repositoryUsuarios.CriarUsuario(usuarioModel)
	// Verifica se houve erro
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"mensagem": "Houve um erro ao criar usuario", "error": err})
	}

	// Retorna um ok
	c.JSON(http.StatusOK, gin.H{"mensagem": "Usuário criado com sucesso"})
}

// Valida se todos os campos foram enviados
//
//	True se houver campos vazios
//	False se estiverem todos preenchidos
func validaSeCamposVazio(c *gin.Context, usuarioRequest UsuarioRequest) bool {
	camposVazios := []string{} // Variável para agrupar nome de campos que estão vazios
	usuarioRequestReflect := reflect.ValueOf(usuarioRequest)

	// Percorre os campos do struct
	for i := 0; i < usuarioRequestReflect.NumField(); i++ {
		valorDoCampo := usuarioRequestReflect.Field(i)
		nomeDoCampo := usuarioRequestReflect.Type().Field(i).Tag.Get("json")
		if valorDoCampo.Kind() == reflect.String && valorDoCampo.Len() == 0 {
			if nomeDoCampo != "cpf" && nomeDoCampo != "cnpj" { // Ignorar CPF e CNPJ nesta verificação
				camposVazios = append(camposVazios, nomeDoCampo)
			}
		}
	}

	// Verificar se pelo menos um dos campos CPF ou CNPJ está preenchido
	if usuarioRequest.CPF == "" && usuarioRequest.CNPJ == "" {
		camposVazios = append(camposVazios, "cpf ou cnpj")
	}

	// Construir mensagem de erro se houver campos vazios
	if len(camposVazios) > 0 {
		camposStr := strings.Join(camposVazios, ", ")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Os campos " + camposStr + " não podem estar vazios"})
		return true
	}

	return false
}

// UsuarioRequest Docs swagger
type UsuarioRequest struct {
	Nome                     string    `json:"nome"`
	Email                    string    `json:"email"`
	Telefone                 int       `json:"telefone"`
	Senha                    string    `json:"senha"`
	Foto                     string    `json:"foto"`
	CidadeDeAtuacao          string    `json:"cidadeDeAtuacao"`
	ServicosPrestados        []string  `json:"tipoDeServico"`
	CPF                      string    `json:"cpf"`
	CNPJ                     string    `json:"cnpj"`
	NomeFantasia             string    `json:"nomeFantasia"`
	QuantidadeDeFuncionarios string    `json:"quantidadeDeFuncionarios"`
	DataDeNascimento         time.Time `json:"dataDeNascimento"`
}
