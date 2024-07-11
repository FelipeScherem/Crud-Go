package util

import (
	"regexp"
	"strconv"
	"strings"
)

// ValidarSenha Valida se senha é segura
//
//	8 ou mais digitos
//	Ao menos 1 caracter maiusculo e minusculo
//	1 numeral
//	1 caracter especial
//
//	True se ok
//	False se há erros
func ValidarSenha(senha string) (string, bool) {
	// Verifica se tem pelo menos 8 caracteres
	if len(senha) < 8 {
		return "A senha deve conter ao menos 8 caracteres", false
	}

	// Verifica se tem pelo menos 1 caractere maiúsculo e 1 minúsculo
	contemMaiuscula := false
	contemMinuscula := false
	for _, char := range senha {
		if 'A' <= char && char <= 'Z' {
			contemMaiuscula = true
		}
		if 'a' <= char && char <= 'z' {
			contemMinuscula = true
		}
	}
	if !contemMaiuscula || !contemMinuscula {
		return "A senha deve conter ao menos uma letra maiuscula e uma minuscula", false
	}

	// Verifica se tem pelo menos 1 numeral
	contemNumeral := false
	for _, char := range senha {
		if '0' <= char && char <= '9' {
			contemNumeral = true
		}
	}
	if !contemNumeral {
		return "A senha deve conter ao menos um numeral", false
	}

	// Verifica se tem pelo menos 1 caracter especial
	regexEspecial := regexp.MustCompile(`[!@#$%^&*()_+{}[\]:;<>,.?/~]`)
	if !regexEspecial.MatchString(senha) {
		return "A senha deve conter ao menos um caracter especial", false
	}

	return "", true
}

// ValidarCPF Valida se senha é segura
//
//	True se ok
//	False se há erros
func ValidarCPF(cpf string) (string, bool) {
	// Remove caracteres não numéricos do CPF
	cpf = strings.ReplaceAll(cpf, ".", "")
	cpf = strings.ReplaceAll(cpf, "-", "")

	// Verifica se o CPF tem 11 dígitos
	if len(cpf) != 11 {
		return "CPF inválido - deve conter exatamente 11 dígitos", false
	}

	// Verifica se todos os caracteres são dígitos
	if _, err := strconv.Atoi(cpf); err != nil {
		return "CPF inválido - deve conter apenas números", false
	}

	// Verifica se todos os dígitos são iguais (caso especial de CPF inválido)
	digitosIguais := true
	for i := 1; i < len(cpf); i++ {
		if cpf[i] != cpf[0] {
			digitosIguais = false
			break
		}
	}
	if digitosIguais {
		return "CPF inválido - todos os dígitos são iguais", false
	}

	// Calcula os dígitos verificadores
	digitosVerificadores := cpf[len(cpf)-2:]

	// Converte os dígitos verificadores para inteiros
	dv1, _ := strconv.Atoi(string(digitosVerificadores[0]))
	dv2, _ := strconv.Atoi(string(digitosVerificadores[1]))

	// Calcula o primeiro dígito verificador
	sum := 0
	for i := 0; i < 9; i++ {
		digit, _ := strconv.Atoi(string(cpf[i]))
		sum += digit * (10 - i)
	}
	resto := sum % 11
	var primeiroDVValido bool
	if resto < 2 {
		primeiroDVValido = dv1 == 0
	} else {
		primeiroDVValido = dv1 == 11-resto
	}

	// Calcula o segundo dígito verificador
	sum = 0
	for i := 0; i < 10; i++ {
		digit, _ := strconv.Atoi(string(cpf[i]))
		sum += digit * (11 - i)
	}
	resto = sum % 11
	var segundoDVValido bool
	if resto < 2 {
		segundoDVValido = dv2 == 0
	} else {
		segundoDVValido = dv2 == 11-resto
	}

	// Verifica se os dígitos verificadores são válidos
	if !primeiroDVValido || !segundoDVValido {
		return "CPF inválido - dígitos verificadores incorretos", false
	}

	// CPF válido
	return "CPF válido", true
}

// ValidarCNPJ Valida se senha é segura
//
//	True se ok
//	False se há erros
func ValidarCNPJ(cnpj string) (string, bool) {
	// Remove caracteres não numéricos do CNPJ
	cnpj = strings.ReplaceAll(cnpj, ".", "")
	cnpj = strings.ReplaceAll(cnpj, "-", "")
	cnpj = strings.ReplaceAll(cnpj, "/", "")

	// Verifica se o CNPJ tem 14 dígitos
	if len(cnpj) != 14 {
		return "CNPJ inválido - deve conter exatamente 14 dígitos", false
	}

	// Verifica se todos os caracteres são dígitos
	if _, err := strconv.Atoi(cnpj); err != nil {
		return "CNPJ inválido - deve conter apenas números", false
	}

	// Calcula o primeiro dígito verificador
	sum := 0
	peso := 5
	for i := 0; i < 12; i++ {
		digit, _ := strconv.Atoi(string(cnpj[i]))
		sum += digit * peso
		peso--
		if peso < 2 {
			peso = 9
		}
	}
	resto := sum % 11
	var primeiroDV int
	if resto < 2 {
		primeiroDV = 0
	} else {
		primeiroDV = 11 - resto
	}
	if primeiroDV != int(cnpj[12]-'0') {
		return "CNPJ inválido - primeiro dígito verificador incorreto", false
	}

	// Calcula o segundo dígito verificador
	sum = 0
	peso = 6
	for i := 0; i < 13; i++ {
		digit, _ := strconv.Atoi(string(cnpj[i]))
		sum += digit * peso
		peso--
		if peso < 2 {
			peso = 9
		}
	}
	resto = sum % 11
	var segundoDV int
	if resto < 2 {
		segundoDV = 0
	} else {
		segundoDV = 11 - resto
	}
	if segundoDV != int(cnpj[13]-'0') {
		return "CNPJ inválido - segundo dígito verificador incorreto", false
	}

	// CNPJ válido
	return "CNPJ válido", true
}
