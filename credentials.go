package main

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strings"
)

type Credentials struct {
	EmailAddress string
	EmailPassword string
	SmtpHost string
	SmtpPort string
	Recipients []string
}

// GetCredentials is used for loading all credentials needed from .env file
func GetCredentials() *Credentials {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("unable to load .env file")
	}

	emailAddress := os.Getenv("EMAIL_FROM_ADDRESS")
	emailPassword := os.Getenv("PASSWORD_FROM_ADDRESS")

	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")

	recipients := strings.Split(os.Getenv("RECIPIENTS"), ",")

	return &Credentials{
		EmailAddress: emailAddress,
		EmailPassword: emailPassword,
		SmtpHost: smtpHost,
		SmtpPort: smtpPort,
		Recipients: recipients,
	}
}
