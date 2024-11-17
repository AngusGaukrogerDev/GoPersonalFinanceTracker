package api

import (
	"Modules/GoFinanceTracker/services"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Inject the service as a dependency
func GetCurrentPortfolioValue(c *gin.Context) {
	quantities := map[string]float64{
		"RAY":  parseEnvFloat("RAY"),
		"BONK": parseEnvFloat("BONK"),
		"SOL":  parseEnvFloat("SOL"),
		"ADA":  parseEnvFloat("ADA"),
		"ETH":  parseEnvFloat("ETH"),
		"COTI": parseEnvFloat("COTI"),
		"ALU":  parseEnvFloat("ALU"),
		"DOGE": parseEnvFloat("DOGE"),
		"JUP":  parseEnvFloat("JUP"),
		"WEN":  parseEnvFloat("WEN"),
		"LINK": parseEnvFloat("LINK"),
		"CAKE": parseEnvFloat("CAKE"),
		"DOT":  parseEnvFloat("DOT"),
		"BNB":  parseEnvFloat("BNB"),
		"CHZ":  parseEnvFloat("CHZ"),
	}

	var cryptoService services.CryptoService = &services.CryptoData{}

	// Fetch the crypto data
	data, err := cryptoService.GetCryptoData()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch cryptocurrency data"})
		return
	}

	// Calculate the total value in USD
	totalValue, err := cryptoService.CalculateTotalValue(data, quantities)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to calculate total value"})
		return
	}

	// Return the total value
	c.JSON(http.StatusOK, gin.H{"total_value_usd": totalValue})
}
func parseEnvFloat(key string) float64 {
	valueStr := os.Getenv(key)
	value, err := strconv.ParseFloat(valueStr, 64)
	if err != nil {
		log.Fatalf("Invalid or missing environment variable: %s. Error: %v", key, err)
	}
	return value
}

// func GetCurrentPortfolioValue(c *gin.Context) {
// 	// Set the base URL and endpoint
// 	baseURL := "https://pro-api.coinmarketcap.com"
// 	endpoint := "/v1/cryptocurrency/listings/latest"
// 	url := baseURL + endpoint

// 	// Prepare the request parameters
// 	params := "id=1"
// 	fullURL := fmt.Sprintf("%s?%s", url, params)

// 	// Create a new GET request
// 	req, err := http.NewRequest("GET", fullURL, nil)
// 	if err != nil {
// 		log.Fatalf("Failed to create request: %v", err)
// 	}

// 	// Set the required headers
// 	req.Header.Set("X-CMC_PRO_API_KEY", os.Getenv("CMC_API"))
// 	req.Header.Set("Accept", "application/json")

// 	// Send the request
// 	resp, err := http.DefaultClient.Do(req)
// 	if err != nil {
// 		fmt.Println("Error sending request to server:", err)
// 		os.Exit(1)
// 	}
// 	defer resp.Body.Close()

// 	// Check the status
// 	if resp.StatusCode != http.StatusOK {
// 		fmt.Println("Error: Received non-OK status:", resp.Status)
// 		return
// 	}

// 	// Read and process the response body
// 	respBody, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		log.Fatalf("Failed to read response body: %v", err)
// 	}

// 	// Print the raw response
// 	fmt.Println(string(respBody))

// 	// Parse the JSON response
// 	var data map[string]interface{}
// 	err = json.Unmarshal(respBody, &data)
// 	if err != nil {
// 		log.Fatalf("Failed to unmarshal response JSON: %v", err)
// 	}

// 	// Write the response to a JSON file
// 	err = ioutil.WriteFile("btc_info.json", respBody, 0644)
// 	if err != nil {
// 		log.Fatalf("Failed to write to file: %v", err)
// 	}

// 	// Optionally, return the response as a JSON object (for API endpoint)
// 	c.JSON(http.StatusOK, data)
// }
