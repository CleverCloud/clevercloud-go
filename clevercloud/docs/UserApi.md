# \UserApi

All URIs are relative to *https://api.clever-cloud.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AskForPasswordResetViaForm**](UserApi.md#AskForPasswordResetViaForm) | **Post** /password_forgotten | 
[**AuthorizePaypalTransaction**](UserApi.md#AuthorizePaypalTransaction) | **Put** /invoice/external/paypal/{bid} | 
[**CancelPaypalTransaction**](UserApi.md#CancelPaypalTransaction) | **Delete** /invoice/external/paypal/{bid} | 
[**ConfirmPasswordResetRequest**](UserApi.md#ConfirmPasswordResetRequest) | **Get** /password_forgotten/{key} | 
[**CreateUserFromForm**](UserApi.md#CreateUserFromForm) | **Post** /users | 
[**DeleteGithubLink**](UserApi.md#DeleteGithubLink) | **Delete** /github/link | 
[**FinsihGithubSignup**](UserApi.md#FinsihGithubSignup) | **Post** /github/signup | 
[**GetApplications**](UserApi.md#GetApplications) | **Get** /users/{id}/applications | 
[**GetEnv**](UserApi.md#GetEnv) | **Get** /application/{appId}/environment | 
[**GetGitInfo**](UserApi.md#GetGitInfo) | **Get** /users/{userId}/git-info | 
[**GetGithub**](UserApi.md#GetGithub) | **Get** /github | 
[**GetGithubApplications**](UserApi.md#GetGithubApplications) | **Get** /github/applications | 
[**GetGithubCallback**](UserApi.md#GetGithubCallback) | **Get** /github/callback | 
[**GetGithubEmails**](UserApi.md#GetGithubEmails) | **Get** /github/emails | 
[**GetGithubKeys**](UserApi.md#GetGithubKeys) | **Get** /github/keys | 
[**GetGithubLink**](UserApi.md#GetGithubLink) | **Get** /github/link | 
[**GetGithubLogin**](UserApi.md#GetGithubLogin) | **Get** /github/login | 
[**GetGithubUsername**](UserApi.md#GetGithubUsername) | **Get** /github/username | 
[**GetLoginForm**](UserApi.md#GetLoginForm) | **Get** /session/login | 
[**GetLoginForm1**](UserApi.md#GetLoginForm1) | **Get** /sessions/login | 
[**GetPasswordForgottenForm**](UserApi.md#GetPasswordForgottenForm) | **Get** /password_forgotten | 
[**GetSignupForm**](UserApi.md#GetSignupForm) | **Get** /session/signup | 
[**GetSignupForm1**](UserApi.md#GetSignupForm1) | **Get** /sessions/signup | 
[**GetUserById**](UserApi.md#GetUserById) | **Get** /users/{id} | 
[**GithubSignup**](UserApi.md#GithubSignup) | **Get** /github/signup | 
[**Login**](UserApi.md#Login) | **Post** /session/login | 
[**Login1**](UserApi.md#Login1) | **Post** /sessions/login | 
[**MfaLogin**](UserApi.md#MfaLogin) | **Post** /session/mfa_login | 
[**MfaLogin1**](UserApi.md#MfaLogin1) | **Post** /sessions/mfa_login | 
[**PostGithubRedeploy**](UserApi.md#PostGithubRedeploy) | **Post** /github/redeploy | 
[**ResetPasswordForgotten**](UserApi.md#ResetPasswordForgotten) | **Post** /password_forgotten/{key} | 
[**UpdateEnv**](UserApi.md#UpdateEnv) | **Put** /application/{appId}/environment | 
[**UpdateInvoice**](UserApi.md#UpdateInvoice) | **Post** /invoice/external/{bid} | 



## AskForPasswordResetViaForm

> string AskForPasswordResetViaForm(ctx).Login(login).DropTokens(dropTokens).CleverFlavor(cleverFlavor).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    login := "login_example" // string |  (optional)
    dropTokens := "dropTokens_example" // string |  (optional)
    cleverFlavor := "cleverFlavor_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserApi.AskForPasswordResetViaForm(context.Background()).Login(login).DropTokens(dropTokens).CleverFlavor(cleverFlavor).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.AskForPasswordResetViaForm``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AskForPasswordResetViaForm`: string
    fmt.Fprintf(os.Stdout, "Response from `UserApi.AskForPasswordResetViaForm`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAskForPasswordResetViaFormRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **login** | **string** |  | 
 **dropTokens** | **string** |  | 
 **cleverFlavor** | **string** |  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthorizePaypalTransaction

> AuthorizePaypalTransaction(ctx, bid).PaymentData(paymentData).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    bid := "bid_example" // string | 
    paymentData := *openapiclient.NewPaymentData() // PaymentData | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserApi.AuthorizePaypalTransaction(context.Background(), bid).PaymentData(paymentData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.AuthorizePaypalTransaction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuthorizePaypalTransactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **paymentData** | [**PaymentData**](PaymentData.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CancelPaypalTransaction

> CancelPaypalTransaction(ctx, bid).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    bid := "bid_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserApi.CancelPaypalTransaction(context.Background(), bid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.CancelPaypalTransaction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelPaypalTransactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConfirmPasswordResetRequest

> string ConfirmPasswordResetRequest(ctx, key).CleverFlavor(cleverFlavor).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    key := "key_example" // string | 
    cleverFlavor := "cleverFlavor_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserApi.ConfirmPasswordResetRequest(context.Background(), key).CleverFlavor(cleverFlavor).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.ConfirmPasswordResetRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ConfirmPasswordResetRequest`: string
    fmt.Fprintf(os.Stdout, "Response from `UserApi.ConfirmPasswordResetRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiConfirmPasswordResetRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cleverFlavor** | **string** |  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateUserFromForm

> CreateUserFromForm(ctx).InvitationKey(invitationKey).AddonBetaInvitationKey(addonBetaInvitationKey).Email(email).Pass(pass).UrlNext(urlNext).Terms(terms).SubscriptionSource(subscriptionSource).CleverFlavor(cleverFlavor).OauthToken(oauthToken).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    invitationKey := "invitationKey_example" // string |  (optional)
    addonBetaInvitationKey := "addonBetaInvitationKey_example" // string |  (optional)
    email := "email_example" // string |  (optional)
    pass := "pass_example" // string |  (optional)
    urlNext := "urlNext_example" // string |  (optional)
    terms := "terms_example" // string |  (optional)
    subscriptionSource := "subscriptionSource_example" // string |  (optional)
    cleverFlavor := "cleverFlavor_example" // string |  (optional)
    oauthToken := "oauthToken_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserApi.CreateUserFromForm(context.Background()).InvitationKey(invitationKey).AddonBetaInvitationKey(addonBetaInvitationKey).Email(email).Pass(pass).UrlNext(urlNext).Terms(terms).SubscriptionSource(subscriptionSource).CleverFlavor(cleverFlavor).OauthToken(oauthToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.CreateUserFromForm``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateUserFromFormRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **invitationKey** | **string** |  | 
 **addonBetaInvitationKey** | **string** |  | 
 **email** | **string** |  | 
 **pass** | **string** |  | 
 **urlNext** | **string** |  | 
 **terms** | **string** |  | 
 **subscriptionSource** | **string** |  | 
 **cleverFlavor** | **string** |  | 
 **oauthToken** | **string** |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteGithubLink

> Message DeleteGithubLink(ctx).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserApi.DeleteGithubLink(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.DeleteGithubLink``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteGithubLink`: Message
    fmt.Fprintf(os.Stdout, "Response from `UserApi.DeleteGithubLink`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGithubLinkRequest struct via the builder pattern


### Return type

[**Message**](Message.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FinsihGithubSignup

> string FinsihGithubSignup(ctx).TransactionId(transactionId).Name(name).OtherId(otherId).OtherEmail(otherEmail).Password(password).AutoLink(autoLink).Terms(terms).InvitationKey(invitationKey).MfaKind(mfaKind).MfaAttempt(mfaAttempt).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    transactionId := "transactionId_example" // string |  (optional)
    name := "name_example" // string |  (optional)
    otherId := "otherId_example" // string |  (optional)
    otherEmail := "otherEmail_example" // string |  (optional)
    password := "password_example" // string |  (optional)
    autoLink := "autoLink_example" // string |  (optional)
    terms := "terms_example" // string |  (optional)
    invitationKey := "invitationKey_example" // string |  (optional)
    mfaKind := "mfaKind_example" // string |  (optional)
    mfaAttempt := "mfaAttempt_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserApi.FinsihGithubSignup(context.Background()).TransactionId(transactionId).Name(name).OtherId(otherId).OtherEmail(otherEmail).Password(password).AutoLink(autoLink).Terms(terms).InvitationKey(invitationKey).MfaKind(mfaKind).MfaAttempt(mfaAttempt).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.FinsihGithubSignup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FinsihGithubSignup`: string
    fmt.Fprintf(os.Stdout, "Response from `UserApi.FinsihGithubSignup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFinsihGithubSignupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **transactionId** | **string** |  | 
 **name** | **string** |  | 
 **otherId** | **string** |  | 
 **otherEmail** | **string** |  | 
 **password** | **string** |  | 
 **autoLink** | **string** |  | 
 **terms** | **string** |  | 
 **invitationKey** | **string** |  | 
 **mfaKind** | **string** |  | 
 **mfaAttempt** | **string** |  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplications

> []ApplicationView GetApplications(ctx, id).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserApi.GetApplications(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.GetApplications``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApplications`: []ApplicationView
    fmt.Fprintf(os.Stdout, "Response from `UserApi.GetApplications`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]ApplicationView**](ApplicationView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEnv

> string GetEnv(ctx, appId).Token(token).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    appId := "appId_example" // string | 
    token := "token_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserApi.GetEnv(context.Background(), appId).Token(token).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.GetEnv``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEnv`: string
    fmt.Fprintf(os.Stdout, "Response from `UserApi.GetEnv`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEnvRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **token** | **string** |  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGitInfo

> string GetGitInfo(ctx, userId).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    userId := "userId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserApi.GetGitInfo(context.Background(), userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.GetGitInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGitInfo`: string
    fmt.Fprintf(os.Stdout, "Response from `UserApi.GetGitInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGitInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGithub

> OAuthTransactionView GetGithub(ctx).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserApi.GetGithub(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.GetGithub``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGithub`: OAuthTransactionView
    fmt.Fprintf(os.Stdout, "Response from `UserApi.GetGithub`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetGithubRequest struct via the builder pattern


### Return type

[**OAuthTransactionView**](OAuthTransactionView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGithubApplications

> []OAuthApplicationView GetGithubApplications(ctx).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserApi.GetGithubApplications(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.GetGithubApplications``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGithubApplications`: []OAuthApplicationView
    fmt.Fprintf(os.Stdout, "Response from `UserApi.GetGithubApplications`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetGithubApplicationsRequest struct via the builder pattern


### Return type

[**[]OAuthApplicationView**](OAuthApplicationView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGithubCallback

> GetGithubCallback(ctx).CcOAuthData(ccOAuthData).Code(code).State(state).Error_(error_).ErrorDescription(errorDescription).ErrorUri(errorUri).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ccOAuthData := "ccOAuthData_example" // string |  (optional)
    code := "code_example" // string |  (optional)
    state := "state_example" // string |  (optional)
    error_ := "error__example" // string |  (optional)
    errorDescription := "errorDescription_example" // string |  (optional)
    errorUri := "errorUri_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserApi.GetGithubCallback(context.Background()).CcOAuthData(ccOAuthData).Code(code).State(state).Error_(error_).ErrorDescription(errorDescription).ErrorUri(errorUri).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.GetGithubCallback``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetGithubCallbackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ccOAuthData** | **string** |  | 
 **code** | **string** |  | 
 **state** | **string** |  | 
 **error_** | **string** |  | 
 **errorDescription** | **string** |  | 
 **errorUri** | **string** |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGithubEmails

> []string GetGithubEmails(ctx).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserApi.GetGithubEmails(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.GetGithubEmails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGithubEmails`: []string
    fmt.Fprintf(os.Stdout, "Response from `UserApi.GetGithubEmails`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetGithubEmailsRequest struct via the builder pattern


### Return type

**[]string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGithubKeys

> []SshKeyView GetGithubKeys(ctx).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserApi.GetGithubKeys(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.GetGithubKeys``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGithubKeys`: []SshKeyView
    fmt.Fprintf(os.Stdout, "Response from `UserApi.GetGithubKeys`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetGithubKeysRequest struct via the builder pattern


### Return type

[**[]SshKeyView**](SshKeyView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGithubLink

> string GetGithubLink(ctx).TransactionId(transactionId).RedirectUrl(redirectUrl).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    transactionId := "transactionId_example" // string |  (optional)
    redirectUrl := "redirectUrl_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserApi.GetGithubLink(context.Background()).TransactionId(transactionId).RedirectUrl(redirectUrl).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.GetGithubLink``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGithubLink`: string
    fmt.Fprintf(os.Stdout, "Response from `UserApi.GetGithubLink`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetGithubLinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **transactionId** | **string** |  | 
 **redirectUrl** | **string** |  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGithubLogin

> GetGithubLogin(ctx).RedirectUrl(redirectUrl).FromAuthorize(fromAuthorize).CliToken(cliToken).CleverFlavor(cleverFlavor).OauthToken(oauthToken).InvitationKey(invitationKey).SubscriptionSource(subscriptionSource).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    redirectUrl := "redirectUrl_example" // string |  (optional)
    fromAuthorize := "fromAuthorize_example" // string |  (optional)
    cliToken := "cliToken_example" // string |  (optional)
    cleverFlavor := "cleverFlavor_example" // string |  (optional)
    oauthToken := "oauthToken_example" // string |  (optional)
    invitationKey := "invitationKey_example" // string |  (optional)
    subscriptionSource := "subscriptionSource_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserApi.GetGithubLogin(context.Background()).RedirectUrl(redirectUrl).FromAuthorize(fromAuthorize).CliToken(cliToken).CleverFlavor(cleverFlavor).OauthToken(oauthToken).InvitationKey(invitationKey).SubscriptionSource(subscriptionSource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.GetGithubLogin``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetGithubLoginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **redirectUrl** | **string** |  | 
 **fromAuthorize** | **string** |  | 
 **cliToken** | **string** |  | 
 **cleverFlavor** | **string** |  | 
 **oauthToken** | **string** |  | 
 **invitationKey** | **string** |  | 
 **subscriptionSource** | **string** |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGithubUsername

> string GetGithubUsername(ctx).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserApi.GetGithubUsername(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.GetGithubUsername``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGithubUsername`: string
    fmt.Fprintf(os.Stdout, "Response from `UserApi.GetGithubUsername`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetGithubUsernameRequest struct via the builder pattern


### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLoginForm

> string GetLoginForm(ctx).SecondaryEmailKey(secondaryEmailKey).DeletionKey(deletionKey).FromAuthorize(fromAuthorize).CliToken(cliToken).CleverFlavor(cleverFlavor).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    secondaryEmailKey := "secondaryEmailKey_example" // string |  (optional)
    deletionKey := "deletionKey_example" // string |  (optional)
    fromAuthorize := "fromAuthorize_example" // string |  (optional)
    cliToken := "cliToken_example" // string |  (optional)
    cleverFlavor := "cleverFlavor_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserApi.GetLoginForm(context.Background()).SecondaryEmailKey(secondaryEmailKey).DeletionKey(deletionKey).FromAuthorize(fromAuthorize).CliToken(cliToken).CleverFlavor(cleverFlavor).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.GetLoginForm``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLoginForm`: string
    fmt.Fprintf(os.Stdout, "Response from `UserApi.GetLoginForm`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLoginFormRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **secondaryEmailKey** | **string** |  | 
 **deletionKey** | **string** |  | 
 **fromAuthorize** | **string** |  | 
 **cliToken** | **string** |  | 
 **cleverFlavor** | **string** |  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLoginForm1

> string GetLoginForm1(ctx).SecondaryEmailKey(secondaryEmailKey).DeletionKey(deletionKey).FromAuthorize(fromAuthorize).CliToken(cliToken).CleverFlavor(cleverFlavor).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    secondaryEmailKey := "secondaryEmailKey_example" // string |  (optional)
    deletionKey := "deletionKey_example" // string |  (optional)
    fromAuthorize := "fromAuthorize_example" // string |  (optional)
    cliToken := "cliToken_example" // string |  (optional)
    cleverFlavor := "cleverFlavor_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserApi.GetLoginForm1(context.Background()).SecondaryEmailKey(secondaryEmailKey).DeletionKey(deletionKey).FromAuthorize(fromAuthorize).CliToken(cliToken).CleverFlavor(cleverFlavor).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.GetLoginForm1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLoginForm1`: string
    fmt.Fprintf(os.Stdout, "Response from `UserApi.GetLoginForm1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLoginForm1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **secondaryEmailKey** | **string** |  | 
 **deletionKey** | **string** |  | 
 **fromAuthorize** | **string** |  | 
 **cliToken** | **string** |  | 
 **cleverFlavor** | **string** |  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPasswordForgottenForm

> string GetPasswordForgottenForm(ctx).CleverFlavor(cleverFlavor).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    cleverFlavor := "cleverFlavor_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserApi.GetPasswordForgottenForm(context.Background()).CleverFlavor(cleverFlavor).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.GetPasswordForgottenForm``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPasswordForgottenForm`: string
    fmt.Fprintf(os.Stdout, "Response from `UserApi.GetPasswordForgottenForm`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPasswordForgottenFormRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cleverFlavor** | **string** |  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSignupForm

> string GetSignupForm(ctx).InvitationKey(invitationKey).UrlNext(urlNext).CleverFlavor(cleverFlavor).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    invitationKey := "invitationKey_example" // string |  (optional)
    urlNext := "urlNext_example" // string |  (optional)
    cleverFlavor := "cleverFlavor_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserApi.GetSignupForm(context.Background()).InvitationKey(invitationKey).UrlNext(urlNext).CleverFlavor(cleverFlavor).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.GetSignupForm``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSignupForm`: string
    fmt.Fprintf(os.Stdout, "Response from `UserApi.GetSignupForm`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSignupFormRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **invitationKey** | **string** |  | 
 **urlNext** | **string** |  | 
 **cleverFlavor** | **string** |  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSignupForm1

> string GetSignupForm1(ctx).InvitationKey(invitationKey).UrlNext(urlNext).CleverFlavor(cleverFlavor).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    invitationKey := "invitationKey_example" // string |  (optional)
    urlNext := "urlNext_example" // string |  (optional)
    cleverFlavor := "cleverFlavor_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserApi.GetSignupForm1(context.Background()).InvitationKey(invitationKey).UrlNext(urlNext).CleverFlavor(cleverFlavor).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.GetSignupForm1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSignupForm1`: string
    fmt.Fprintf(os.Stdout, "Response from `UserApi.GetSignupForm1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSignupForm1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **invitationKey** | **string** |  | 
 **urlNext** | **string** |  | 
 **cleverFlavor** | **string** |  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserById

> UserView GetUserById(ctx, id).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserApi.GetUserById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.GetUserById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUserById`: UserView
    fmt.Fprintf(os.Stdout, "Response from `UserApi.GetUserById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UserView**](UserView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GithubSignup

> GithubSignup(ctx).RedirectUrl(redirectUrl).FromAuthorize(fromAuthorize).CliToken(cliToken).CleverFlavor(cleverFlavor).OauthToken(oauthToken).InvitationKey(invitationKey).SubscriptionSource(subscriptionSource).Terms(terms).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    redirectUrl := "redirectUrl_example" // string |  (optional)
    fromAuthorize := "fromAuthorize_example" // string |  (optional)
    cliToken := "cliToken_example" // string |  (optional)
    cleverFlavor := "cleverFlavor_example" // string |  (optional)
    oauthToken := "oauthToken_example" // string |  (optional)
    invitationKey := "invitationKey_example" // string |  (optional)
    subscriptionSource := "subscriptionSource_example" // string |  (optional)
    terms := "terms_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserApi.GithubSignup(context.Background()).RedirectUrl(redirectUrl).FromAuthorize(fromAuthorize).CliToken(cliToken).CleverFlavor(cleverFlavor).OauthToken(oauthToken).InvitationKey(invitationKey).SubscriptionSource(subscriptionSource).Terms(terms).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.GithubSignup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGithubSignupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **redirectUrl** | **string** |  | 
 **fromAuthorize** | **string** |  | 
 **cliToken** | **string** |  | 
 **cleverFlavor** | **string** |  | 
 **oauthToken** | **string** |  | 
 **invitationKey** | **string** |  | 
 **subscriptionSource** | **string** |  | 
 **terms** | **string** |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Login

> string Login(ctx).Email(email).Pass(pass).FromAuthorize(fromAuthorize).CliToken(cliToken).CleverFlavor(cleverFlavor).OauthToken(oauthToken).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    email := "email_example" // string |  (optional)
    pass := "pass_example" // string |  (optional)
    fromAuthorize := "fromAuthorize_example" // string |  (optional)
    cliToken := "cliToken_example" // string |  (optional)
    cleverFlavor := "cleverFlavor_example" // string |  (optional)
    oauthToken := "oauthToken_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserApi.Login(context.Background()).Email(email).Pass(pass).FromAuthorize(fromAuthorize).CliToken(cliToken).CleverFlavor(cleverFlavor).OauthToken(oauthToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.Login``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Login`: string
    fmt.Fprintf(os.Stdout, "Response from `UserApi.Login`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLoginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **email** | **string** |  | 
 **pass** | **string** |  | 
 **fromAuthorize** | **string** |  | 
 **cliToken** | **string** |  | 
 **cleverFlavor** | **string** |  | 
 **oauthToken** | **string** |  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Login1

> string Login1(ctx).Email(email).Pass(pass).FromAuthorize(fromAuthorize).CliToken(cliToken).CleverFlavor(cleverFlavor).OauthToken(oauthToken).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    email := "email_example" // string |  (optional)
    pass := "pass_example" // string |  (optional)
    fromAuthorize := "fromAuthorize_example" // string |  (optional)
    cliToken := "cliToken_example" // string |  (optional)
    cleverFlavor := "cleverFlavor_example" // string |  (optional)
    oauthToken := "oauthToken_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserApi.Login1(context.Background()).Email(email).Pass(pass).FromAuthorize(fromAuthorize).CliToken(cliToken).CleverFlavor(cleverFlavor).OauthToken(oauthToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.Login1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Login1`: string
    fmt.Fprintf(os.Stdout, "Response from `UserApi.Login1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLogin1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **email** | **string** |  | 
 **pass** | **string** |  | 
 **fromAuthorize** | **string** |  | 
 **cliToken** | **string** |  | 
 **cleverFlavor** | **string** |  | 
 **oauthToken** | **string** |  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MfaLogin

> []OAuthApplicationView MfaLogin(ctx).MfaKind(mfaKind).MfaAttempt(mfaAttempt).Email(email).AuthId(authId).FromAuthorize(fromAuthorize).CliToken(cliToken).CleverFlavor(cleverFlavor).OauthToken(oauthToken).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    mfaKind := "mfaKind_example" // string |  (optional)
    mfaAttempt := "mfaAttempt_example" // string |  (optional)
    email := "email_example" // string |  (optional)
    authId := "authId_example" // string |  (optional)
    fromAuthorize := "fromAuthorize_example" // string |  (optional)
    cliToken := "cliToken_example" // string |  (optional)
    cleverFlavor := "cleverFlavor_example" // string |  (optional)
    oauthToken := "oauthToken_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserApi.MfaLogin(context.Background()).MfaKind(mfaKind).MfaAttempt(mfaAttempt).Email(email).AuthId(authId).FromAuthorize(fromAuthorize).CliToken(cliToken).CleverFlavor(cleverFlavor).OauthToken(oauthToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.MfaLogin``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MfaLogin`: []OAuthApplicationView
    fmt.Fprintf(os.Stdout, "Response from `UserApi.MfaLogin`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMfaLoginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mfaKind** | **string** |  | 
 **mfaAttempt** | **string** |  | 
 **email** | **string** |  | 
 **authId** | **string** |  | 
 **fromAuthorize** | **string** |  | 
 **cliToken** | **string** |  | 
 **cleverFlavor** | **string** |  | 
 **oauthToken** | **string** |  | 

### Return type

[**[]OAuthApplicationView**](OAuthApplicationView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MfaLogin1

> []OAuthApplicationView MfaLogin1(ctx).MfaKind(mfaKind).MfaAttempt(mfaAttempt).Email(email).AuthId(authId).FromAuthorize(fromAuthorize).CliToken(cliToken).CleverFlavor(cleverFlavor).OauthToken(oauthToken).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    mfaKind := "mfaKind_example" // string |  (optional)
    mfaAttempt := "mfaAttempt_example" // string |  (optional)
    email := "email_example" // string |  (optional)
    authId := "authId_example" // string |  (optional)
    fromAuthorize := "fromAuthorize_example" // string |  (optional)
    cliToken := "cliToken_example" // string |  (optional)
    cleverFlavor := "cleverFlavor_example" // string |  (optional)
    oauthToken := "oauthToken_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserApi.MfaLogin1(context.Background()).MfaKind(mfaKind).MfaAttempt(mfaAttempt).Email(email).AuthId(authId).FromAuthorize(fromAuthorize).CliToken(cliToken).CleverFlavor(cleverFlavor).OauthToken(oauthToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.MfaLogin1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MfaLogin1`: []OAuthApplicationView
    fmt.Fprintf(os.Stdout, "Response from `UserApi.MfaLogin1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMfaLogin1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mfaKind** | **string** |  | 
 **mfaAttempt** | **string** |  | 
 **email** | **string** |  | 
 **authId** | **string** |  | 
 **fromAuthorize** | **string** |  | 
 **cliToken** | **string** |  | 
 **cleverFlavor** | **string** |  | 
 **oauthToken** | **string** |  | 

### Return type

[**[]OAuthApplicationView**](OAuthApplicationView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostGithubRedeploy

> Message PostGithubRedeploy(ctx).GithubWebhookPayload(githubWebhookPayload).UserAgent(userAgent).XGithubEvent(xGithubEvent).XHubSignature(xHubSignature).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    githubWebhookPayload := *openapiclient.NewGithubWebhookPayload() // GithubWebhookPayload | 
    userAgent := "userAgent_example" // string |  (optional)
    xGithubEvent := "xGithubEvent_example" // string |  (optional)
    xHubSignature := "xHubSignature_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserApi.PostGithubRedeploy(context.Background()).GithubWebhookPayload(githubWebhookPayload).UserAgent(userAgent).XGithubEvent(xGithubEvent).XHubSignature(xHubSignature).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.PostGithubRedeploy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostGithubRedeploy`: Message
    fmt.Fprintf(os.Stdout, "Response from `UserApi.PostGithubRedeploy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostGithubRedeployRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **githubWebhookPayload** | [**GithubWebhookPayload**](GithubWebhookPayload.md) |  | 
 **userAgent** | **string** |  | 
 **xGithubEvent** | **string** |  | 
 **xHubSignature** | **string** |  | 

### Return type

[**Message**](Message.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResetPasswordForgotten

> string ResetPasswordForgotten(ctx, key).Pass(pass).Pass2(pass2).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    key := "key_example" // string | 
    pass := "pass_example" // string |  (optional)
    pass2 := "pass2_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserApi.ResetPasswordForgotten(context.Background(), key).Pass(pass).Pass2(pass2).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.ResetPasswordForgotten``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ResetPasswordForgotten`: string
    fmt.Fprintf(os.Stdout, "Response from `UserApi.ResetPasswordForgotten`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiResetPasswordForgottenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pass** | **string** |  | 
 **pass2** | **string** |  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateEnv

> Message UpdateEnv(ctx, appId).Body(body).Token(token).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    appId := "appId_example" // string | 
    body := "body_example" // string | 
    token := "token_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserApi.UpdateEnv(context.Background(), appId).Body(body).Token(token).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.UpdateEnv``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateEnv`: Message
    fmt.Fprintf(os.Stdout, "Response from `UserApi.UpdateEnv`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateEnvRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **string** |  | 
 **token** | **string** |  | 

### Return type

[**Message**](Message.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateInvoice

> UpdateInvoice(ctx, bid).EndOfInvoiceResponse(endOfInvoiceResponse).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    bid := "bid_example" // string | 
    endOfInvoiceResponse := *openapiclient.NewEndOfInvoiceResponse() // EndOfInvoiceResponse | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserApi.UpdateInvoice(context.Background(), bid).EndOfInvoiceResponse(endOfInvoiceResponse).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.UpdateInvoice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateInvoiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **endOfInvoiceResponse** | [**EndOfInvoiceResponse**](EndOfInvoiceResponse.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

