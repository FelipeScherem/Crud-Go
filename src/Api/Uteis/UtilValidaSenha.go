package util

import "regexp"

// ValidaSenha Valida se senha é segura
//
//	8 ou mais digitos
//	Ao menos 1 caracter maiusculo e minusculo
//	1 numeral
//	1 caracter especial
func ValidaSenha(password string) bool {
	// Verifica se tem pelo menos 8 caracteres
	if len(password) < 8 {
		return false
	}

	// Verifica se tem pelo menos 1 caractere maiúsculo e 1 minúsculo
	hasUpper := false
	hasLower := false
	for _, char := range password {
		if 'A' <= char && char <= 'Z' {
			hasUpper = true
		}
		if 'a' <= char && char <= 'z' {
			hasLower = true
		}
	}
	if !hasUpper || !hasLower {
		return false
	}

	// Verifica se tem pelo menos 1 numeral
	hasDigit := false
	for _, char := range password {
		if '0' <= char && char <= '9' {
			hasDigit = true
		}
	}
	if !hasDigit {
		return false
	}

	// Verifica se tem pelo menos 1 caracter especial
	regexSpecial := regexp.MustCompile(`[!@#$%^&*()_+{}[\]:;<>,.?/~]`)
	if !regexSpecial.MatchString(password) {
		return false
	}

	return true
}
