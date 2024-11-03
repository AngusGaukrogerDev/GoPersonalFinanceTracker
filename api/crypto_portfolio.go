package api

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func GetCurrentPortfolioValue(c *gin.Context) {
	url := "https://pro-api.coinmarketcap.com/v1/cryptocurrency/map"

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		log.Fatalf("Failed to create request: %v", err)
	}

	req.Header.Set("X-CMC_PRO_API_KEY", os.Getenv("CMC_KEY"))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Error sending request to server")
		os.Exit(1)
	}
	fmt.Println(resp.Status)
	respBody, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(respBody))
}
