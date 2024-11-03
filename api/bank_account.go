// api/bank_account.go
package api

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func GetBankAccountBalance(c *gin.Context) {
	url := os.Getenv("STARLING_URL") + "/api/v2/accounts/" + os.Getenv("STARLING_USER_ID") + "/balance"
	// "/api/v2/accounts"
	// "/api/v2/accounts/{accountUid}/balance"

	// Create the HTTP GET request
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatalf("Failed to create request: %v", err)
	}

	req.Header.Set("Authorization", "Bearer "+os.Getenv("STARLING_ACCESS_TOKEN"))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("Failed to make GET request: %v", err)
	}
	defer resp.Body.Close()

	// Check if the response status is OK
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Error response from server: %s", resp.Status)
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %v", err)
	}

	// Log the output as the message variable
	response := string(body)
	log.Println("Response Message:", response)

	// Return the message as JSON
	c.JSON(http.StatusOK, response)
}
