package dbpostgres

//Melhor os erros no futuro
func TrataErroPostgres(code string) string {
	switch code {
	case "23505":
		return "Dado duplicado"
	case "23514":
		return "Dado não está de acordo com as regras estipuladas no banco"
	case "23503":
		return "A INSERÇÃO OU ATUALIZAÇÃO DO VALOR DA CHAVE ESTRANGEIRA É INVÁLIDA"
	}
	return ""
}
