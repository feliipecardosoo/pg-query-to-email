package email

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"net/smtp"
	"strings"

	"pg-query-to-email/internal/model"
)

// EmailSender é responsável por enviar emails, construindo a mensagem com base na configuração recebida e utilizando o protocolo SMTP para envio.
type EmailSender struct {
	from     string
	password string
}

// NewEmailSender é um construtor para EmailSender, recebendo o email de origem e a senha como parâmetros, que serão utilizados para autenticação no servidor SMTP.
func NewEmailSender(from, password string) EmailSender {
	return EmailSender{
		from:     from,
		password: password,
	}
}

// Send envia o email com base na configuração recebida
func (e *EmailSender) Send(config model.EmailConfig) error {

	msg := e.buildMessage(config)

	auth := smtp.PlainAuth("", e.from, e.password, "smtp.gmail.com")

	return smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		e.from,
		config.To,
		msg.Bytes(),
	)
}

// buildMessage constrói a mensagem do email, incluindo headers, corpo e anexos, seguindo o formato MIME multipart.
func (e *EmailSender) buildMessage(config model.EmailConfig) *bytes.Buffer {

	boundary := "my-boundary-779"
	var msg bytes.Buffer

	// Headers
	msg.WriteString(fmt.Sprintf("From: %s\r\n", e.from))
	msg.WriteString(fmt.Sprintf("To: %s\r\n", strings.Join(config.To, ",")))
	msg.WriteString(fmt.Sprintf("Subject: %s\r\n", config.Subject))
	msg.WriteString("MIME-Version: 1.0\r\n")
	msg.WriteString("Content-Type: multipart/mixed; boundary=" + boundary + "\r\n\r\n")

	// Body
	msg.WriteString(e.buildBody(boundary, config.Body))

	// Attachments
	for _, att := range config.Attachments {
		msg.WriteString(e.buildAttachment(boundary, att))
	}

	msg.WriteString("--" + boundary + "--")

	return &msg
}

// buildBody constrói a parte do corpo do email seguindo o formato MIME multipart.
func (e *EmailSender) buildBody(boundary, body string) string {
	return fmt.Sprintf(
		"--%s\r\nContent-Type: text/html; charset=\"UTF-8\"\r\n\r\n%s\r\n",
		boundary,
		body,
	)
}

// buildAttachment constrói a parte do anexo do email, codificando o arquivo em base64 e seguindo o formato MIME multipart.
func (e *EmailSender) buildAttachment(boundary string, att model.Attachment) string {

	encodedFile := base64.StdEncoding.EncodeToString(att.Data)

	return fmt.Sprintf(
		"--%s\r\nContent-Type: text/csv\r\nContent-Transfer-Encoding: base64\r\nContent-Disposition: attachment; filename=\"%s\"\r\n\r\n%s\r\n",
		boundary,
		att.FileName,
		encodedFile,
	)
}
