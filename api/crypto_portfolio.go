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
	c.JSON(http.StatusOK, gin.H{"total_value_gbp": totalValue})
}
func parseEnvFloat(key string) float64 {
	valueStr := os.Getenv(key)
	value, err := strconv.ParseFloat(valueStr, 64)
	if err != nil {
		log.Fatalf("Invalid or missing environment variable: %s. Error: %v", key, err)
	}
	return value
}
