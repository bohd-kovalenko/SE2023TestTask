package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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
