package repositorysProdutos

import (
	db "projeto404/src/Api/Database"
	modelProdutos "projeto404/src/Api/Models/ModelProdutos"
)

// CriarProduto insere os dados do produto no banco e retorna
func CriarProduto(produto modelProdutos.ProdutoStruct) error {
	database := db.ConectaDB() // Abre a conexão com o banco de dados
	defer db.FechaDB(database) // Fecha conexão com o banco de dados no final da função

	query := database.Create(&produto) // Faz a inserção
	// Verifica se houve erro
	if query.Error != nil {
		return query.Error
	}

	// Se esta tudo ok, retorna
	return nil
}
