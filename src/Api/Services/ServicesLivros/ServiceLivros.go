package services

import (
	modelLivros "projeto404/src/Api/Models/ModelLivros"
	db "projeto404/src/Api/Uteis"
)

func CriarLivro(livro *modelLivros.Livros) error {
	database := db.ConectaDB() // Abre a conexão com o banco de dados
	defer db.FechaDB(database) // Fecha conexão com o banco de dados no final da função

	query := database.Create(livro) // Faz a inserção 
	// Verifica se houve erro
	if query.Error != nil {
		return query.Error
	}

	// Se esta tudo ok, retorna 
	return nil
}
