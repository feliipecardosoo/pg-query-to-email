package model

import (
	"time"
)

// User representa a estrutura de um usuário retornado pela API
type User struct {
	Nome, Email, Telefone string
	CreatedAt             time.Time
}

// APIResponse representa a estrutura da resposta da API, contendo os dados dos usuários e uma mensagem
type APIResponse struct {
	Data    []User `json:"data"`
	Message string `json:"message"`
}

// Attachment representa um arquivo a ser anexado no email, contendo o nome do arquivo e os dados em bytes
type Attachment struct {
	FileName string
	Data     []byte
}

// EmailConfig representa a configuração necessária para enviar um email, incluindo destinatários, assunto, corpo e anexos
type EmailConfig struct {
	To          []string
	Subject     string
	Body        string
	Attachments []Attachment
}
