package util

import (
	"bytes"
	"fmt"
	"html/template"
	"io/ioutil"
	"strconv"
	"time"

	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/config"
	dtousecase "git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/dto/usecase"
	"gopkg.in/gomail.v2"
)

func SendMail(payload dtousecase.SendEmailPayload) error {
	data := struct {
		Link      string
		Recipient string
		ExpiresAt time.Time
		Email     string
	}{
		Link:      fmt.Sprintf("%v%v", config.GetEnv("CLIENT_RESET_PASSWORD_URL"), payload.Token),
		Recipient: payload.RecipientName,
		ExpiresAt: payload.ExpiresAt,
		Email:     payload.RecipientEmail,
	}

	htmlTemplatePath := "template_forget_password.html"
	htmlTemplate, err := ioutil.ReadFile(htmlTemplatePath)
	if err != nil {
		return err
	}

	tmpl, err := template.New("emailTemplate").Parse(string(htmlTemplate))
	if err != nil {
		return err
	}

	var tplBuffer bytes.Buffer
	if err := tmpl.Execute(&tplBuffer, data); err != nil {
		return err
	}

	emailBody := tplBuffer.String()

	mailer := gomail.NewMessage()
	mailer.SetHeader("From", config.GetEnv("EMAIL_SENDER_NAME"))
	mailer.SetHeader("To", payload.RecipientEmail)
	mailer.SetAddressHeader("Cc", "tralalala@gmail.com", "Tra Lala La")
	mailer.SetHeader("Subject", "Request Forget Password")
	mailer.SetBody("text/html", emailBody)

	port, err := strconv.Atoi(config.GetEnv("EMAIL_SMTP_PORT"))
	if err != nil {
		return err
	}

	dialer := gomail.NewDialer(
		config.GetEnv("EMAIL_SMTP_HOST"),
		port,
		config.GetEnv("EMAIL_AUTH"),
		config.GetEnv("EMAIL_AUTH_PASSWORD"),
	)

	err = dialer.DialAndSend(mailer)
	if err != nil {
		return err
	}

	return nil
}
