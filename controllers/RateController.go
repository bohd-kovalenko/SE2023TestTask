package controllers

import (
	"SE2023/services"
	"fmt"
	"net/http"
	"strconv"
)

const OkStatus = 200
const BadRequestStatus = 400

func RateHandler(writer http.ResponseWriter, _ *http.Request) {
	price, err := services.ExtractPrice()
	if err != nil {
		fmt.Println("Something wrong when extracting price of BTC", err)
		writer.WriteHeader(BadRequestStatus)
		return
	}
	writer.WriteHeader(OkStatus)
	_, _ = writer.Write([]byte(strconv.Itoa(price)))
}
