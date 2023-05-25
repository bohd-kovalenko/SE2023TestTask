package controllers

import (
	"SE2023/repositories"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const ConflictStatus = 409

type EmailRequest struct {
	Email string `json:"email"`
}

func SubscriptionNewMailHandler(writer http.ResponseWriter, request *http.Request) {
	var emailRequest EmailRequest
	defer request.Body.Close()
	body, err := io.ReadAll(request.Body)
	if err != nil {
		fmt.Println("Error, when reading request body ", err)
		return
	}
	err = json.Unmarshal(body, &emailRequest)
	if err != nil {
		fmt.Println("Error, when unmarshalling request body ", err)
		return
	}
	isExist, err := repositories.FindEmailInFile(emailRequest.Email)
	if err != nil {
		fmt.Println("Error, when reading email base ", err)
		return
	}
	if isExist {
		fmt.Println("Email already exists in email base")
		writer.WriteHeader(ConflictStatus)
		return
	}
	err = repositories.SaveEmailToFile(emailRequest.Email)
	if err != nil {
		if err == repositories.ErrInvalidEmail {
			fmt.Println("Invalid email ", err)
			writer.WriteHeader(BadRequestStatus)
			return
		} else {
			fmt.Println("Error, when saving email to file ", err)
			writer.WriteHeader(InternalServerErrorStatus)
			return
		}
	}
	writer.WriteHeader(OkStatus)
}
