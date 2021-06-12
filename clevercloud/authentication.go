package clevercloud

import (
	"os"

	"github.com/dghubble/oauth1"
)

const (
	APIURL         = "https://api.clever-cloud.com/v2"
	consumerKey    = "x9Dv3WBvHcZZgg3sLmzXTyi2FFhfSu"
	consumerSecret = "RnUvL43r4RoUgH9cHqTzeeCh2v0nbv"
)

type OAuthClient struct {
	OAuthConfig *oauth1.Config
	OAuthToken  *oauth1.Token
}

func NewOAuthClient() *OAuthClient {
	client := &OAuthClient{
		OAuthConfig: &oauth1.Config{
			ConsumerKey:    consumerKey,
			ConsumerSecret: consumerSecret,
			Signer: &oauth1.HMAC256Signer{
				ConsumerSecret: consumerSecret,
			},
			Endpoint: oauth1.Endpoint{
				RequestTokenURL: APIURL + "/oauth/request_token",
				AuthorizeURL:    APIURL + "/oauth/authorize",
				AccessTokenURL:  APIURL + "/oauth/access_token",
			},
		},
	}

	if client.OAuthToken == nil {
		client.SetOAuthTokensFromEnv()
	}

	return client
}

func NewOAuthAPIClient(oc *OAuthClient, cfg *Configuration) *APIClient {
	if cfg.HTTPClient == nil {
		cfg.HTTPClient = oc.OAuthConfig.Client(oauth1.NoContext, oc.OAuthToken)
	}

	c := &APIClient{}
	c.cfg = cfg
	c.common.client = c

	// API Services
	c.AuthApi = (*AuthApiService)(&c.common)
	c.DefaultApi = (*DefaultApiService)(&c.common)
	c.OrganisationApi = (*OrganisationApiService)(&c.common)
	c.PaymentApi = (*PaymentApiService)(&c.common)
	c.ProductsApi = (*ProductsApiService)(&c.common)
	c.SelfApi = (*SelfApiService)(&c.common)
	c.UserApi = (*UserApiService)(&c.common)

	return c
}

// SetOAuthTokensFromEnv set oauth tokens from env
func (o *OAuthClient) SetOAuthTokensFromEnv() {
	o.SetOAuthTokens(os.Getenv("CLEVER_TOKEN"), os.Getenv("CLEVER_SECRET"))
}

// SetOAuthTokens set oauth tokens in configuration
func (o *OAuthClient) SetOAuthTokens(token, secret string) {
	o.OAuthToken = &oauth1.Token{
		Token:       token,
		TokenSecret: secret,
	}
}
