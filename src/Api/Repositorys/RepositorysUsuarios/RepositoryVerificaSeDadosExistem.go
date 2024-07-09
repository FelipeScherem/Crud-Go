package repositoryUsuarios

import (
	"fmt"
	db "projeto404/src/Api/Uteis"
)

// verificaSeDadosExistem Verifica se os seguintes dados já estão no banco e retorna um array com os dados repetidos
//
//	CPF
//	CNPJ
//	Email
//	Telefone
func VerificaSeDadosExistem(CPF string, CNPJ string, Email string, Telefone int) ([]string, error) {
	database := db.ConectaDB() // Abre a conexão com o banco de dados
	defer db.FechaDB(database) // Fecha conexão com o banco de dados no final da função

	var camposDuplicados []string
	var registros int64

	// Verifica duplicidade do cpf
	if err := database.
		Where("cpf = ?", CPF).
		Count(&registros).Error; err != nil {
		return nil, fmt.Errorf("erro ao verificar duplicidade do CPF: %v", err)
	}
	if registros > 0 {
		camposDuplicados = append(camposDuplicados, "CPF")
	}

	// Verifica duplicidade do cnpj
	if err := database.
		Where("cnpj = ?", CNPJ).
		Count(&registros).Error; err != nil {
		return nil, fmt.Errorf("erro ao verificar duplicidade do CNPJ: %v", err)
	}
	if registros > 0 {
		camposDuplicados = append(camposDuplicados, "CNPJ")
	}

	// Verifica duplicidade do email
	if err := database.
		Where("email = ?", Email).
		Count(&registros).Error; err != nil {
		return nil, fmt.Errorf("erro ao verificar duplicidade do Email: %v", err)
	}
	if registros > 0 {
		camposDuplicados = append(camposDuplicados, "Email")
	}

	// Verifica duplicidade do telefone
	if err := database.
		Where("telefone = ?", Telefone).
		Count(&registros).Error; err != nil {
		return nil, fmt.Errorf("erro ao verificar duplicidade do Telefone: %v", err)
	}
	if registros > 0 {
		camposDuplicados = append(camposDuplicados, "Telefone")
	}

	return camposDuplicados, nil
}
