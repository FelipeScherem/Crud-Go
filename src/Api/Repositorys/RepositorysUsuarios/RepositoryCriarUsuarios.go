package repositoryUsuarios

import (
	modelUsuario "projeto404/src/Api/Models/ModelUsers"
	db "projeto404/src/Api/Uteis"
)

// CriarUsuario insere os dados do cliente no banco
func CriarUsuario(usuarioStruct modelUsuario.UsuarioStruct) error {
	database := db.ConectaDB() // Abre a conexão com o banco de dados
	defer db.FechaDB(database) // Fecha conexão com o banco de dados no final da função

	// Faz a inserção
	query := database.Create(&usuarioStruct)
	// Verifica se houve erro
	if query.Error != nil {
		return query.Error
	}

	// Se esta tudo ok, retorna
	return nil
}
