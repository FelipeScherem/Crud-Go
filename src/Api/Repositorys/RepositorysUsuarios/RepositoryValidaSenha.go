package repositoryUsuarios

import (
	db "projeto404/src/Api/Database"
	modelUsuario "projeto404/src/Api/Models/ModelUsers"
)

// ValidarSenhaAntiga verifica se a senha antiga está correta para o usuário
func ValidarSenha(idUsuario int, senhaAtual string) (bool, error) {
	database := db.ConectaDB()
	defer db.FechaDB(database)

	var senha string
	query := database.Model(&modelUsuario.UsuarioStruct{}).
		Select("senha").
		Where("id = ?", idUsuario).
		Scan(&senha)
	if query.Error != nil {
		return false, query.Error
	}
	if senhaAtual != senha {
		return false, nil
	} else {
		return true, nil
	}
}
