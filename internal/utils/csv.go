package utils

import (
	"bytes"
	"encoding/csv"
	"pg-query-to-email/internal/model"
)

// GenerateCSVInMemory recebe uma lista de usuários e gera um arquivo CSV em memória, retornando os bytes do arquivo ou um erro caso ocorra algum problema durante a geração.
func GenerateCSVInMemory(users []model.User) ([]byte, error) {

	var buffer bytes.Buffer
	writer := csv.NewWriter(&buffer)

	// Header
	if err := writer.Write(CsvHeader); err != nil {
		return nil, err
	}

	// Dados
	for _, u := range users {
		err := writer.Write([]string{
			u.Nome,
			u.Email,
			u.Telefone,
			u.CreatedAt.Format("2006-01-02 15:04:05"),
		})
		if err != nil {
			return nil, err
		}
	}

	writer.Flush()

	if err := writer.Error(); err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}
