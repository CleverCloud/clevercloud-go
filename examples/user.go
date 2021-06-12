package main

import (
	"context"
	"fmt"
	"os"

	"../clevercloud"
)

func main() {
	oauth := clevercloud.NewOAuthClient()
	config := clevercloud.NewConfiguration()
	api := clevercloud.NewOAuthAPIClient(oauth, config)

	resp, r, err := api.SelfApi.GetUser(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SelfApi.GetUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUser`: UserView
	fmt.Fprintf(os.Stdout, "Response from `SelfApi.GetUser`: %v\n", resp)
}
