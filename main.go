package main

import (
	"fmt"
	"log"
	"net/smtp"
)

func main() {

	// load secrets with environment vars with godotenv
	credentials := GetCredentials()

	// smtp config
	smtpServerConfig := NewSmtpServer(credentials.SmtpHost, credentials.SmtpPort)

	// authentication
	auth := smtp.PlainAuth("", credentials.EmailAddress, credentials.EmailPassword, smtpServerConfig.Host)

	// new email template
	body := NewEmailTemplate()

	// sending email
	err := smtp.SendMail(smtpServerConfig.Address(), auth, credentials.EmailAddress, credentials.Recipients, body.Bytes())

	if err != nil {
		log.Fatalf("unable to send the mail: %s", err)
	}

	fmt.Println("Email sent!")
}

