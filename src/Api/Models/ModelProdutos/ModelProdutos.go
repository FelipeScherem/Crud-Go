// Package modelProduto projeto404/src/Api/Models/ModelsProdutos/ModelsProdutos.go
package modelProduto

import (
	"gorm.io/gorm"
)

type ProdutoStruct struct {
	gorm.Model
	Nome           string  `gorm:"column:nome"`
	Descricao      string  `gorm:"column:descricao"`
	Preco          float64 `gorm:"column:preco"`
	TipoDeDesconto string  `gorm:"column:tipo_de_desconto"`
	Valor          float64 `gorm:"column:desconto"`
	Estoque        int     `gorm:"column:estoque"`
	Disponivel     bool    `gorm:"column:disponivel"`
	Categorias     string  `gorm:"type:jsonb;column:categorias"`
	Imagens        string  `gorm:"type:jsonb;column:imagens"`
	Tamanho        string  `gorm:"type:jsonb;column:tamanho"`
	Cor            string  `gorm:"type:jsonb;column:cor"`
}

// TableName define o nome da tabela no banco de dados
func (ProdutoStruct) TableName() string {
	return "produtos"
}
