package repositoryUsuarios

import (
	db "projeto404/src/Api/Database"
	modelUsuario "projeto404/src/Api/Models/ModelUsers"
)

func VerificarSoftdelete(email string) (bool, error) {
	database := db.ConectaDB() // Abre a conexão com o banco de dados
	defer db.FechaDB(database) // Fecha conexão com o banco de dados no final da função

	var resultado modelUsuario.UsuarioStruct
	database.Unscoped().
		Where("email = ?", email).
		First(&resultado)

	if resultado.DeletedAt.Valid {
		return true, nil
	} else {
		return false, nil
	}

}
