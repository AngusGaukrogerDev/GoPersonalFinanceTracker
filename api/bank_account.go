// api/bank_account.go
package api

import (
	"Modules/GoFinanceTracker/services"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func GetBankAccountBalance(c *gin.Context) {
	url := os.Getenv("STARLING_URL") + "/api/v2/accounts/" + os.Getenv("STARLING_USER_ID") + "/balance"
	// url := "https://api.starlingbank.com/api/v2/account/941638de-ccd2-4f02-9541-f47fa5286a4e/spaces"
	// "/api/v2/accounts"
	// "/api/v2/accounts/{accountUid}/balance"

	bankAccountService := &services.AccountBalance{}

	response := bankAccountService.GetBankAccountBalance(url)

	log.Println("Response Message:", response)

	c.JSON(http.StatusOK, gin.H{"bankAccountBalance": response})
}

func GetFootballKittyBalance(c *gin.Context) {
	// url := os.Getenv("STARLING_URL") + "/api/v2/accounts/" + os.Getenv("STARLING_USER_ID") + "/spaces"
	url := "https://api.starlingbank.com/api/v2/account/941638de-ccd2-4f02-9541-f47fa5286a4e/spaces"

	bankAccountService := &services.AccountBalance{}
	response := bankAccountService.GetFootballKittyBalance(url)
	log.Println("Football Kitty Response Message:", response)
	c.JSON(http.StatusOK, gin.H{"FootballKittyBalance": response})

}
