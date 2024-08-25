package main

import (
	"fingo/fingo"
	"fmt"
	"log"
	"os"
)

/*
How should this work
setupToken: The token that a user recieves from simple fin and presumably puts into the application.
accessToken: Obtained by sending a post request to the base64 decoded setupToken. This is the token that the application holds on to in order to access the information.
*/

func main() {
	accessToken := os.Getenv("SIMPLEFIN_ACCESS_TOKEN")
	fmt.Println("Access token: " + accessToken)
	fingo.ConfigureAccessToken(accessToken)

	gnto := fingo.NewGetTransactionsOptions()

	resp, err := fingo.GetTransactions(gnto)
	if err != nil {
		log.Fatalf("oof")
	}

	for _, account := range resp.Accounts {
		fmt.Printf("%s %s: %s\n", account.Org.Name, account.Name, account.Balance)
		for _, transaction := range account.Transactions {
			fmt.Printf("\t transaction: %s for $%s\n", transaction.Payee, transaction.Amount)
		}
	}

}
