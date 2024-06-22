package util

// Função auxiliar para verificar se um valor é do tipo int8
func ValidaInt8(int interface{}) bool {
	switch int.(type) {
	case int8:
		return true
	default:
		return false
	}
}
