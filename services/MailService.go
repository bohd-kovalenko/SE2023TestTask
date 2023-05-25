package services

import (
	"SE2023/repositories"
	"fmt"
	"net/smtp"
)

const SMTPServer = "smtp.gmail.com"
const SMTPPort = ":587"
const SMTPEmail = "higuanda.john@gmail.com"
const SMTPPassword = "thammbqgfzbefjzu"

const EmailSubject = "BTC Price"

func SendAllEmails() error {
	BTCPrice, err := ExtractPrice()
	if err != nil {
		fmt.Println("Error, when obtaining BTC price ", err)
		return err
	}
	emails, err := repositories.FindAllEmails()
	if err != nil {
		fmt.Println("Error when trying to find all emails ", err)
		return err
	}
	emailMessage := fmt.Sprintf("Subject: %s\r\n\r\n%d", EmailSubject, BTCPrice)
	err = smtp.SendMail(SMTPServer+SMTPPort,
		smtp.PlainAuth("higuanda.john@gmail.com", SMTPEmail, SMTPPassword, SMTPServer),
		SMTPEmail,
		emails,
		[]byte(emailMessage))
	if err != nil {
		fmt.Println("Error, when carrying email ", err)
		return err
	}
	return nil
}
