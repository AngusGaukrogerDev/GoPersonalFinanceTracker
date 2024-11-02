// handlers/bank_account.go
package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetBankAccountBalance(c *gin.Context) {
	message := "AccountBalance"
	c.JSON(http.StatusOK, message)
}
