# \AuthApi

All URIs are relative to *https://api.clever-cloud.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuthorizeForm**](AuthApi.md#AuthorizeForm) | **Get** /oauth/authorize | 
[**AuthorizeToken**](AuthApi.md#AuthorizeToken) | **Post** /oauth/authorize | 
[**GetAvailableRights**](AuthApi.md#GetAvailableRights) | **Get** /oauth/rights | 
[**GetLoginData**](AuthApi.md#GetLoginData) | **Get** /oauth/login_data | 
[**PostAccessTokenRequest**](AuthApi.md#PostAccessTokenRequest) | **Post** /oauth/access_token | 
[**PostAccessTokenRequestQuery**](AuthApi.md#PostAccessTokenRequestQuery) | **Post** /oauth/access_token_query | 
[**PostAuthorize**](AuthApi.md#PostAuthorize) | **Post** /authorize | 
[**PostReqTokenRequest**](AuthApi.md#PostReqTokenRequest) | **Post** /oauth/request_token | 
[**PostReqTokenRequestQueryString**](AuthApi.md#PostReqTokenRequestQueryString) | **Post** /oauth/request_token_query | 



## AuthorizeForm

> string AuthorizeForm(ctx).Ccid(ccid).Cctk(cctk).OauthToken(oauthToken).Ccid2(ccid2).CliToken(cliToken).FromOauth(fromOauth).Execute()



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
    ccid := "ccid_example" // string |  (optional)
    cctk := "cctk_example" // string |  (optional)
    oauthToken := "oauthToken_example" // string |  (optional)
    ccid2 := "ccid_example" // string |  (optional)
    cliToken := "cliToken_example" // string |  (optional)
    fromOauth := "fromOauth_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthApi.AuthorizeForm(context.Background()).Ccid(ccid).Cctk(cctk).OauthToken(oauthToken).Ccid2(ccid2).CliToken(cliToken).FromOauth(fromOauth).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthApi.AuthorizeForm``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthorizeForm`: string
    fmt.Fprintf(os.Stdout, "Response from `AuthApi.AuthorizeForm`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthorizeFormRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ccid** | **string** |  | 
 **cctk** | **string** |  | 
 **oauthToken** | **string** |  | 
 **ccid2** | **string** |  | 
 **cliToken** | **string** |  | 
 **fromOauth** | **string** |  | 

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


## AuthorizeToken

> AuthorizeToken(ctx).Ccid(ccid).Cctk(cctk).Almighty(almighty).AccessOrganisations(accessOrganisations).ManageOrganisations(manageOrganisations).ManageOrganisationsServices(manageOrganisationsServices).ManageOrganisationsApplications(manageOrganisationsApplications).ManageOrganisationsMembers(manageOrganisationsMembers).AccessOrganisationsBills(accessOrganisationsBills).AccessOrganisationsCreditCount(accessOrganisationsCreditCount).AccessOrganisationsConsumptionStatistics(accessOrganisationsConsumptionStatistics).AccessPersonalInformation(accessPersonalInformation).ManagePersonalInformation(managePersonalInformation).ManageSshKeys(manageSshKeys).Execute()



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
    ccid := "ccid_example" // string |  (optional)
    cctk := "cctk_example" // string |  (optional)
    almighty := "almighty_example" // string |  (optional)
    accessOrganisations := "accessOrganisations_example" // string |  (optional)
    manageOrganisations := "manageOrganisations_example" // string |  (optional)
    manageOrganisationsServices := "manageOrganisationsServices_example" // string |  (optional)
    manageOrganisationsApplications := "manageOrganisationsApplications_example" // string |  (optional)
    manageOrganisationsMembers := "manageOrganisationsMembers_example" // string |  (optional)
    accessOrganisationsBills := "accessOrganisationsBills_example" // string |  (optional)
    accessOrganisationsCreditCount := "accessOrganisationsCreditCount_example" // string |  (optional)
    accessOrganisationsConsumptionStatistics := "accessOrganisationsConsumptionStatistics_example" // string |  (optional)
    accessPersonalInformation := "accessPersonalInformation_example" // string |  (optional)
    managePersonalInformation := "managePersonalInformation_example" // string |  (optional)
    manageSshKeys := "manageSshKeys_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthApi.AuthorizeToken(context.Background()).Ccid(ccid).Cctk(cctk).Almighty(almighty).AccessOrganisations(accessOrganisations).ManageOrganisations(manageOrganisations).ManageOrganisationsServices(manageOrganisationsServices).ManageOrganisationsApplications(manageOrganisationsApplications).ManageOrganisationsMembers(manageOrganisationsMembers).AccessOrganisationsBills(accessOrganisationsBills).AccessOrganisationsCreditCount(accessOrganisationsCreditCount).AccessOrganisationsConsumptionStatistics(accessOrganisationsConsumptionStatistics).AccessPersonalInformation(accessPersonalInformation).ManagePersonalInformation(managePersonalInformation).ManageSshKeys(manageSshKeys).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthApi.AuthorizeToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthorizeTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ccid** | **string** |  | 
 **cctk** | **string** |  | 
 **almighty** | **string** |  | 
 **accessOrganisations** | **string** |  | 
 **manageOrganisations** | **string** |  | 
 **manageOrganisationsServices** | **string** |  | 
 **manageOrganisationsApplications** | **string** |  | 
 **manageOrganisationsMembers** | **string** |  | 
 **accessOrganisationsBills** | **string** |  | 
 **accessOrganisationsCreditCount** | **string** |  | 
 **accessOrganisationsConsumptionStatistics** | **string** |  | 
 **accessPersonalInformation** | **string** |  | 
 **managePersonalInformation** | **string** |  | 
 **manageSshKeys** | **string** |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: text/html, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAvailableRights

> GetAvailableRights(ctx).Execute()



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
    resp, r, err := api_client.AuthApi.GetAvailableRights(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthApi.GetAvailableRights``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAvailableRightsRequest struct via the builder pattern


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


## GetLoginData

> GetLoginData(ctx).OauthKey(oauthKey).Execute()



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
    oauthKey := "oauthKey_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthApi.GetLoginData(context.Background()).OauthKey(oauthKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthApi.GetLoginData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLoginDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oauthKey** | **string** |  | 

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


## PostAccessTokenRequest

> PostAccessTokenRequest(ctx).OauthConsumerKey(oauthConsumerKey).OauthToken(oauthToken).OauthSignatureMethod(oauthSignatureMethod).OauthSignature(oauthSignature).OauthTimestamp(oauthTimestamp).OauthNonce(oauthNonce).OauthVersion(oauthVersion).OauthVerifier(oauthVerifier).OauthCallback(oauthCallback).OauthTokenSecret(oauthTokenSecret).OauthCallbackConfirmed(oauthCallbackConfirmed).Execute()



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
    oauthConsumerKey := "oauthConsumerKey_example" // string |  (optional)
    oauthToken := "oauthToken_example" // string |  (optional)
    oauthSignatureMethod := "oauthSignatureMethod_example" // string |  (optional)
    oauthSignature := "oauthSignature_example" // string |  (optional)
    oauthTimestamp := "oauthTimestamp_example" // string |  (optional)
    oauthNonce := "oauthNonce_example" // string |  (optional)
    oauthVersion := "oauthVersion_example" // string |  (optional)
    oauthVerifier := "oauthVerifier_example" // string |  (optional)
    oauthCallback := "oauthCallback_example" // string |  (optional)
    oauthTokenSecret := "oauthTokenSecret_example" // string |  (optional)
    oauthCallbackConfirmed := "oauthCallbackConfirmed_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthApi.PostAccessTokenRequest(context.Background()).OauthConsumerKey(oauthConsumerKey).OauthToken(oauthToken).OauthSignatureMethod(oauthSignatureMethod).OauthSignature(oauthSignature).OauthTimestamp(oauthTimestamp).OauthNonce(oauthNonce).OauthVersion(oauthVersion).OauthVerifier(oauthVerifier).OauthCallback(oauthCallback).OauthTokenSecret(oauthTokenSecret).OauthCallbackConfirmed(oauthCallbackConfirmed).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthApi.PostAccessTokenRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAccessTokenRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oauthConsumerKey** | **string** |  | 
 **oauthToken** | **string** |  | 
 **oauthSignatureMethod** | **string** |  | 
 **oauthSignature** | **string** |  | 
 **oauthTimestamp** | **string** |  | 
 **oauthNonce** | **string** |  | 
 **oauthVersion** | **string** |  | 
 **oauthVerifier** | **string** |  | 
 **oauthCallback** | **string** |  | 
 **oauthTokenSecret** | **string** |  | 
 **oauthCallbackConfirmed** | **string** |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/x-www-form-urlencoded

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAccessTokenRequestQuery

> PostAccessTokenRequestQuery(ctx).OauthConsumerKey(oauthConsumerKey).OauthToken(oauthToken).OauthSignatureMethod(oauthSignatureMethod).OauthSignature(oauthSignature).OauthTimestamp(oauthTimestamp).OauthNonce(oauthNonce).OauthVersion(oauthVersion).OauthVerifier(oauthVerifier).OauthCallback(oauthCallback).OauthTokenSecret(oauthTokenSecret).OauthCallbackConfirmed(oauthCallbackConfirmed).Execute()



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
    oauthConsumerKey := "oauthConsumerKey_example" // string |  (optional)
    oauthToken := "oauthToken_example" // string |  (optional)
    oauthSignatureMethod := "oauthSignatureMethod_example" // string |  (optional)
    oauthSignature := "oauthSignature_example" // string |  (optional)
    oauthTimestamp := "oauthTimestamp_example" // string |  (optional)
    oauthNonce := "oauthNonce_example" // string |  (optional)
    oauthVersion := "oauthVersion_example" // string |  (optional)
    oauthVerifier := "oauthVerifier_example" // string |  (optional)
    oauthCallback := "oauthCallback_example" // string |  (optional)
    oauthTokenSecret := "oauthTokenSecret_example" // string |  (optional)
    oauthCallbackConfirmed := "oauthCallbackConfirmed_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthApi.PostAccessTokenRequestQuery(context.Background()).OauthConsumerKey(oauthConsumerKey).OauthToken(oauthToken).OauthSignatureMethod(oauthSignatureMethod).OauthSignature(oauthSignature).OauthTimestamp(oauthTimestamp).OauthNonce(oauthNonce).OauthVersion(oauthVersion).OauthVerifier(oauthVerifier).OauthCallback(oauthCallback).OauthTokenSecret(oauthTokenSecret).OauthCallbackConfirmed(oauthCallbackConfirmed).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthApi.PostAccessTokenRequestQuery``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAccessTokenRequestQueryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oauthConsumerKey** | **string** |  | 
 **oauthToken** | **string** |  | 
 **oauthSignatureMethod** | **string** |  | 
 **oauthSignature** | **string** |  | 
 **oauthTimestamp** | **string** |  | 
 **oauthNonce** | **string** |  | 
 **oauthVersion** | **string** |  | 
 **oauthVerifier** | **string** |  | 
 **oauthCallback** | **string** |  | 
 **oauthTokenSecret** | **string** |  | 
 **oauthCallbackConfirmed** | **string** |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/x-www-form-urlencoded

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAuthorize

> Message PostAuthorize(ctx).WannabeAuthorization(wannabeAuthorization).Execute()



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
    wannabeAuthorization := *openapiclient.NewWannabeAuthorization() // WannabeAuthorization | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthApi.PostAuthorize(context.Background()).WannabeAuthorization(wannabeAuthorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthApi.PostAuthorize``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostAuthorize`: Message
    fmt.Fprintf(os.Stdout, "Response from `AuthApi.PostAuthorize`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAuthorizeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **wannabeAuthorization** | [**WannabeAuthorization**](WannabeAuthorization.md) |  | 

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


## PostReqTokenRequest

> string PostReqTokenRequest(ctx).CleverFlavor(cleverFlavor).OauthConsumerKey(oauthConsumerKey).OauthToken(oauthToken).OauthSignatureMethod(oauthSignatureMethod).OauthSignature(oauthSignature).OauthTimestamp(oauthTimestamp).OauthNonce(oauthNonce).OauthVersion(oauthVersion).OauthVerifier(oauthVerifier).OauthCallback(oauthCallback).OauthTokenSecret(oauthTokenSecret).OauthCallbackConfirmed(oauthCallbackConfirmed).Execute()



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
    oauthConsumerKey := "oauthConsumerKey_example" // string |  (optional)
    oauthToken := "oauthToken_example" // string |  (optional)
    oauthSignatureMethod := "oauthSignatureMethod_example" // string |  (optional)
    oauthSignature := "oauthSignature_example" // string |  (optional)
    oauthTimestamp := "oauthTimestamp_example" // string |  (optional)
    oauthNonce := "oauthNonce_example" // string |  (optional)
    oauthVersion := "oauthVersion_example" // string |  (optional)
    oauthVerifier := "oauthVerifier_example" // string |  (optional)
    oauthCallback := "oauthCallback_example" // string |  (optional)
    oauthTokenSecret := "oauthTokenSecret_example" // string |  (optional)
    oauthCallbackConfirmed := "oauthCallbackConfirmed_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthApi.PostReqTokenRequest(context.Background()).CleverFlavor(cleverFlavor).OauthConsumerKey(oauthConsumerKey).OauthToken(oauthToken).OauthSignatureMethod(oauthSignatureMethod).OauthSignature(oauthSignature).OauthTimestamp(oauthTimestamp).OauthNonce(oauthNonce).OauthVersion(oauthVersion).OauthVerifier(oauthVerifier).OauthCallback(oauthCallback).OauthTokenSecret(oauthTokenSecret).OauthCallbackConfirmed(oauthCallbackConfirmed).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthApi.PostReqTokenRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostReqTokenRequest`: string
    fmt.Fprintf(os.Stdout, "Response from `AuthApi.PostReqTokenRequest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostReqTokenRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cleverFlavor** | **string** |  | 
 **oauthConsumerKey** | **string** |  | 
 **oauthToken** | **string** |  | 
 **oauthSignatureMethod** | **string** |  | 
 **oauthSignature** | **string** |  | 
 **oauthTimestamp** | **string** |  | 
 **oauthNonce** | **string** |  | 
 **oauthVersion** | **string** |  | 
 **oauthVerifier** | **string** |  | 
 **oauthCallback** | **string** |  | 
 **oauthTokenSecret** | **string** |  | 
 **oauthCallbackConfirmed** | **string** |  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/x-www-form-urlencoded

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostReqTokenRequestQueryString

> string PostReqTokenRequestQueryString(ctx).CleverFlavor(cleverFlavor).OauthConsumerKey(oauthConsumerKey).OauthToken(oauthToken).OauthSignatureMethod(oauthSignatureMethod).OauthSignature(oauthSignature).OauthTimestamp(oauthTimestamp).OauthNonce(oauthNonce).OauthVersion(oauthVersion).OauthVerifier(oauthVerifier).OauthCallback(oauthCallback).OauthTokenSecret(oauthTokenSecret).OauthCallbackConfirmed(oauthCallbackConfirmed).Execute()



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
    oauthConsumerKey := "oauthConsumerKey_example" // string |  (optional)
    oauthToken := "oauthToken_example" // string |  (optional)
    oauthSignatureMethod := "oauthSignatureMethod_example" // string |  (optional)
    oauthSignature := "oauthSignature_example" // string |  (optional)
    oauthTimestamp := "oauthTimestamp_example" // string |  (optional)
    oauthNonce := "oauthNonce_example" // string |  (optional)
    oauthVersion := "oauthVersion_example" // string |  (optional)
    oauthVerifier := "oauthVerifier_example" // string |  (optional)
    oauthCallback := "oauthCallback_example" // string |  (optional)
    oauthTokenSecret := "oauthTokenSecret_example" // string |  (optional)
    oauthCallbackConfirmed := "oauthCallbackConfirmed_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthApi.PostReqTokenRequestQueryString(context.Background()).CleverFlavor(cleverFlavor).OauthConsumerKey(oauthConsumerKey).OauthToken(oauthToken).OauthSignatureMethod(oauthSignatureMethod).OauthSignature(oauthSignature).OauthTimestamp(oauthTimestamp).OauthNonce(oauthNonce).OauthVersion(oauthVersion).OauthVerifier(oauthVerifier).OauthCallback(oauthCallback).OauthTokenSecret(oauthTokenSecret).OauthCallbackConfirmed(oauthCallbackConfirmed).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthApi.PostReqTokenRequestQueryString``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostReqTokenRequestQueryString`: string
    fmt.Fprintf(os.Stdout, "Response from `AuthApi.PostReqTokenRequestQueryString`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostReqTokenRequestQueryStringRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cleverFlavor** | **string** |  | 
 **oauthConsumerKey** | **string** |  | 
 **oauthToken** | **string** |  | 
 **oauthSignatureMethod** | **string** |  | 
 **oauthSignature** | **string** |  | 
 **oauthTimestamp** | **string** |  | 
 **oauthNonce** | **string** |  | 
 **oauthVersion** | **string** |  | 
 **oauthVerifier** | **string** |  | 
 **oauthCallback** | **string** |  | 
 **oauthTokenSecret** | **string** |  | 
 **oauthCallbackConfirmed** | **string** |  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/x-www-form-urlencoded

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

