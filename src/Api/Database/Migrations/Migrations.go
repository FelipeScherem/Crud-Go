package Migrations

import (
	"gorm.io/gorm"
	modelUsuario "projeto404/src/Api/Models/ModelUsers"
)

// RodarMigrations recebe o banco de dados e cria as tabelas
func RodarMigrations(db *gorm.DB) (string, error) {
	if err := db.AutoMigrate(&modelUsuario.UsuarioStruct{}); err != nil {
		return "Houve um erro ao criar tabela usuario", err
	}

	return "Tabelas atualizadas com sucesso", nil
}
