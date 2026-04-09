# pg-query-to-email

Microserviço em Go que consulta uma API, gera um relatório CSV com os usuários retornados e envia esse relatório por e-mail.

## Visão geral

O serviço executa o seguinte fluxo:

1. Carrega variáveis de ambiente a partir de `internal/config/.env`.
2. Faz uma requisição HTTP POST para a API configurada em `API_URL` no caminho `/api/notifier/send`.
3. Autentica a requisição usando o token JWT/Bearer em `NOTIFIER_API_KEY`.
4. Recebe a resposta JSON com a lista de usuários.
5. Converte os dados em um arquivo CSV em memória.
6. Envia um e-mail com o CSV em anexo usando SMTP.

## Estrutura do projeto

- `cmd/app/main.go` - ponto de entrada da aplicação.
- `internal/config/env.go` - carrega variáveis de ambiente.
- `internal/client/notifier_client.go` - faz a requisição HTTP para a API externa.
- `internal/service/notifier_service.go` - orquestra o fluxo de consulta, geração de CSV e envio de e-mail.
- `internal/email/sender.go` - constrói e envia o e-mail via SMTP.
- `internal/model/models.go` - definições de modelos de dados usados pela aplicação.
- `internal/utils/csv.go` - geração de CSV a partir dos dados de usuários.

## Variáveis de ambiente

A configuração do serviço depende das seguintes variáveis:

- `API_URL` - URL base da API que fornece os dados de usuários.
- `NOTIFIER_API_KEY` - token de autenticação para a API.
- `EMAIL_FROM` - endereço de e-mail remetente para envio via SMTP.
- `EMAIL_PASSWORD` - senha ou app password do SMTP do remetente.
- `EMAIL_TO` - endereço de e-mail de destino.

> Observação: o código atual carrega o arquivo de ambiente em `internal/config/.env`. Se quiser usar um `.env` na raiz do projeto, adapte o caminho em `internal/config/env.go`.

## Requisitos

- Go 1.25+
- Acesso à API em `API_URL`.
- Conta de e-mail configurada para envio SMTP.

## Como executar

1. Copie ou configure o arquivo `.env` em `internal/config/.env` com os valores corretos.
2. No diretório do projeto, execute:

```bash
go run ./cmd/app
```

## Executar com Docker

Caso queira rodar em container, use o `dockerfile` do projeto.

1. Monte a imagem:

```bash
docker build -t pg-query-to-email .
```

2. Execute a imagem passando as variáveis de ambiente ou com um arquivo `.env`:

```bash
docker run --rm --env-file internal/config/.env pg-query-to-email
```

## Observações importantes

- O envio de e-mail usa `smtp.gmail.com:587` no código atual.
- Para Gmail, é recomendado usar uma senha de app ou habilitar o acesso de apps menos seguros conforme as políticas da conta.
- Verifique se a API em `API_URL` aceita o método POST no caminho `/api/notifier/send` e retorna um JSON compatível com o modelo:

```json
{
  "data": [
    {
      "Nome": "Fulano",
      "Email": "fulano@example.com",
      "Telefone": "99999-9999",
      "CreatedAt": "2024-01-01T00:00:00Z"
    }
  ],
  "message": ""
}
```
