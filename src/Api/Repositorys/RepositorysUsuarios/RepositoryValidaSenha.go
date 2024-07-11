package repositoryUsuarios

import (
	db "projeto404/src/Api/Database"
	modelUsuario "projeto404/src/Api/Models/ModelUsers"
)

// ValidarSenhaAntiga verifica se a senha antiga está correta para o usuário
//
//	True se estiver ok
//	False se estiver errada
func ValidarSenha(idUsuario int, senhaAtual string) (bool, error) {
	database := db.ConectaDB() // Abre a conexão com o banco de dados
	defer db.FechaDB(database) // Fecha conexão com o banco de dados no final da função

	var senhaNoBanco string
	query := database.Model(&modelUsuario.UsuarioStruct{}).
		Select("senha").
		Where("id = ?", idUsuario).
		Scan(&senhaNoBanco)
	// Verifica se houve erro na query
	if query.Error != nil {
		return false, query.Error
	}
	//Verifica se senha informada é igual a senha do banco
	if senhaAtual != senhaNoBanco {
		return false, nil
	} else {
		return true, nil
	}
}
