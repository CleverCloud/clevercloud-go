package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/clevercloud/clevercloud-go/clevercloud"
	"github.com/joho/godotenv"
	"moul.io/http2curl"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	oauth := clevercloud.NewOAuthClient(os.Getenv("OAUTH_CONSUMER_KEY"), os.Getenv("OAUTH_CONSUMER_SECRET"))
	oauth.SetTokens(os.Getenv("CLEVER_TOKEN"), os.Getenv("CLEVER_SECRET"))

	config := clevercloud.NewConfiguration()
	api := clevercloud.NewOAuthAPIClient(oauth, config)

	resp, r, err := api.SelfApi.GetUser(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SelfApi.GetUser``: %v\n", err)

		command, _ := http2curl.GetCurlCommand(r.Request)
		fmt.Fprint(os.Stderr, command)

		return
	}

	// response from `GetUser`: UserView
	fmt.Fprintf(os.Stdout, "Response from `SelfApi.GetUser`: %v\n", resp)
}
