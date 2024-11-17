package api

import (
	"Modules/GoFinanceTracker/services"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func GetVanguard2060FundCurrentPrice(c *gin.Context) {
	// Initialize the financial data service
	service := services.NewFinancialDataService(os.Getenv("FINANCE_API_KEY"))

	// Fetch the Vanguard Fund price and currency exchange rate
	stockPriceUSD, stockPriceGBP, currencyRate, err := service.GetVanguardFundPrice()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	vanguard_balance := stockPriceGBP * parseEnvFloat("VANGUARD_2060_QUANTITY")

	// Respond with the stock price in USD, GBP, and the currency exchange rate
	c.JSON(http.StatusOK, gin.H{
		"stock_price_usd":  stockPriceUSD,
		"stock_price_gbp":  stockPriceGBP,
		"currency_rate":    currencyRate,
		"vanguard_balance": vanguard_balance,
	})
}
