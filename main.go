package main

import (
	"Modules/GoFinanceTracker/api"
	"Modules/GoFinanceTracker/middleware"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	r := gin.Default()

	r.Use(middleware.CORSMiddleware())

	r.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "It Works!",
		})
	})

	r.POST("/login", api.Login)

	bankAccount := r.Group("/bank_account")
	bankAccount.Use(middleware.AuthMiddleware())
	{
		bankAccount.GET("/balance", api.GetBankAccountBalance)
		bankAccount.GET("/kitty", api.GetFootballKittyBalance)
	}
	vanguardPortfolio := r.Group("/vanguard_portfolio")
	vanguardPortfolio.Use(middleware.AuthMiddleware())
	{
		vanguardPortfolio.GET("/balance", api.GetVanguard2060FundCurrentPrice)
	}

	cryptoPortfolio := r.Group("/crypto_portfolio")
	cryptoPortfolio.Use(middleware.AuthMiddleware())
	{
		cryptoPortfolio.GET("/balance", api.GetCurrentPortfolioValue)
	}

	totalBalance := r.Group("/total")
	totalBalance.Use(middleware.AuthMiddleware())
	{
		totalBalance.GET("/balance", api.GetTotalBalance)
	}
	// r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run() // listen and serve on :8080
}
