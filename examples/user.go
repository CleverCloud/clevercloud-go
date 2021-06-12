package examples

import (
	"context"
	"fmt"
	"os"

	"github.com/clevercloud/clevercloud-go/clevercloud"
)

func main() {
	config := clevercloud.NewConfiguration()
	api := clevercloud.NewAPIClient(config)

	resp, r, err := api.SelfApi.GetUser(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SelfApi.GetUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUser`: UserView
	fmt.Fprintf(os.Stdout, "Response from `SelfApi.GetUser`: %v\n", resp)
}
