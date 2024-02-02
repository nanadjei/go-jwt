package mailer

import (
	"os"
	"fmt"
	"bytes"
	"net/smtp"
	"io"
)

type Mailer interface {
	Send(to string, subject string, message string) error
}

type EmailTemplate interface {
   Execute(wr io.Writer, data interface{}) error
}

// Structure of the mail to be sent
type SMTPMailer struct {
	from string
	to string
	subject string
	message string
	smtpHost string 
	smtpPort string 
}

func NewSMTPMail() *SMTPMailer {
	return &SMTPMailer{
		smtpHost: os.Getenv("MAIL_HOST"),
		smtpPort: os.Getenv("MAIL_PORT"),
		from: os.Getenv("MAIL_HOST_EMAIL"),
	}
}

func (s *SMTPMailer) Send(to string, subject string, template EmailTemplate, data interface{}) error {

	// Authentication
	auth := smtp.PlainAuth("", os.Getenv("MAIL_USERNAME"), os.Getenv("MAIL_PASSWORD"), os.Getenv("MAIL_HOST"))

	var body bytes.Buffer
	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body.Write([]byte(fmt.Sprintf("Subject: %s \n%s\n\n", subject, mimeHeaders)))

	template.Execute(&body, data)

	newMailer := NewSMTPMail()

	// Send mail
	err := smtp.SendMail(newMailer.smtpHost+":"+newMailer.smtpPort, auth, newMailer.from, []string{to}, body.Bytes())

	if err != nil {
    println(err)
	}
	return err
}