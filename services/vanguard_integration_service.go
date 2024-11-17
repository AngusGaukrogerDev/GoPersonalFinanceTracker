package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

type FinancialDataService interface {
	// Get the current price of the Vanguard Fund in USD and GBP
	GetVanguardFundPrice() (stockPriceUSD float64, stockPriceGBP float64, currencyRate float64, err error)
}

// FinancialDataService implementation
type financialDataService struct {
	client *http.Client
	apiKey string
}

// NewFinancialDataService creates a new instance of the financial data service.
func NewFinancialDataService(apiKey string) FinancialDataService {
	return &financialDataService{
		client: &http.Client{},
		apiKey: apiKey,
	}
}

// GetVanguardFundPrice retrieves the current price of the Vanguard Fund in USD and GBP
func (f *financialDataService) GetVanguardFundPrice() (float64, float64, float64, error) {
	// URLs for the Vanguard Fund and Currency Exchange Rate
	url := "https://www.alphavantage.co/query?function=GLOBAL_QUOTE&symbol=VTTSX&apikey=" + f.apiKey
	url2 := "https://www.alphavantage.co/query?function=CURRENCY_EXCHANGE_RATE&from_currency=USD&to_currency=GBP&apikey=" + f.apiKey

	// Fetch Vanguard Fund data (stock price)
	stockPrice, err := f.fetchStockPrice(url)
	if err != nil {
		return 0, 0, 0, fmt.Errorf("error fetching Vanguard Fund price: %v", err)
	}

	// Fetch Currency Exchange Rate (USD to GBP)
	currencyRate, err := f.fetchCurrencyRate(url2)
	if err != nil {
		return 0, 0, 0, fmt.Errorf("error fetching currency exchange rate: %v", err)
	}

	// Convert the stock price to GBP
	stockPriceInGBP := stockPrice * currencyRate

	return stockPrice, stockPriceInGBP, currencyRate, nil
}

// fetchStockPrice makes a request to fetch the Vanguard Fund's stock price.
func (f *financialDataService) fetchStockPrice(url string) (float64, error) {
	resp, err := f.makeRequest(url)
	if err != nil {
		return 0, err
	}

	var stockData map[string]interface{}
	err = json.Unmarshal(resp, &stockData)
	if err != nil {
		return 0, fmt.Errorf("error parsing stock data: %v", err)
	}

	stockPriceStr := stockData["Global Quote"].(map[string]interface{})["05. price"].(string)
	stockPrice, err := strconv.ParseFloat(stockPriceStr, 64)
	if err != nil {
		return 0, fmt.Errorf("error parsing stock price: %v", err)
	}

	return stockPrice, nil
}

// fetchCurrencyRate makes a request to fetch the currency exchange rate from USD to GBP.
func (f *financialDataService) fetchCurrencyRate(url string) (float64, error) {
	resp, err := f.makeRequest(url)
	if err != nil {
		return 0, err
	}

	var currencyData map[string]interface{}
	err = json.Unmarshal(resp, &currencyData)
	if err != nil {
		return 0, fmt.Errorf("error parsing currency data: %v", err)
	}

	currencyRateStr := currencyData["Realtime Currency Exchange Rate"].(map[string]interface{})["5. Exchange Rate"].(string)
	currencyRate, err := strconv.ParseFloat(currencyRateStr, 64)
	if err != nil {
		return 0, fmt.Errorf("error parsing currency rate: %v", err)
	}

	return currencyRate, nil
}

// makeRequest is a helper function to handle HTTP requests.
func (f *financialDataService) makeRequest(url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	resp, err := f.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}

	return body, nil
}
