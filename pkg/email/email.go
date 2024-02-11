package email

import (
	"bytes"
	"fmt"
	"html/template"
	"net/smtp"
	"os"
)

// Request struct

type Request struct {
	to, subject, body string
}

func NewRequest(to []string, subject, body string) *Request {
	return &Request{
		to:      to[0],
		subject: subject,
		body:    body,
	}
}

func (request *Request) SendEmail() (bool, error) {
	// Sender and recipient email addresses
	from := "go_be@gmail.com"

	// SMTP server configuration
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")
	smtpUsername := os.Getenv("SMTP_USER")
	smtpPassword := os.Getenv("SMTP_PASSWORD")

	auth := smtp.PlainAuth("", smtpUsername, smtpPassword, smtpHost)

	// Construct MIME message
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	msg := []byte("To: " + request.to + "\r\n" +
		"From: " + from + "\r\n" +
		"Subject: " + request.subject + "\r\n" +
		mime + "\r\n" +
		request.body)

	// Send email
	if err := smtp.SendMail(fmt.Sprintf("%s:%s", smtpHost, smtpPort), auth, from, []string{request.to}, msg); err != nil {
		return false, err
	}
	return true, nil
}

func (r *Request) ParseTemplate(templateFileName string, data interface{}) error {
	t, err := template.ParseFiles(templateFileName)
	if err != nil {
		return err
	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		return err
	}
	r.body = buf.String()
	return nil
}
