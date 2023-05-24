package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

type BTCPriceResponse struct {
	Bitcoin struct {
		UAH int `json:"uah"`
	} `json:"bitcoin"`
}

func extractPrice() (int, error) {
	response, err := http.Get("https://api.coingecko.com/api/v3/simple/price?ids=bitcoin&vs_currencies=uah")
	if err != nil {
		fmt.Println("Error when request executing ", err)
		return 0, err
	}
	defer response.Body.Close()

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error when reading the response body ", err)
		return 0, err
	}
	var requestRes BTCPriceResponse
	err = json.Unmarshal(responseBody, &requestRes)
	if err != nil {
		fmt.Println("Error when unmarshalling response body ", err)
		return 0, err
	}
	return requestRes.Bitcoin.UAH, nil
}

func RateHandler(writer http.ResponseWriter, request *http.Request) {
	price, err := extractPrice()
	if err != nil {
		fmt.Println("Something wrong when extracting price of BTC", err)
		writer.WriteHeader(400)
		return
	}
	writer.WriteHeader(200)
	writer.Write([]byte(strconv.Itoa(price)))
}
