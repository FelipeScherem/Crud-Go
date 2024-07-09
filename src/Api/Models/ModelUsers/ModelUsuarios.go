package modelUsuario

import (
	"gorm.io/gorm"
	"time"
)

// UsuarioStruct Struct com os dados de usuarios
type UsuarioStruct struct {
	gorm.Model
	Nome                     string    `gorm:"column:nome"`
	Email                    string    `gorm:"column:email"`
	Telefone                 int       `gorm:"column:telefone"`
	Senha                    string    `gorm:"column:senha"`
	DataDeNascimento         time.Time `gorm:"column:dataDeNascimento"`
	Foto                     string    `gorm:"column:foto"`
	CidadeDeAtuacao          string    `gorm:"column:cidadeDeAtuacao"`
	ServicosPrestados        []string  `gorm:"column:tipoDeServico"`
	CPF                      string    `gorm:"column:cpf"`
	CNPJ                     string    `gorm:"column:cnpj"`
	NomeFantasia             string    `gorm:"column:nomeFantasia"`
	QuantidadeDeFuncionarios string    `gorm:"column:quantidadeDeFuncionarios"`
}

// TableName define o nome da tabela no banco de dados
func (UsuarioStruct) TableName() string {
	return "usuarios"
}
