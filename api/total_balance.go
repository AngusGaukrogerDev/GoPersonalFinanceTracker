package api

import (
	"Modules/GoFinanceTracker/services"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func GetTotalBalance(c *gin.Context) {
	// Initialize the service dependencies
	bankAccountService := &services.AccountBalance{}
	cryptoService := &services.CryptoData{}

	// Fetch bank account balance
	bankAccountURL := os.Getenv("STARLING_URL") + "/api/v2/accounts/" + os.Getenv("STARLING_USER_ID") + "/balance"
	bankAccountBalanceRaw := bankAccountService.GetBankAccountBalance(bankAccountURL)
	bankAccountBalance := float64(bankAccountBalanceRaw) / 100.0
	log.Println("Bank Account Balance (as float):", bankAccountBalance)

	// Fetch football kitty balance
	footballKittyURL := os.Getenv("STARLING_URL") + "/api/v2/account/" + os.Getenv("STARLING_USER_ID") + "/spaces"
	footballKittyBalanceRaw := bankAccountService.GetFootballKittyBalance(footballKittyURL)
	footballKittyBalance := float64(footballKittyBalanceRaw) / 100.0
	log.Println("Football Kitty Balance (as float):", footballKittyBalance)

	// Fetch and calculate cryptocurrency portfolio value
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

	cryptoData, err := cryptoService.GetCryptoData()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch cryptocurrency data"})
		return
	}

	cryptoPortfolioValue, err := cryptoService.CalculateTotalValue(cryptoData, quantities)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to calculate cryptocurrency portfolio value"})
		return
	}
	log.Println("Cryptocurrency Portfolio Value:", cryptoPortfolioValue)

	// Calculate total balance
	totalBalance := bankAccountBalance + cryptoPortfolioValue - footballKittyBalance
	log.Println("Total Balance:", totalBalance)

	// Return the total balance
	c.JSON(http.StatusOK, gin.H{
		"total_balance":          totalBalance,
		"bank_account_balance":   bankAccountBalance,
		"crypto_portfolio_value": cryptoPortfolioValue,
		"football_kitty_balance": footballKittyBalance,
	})
}
