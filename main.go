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
	// r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run() // listen and serve on :8080
}
