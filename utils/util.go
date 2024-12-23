package utils

import (
	"github.com/rs/zerolog/log"
	gomail "gopkg.in/gomail.v2"
)

func SendEmail(name, email, message string) {
	msg := gomail.NewMessage()
	msg.SetHeader("From", email)
	msg.SetHeader("To", "ricardo13work@gmail.com")
	msg.SetHeader("Subject", name+"Contacto desde la web")
	msg.SetBody("text/html", message)

	n := gomail.NewDialer("sandbox.smtp.mailtrap.io", 2525, "b157e23b893845", "77c91ed89fe266")
	if err := n.DialAndSend(msg); err != nil {
		log.Error().Msg("Error sending email")
	}

}
