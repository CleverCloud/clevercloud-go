package clevercloud

import (
	"github.com/dghubble/oauth1"
)

const (
	APIURL     = "https://api.clever-cloud.com/v2"
	configPath = "/.config/clever-cloud"
)

type Config struct {
	Token  string `json:"token"`
	Secret string `json:"secret"`
}

type OAuthClient struct {
	OAuthConfig *oauth1.Config
	OAuthToken  *oauth1.Token
}

func NewOAuthClient(consumerKey, consumerSecret string) *OAuthClient {
	client := &OAuthClient{
		OAuthConfig: &oauth1.Config{
			ConsumerKey:    consumerKey,
			ConsumerSecret: consumerSecret,
			Endpoint: oauth1.Endpoint{
				RequestTokenURL: APIURL + "/oauth/request_token",
				AuthorizeURL:    APIURL + "/oauth/authorize",
				AccessTokenURL:  APIURL + "/oauth/access_token",
			},
		},
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

// SetOAuthTokens set oauth tokens in configuration
func (o *OAuthClient) SetTokens(token, secret string) {
	o.OAuthToken = &oauth1.Token{
		Token:       token,
		TokenSecret: secret,
	}
}
