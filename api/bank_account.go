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
	// "/api/v2/accounts"
	// "/api/v2/accounts/{accountUid}/balance"

	var bankAccountService services.BankAccountService = &services.AccountBalance{}

	response := bankAccountService.GetBankAccountBalance(url)

	log.Println("Response Message:", response)

	// Return the message as JSON
	c.JSON(http.StatusOK, gin.H{"bankAccountBalance": response})
}
