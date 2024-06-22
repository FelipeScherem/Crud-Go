package modelProduto

type ProdutoStruct struct {
	Nome            string     `gorm:"nome"`
	Descricao       string     `gorm:"descricao"`
	Preco           float64    `gorm:"preco"`
	Desconto        []struct{} `gorm:"Desconto"`
	Estoque         int        `gorm:"estoque"`
	Disponivel      bool       `gorm:"disponivel"`
	Categorias      []string   `gorm:"categorias"`
	Imagens         []string   `gorm:"imagens"`
	Tamanho         []string   `gorm:"tamanho"`
	Cor             []string   `gorm:"cor"`
	DataCriacao     string     `gorm:"data_criacao"`
	DataAtualizacao string     `gorm:"data_atualizacao"`
}

// Struct para representar uma variante do produto (combinação de tamanho e cor)
type Desconto struct {
	TipoDeDesconto string  `gorm:"tipo_de_desconto"`
	Desconto       float64 `gorm:"desconto"`
}
