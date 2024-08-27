package fingo

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Client struct {
	accessTokenUrl string
}

type GetTransactionResponse struct {
	Errors   []interface{} `json:"errors"`
	Accounts []Account     `json:"accounts"`
}

// Account struct to represent each account
type Account struct {
	Org              Organization  `json:"org"`
	ID               string        `json:"id"`
	Name             string        `json:"name"`
	Currency         string        `json:"currency"`
	Balance          string        `json:"balance"`
	AvailableBalance string        `json:"available-balance"`
	BalanceDate      int           `json:"balance-date"`
	Transactions     []Transaction `json:"transactions"`
	Holdings         []interface{} `json:"holdings"`
}

// Organization struct to represent the organization associated with an account
type Organization struct {
	Domain  string `json:"domain"`
	Name    string `json:"name"`
	SfinURL string `json:"sfin-url"`
	URL     string `json:"url"`
	ID      string `json:"id"`
}

// Transaction struct to represent each transaction within an account
type Transaction struct {
	ID           string `json:"id"`
	Posted       int    `json:"posted"`
	Amount       string `json:"amount"`
	Description  string `json:"description"`
	Payee        string `json:"payee"`
	Memo         string `json:"memo"`
	TransactedAt int    `json:"transacted_at"`
}

type GetTransactionsOptions struct {
	StartDate int
	EndDate   int
	Pending   int
}

const accountsEndpoint = "/accounts"

var client Client

// Get an access token from the setup token obtained from
// simplefin
func GetAccessTokenFromSetupToken(setupToken string) string {
	decodedBytes, err := base64.StdEncoding.DecodeString(setupToken)
	if err != nil {
		log.Fatalf("Failed to decode Base64 string: %v", err)
	}

	accessTokenUrl := string(decodedBytes)

	return accessTokenUrl
	//send a post request token to the accessTokenUrl
	//receive a setupToken if the response type is 200
}

func NewGetTransactionsOptions() GetTransactionsOptions {
	var getTransactionsOptions GetTransactionsOptions

	getTransactionsOptions.StartDate = -1
	getTransactionsOptions.EndDate = -1
	getTransactionsOptions.Pending = -1

	return getTransactionsOptions
}

func ConfigureAccessToken(accessToken string) {
	client.accessTokenUrl = accessToken
}

func GetTransactions(gto GetTransactionsOptions) (GetTransactionResponse, error) {
	var apiResponse GetTransactionResponse
	var queryStringParamsArr []string
	var param string

	if accessTokenUrl == "" {
		return apiResponse, errors.New("No access token supplied")
	}

	if gto.StartDate != -1 {
		param = "start-date=" + strconv.Itoa(gto.StartDate)
		queryStringParamsArr = append(queryStringParamsArr, param)
	}

	if gto.EndDate != -1 {
		param = "end-date=" + strconv.Itoa(gto.EndDate)
		queryStringParamsArr = append(queryStringParamsArr, param)
	}

	if gto.Pending != -1 {
		param = "pending=" + strconv.Itoa(gto.Pending)
		queryStringParamsArr = append(queryStringParamsArr, param)
	}

	queryStringParams := strings.Join(queryStringParamsArr, "&")
	if len(queryStringParams) > 0 {
		queryStringParams = "?" + queryStringParams
	}

	url := client.accessTokenUrl + accountsEndpoint + queryStringParams

	resp, err := http.Get(url)

	if err != nil {
		fmt.Printf("Error making GET request: %v\n", err)
		return apiResponse, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %v\n", err)
		return apiResponse, err
	}

	// Parse JSON into the struct
	err = json.Unmarshal(body, &apiResponse)
	if err != nil {
		fmt.Printf("Error parsing JSON: %v\n", err)
		return apiResponse, err
	}

	return apiResponse, nil
}

/*
func GetInfo() {
}
*/
