package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"os"
)

var setupToken string
var accessToken string 

/*
How should this work
setupToken: The token that a user recieves from simple fin and presumably puts into the application.
accessToken: Obtained by sending a post request to the base64 decoded setupToken. This is the token that the application holds on to in order to access the information.
*/

func getAccessToken(string setupToken) string{
	decodedBytes, err := base64.StdEncoding.DecodeString(setupToken)
	if err != nil {
		log.Fatalf("Failed to decode Base64 string: %v", err)
	}

  accessTokenUrl := string(decodedBytes)

	//send a post request token to the accessTokenUrl
	//receive a setupToken if the response type is 200
}

func main() {
  // No functionality should be kept in this method.
  // This is just for testing.

	fmt.Println("oof")
	setupToken = os.Getenv("SIMPLEFIN_SETUP_TOKEN")
	accessToken =
	if setupToken == "" {
		log.Fatal("SIMPLEFIN_SETUP_TOKEN is not available in the environment")
	}

	fmt.Println("Encoded string: " + setupToken)

	// Print the decoded string
	fmt.Println("Decoded string: ", decodedStr)

}
