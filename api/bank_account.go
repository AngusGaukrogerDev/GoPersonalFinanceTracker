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

	bankAccountService := &services.AccountBalance{}

	response := bankAccountService.GetBankAccountBalance(url)

	log.Println("Response Message:", response)

	c.JSON(http.StatusOK, gin.H{"bankAccountBalance": response})
}

func GetFootballKittyBalance(c *gin.Context) {
	// url := os.Getenv("STARLING_URL") + "/api/v2/accounts/" + os.Getenv("STARLING_USER_ID") + "/spaces"
	url := os.Getenv("STARLING_URL") + "/api/v2/account/" + os.Getenv("STARLING_USER_ID") + "/spaces"

	bankAccountService := &services.AccountBalance{}
	response := bankAccountService.GetFootballKittyBalance(url)
	log.Println("Football Kitty Response Message:", response)
	c.JSON(http.StatusOK, gin.H{"FootballKittyBalance": response})

}
