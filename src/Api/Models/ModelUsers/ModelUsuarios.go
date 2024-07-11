package modelUsuario

import (
	"gorm.io/gorm"
	"time"
)

// UsuarioStruct Struct com os dados de usuarios
type UsuarioStruct struct {
	gorm.Model
	ID                       uint      `gorm:"primaryKey;autoIncrement"`
	Nome                     string    `gorm:"column:nome"`
	Email                    string    `gorm:"column:email;unique"`
	Telefone                 int       `gorm:"column:telefone;unique"`
	Senha                    string    `gorm:"column:senha"`
	DataDeNascimento         time.Time `gorm:"column:dataDeNascimento"`
	Foto                     string    `gorm:"column:foto"`
	CidadeDeAtuacao          string    `gorm:"column:cidadeDeAtuacao"`
	ServicosPrestados        string    `gorm:"column:tipoDeServico"`
	CPF                      string    `gorm:"column:cpf;unique"`
	CNPJ                     string    `gorm:"column:cnpj;unique"`
	RazaoSocial              string    `gorm:"column:razaoSocial"`
	NomeFantasia             string    `gorm:"column:nomeFantasia"`
	QuantidadeDeFuncionarios string    `gorm:"column:quantidadeDeFuncionarios"`
	CreatedAt                time.Time
	UpdatedAt                time.Time
	DeletedAt                gorm.DeletedAt
}

// TableName define o nome da tabela no banco de dados
func (UsuarioStruct) TableName() string {
	return "usuarios"
}
