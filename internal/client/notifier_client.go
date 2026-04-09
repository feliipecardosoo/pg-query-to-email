package client

import (
	"context"
	"net/http"
	"os"
)

// NotifierClient é responsável por fazer a requisição HTTP para a API externa, utilizando as variáveis de ambiente para configurar a URL e a chave de autenticação.
type NotifierClient struct {
	baseURL string
	apiKey  string
}

// NewNotifierClient é um construtor para NotifierClient, que lê as variáveis de ambiente e retorna uma instância configurada do cliente.
func NewNotifierClient() *NotifierClient {
	return &NotifierClient{
		baseURL: os.Getenv("API_URL"),
		apiKey:  os.Getenv("NOTIFIER_API_KEY"),
	}
}

// SendRequest é o método que realiza a requisição HTTP para a API, utilizando o contexto para controle de timeout e cancelamento, e retornando a resposta ou um erro caso ocorra algum problema durante a requisição.
func (c *NotifierClient) SendRequest(ctx context.Context) (*http.Response, error) {

	req, err := http.NewRequestWithContext(
		ctx,
		"POST",
		c.baseURL+"/api/notifier/send",
		nil, // 👈 sem body
	)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+c.apiKey)

	client := &http.Client{}
	return client.Do(req)
}
