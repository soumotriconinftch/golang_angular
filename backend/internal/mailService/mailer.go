package mailservice

import (
	"bytes"
	"html/template"
	"log"
	"path/filepath"
	"time"

	gomail "gopkg.in/mail.v2"
)

type MailService struct {
	From         string
	To           string
	Subject      string
	Body         string
	TemplateName string
	Data         interface{}
}

func (m *MailService) MailerFunc() error {
	start := time.Now()

	log.Println("[MAIL] init send")
	log.Printf("[MAIL] from=%s to=%s subject=%s", m.From, m.To, m.Subject)

	message := gomail.NewMessage()
	message.SetHeader("From", m.From)
	message.SetHeader("To", m.To)
	message.SetHeader("Subject", m.Subject)

	if m.TemplateName != "" {
		templatePath := filepath.Join("internal", "templates", m.TemplateName)
		log.Printf("[MAIL] loading template: %s", templatePath)

		t, err := template.ParseFiles(templatePath)
		if err != nil {
			log.Printf("[MAIL][ERROR] template parse failed: %v", err)
			return err
		}

		var body bytes.Buffer
		if err := t.Execute(&body, m.Data); err != nil {
			log.Printf("[MAIL][ERROR] template execute failed: %v", err)
			return err
		}

		message.SetBody("text/html", body.String())
		log.Println("[MAIL] template rendered")
	} else {
		message.SetBody("text/html", m.Body)
		log.Println("[MAIL] raw body used")
	}

	log.Println("[MAIL] connecting to smtp.mailtrap.io")

	dialer := gomail.NewDialer(
		"sandbox.smtp.mailtrap.io",
		587,
		"8c0c550675150a",
		"c226cea1e4b517",
	)

	if err := dialer.DialAndSend(message); err != nil {
		log.Printf("[MAIL][ERROR] send failed: %v", err)
		return err
	}
	log.Printf("[MAIL] send success (duration=%s)", time.Since(start))
	return nil

}
