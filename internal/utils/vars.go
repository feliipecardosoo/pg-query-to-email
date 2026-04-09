package utils

// Variáveis globais relacionadas ao CSV, como nome do arquivo e header, para evitar hardcoding e facilitar manutenção.
var (
	CsvFileName = "users.csv"

	CsvHeader = []string{
		"Nome",
		"Email",
		"Telefone",
		"CreatedAt",
	}
)
