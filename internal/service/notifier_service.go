package service

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"

	"pg-query-to-email/internal/client"
	"pg-query-to-email/internal/email"
	"pg-query-to-email/internal/model"
	"pg-query-to-email/internal/utils"
)

// NotifierService é responsável por orquestrar todo o processo: chamar a API, processar os dados, gerar o CSV e enviar o email.
type NotifierService struct {
	client *client.NotifierClient
}

// NewNotifierService é um construtor para NotifierService, recebendo um cliente de API como dependência.
func NewNotifierService(c *client.NotifierClient) *NotifierService {
	return &NotifierService{client: c}
}

// Execute é o método principal que executa toda a lógica do serviço, seguindo os passos descritos.
func (s *NotifierService) Execute(ctx context.Context) error {

	// Chama API
	resp, err := s.client.SendRequest(ctx)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// Converter JSON → struct
	var apiResp model.APIResponse
	if err := json.Unmarshal(bodyBytes, &apiResp); err != nil {
		return err
	}
	users := apiResp.Data

	// Gerar CSV em memória
	csvBytes, err := utils.GenerateCSVInMemory(users)
	if err != nil {
		return err
	}

	// Montar email
	emailConfig := model.EmailConfig{
		To:      []string{os.Getenv("EMAIL_TO")},
		Subject: "Relatório de usuários",
		Body:    "Segue em anexo o relatório em CSV 📎",
		Attachments: []model.Attachment{
			{
				FileName: utils.CsvFileName,
				Data:     csvBytes,
			},
		},
	}

	// Enviar email
	sender := email.NewEmailSender(
		os.Getenv("EMAIL_FROM"),
		os.Getenv("EMAIL_PASSWORD"),
	)

	if err := sender.Send(emailConfig); err != nil {
		return err
	}

	fmt.Println("Email enviado com sucesso")

	return nil
}
