package main

import (
	"bytes"
	"fmt"
	"log"
	"text/template"
)

func NewEmailTemplate() bytes.Buffer {
	// Basic mail content
	mailContent := NewEmailContent("Daily Dose of Challenges and Cheats")

	// email template setup
	mailTemplate, _ := template.ParseFiles("template.html")

	var body bytes.Buffer

	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body.Write([]byte(fmt.Sprintf("Subject: %s \n%s\n\n", mailContent.Subject, mimeHeaders)))

	err := mailTemplate.Execute(&body, mailContent)

	if err != nil {
		log.Fatalf("unable to load HTML mail template correctly: %s", err)
	}

	return body
}
