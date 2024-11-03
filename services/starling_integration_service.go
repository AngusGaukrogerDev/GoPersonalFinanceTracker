package services

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
)

type BalanceResponse struct {
	TotalEffectiveBalance struct {
		Currency   string `json:"currency"`
		MinorUnits int    `json:"minorUnits"`
	} `json:"totalEffectiveBalance"`
}

type FootballKittyBalanceResponse struct {
	SavingsGoals []struct {
		SavingsGoalUid string `json:"savingsGoalUid"`
		Name           string `json:"name"`
		TotalSaved     struct {
			Currency   string `json:"currency"`
			MinorUnits int    `json:"minorUnits"`
		} `json:"totalSaved"`
		SortOrder int    `json:"sortOrder"`
		State     string `json:"state"`
	} `json:"savingsGoals"`
}

type BankAccountService interface {
	GetBankAccountBalance(url string) int
	GetFootballKittyBalance(url string) int
}

type AccountBalance struct {
	totalBalance int
	kittyBalance int
}

func (accountBalance *AccountBalance) GetBankAccountBalance(url string) int {
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

	// Unmarshal the response JSON into BalanceResponse
	var balanceResponse BalanceResponse
	err = json.Unmarshal(body, &balanceResponse)
	if err != nil {
		log.Fatalf("Failed to unmarshal response: %v", err)
	}

	// Set and return the totalEffectiveBalance.minorUnits
	accountBalance.totalBalance = balanceResponse.TotalEffectiveBalance.MinorUnits
	return accountBalance.totalBalance
}

func (accountBalance *AccountBalance) GetFootballKittyBalance(url string) int {
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
	// if resp.StatusCode != http.StatusOK {
	// 	log.Fatalf("Error response from server: %s", resp.Status)
	// }

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %v", err)
	}

	// Unmarshal the response JSON into BalanceResponse
	var footballKittyBalanceResponse FootballKittyBalanceResponse
	err = json.Unmarshal(body, &footballKittyBalanceResponse)
	if err != nil {
		log.Fatalf("Failed to unmarshal response: %v", err)
	}

	for _, goal := range footballKittyBalanceResponse.SavingsGoals {
		if goal.Name == "Football Kitty" {
			accountBalance.kittyBalance = goal.TotalSaved.MinorUnits
			return accountBalance.kittyBalance
		}
	}

	log.Println("Football Kitty savings goal not found.")
	return 0
}
