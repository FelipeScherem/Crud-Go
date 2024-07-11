package repositoryUsuarios

import (
	"github.com/jackc/pgx/v5/pgconn"
	db "projeto404/src/Api/Database"
	modelUsuario "projeto404/src/Api/Models/ModelUsers"
	"strings"
)

// AtualizarUsuario insere os dados do cliente no banco
func AtualizarUsuario(usuarioStruct modelUsuario.UsuarioStruct, idUsuario int) (string, error) {
	database := db.ConectaDB() // Abre a conexão com o banco de dados
	defer db.FechaDB(database) // Fecha conexão com o banco de dados no final da função

	// Tenta fazer a inserção, se houver dados repetidos, ele retorna com os campos repetidos
	result := database.
		Model(&modelUsuario.UsuarioStruct{}).
		Where("id = ?", idUsuario).
		Updates(&usuarioStruct)
	if err := result.Error; err != nil {
		// Verifica se houve erro de violação de unicidade (unique_violation)
		if pgErr, ok := err.(*pgconn.PgError); ok && pgErr.Code == "23505" {
			camposRepetidos := "Campos repetidos: "
			if strings.Contains(pgErr.Detail, "email") {
				camposRepetidos += "email, "
			}
			if strings.Contains(pgErr.Detail, "telefone") {
				camposRepetidos += "telefone, "
			}
			if strings.Contains(pgErr.Detail, "cpf") {
				camposRepetidos += "cpf, "
			}
			if strings.Contains(pgErr.Detail, "cnpj") {
				camposRepetidos += "cnpj, "
			}
			// Remove a vírgula extra no final, se houver
			camposRepetidos = strings.TrimSuffix(camposRepetidos, ", ")

			return camposRepetidos, err
		}
		return "", err
	}

	// Se esta tudo ok, retorna
	return "", nil
}
