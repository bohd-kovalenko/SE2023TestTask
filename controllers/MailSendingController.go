package controllers

import (
	"SE2023/services"
	"fmt"
	"net/http"
)

const InternalServerErrorStatus = 500

func MailSendingHandler(writer http.ResponseWriter, _ *http.Request) {
	err := services.SendAllEmails()
	if err != nil {
		fmt.Println("Error, when tried to send all emails ", err)
		writer.WriteHeader(InternalServerErrorStatus)
		return
	}
	writer.WriteHeader(OkStatus)
}
