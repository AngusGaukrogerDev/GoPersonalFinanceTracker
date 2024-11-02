package main

import (
	"Modules/GoFinanceTracker/handlers"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	myVar := os.Getenv("API")
	fmt.Println("API:", myVar)

	r := gin.Default()
	r.GET("/HelloWorld", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Hello World",
		})
	})
	r.POST("/login", handlers.Login)
	bankAccount := r.Group("/bank_account")
	{
		bankAccount.GET("", handlers.GetBankAccountBalance)
	}
	// r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run() // listen and serve on :8080
}
