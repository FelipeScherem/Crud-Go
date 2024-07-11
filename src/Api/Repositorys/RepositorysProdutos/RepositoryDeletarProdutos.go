package repositorysProdutos

import (
	db "projeto404/src/Api/Database"
)

// DeletarProduto Faz um Soft-Delete. No banco existe uma rotina onde excluir automaticamente após 30 dias
func DeletarProduto(ids []int8) error {
	database := db.ConectaDB() // Abre a conexão com o banco de dados
	defer db.FechaDB(database) // Fecha conexão com o banco de dados no final da função

	// Faz a operação no banco
	query := database.Where("id IN ?", ids).Update("data_de_cadastro", "")
	//Verifica se houve erro, se houve, retorna
	if query.Error != nil {
		return query.Error
	}
	// Se esta tudo ok, retorna
	return nil
}
