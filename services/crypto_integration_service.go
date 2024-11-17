package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

type CryptoService interface {
	GetCryptoData() ([]byte, error)
	CalculateTotalValue(response []byte, quantities map[string]float64) (float64, error)
}

// Define the structure of the JSON response
type Quote struct {
	USD struct {
		Price float64 `json:"price"`
	} `json:"USD"`
}

type CryptoData struct{}

// GetCryptoData fetches cryptocurrency data from the API
func (c *CryptoData) GetCryptoData() ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://pro-api.coinmarketcap.com/v1/cryptocurrency/quotes/latest", nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	q := url.Values{}
	q.Add("symbol", "RAY,BONK,SOL,ADA,ETH,COTI,ALU,DOGE,JUP,WEN,LINK,CAKE,DOT,BNB,CHZ")
	req.Header.Set("Accepts", "application/json")
	req.Header.Add("X-CMC_PRO_API_KEY", os.Getenv("CMC_KEY"))
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request to server: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}

	return body, nil
}

// CalculateTotalValue calculates the total value of all cryptocurrencies in USD
// taking into account the quantity of each symbol.
func (c *CryptoData) CalculateTotalValue(response []byte, quantities map[string]float64) (float64, error) {
	var data struct {
		Data map[string]struct {
			Quote Quote `json:"quote"`
		} `json:"data"`
	}

	if err := json.Unmarshal(response, &data); err != nil {
		return 0, fmt.Errorf("error unmarshaling JSON: %v", err)
	}

	var totalValue float64
	for symbol, crypto := range data.Data {
		quantity, exists := quantities[symbol]
		if !exists {
			continue // Skip symbols not present in the quantities map
		}
		totalValue += crypto.Quote.USD.Price * quantity
	}

	return totalValue, nil
}
