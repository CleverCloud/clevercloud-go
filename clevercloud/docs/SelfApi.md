# \SelfApi

All URIs are relative to *https://api.clever-cloud.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddEmailAddress**](SelfApi.md#AddEmailAddress) | **Put** /self/emails/{email} | 
[**AddSelfAddonTagByAddonId**](SelfApi.md#AddSelfAddonTagByAddonId) | **Put** /self/addons/{addonId}/tags/{tag} | 
[**AddSelfApplication**](SelfApi.md#AddSelfApplication) | **Post** /self/applications | 
[**AddSelfApplicationDependencyByAppId**](SelfApi.md#AddSelfApplicationDependencyByAppId) | **Put** /self/applications/{appId}/dependencies/{dependencyId} | 
[**AddSelfApplicationTagByAppId**](SelfApi.md#AddSelfApplicationTagByAppId) | **Put** /self/applications/{appId}/tags/{tag} | 
[**AddSelfPaymentMethod**](SelfApi.md#AddSelfPaymentMethod) | **Post** /self/payments/methods | 
[**AddSelfVhostByAppId**](SelfApi.md#AddSelfVhostByAppId) | **Put** /self/applications/{appId}/vhosts/{domain} | 
[**AddSshKey**](SelfApi.md#AddSshKey) | **Put** /self/keys/{key} | 
[**BuySelfDrops**](SelfApi.md#BuySelfDrops) | **Post** /self/payments/billings | 
[**CancelDeploy**](SelfApi.md#CancelDeploy) | **Delete** /self/applications/{appId}/deployments/{deploymentId}/instances | 
[**ChangeSelfAddonPlanByAddonId**](SelfApi.md#ChangeSelfAddonPlanByAddonId) | **Put** /self/addons/{addonId}/plan | 
[**ChangeUserPassword**](SelfApi.md#ChangeUserPassword) | **Put** /self/change_password | 
[**ChooseSelfPaymentProvider**](SelfApi.md#ChooseSelfPaymentProvider) | **Put** /self/payments/billings/{bid} | 
[**CreateMFA**](SelfApi.md#CreateMFA) | **Post** /self/mfa/{kind} | 
[**CreateSelfConsumer**](SelfApi.md#CreateSelfConsumer) | **Post** /self/consumers | 
[**DeleteMFA**](SelfApi.md#DeleteMFA) | **Delete** /self/mfa/{kind} | 
[**DeleteSelfAddonTagByAddonId**](SelfApi.md#DeleteSelfAddonTagByAddonId) | **Delete** /self/addons/{addonId}/tags/{tag} | 
[**DeleteSelfApplicationByAppId**](SelfApi.md#DeleteSelfApplicationByAppId) | **Delete** /self/applications/{appId} | 
[**DeleteSelfApplicationDependencyByAppId**](SelfApi.md#DeleteSelfApplicationDependencyByAppId) | **Delete** /self/applications/{appId}/dependencies/{dependencyId} | 
[**DeleteSelfApplicationTagAppId**](SelfApi.md#DeleteSelfApplicationTagAppId) | **Delete** /self/applications/{appId}/tags/{tag} | 
[**DeleteSelfCard**](SelfApi.md#DeleteSelfCard) | **Delete** /self/payments/methods/{mId} | 
[**DeleteSelfConsumer**](SelfApi.md#DeleteSelfConsumer) | **Delete** /self/consumers/{key} | 
[**DeleteSelfPurchaseOrder**](SelfApi.md#DeleteSelfPurchaseOrder) | **Delete** /self/payments/billings/{bid} | 
[**DeleteSelfRecurrentPayment**](SelfApi.md#DeleteSelfRecurrentPayment) | **Delete** /self/payments/recurring | 
[**DeleteUser**](SelfApi.md#DeleteUser) | **Delete** /self | 
[**DeprovisionSelfAddonById**](SelfApi.md#DeprovisionSelfAddonById) | **Delete** /self/addons/{addonId} | 
[**EditSelfApplicationByAppId**](SelfApi.md#EditSelfApplicationByAppId) | **Put** /self/applications/{appId} | 
[**EditSelfApplicationEnvByAppId**](SelfApi.md#EditSelfApplicationEnvByAppId) | **Put** /self/applications/{appId}/env | 
[**EditSelfApplicationEnvByAppIdAndEnvName**](SelfApi.md#EditSelfApplicationEnvByAppIdAndEnvName) | **Put** /self/applications/{appId}/env/{envName} | 
[**EditUser**](SelfApi.md#EditUser) | **Put** /self | 
[**FavMFA**](SelfApi.md#FavMFA) | **Put** /self/mfa/{kind} | 
[**GetAddonSSOData**](SelfApi.md#GetAddonSSOData) | **Get** /self/addons/{addonId}/sso | 
[**GetApplicationDeployment**](SelfApi.md#GetApplicationDeployment) | **Get** /self/applications/{appId}/deployments/{deploymentId} | 
[**GetApplicationDeployments**](SelfApi.md#GetApplicationDeployments) | **Get** /self/applications/{appId}/deployments | 
[**GetBackupCodes**](SelfApi.md#GetBackupCodes) | **Get** /self/mfa/{kind}/backupcodes | 
[**GetConfirmationEmail**](SelfApi.md#GetConfirmationEmail) | **Get** /self/confirmation_email | 
[**GetConsumptions**](SelfApi.md#GetConsumptions) | **Get** /self/consumptions | 
[**GetEmailAddresses**](SelfApi.md#GetEmailAddresses) | **Get** /self/emails | 
[**GetId**](SelfApi.md#GetId) | **Get** /self/id | 
[**GetSelfAddonById**](SelfApi.md#GetSelfAddonById) | **Get** /self/addons/{addonId} | 
[**GetSelfAddonEnvByAddonId**](SelfApi.md#GetSelfAddonEnvByAddonId) | **Get** /self/addons/{addonId}/env | 
[**GetSelfAddonTagsByAddonId**](SelfApi.md#GetSelfAddonTagsByAddonId) | **Get** /self/addons/{addonId}/tags | 
[**GetSelfAddons**](SelfApi.md#GetSelfAddons) | **Get** /self/addons | 
[**GetSelfAddonsLinkedToApplicationByAppId**](SelfApi.md#GetSelfAddonsLinkedToApplicationByAppId) | **Get** /self/applications/{appId}/addons | 
[**GetSelfAmount**](SelfApi.md#GetSelfAmount) | **Get** /self/credits | 
[**GetSelfApplicationBranchesByAppId**](SelfApi.md#GetSelfApplicationBranchesByAppId) | **Get** /self/applications/{appId}/branches | 
[**GetSelfApplicationByAppId**](SelfApi.md#GetSelfApplicationByAppId) | **Get** /self/applications/{appId} | 
[**GetSelfApplicationDependenciesByAppId**](SelfApi.md#GetSelfApplicationDependenciesByAppId) | **Get** /self/applications/{appId}/dependencies | 
[**GetSelfApplicationDependenciesEnvByAppId**](SelfApi.md#GetSelfApplicationDependenciesEnvByAppId) | **Get** /self/applications/{appId}/dependencies/env | 
[**GetSelfApplicationDependents**](SelfApi.md#GetSelfApplicationDependents) | **Get** /self/applications/{appId}/dependents | 
[**GetSelfApplicationEnvByAppId**](SelfApi.md#GetSelfApplicationEnvByAppId) | **Get** /self/applications/{appId}/env | 
[**GetSelfApplicationInstanceByAppAndInstanceId**](SelfApi.md#GetSelfApplicationInstanceByAppAndInstanceId) | **Get** /self/applications/{appId}/instances/{instanceId} | 
[**GetSelfApplicationInstancesByAppId**](SelfApi.md#GetSelfApplicationInstancesByAppId) | **Get** /self/applications/{appId}/instances | 
[**GetSelfApplicationTagsByAppId**](SelfApi.md#GetSelfApplicationTagsByAppId) | **Get** /self/applications/{appId}/tags | 
[**GetSelfApplications**](SelfApi.md#GetSelfApplications) | **Get** /self/applications | 
[**GetSelfApplicationsLinkedToAddonByAddonId**](SelfApi.md#GetSelfApplicationsLinkedToAddonByAddonId) | **Get** /self/addons/{addonId}/applications | 
[**GetSelfCliTokens**](SelfApi.md#GetSelfCliTokens) | **Get** /self/cli_tokens | 
[**GetSelfConsumer**](SelfApi.md#GetSelfConsumer) | **Get** /self/consumers/{key} | 
[**GetSelfConsumerSecret**](SelfApi.md#GetSelfConsumerSecret) | **Get** /self/consumers/{key}/secret | 
[**GetSelfConsumers**](SelfApi.md#GetSelfConsumers) | **Get** /self/consumers | 
[**GetSelfDefaultMethod**](SelfApi.md#GetSelfDefaultMethod) | **Get** /self/payments/methods/default | 
[**GetSelfEnvOfAddonsLinkedToApplicationByAppId**](SelfApi.md#GetSelfEnvOfAddonsLinkedToApplicationByAppId) | **Get** /self/applications/{appId}/addons/env | 
[**GetSelfExposedEnvByAppId**](SelfApi.md#GetSelfExposedEnvByAppId) | **Get** /self/applications/{appId}/exposed_env | 
[**GetSelfFavouriteVhostByAppId**](SelfApi.md#GetSelfFavouriteVhostByAppId) | **Get** /self/applications/{appId}/vhosts/favourite | 
[**GetSelfInstancesForAllApps**](SelfApi.md#GetSelfInstancesForAllApps) | **Get** /self/instances | 
[**GetSelfInvoiceById**](SelfApi.md#GetSelfInvoiceById) | **Get** /self/payments/billings/{bid} | 
[**GetSelfInvoices**](SelfApi.md#GetSelfInvoices) | **Get** /self/payments/billings | 
[**GetSelfMonthlyInvoice**](SelfApi.md#GetSelfMonthlyInvoice) | **Get** /self/payments/monthlyinvoice | 
[**GetSelfPaymentInfo**](SelfApi.md#GetSelfPaymentInfo) | **Get** /self/payment-info | 
[**GetSelfPaymentMethods**](SelfApi.md#GetSelfPaymentMethods) | **Get** /self/payments/methods | 
[**GetSelfPdfInvoiceById**](SelfApi.md#GetSelfPdfInvoiceById) | **Get** /self/payments/billings/{bid}.pdf | 
[**GetSelfPriceWithTax**](SelfApi.md#GetSelfPriceWithTax) | **Get** /self/payments/fullprice/{price} | 
[**GetSelfRecurrentPayment**](SelfApi.md#GetSelfRecurrentPayment) | **Get** /self/payments/recurring | 
[**GetSelfStripeToken**](SelfApi.md#GetSelfStripeToken) | **Get** /self/payments/tokens/stripe | 
[**GetSelfTokens**](SelfApi.md#GetSelfTokens) | **Get** /self/tokens | 
[**GetSelfVhostByAppId**](SelfApi.md#GetSelfVhostByAppId) | **Get** /self/applications/{appId}/vhosts | 
[**GetSshKeys**](SelfApi.md#GetSshKeys) | **Get** /self/keys | 
[**GetSummary**](SelfApi.md#GetSummary) | **Get** /summary | 
[**GetUser**](SelfApi.md#GetUser) | **Get** /self | 
[**LinkSelfAddonToApplicationByAppId**](SelfApi.md#LinkSelfAddonToApplicationByAppId) | **Post** /self/applications/{appId}/addons | 
[**MarkSelfFavouriteVhostByAppId**](SelfApi.md#MarkSelfFavouriteVhostByAppId) | **Put** /self/applications/{appId}/vhosts/favourite | 
[**PreorderSelfAddon**](SelfApi.md#PreorderSelfAddon) | **Post** /self/addons/preorders | 
[**ProvisionSelfAddon**](SelfApi.md#ProvisionSelfAddon) | **Post** /self/addons | 
[**RedeploySelfApplicationByAppId**](SelfApi.md#RedeploySelfApplicationByAppId) | **Post** /self/applications/{appId}/instances | 
[**RemoveEmailAddress**](SelfApi.md#RemoveEmailAddress) | **Delete** /self/emails/{email} | 
[**RemoveSelfApplicationEnvByAppIdAndEnvName**](SelfApi.md#RemoveSelfApplicationEnvByAppIdAndEnvName) | **Delete** /self/applications/{appId}/env/{envName} | 
[**RemoveSelfVhostByAppId**](SelfApi.md#RemoveSelfVhostByAppId) | **Delete** /self/applications/{appId}/vhosts/{domain} | 
[**RemoveSshKey**](SelfApi.md#RemoveSshKey) | **Delete** /self/keys/{key} | 
[**RenameAddon**](SelfApi.md#RenameAddon) | **Put** /self/addons/{addonId} | 
[**RevokeAllTokens**](SelfApi.md#RevokeAllTokens) | **Delete** /self/tokens | 
[**RevokeToken**](SelfApi.md#RevokeToken) | **Delete** /self/tokens/{token} | 
[**SetSelfApplicationBranchByAppId**](SelfApi.md#SetSelfApplicationBranchByAppId) | **Put** /self/applications/{appId}/branch | 
[**SetSelfBuildInstanceFlavorByAppId**](SelfApi.md#SetSelfBuildInstanceFlavorByAppId) | **Put** /self/applications/{appId}/buildflavor | 
[**SetSelfDefaultMethod**](SelfApi.md#SetSelfDefaultMethod) | **Put** /self/payments/methods/default | 
[**SetSelfMaxCreditsPerMonth**](SelfApi.md#SetSelfMaxCreditsPerMonth) | **Put** /self/payments/monthlyinvoice/maxcredit | 
[**SetUserAvatarFromString**](SelfApi.md#SetUserAvatarFromString) | **Put** /self/avatar | 
[**UndeploySelfApplicationByAppId**](SelfApi.md#UndeploySelfApplicationByAppId) | **Delete** /self/applications/{appId}/instances | 
[**UnlinkSelfddonFromApplicationByAppAndAddonId**](SelfApi.md#UnlinkSelfddonFromApplicationByAppAndAddonId) | **Delete** /self/applications/{appId}/addons/{addonId} | 
[**UnmarkSelfFavouriteVhostByAppId**](SelfApi.md#UnmarkSelfFavouriteVhostByAppId) | **Delete** /self/applications/{appId}/vhosts/favourite | 
[**UpdateSelfConsumer**](SelfApi.md#UpdateSelfConsumer) | **Put** /self/consumers/{key} | 
[**UpdateSelfExposedEnvByAppId**](SelfApi.md#UpdateSelfExposedEnvByAppId) | **Put** /self/applications/{appId}/exposed_env | 
[**ValidateEmail**](SelfApi.md#ValidateEmail) | **Get** /self/validate_email | 
[**ValidateMFA**](SelfApi.md#ValidateMFA) | **Post** /self/mfa/{kind}/confirmation | 



## AddEmailAddress

> Message AddEmailAddress(ctx, email).Body(body).Execute()



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
    email := "email_example" // string | 
    body := "body_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SelfApi.AddEmailAddress(context.Background(), email).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SelfApi.AddEmailAddress``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddEmailAddress`: Message
    fmt.Fprintf(os.Stdout, "Response from `SelfApi.AddEmailAddress`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**email** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddEmailAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **string** |  | 

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


## AddSelfAddonTagByAddonId

> []string AddSelfAddonTagByAddonId(ctx, addonId, tag).Execute()



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
    addonId := "addonId_example" // string | 
    tag := "tag_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SelfApi.AddSelfAddonTagByAddonId(context.Background(), addonId, tag).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SelfApi.AddSelfAddonTagByAddonId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddSelfAddonTagByAddonId`: []string
    fmt.Fprintf(os.Stdout, "Response from `SelfApi.AddSelfAddonTagByAddonId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**addonId** | **string** |  | 
**tag** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddSelfAddonTagByAddonIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## AddSelfApplication

> ApplicationView AddSelfApplication(ctx).WannabeApplication(wannabeApplication).Execute()



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
    wannabeApplication := *openapiclient.NewWannabeApplication() // WannabeApplication | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SelfApi.AddSelfApplication(context.Background()).WannabeApplication(wannabeApplication).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SelfApi.AddSelfApplication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddSelfApplication`: ApplicationView
    fmt.Fprintf(os.Stdout, "Response from `SelfApi.AddSelfApplication`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddSelfApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **wannabeApplication** | [**WannabeApplication**](WannabeApplication.md) |  | 

### Return type

[**ApplicationView**](ApplicationView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddSelfApplicationDependencyByAppId

> Message AddSelfApplicationDependencyByAppId(ctx, appId, dependencyId).Execute()



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
    dependencyId := "dependencyId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SelfApi.AddSelfApplicationDependencyByAppId(context.Background(), appId, dependencyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SelfApi.AddSelfApplicationDependencyByAppId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddSelfApplicationDependencyByAppId`: Message
    fmt.Fprintf(os.Stdout, "Response from `SelfApi.AddSelfApplicationDependencyByAppId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** |  | 
**dependencyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddSelfApplicationDependencyByAppIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## AddSelfApplicationTagByAppId

> []string AddSelfApplicationTagByAppId(ctx, appId, tag).Execute()



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
    tag := "tag_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SelfApi.AddSelfApplicationTagByAppId(context.Background(), appId, tag).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SelfApi.AddSelfApplicationTagByAppId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddSelfApplicationTagByAppId`: []string
    fmt.Fprintf(os.Stdout, "Response from `SelfApi.AddSelfApplicationTagByAppId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** |  | 
**tag** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddSelfApplicationTagByAppIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## AddSelfPaymentMethod

> PaymentMethodView AddSelfPaymentMethod(ctx).PaymentData(paymentData).Execute()



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
    paymentData := *openapiclient.NewPaymentData() // PaymentData | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SelfApi.AddSelfPaymentMethod(context.Background()).PaymentData(paymentData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SelfApi.AddSelfPaymentMethod``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddSelfPaymentMethod`: PaymentMethodView
    fmt.Fprintf(os.Stdout, "Response from `SelfApi.AddSelfPaymentMethod`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddSelfPaymentMethodRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **paymentData** | [**PaymentData**](PaymentData.md) |  | 

### Return type

[**PaymentMethodView**](PaymentMethodView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddSelfVhostByAppId

> Message AddSelfVhostByAppId(ctx, appId, domain).Execute()



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
    domain := "domain_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SelfApi.AddSelfVhostByAppId(context.Background(), appId, domain).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SelfApi.AddSelfVhostByAppId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddSelfVhostByAppId`: Message
    fmt.Fprintf(os.Stdout, "Response from `SelfApi.AddSelfVhostByAppId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** |  | 
**domain** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddSelfVhostByAppIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## AddSshKey

> Message AddSshKey(ctx, key).Execute()



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SelfApi.AddSshKey(context.Background(), key).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SelfApi.AddSshKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddSshKey`: Message
    fmt.Fprintf(os.Stdout, "Response from `SelfApi.AddSshKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddSshKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## BuySelfDrops

> InvoiceRendering BuySelfDrops(ctx).WannaBuyPackage(wannaBuyPackage).Execute()



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
    wannaBuyPackage := *openapiclient.NewWannaBuyPackage() // WannaBuyPackage | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SelfApi.BuySelfDrops(context.Background()).WannaBuyPackage(wannaBuyPackage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SelfApi.BuySelfDrops``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BuySelfDrops`: InvoiceRendering
    fmt.Fprintf(os.Stdout, "Response from `SelfApi.BuySelfDrops`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBuySelfDropsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **wannaBuyPackage** | [**WannaBuyPackage**](WannaBuyPackage.md) |  | 

### Return type

[**InvoiceRendering**](InvoiceRendering.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CancelDeploy

> Message CancelDeploy(ctx, appId, deploymentId).Execute()



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
    deploymentId := "deploymentId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SelfApi.CancelDeploy(context.Background(), appId, deploymentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SelfApi.CancelDeploy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CancelDeploy`: Message
    fmt.Fprintf(os.Stdout, "Response from `SelfApi.CancelDeploy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** |  | 
**deploymentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelDeployRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## ChangeSelfAddonPlanByAddonId

> AddonView ChangeSelfAddonPlanByAddonId(ctx, addonId).WannabePlanChange(wannabePlanChange).Execute()



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
    addonId := "addonId_example" // string | 
    wannabePlanChange := *openapiclient.NewWannabePlanChange() // WannabePlanChange | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SelfApi.ChangeSelfAddonPlanByAddonId(context.Background(), addonId).WannabePlanChange(wannabePlanChange).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SelfApi.ChangeSelfAddonPlanByAddonId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ChangeSelfAddonPlanByAddonId`: AddonView
    fmt.Fprintf(os.Stdout, "Response from `SelfApi.ChangeSelfAddonPlanByAddonId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**addonId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiChangeSelfAddonPlanByAddonIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **wannabePlanChange** | [**WannabePlanChange**](WannabePlanChange.md) |  | 

### Return type

[**AddonView**](AddonView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChangeUserPassword

> Message ChangeUserPassword(ctx).WannabePassword(wannabePassword).Execute()



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
    wannabePassword := *openapiclient.NewWannabePassword() // WannabePassword | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SelfApi.ChangeUserPassword(context.Background()).WannabePassword(wannabePassword).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SelfApi.ChangeUserPassword``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ChangeUserPassword`: Message
    fmt.Fprintf(os.Stdout, "Response from `SelfApi.ChangeUserPassword`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiChangeUserPasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **wannabePassword** | [**WannabePassword**](WannabePassword.md) |  | 

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


## ChooseSelfPaymentProvider

> NextInPaymentFlow ChooseSelfPaymentProvider(ctx, bid).PaymentProviderSelection(paymentProviderSelection).Execute()



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
    paymentProviderSelection := *openapiclient.NewPaymentProviderSelection() // PaymentProviderSelection | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SelfApi.ChooseSelfPaymentProvider(context.Background(), bid).PaymentProviderSelection(paymentProviderSelection).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SelfApi.ChooseSelfPaymentProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ChooseSelfPaymentProvider`: NextInPaymentFlow
    fmt.Fprintf(os.Stdout, "Response from `SelfApi.ChooseSelfPaymentProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiChooseSelfPaymentProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **paymentProviderSelection** | [**PaymentProviderSelection**](PaymentProviderSelection.md) |  | 

### Return type

[**NextInPaymentFlow**](NextInPaymentFlow.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateMFA

> CreateMFA(ctx, kind).Execute()



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
    kind := "kind_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SelfApi.CreateMFA(context.Background(), kind).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SelfApi.CreateMFA``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**kind** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateMFARequest struct via the builder pattern


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


## CreateSelfConsumer

> OAuth1ConsumerView CreateSelfConsumer(ctx).WannabeOAuth1Consumer(wannabeOAuth1Consumer).Execute()



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
    wannabeOAuth1Consumer := *openapiclient.NewWannabeOAuth1Consumer() // WannabeOAuth1Consumer | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SelfApi.CreateSelfConsumer(context.Background()).WannabeOAuth1Consumer(wannabeOAuth1Consumer).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SelfApi.CreateSelfConsumer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSelfConsumer`: OAuth1ConsumerView
    fmt.Fprintf(os.Stdout, "Response from `SelfApi.CreateSelfConsumer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSelfConsumerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **wannabeOAuth1Consumer** | [**WannabeOAuth1Consumer**](WannabeOAuth1Consumer.md) |  | 

### Return type

[**OAuth1ConsumerView**](OAuth1ConsumerView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMFA

> DeleteMFA(ctx, kind).Execute()



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
    kind := "kind_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SelfApi.DeleteMFA(context.Background(), kind).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SelfApi.DeleteMFA``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**kind** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMFARequest struct via the builder pattern


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


## DeleteSelfAddonTagByAddonId

> []string DeleteSelfAddonTagByAddonId(ctx, addonId, tag).Execute()



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
    addonId := "addonId_example" // string | 
    tag := "tag_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SelfApi.DeleteSelfAddonTagByAddonId(context.Background(), addonId, tag).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SelfApi.DeleteSelfAddonTagByAddonId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteSelfAddonTagByAddonId`: []string
    fmt.Fprintf(os.Stdout, "Response from `SelfApi.DeleteSelfAddonTagByAddonId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**addonId** | **string** |  | 
**tag** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSelfAddonTagByAddonIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## DeleteSelfApplicationByAppId

> Message DeleteSelfApplicationByAppId(ctx, appId).Execute()



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SelfApi.DeleteSelfApplicationByAppId(context.Background(), appId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SelfApi.DeleteSelfApplicationByAppId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteSelfApplicationByAppId`: Message
    fmt.Fprintf(os.Stdout, "Response from `SelfApi.DeleteSelfApplicationByAppId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSelfApplicationByAppIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## DeleteSelfApplicationDependencyByAppId

> DeleteSelfApplicationDependencyByAppId(ctx, appId, dependencyId).Execute()



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
    dependencyId := "dependencyId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SelfApi.DeleteSelfApplicationDependencyByAppId(context.Background(), appId, dependencyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SelfApi.DeleteSelfApplicationDependencyByAppId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** |  | 
**dependencyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSelfApplicationDependencyByAppIdRequest struct via the builder pattern


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


## DeleteSelfApplicationTagAppId

> []string DeleteSelfApplicationTagAppId(ctx, appId, tag).Execute()



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
    tag := "tag_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SelfApi.DeleteSelfApplicationTagAppId(context.Background(), appId, tag).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SelfApi.DeleteSelfApplicationTagAppId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteSelfApplicationTagAppId`: []string
    fmt.Fprintf(os.Stdout, "Response from `SelfApi.DeleteSelfApplicationTagAppId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** |  | 
**tag** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSelfApplicationTagAppIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## DeleteSelfCard

> DeleteSelfCard(ctx, mId).Execute()



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
    mId := "mId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SelfApi.DeleteSelfCard(context.Background(), mId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SelfApi.DeleteSelfCard``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSelfCardRequest struct via the builder pattern


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


## DeleteSelfConsumer

> DeleteSelfConsumer(ctx, key).Execute()



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SelfApi.DeleteSelfConsumer(context.Background(), key).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SelfApi.DeleteSelfConsumer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSelfConsumerRequest struct via the builder pattern


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


## DeleteSelfPurchaseOrder

> DeleteSelfPurchaseOrder(ctx, bid).Execute()



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
    resp, r, err := api_client.SelfApi.DeleteSelfPurchaseOrder(context.Background(), bid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SelfApi.DeleteSelfPurchaseOrder``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteSelfPurchaseOrderRequest struct via the builder pattern


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


## DeleteSelfRecurrentPayment

> DeleteSelfRecurrentPayment(ctx).Execute()



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
    resp, r, err := api_client.SelfApi.DeleteSelfRecurrentPayment(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SelfApi.DeleteSelfRecurrentPayment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSelfRecurrentPaymentRequest struct via the builder pattern


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


## DeleteUser

> Message DeleteUser(ctx).Execute()



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
    resp, r, err := api_client.SelfApi.DeleteUser(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SelfApi.DeleteUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteUser`: Message
    fmt.Fprintf(os.Stdout, "Response from `SelfApi.DeleteUser`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUserRequest struct via the builder pattern


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


## DeprovisionSelfAddonById

> Message DeprovisionSelfAddonById(ctx, addonId).Execute()



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
    addonId := "addonId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SelfApi.DeprovisionSelfAddonById(context.Background(), addonId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SelfApi.DeprovisionSelfAddonById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeprovisionSelfAddonById`: Message
    fmt.Fprintf(os.Stdout, "Response from `SelfApi.DeprovisionSelfAddonById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**addonId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeprovisionSelfAddonByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## EditSelfApplicationByAppId

> ApplicationView EditSelfApplicationByAppId(ctx, appId).WannabeApplication(wannabeApplication).Execute()



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
    wannabeApplication := *openapiclient.NewWannabeApplication() // WannabeApplication | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SelfApi.EditSelfApplicationByAppId(context.Background(), appId).WannabeApplication(wannabeApplication).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SelfApi.EditSelfApplicationByAppId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditSelfApplicationByAppId`: ApplicationView
    fmt.Fprintf(os.Stdout, "Response from `SelfApi.EditSelfApplicationByAppId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditSelfApplicationByAppIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **wannabeApplication** | [**WannabeApplication**](WannabeApplication.md) |  | 

### Return type

[**ApplicationView**](ApplicationView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditSelfApplicationEnvByAppId

> ApplicationView EditSelfApplicationEnvByAppId(ctx, appId).Body(body).Execute()



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SelfApi.EditSelfApplicationEnvByAppId(context.Background(), appId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SelfApi.EditSelfApplicationEnvByAppId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditSelfApplicationEnvByAppId`: ApplicationView
    fmt.Fprintf(os.Stdout, "Response from `SelfApi.EditSelfApplicationEnvByAppId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditSelfApplicationEnvByAppIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **string** |  | 

### Return type

[**ApplicationView**](ApplicationView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditSelfApplicationEnvByAppIdAndEnvName

> ApplicationView EditSelfApplicationEnvByAppIdAndEnvName(ctx, appId, envName).WannabeValue(wannabeValue).Execute()



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
    envName := "envName_example" // string | 
    wannabeValue := *openapiclient.NewWannabeValue() // WannabeValue | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SelfApi.EditSelfApplicationEnvByAppIdAndEnvName(context.Background(), appId, envName).WannabeValue(wannabeValue).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SelfApi.EditSelfApplicationEnvByAppIdAndEnvName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditSelfApplicationEnvByAppIdAndEnvName`: ApplicationView
    fmt.Fprintf(os.Stdout, "Response from `SelfApi.EditSelfApplicationEnvByAppIdAndEnvName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** |  | 
**envName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditSelfApplicationEnvByAppIdAndEnvNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **wannabeValue** | [**WannabeValue**](WannabeValue.md) |  | 

### Return type

[**ApplicationView**](ApplicationView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditUser

> UserView EditUser(ctx).WannabeUser(wannabeUser).Execute()



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
    wannabeUser := *openapiclient.NewWannabeUser() // WannabeUser | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SelfApi.EditUser(context.Background()).WannabeUser(wannabeUser).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SelfApi.EditUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditUser`: UserView
    fmt.Fprintf(os.Stdout, "Response from `SelfApi.EditUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEditUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **wannabeUser** | [**WannabeUser**](WannabeUser.md) |  | 

### Return type

[**UserView**](UserView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FavMFA

> FavMFA(ctx, kind).WannabeMFAFav(wannabeMFAFav).Execute()



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
    kind := "kind_example" // string | 
    wannabeMFAFav := *openapiclient.NewWannabeMFAFav() // WannabeMFAFav | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SelfApi.FavMFA(context.Background(), kind).WannabeMFAFav(wannabeMFAFav).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SelfApi.FavMFA``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**kind** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFavMFARequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **wannabeMFAFav** | [**WannabeMFAFav**](WannabeMFAFav.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAddonSSOData

> AddonSSOData GetAddonSSOData(ctx, addonId).Execute()



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
    addonId := "addonId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SelfApi.GetAddonSSOData(context.Background(), addonId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SelfApi.GetAddonSSOData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAddonSSOData`: AddonSSOData
    fmt.Fprintf(os.Stdout, "Response from `SelfApi.GetAddonSSOData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**addonId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAddonSSODataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AddonSSOData**](AddonSSOData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplicationDeployment

> GetApplicationDeployment(ctx, appId, deploymentId).Execute()



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
    deploymentId := "deploymentId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SelfApi.GetApplicationDeployment(context.Background(), appId, deploymentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SelfApi.GetApplicationDeployment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** |  | 
**deploymentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationDeploymentRequest struct via the builder pattern


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


## GetApplicationDeployments

> []DeploymentView GetApplicationDeployments(ctx, appId).Limit(limit).Offset(offset).Action(action).Execute()



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
    limit := "limit_example" // string |  (optional)
    offset := "offset_example" // string |  (optional)
    action := "action_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SelfApi.GetApplicationDeployments(context.Background(), appId).Limit(limit).Offset(offset).Action(action).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SelfApi.GetApplicationDeployments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApplicationDeployments`: []DeploymentView
    fmt.Fprintf(os.Stdout, "Response from `SelfApi.GetApplicationDeployments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationDeploymentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **string** |  | 
 **offset** | **string** |  | 
 **action** | **string** |  | 

### Return type

[**[]DeploymentView**](DeploymentView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBackupCodes

> []MFARecoveryCode GetBackupCodes(ctx, kind).Execute()



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
    kind := "kind_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SelfApi.GetBackupCodes(context.Background(), kind).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SelfApi.GetBackupCodes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBackupCodes`: []MFARecoveryCode
    fmt.Fprintf(os.Stdout, "Response from `SelfApi.GetBackupCodes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**kind** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBackupCodesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]MFARecoveryCode**](MFARecoveryCode.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConfirmationEmail

> Message GetConfirmationEmail(ctx).Execute()



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
    resp, r, err := api_client.SelfApi.GetConfirmationEmail(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SelfApi.GetConfirmationEmail``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetConfirmationEmail`: Message
    fmt.Fprintf(os.Stdout, "Response from `SelfApi.GetConfirmationEmail`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetConfirmationEmailRequest struct via the builder pattern


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


## GetConsumptions

> string GetConsumptions(ctx).AppId(appId).From(from).To(to).For_(for_).Execute()



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
    appId := "appId_example" // string |  (optional)
    from := "from_example" // string |  (optional)
    to := "to_example" // string |  (optional)
    for_ := "for__example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SelfApi.GetConsumptions(context.Background()).AppId(appId).From(from).To(to).For_(for_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SelfApi.GetConsumptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetConsumptions`: string
    fmt.Fprintf(os.Stdout, "Response from `SelfApi.GetConsumptions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetConsumptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **appId** | **string** |  | 
 **from** | **string** |  | 
 **to** | **string** |  | 
 **for_** | **string** |  | 

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


## GetEmailAddresses

> []string GetEmailAddresses(ctx).Execute()



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
    resp, r, err := api_client.SelfApi.GetEmailAddresses(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SelfApi.GetEmailAddresses``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEmailAddresses`: []string
    fmt.Fprintf(os.Stdout, "Response from `SelfApi.GetEmailAddresses`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetEmailAddressesRequest struct via the builder pattern


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


## GetId

> string GetId(ctx).Execute()



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
    resp, r, err := api_client.SelfApi.GetId(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SelfApi.GetId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetId`: string
    fmt.Fprintf(os.Stdout, "Response from `SelfApi.GetId`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetIdRequest struct via the builder pattern


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


## GetSelfAddonById

> AddonView GetSelfAddonById(ctx, addonId).Execute()



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
    addonId := "addonId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SelfApi.GetSelfAddonById(context.Background(), addonId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SelfApi.GetSelfAddonById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSelfAddonById`: AddonView
    fmt.Fprintf(os.Stdout, "Response from `SelfApi.GetSelfAddonById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**addonId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSelfAddonByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AddonView**](AddonView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSelfAddonEnvByAddonId

> []AddonEnvironmentView GetSelfAddonEnvByAddonId(ctx, addonId).Execute()



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
    addonId := "addonId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SelfApi.GetSelfAddonEnvByAddonId(context.Background(), addonId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SelfApi.GetSelfAddonEnvByAddonId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSelfAddonEnvByAddonId`: []AddonEnvironmentView
    fmt.Fprintf(os.Stdout, "Response from `SelfApi.GetSelfAddonEnvByAddonId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**addonId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSelfAddonEnvByAddonIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]AddonEnvironmentView**](AddonEnvironmentView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSelfAddonTagsByAddonId

> []string GetSelfAddonTagsByAddonId(ctx, addonId).Execute()



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
    addonId := "addonId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SelfApi.GetSelfAddonTagsByAddonId(context.Background(), addonId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SelfApi.GetSelfAddonTagsByAddonId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSelfAddonTagsByAddonId`: []string
    fmt.Fprintf(os.Stdout, "Response from `SelfApi.GetSelfAddonTagsByAddonId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**addonId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSelfAddonTagsByAddonIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## GetSelfAddons

> []AddonView GetSelfAddons(ctx).Execute()



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
    resp, r, err := api_client.SelfApi.GetSelfAddons(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SelfApi.GetSelfAddons``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSelfAddons`: []AddonView
    fmt.Fprintf(os.Stdout, "Response from `SelfApi.GetSelfAddons`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSelfAddonsRequest struct via the builder pattern


### Return type

[**[]AddonView**](AddonView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSelfAddonsLinkedToApplicationByAppId

> []AddonView GetSelfAddonsLinkedToApplicationByAppId(ctx, appId).Execute()



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SelfApi.GetSelfAddonsLinkedToApplicationByAppId(context.Background(), appId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SelfApi.GetSelfAddonsLinkedToApplicationByAppId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSelfAddonsLinkedToApplicationByAppId`: []AddonView
    fmt.Fprintf(os.Stdout, "Response from `SelfApi.GetSelfAddonsLinkedToApplicationByAppId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSelfAddonsLinkedToApplicationByAppIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]AddonView**](AddonView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSelfAmount

> DropCountView GetSelfAmount(ctx).Execute()



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
    resp, r, err := api_client.SelfApi.GetSelfAmount(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SelfApi.GetSelfAmount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSelfAmount`: DropCountView
    fmt.Fprintf(os.Stdout, "Response from `SelfApi.GetSelfAmount`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSelfAmountRequest struct via the builder pattern


### Return type

[**DropCountView**](DropCountView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSelfApplicationBranchesByAppId

> []string GetSelfApplicationBranchesByAppId(ctx, appId).Execute()



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SelfApi.GetSelfApplicationBranchesByAppId(context.Background(), appId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SelfApi.GetSelfApplicationBranchesByAppId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSelfApplicationBranchesByAppId`: []string
    fmt.Fprintf(os.Stdout, "Response from `SelfApi.GetSelfApplicationBranchesByAppId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSelfApplicationBranchesByAppIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## GetSelfApplicationByAppId

> ApplicationView GetSelfApplicationByAppId(ctx, appId).Execute()



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SelfApi.GetSelfApplicationByAppId(context.Background(), appId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SelfApi.GetSelfApplicationByAppId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSelfApplicationByAppId`: ApplicationView
    fmt.Fprintf(os.Stdout, "Response from `SelfApi.GetSelfApplicationByAppId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSelfApplicationByAppIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApplicationView**](ApplicationView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSelfApplicationDependenciesByAppId

> []ApplicationView GetSelfApplicationDependenciesByAppId(ctx, appId).Execute()



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SelfApi.GetSelfApplicationDependenciesByAppId(context.Background(), appId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SelfApi.GetSelfApplicationDependenciesByAppId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSelfApplicationDependenciesByAppId`: []ApplicationView
    fmt.Fprintf(os.Stdout, "Response from `SelfApi.GetSelfApplicationDependenciesByAppId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSelfApplicationDependenciesByAppIdRequest struct via the builder pattern


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


## GetSelfApplicationDependenciesEnvByAppId

> GetSelfApplicationDependenciesEnvByAppId(ctx, id, appId).Execute()



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
    appId := "appId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SelfApi.GetSelfApplicationDependenciesEnvByAppId(context.Background(), id, appId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SelfApi.GetSelfApplicationDependenciesEnvByAppId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**appId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSelfApplicationDependenciesEnvByAppIdRequest struct via the builder pattern


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


## GetSelfApplicationDependents

> []ApplicationView GetSelfApplicationDependents(ctx, appId).Execute()



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SelfApi.GetSelfApplicationDependents(context.Background(), appId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SelfApi.GetSelfApplicationDependents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSelfApplicationDependents`: []ApplicationView
    fmt.Fprintf(os.Stdout, "Response from `SelfApi.GetSelfApplicationDependents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSelfApplicationDependentsRequest struct via the builder pattern


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


## GetSelfApplicationEnvByAppId

> []AddonEnvironmentView GetSelfApplicationEnvByAppId(ctx, appId).Execute()



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SelfApi.GetSelfApplicationEnvByAppId(context.Background(), appId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SelfApi.GetSelfApplicationEnvByAppId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSelfApplicationEnvByAppId`: []AddonEnvironmentView
    fmt.Fprintf(os.Stdout, "Response from `SelfApi.GetSelfApplicationEnvByAppId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSelfApplicationEnvByAppIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]AddonEnvironmentView**](AddonEnvironmentView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSelfApplicationInstanceByAppAndInstanceId

> SuperNovaInstanceView GetSelfApplicationInstanceByAppAndInstanceId(ctx, appId, instanceId).Execute()



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
    instanceId := "instanceId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SelfApi.GetSelfApplicationInstanceByAppAndInstanceId(context.Background(), appId, instanceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SelfApi.GetSelfApplicationInstanceByAppAndInstanceId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSelfApplicationInstanceByAppAndInstanceId`: SuperNovaInstanceView
    fmt.Fprintf(os.Stdout, "Response from `SelfApi.GetSelfApplicationInstanceByAppAndInstanceId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** |  | 
**instanceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSelfApplicationInstanceByAppAndInstanceIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**SuperNovaInstanceView**](SuperNovaInstanceView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSelfApplicationInstancesByAppId

> []SuperNovaInstanceView GetSelfApplicationInstancesByAppId(ctx, appId).DeploymentId(deploymentId).WithDeleted(withDeleted).Execute()



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
    deploymentId := "deploymentId_example" // string |  (optional)
    withDeleted := "withDeleted_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SelfApi.GetSelfApplicationInstancesByAppId(context.Background(), appId).DeploymentId(deploymentId).WithDeleted(withDeleted).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SelfApi.GetSelfApplicationInstancesByAppId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSelfApplicationInstancesByAppId`: []SuperNovaInstanceView
    fmt.Fprintf(os.Stdout, "Response from `SelfApi.GetSelfApplicationInstancesByAppId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSelfApplicationInstancesByAppIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deploymentId** | **string** |  | 
 **withDeleted** | **string** |  | 

### Return type

[**[]SuperNovaInstanceView**](SuperNovaInstanceView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSelfApplicationTagsByAppId

> []string GetSelfApplicationTagsByAppId(ctx, appId).Execute()



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SelfApi.GetSelfApplicationTagsByAppId(context.Background(), appId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SelfApi.GetSelfApplicationTagsByAppId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSelfApplicationTagsByAppId`: []string
    fmt.Fprintf(os.Stdout, "Response from `SelfApi.GetSelfApplicationTagsByAppId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSelfApplicationTagsByAppIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## GetSelfApplications

> []ApplicationView GetSelfApplications(ctx).Execute()



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
    resp, r, err := api_client.SelfApi.GetSelfApplications(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SelfApi.GetSelfApplications``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSelfApplications`: []ApplicationView
    fmt.Fprintf(os.Stdout, "Response from `SelfApi.GetSelfApplications`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSelfApplicationsRequest struct via the builder pattern


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


## GetSelfApplicationsLinkedToAddonByAddonId

> []ApplicationView GetSelfApplicationsLinkedToAddonByAddonId(ctx, addonId).Execute()



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
    addonId := "addonId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SelfApi.GetSelfApplicationsLinkedToAddonByAddonId(context.Background(), addonId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SelfApi.GetSelfApplicationsLinkedToAddonByAddonId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSelfApplicationsLinkedToAddonByAddonId`: []ApplicationView
    fmt.Fprintf(os.Stdout, "Response from `SelfApi.GetSelfApplicationsLinkedToAddonByAddonId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**addonId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSelfApplicationsLinkedToAddonByAddonIdRequest struct via the builder pattern


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


## GetSelfCliTokens

> []CliTokenView GetSelfCliTokens(ctx).CliToken(cliToken).Execute()



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
    cliToken := "cliToken_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SelfApi.GetSelfCliTokens(context.Background()).CliToken(cliToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SelfApi.GetSelfCliTokens``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSelfCliTokens`: []CliTokenView
    fmt.Fprintf(os.Stdout, "Response from `SelfApi.GetSelfCliTokens`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSelfCliTokensRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cliToken** | **string** |  | 

### Return type

[**[]CliTokenView**](CliTokenView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSelfConsumer

> OAuth1ConsumerView GetSelfConsumer(ctx, key).Execute()



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SelfApi.GetSelfConsumer(context.Background(), key).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SelfApi.GetSelfConsumer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSelfConsumer`: OAuth1ConsumerView
    fmt.Fprintf(os.Stdout, "Response from `SelfApi.GetSelfConsumer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSelfConsumerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OAuth1ConsumerView**](OAuth1ConsumerView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSelfConsumerSecret

> SecretView GetSelfConsumerSecret(ctx, key).Execute()



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SelfApi.GetSelfConsumerSecret(context.Background(), key).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SelfApi.GetSelfConsumerSecret``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSelfConsumerSecret`: SecretView
    fmt.Fprintf(os.Stdout, "Response from `SelfApi.GetSelfConsumerSecret`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSelfConsumerSecretRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SecretView**](SecretView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSelfConsumers

> GetSelfConsumers(ctx).Execute()



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
    resp, r, err := api_client.SelfApi.GetSelfConsumers(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SelfApi.GetSelfConsumers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSelfConsumersRequest struct via the builder pattern


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


## GetSelfDefaultMethod

> PaymentMethodView GetSelfDefaultMethod(ctx).Execute()



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
    resp, r, err := api_client.SelfApi.GetSelfDefaultMethod(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SelfApi.GetSelfDefaultMethod``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSelfDefaultMethod`: PaymentMethodView
    fmt.Fprintf(os.Stdout, "Response from `SelfApi.GetSelfDefaultMethod`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSelfDefaultMethodRequest struct via the builder pattern


### Return type

[**PaymentMethodView**](PaymentMethodView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSelfEnvOfAddonsLinkedToApplicationByAppId

> []LinkedAddonEnvironmentView GetSelfEnvOfAddonsLinkedToApplicationByAppId(ctx, appId).Execute()



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SelfApi.GetSelfEnvOfAddonsLinkedToApplicationByAppId(context.Background(), appId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SelfApi.GetSelfEnvOfAddonsLinkedToApplicationByAppId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSelfEnvOfAddonsLinkedToApplicationByAppId`: []LinkedAddonEnvironmentView
    fmt.Fprintf(os.Stdout, "Response from `SelfApi.GetSelfEnvOfAddonsLinkedToApplicationByAppId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSelfEnvOfAddonsLinkedToApplicationByAppIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]LinkedAddonEnvironmentView**](LinkedAddonEnvironmentView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSelfExposedEnvByAppId

> string GetSelfExposedEnvByAppId(ctx, appId).Execute()



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SelfApi.GetSelfExposedEnvByAppId(context.Background(), appId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SelfApi.GetSelfExposedEnvByAppId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSelfExposedEnvByAppId`: string
    fmt.Fprintf(os.Stdout, "Response from `SelfApi.GetSelfExposedEnvByAppId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSelfExposedEnvByAppIdRequest struct via the builder pattern


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


## GetSelfFavouriteVhostByAppId

> VhostView GetSelfFavouriteVhostByAppId(ctx, appId).Execute()



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SelfApi.GetSelfFavouriteVhostByAppId(context.Background(), appId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SelfApi.GetSelfFavouriteVhostByAppId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSelfFavouriteVhostByAppId`: VhostView
    fmt.Fprintf(os.Stdout, "Response from `SelfApi.GetSelfFavouriteVhostByAppId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSelfFavouriteVhostByAppIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**VhostView**](VhostView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSelfInstancesForAllApps

> SuperNovaInstanceMap GetSelfInstancesForAllApps(ctx).Execute()



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
    resp, r, err := api_client.SelfApi.GetSelfInstancesForAllApps(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SelfApi.GetSelfInstancesForAllApps``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSelfInstancesForAllApps`: SuperNovaInstanceMap
    fmt.Fprintf(os.Stdout, "Response from `SelfApi.GetSelfInstancesForAllApps`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSelfInstancesForAllAppsRequest struct via the builder pattern


### Return type

[**SuperNovaInstanceMap**](SuperNovaInstanceMap.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSelfInvoiceById

> InvoiceRendering GetSelfInvoiceById(ctx, bid).Execute()



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
    resp, r, err := api_client.SelfApi.GetSelfInvoiceById(context.Background(), bid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SelfApi.GetSelfInvoiceById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSelfInvoiceById`: InvoiceRendering
    fmt.Fprintf(os.Stdout, "Response from `SelfApi.GetSelfInvoiceById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSelfInvoiceByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InvoiceRendering**](InvoiceRendering.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSelfInvoices

> []InvoiceRendering GetSelfInvoices(ctx).Execute()



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
    resp, r, err := api_client.SelfApi.GetSelfInvoices(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SelfApi.GetSelfInvoices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSelfInvoices`: []InvoiceRendering
    fmt.Fprintf(os.Stdout, "Response from `SelfApi.GetSelfInvoices`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSelfInvoicesRequest struct via the builder pattern


### Return type

[**[]InvoiceRendering**](InvoiceRendering.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSelfMonthlyInvoice

> string GetSelfMonthlyInvoice(ctx).Execute()



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
    resp, r, err := api_client.SelfApi.GetSelfMonthlyInvoice(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SelfApi.GetSelfMonthlyInvoice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSelfMonthlyInvoice`: string
    fmt.Fprintf(os.Stdout, "Response from `SelfApi.GetSelfMonthlyInvoice`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSelfMonthlyInvoiceRequest struct via the builder pattern


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


## GetSelfPaymentInfo

> PaymentInfoView GetSelfPaymentInfo(ctx).Execute()



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
    resp, r, err := api_client.SelfApi.GetSelfPaymentInfo(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SelfApi.GetSelfPaymentInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSelfPaymentInfo`: PaymentInfoView
    fmt.Fprintf(os.Stdout, "Response from `SelfApi.GetSelfPaymentInfo`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSelfPaymentInfoRequest struct via the builder pattern


### Return type

[**PaymentInfoView**](PaymentInfoView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSelfPaymentMethods

> []PaymentMethodView GetSelfPaymentMethods(ctx).Execute()



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
    resp, r, err := api_client.SelfApi.GetSelfPaymentMethods(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SelfApi.GetSelfPaymentMethods``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSelfPaymentMethods`: []PaymentMethodView
    fmt.Fprintf(os.Stdout, "Response from `SelfApi.GetSelfPaymentMethods`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSelfPaymentMethodsRequest struct via the builder pattern


### Return type

[**[]PaymentMethodView**](PaymentMethodView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSelfPdfInvoiceById

> GetSelfPdfInvoiceById(ctx, bid).Token(token).Execute()



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
    token := "token_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SelfApi.GetSelfPdfInvoiceById(context.Background(), bid).Token(token).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SelfApi.GetSelfPdfInvoiceById``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetSelfPdfInvoiceByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **token** | **string** |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/pdf

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSelfPriceWithTax

> PriceWithTaxInfo GetSelfPriceWithTax(ctx, price).Execute()



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
    price := "price_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SelfApi.GetSelfPriceWithTax(context.Background(), price).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SelfApi.GetSelfPriceWithTax``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSelfPriceWithTax`: PriceWithTaxInfo
    fmt.Fprintf(os.Stdout, "Response from `SelfApi.GetSelfPriceWithTax`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**price** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSelfPriceWithTaxRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PriceWithTaxInfo**](PriceWithTaxInfo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSelfRecurrentPayment

> RecurrentPaymentView GetSelfRecurrentPayment(ctx).Execute()



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
    resp, r, err := api_client.SelfApi.GetSelfRecurrentPayment(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SelfApi.GetSelfRecurrentPayment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSelfRecurrentPayment`: RecurrentPaymentView
    fmt.Fprintf(os.Stdout, "Response from `SelfApi.GetSelfRecurrentPayment`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSelfRecurrentPaymentRequest struct via the builder pattern


### Return type

[**RecurrentPaymentView**](RecurrentPaymentView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSelfStripeToken

> BraintreeToken GetSelfStripeToken(ctx).Execute()



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
    resp, r, err := api_client.SelfApi.GetSelfStripeToken(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SelfApi.GetSelfStripeToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSelfStripeToken`: BraintreeToken
    fmt.Fprintf(os.Stdout, "Response from `SelfApi.GetSelfStripeToken`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSelfStripeTokenRequest struct via the builder pattern


### Return type

[**BraintreeToken**](BraintreeToken.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSelfTokens

> []OAuth1AccessTokenView GetSelfTokens(ctx).Execute()



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
    resp, r, err := api_client.SelfApi.GetSelfTokens(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SelfApi.GetSelfTokens``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSelfTokens`: []OAuth1AccessTokenView
    fmt.Fprintf(os.Stdout, "Response from `SelfApi.GetSelfTokens`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSelfTokensRequest struct via the builder pattern


### Return type

[**[]OAuth1AccessTokenView**](OAuth1AccessTokenView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSelfVhostByAppId

> []VhostView GetSelfVhostByAppId(ctx, appId).Execute()



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SelfApi.GetSelfVhostByAppId(context.Background(), appId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SelfApi.GetSelfVhostByAppId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSelfVhostByAppId`: []VhostView
    fmt.Fprintf(os.Stdout, "Response from `SelfApi.GetSelfVhostByAppId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSelfVhostByAppIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]VhostView**](VhostView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSshKeys

> []SshKeyView GetSshKeys(ctx).Execute()



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
    resp, r, err := api_client.SelfApi.GetSshKeys(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SelfApi.GetSshKeys``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSshKeys`: []SshKeyView
    fmt.Fprintf(os.Stdout, "Response from `SelfApi.GetSshKeys`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSshKeysRequest struct via the builder pattern


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


## GetSummary

> Summary GetSummary(ctx).Full(full).Execute()



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
    full := "full_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SelfApi.GetSummary(context.Background()).Full(full).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SelfApi.GetSummary``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSummary`: Summary
    fmt.Fprintf(os.Stdout, "Response from `SelfApi.GetSummary`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSummaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **full** | **string** |  | 

### Return type

[**Summary**](Summary.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUser

> UserView GetUser(ctx).Execute()



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
    resp, r, err := api_client.SelfApi.GetUser(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SelfApi.GetUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUser`: UserView
    fmt.Fprintf(os.Stdout, "Response from `SelfApi.GetUser`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserRequest struct via the builder pattern


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


## LinkSelfAddonToApplicationByAppId

> LinkSelfAddonToApplicationByAppId(ctx, appId).Execute()



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SelfApi.LinkSelfAddonToApplicationByAppId(context.Background(), appId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SelfApi.LinkSelfAddonToApplicationByAppId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiLinkSelfAddonToApplicationByAppIdRequest struct via the builder pattern


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


## MarkSelfFavouriteVhostByAppId

> VhostView MarkSelfFavouriteVhostByAppId(ctx, appId).VhostView(vhostView).Execute()



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
    vhostView := *openapiclient.NewVhostView() // VhostView | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SelfApi.MarkSelfFavouriteVhostByAppId(context.Background(), appId).VhostView(vhostView).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SelfApi.MarkSelfFavouriteVhostByAppId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MarkSelfFavouriteVhostByAppId`: VhostView
    fmt.Fprintf(os.Stdout, "Response from `SelfApi.MarkSelfFavouriteVhostByAppId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMarkSelfFavouriteVhostByAppIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **vhostView** | [**VhostView**](VhostView.md) |  | 

### Return type

[**VhostView**](VhostView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PreorderSelfAddon

> InvoiceRendering PreorderSelfAddon(ctx).WannabeAddonProvision(wannabeAddonProvision).Execute()



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
    wannabeAddonProvision := *openapiclient.NewWannabeAddonProvision("ProviderId_example", "Plan_example", "Region_example") // WannabeAddonProvision | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SelfApi.PreorderSelfAddon(context.Background()).WannabeAddonProvision(wannabeAddonProvision).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SelfApi.PreorderSelfAddon``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PreorderSelfAddon`: InvoiceRendering
    fmt.Fprintf(os.Stdout, "Response from `SelfApi.PreorderSelfAddon`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPreorderSelfAddonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **wannabeAddonProvision** | [**WannabeAddonProvision**](WannabeAddonProvision.md) |  | 

### Return type

[**InvoiceRendering**](InvoiceRendering.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProvisionSelfAddon

> AddonView ProvisionSelfAddon(ctx).WannabeAddonProvision(wannabeAddonProvision).Execute()



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
    wannabeAddonProvision := *openapiclient.NewWannabeAddonProvision("ProviderId_example", "Plan_example", "Region_example") // WannabeAddonProvision | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SelfApi.ProvisionSelfAddon(context.Background()).WannabeAddonProvision(wannabeAddonProvision).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SelfApi.ProvisionSelfAddon``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProvisionSelfAddon`: AddonView
    fmt.Fprintf(os.Stdout, "Response from `SelfApi.ProvisionSelfAddon`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProvisionSelfAddonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **wannabeAddonProvision** | [**WannabeAddonProvision**](WannabeAddonProvision.md) |  | 

### Return type

[**AddonView**](AddonView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RedeploySelfApplicationByAppId

> Message RedeploySelfApplicationByAppId(ctx, appId).Commit(commit).UseCache(useCache).Execute()



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
    commit := "commit_example" // string |  (optional)
    useCache := "useCache_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SelfApi.RedeploySelfApplicationByAppId(context.Background(), appId).Commit(commit).UseCache(useCache).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SelfApi.RedeploySelfApplicationByAppId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RedeploySelfApplicationByAppId`: Message
    fmt.Fprintf(os.Stdout, "Response from `SelfApi.RedeploySelfApplicationByAppId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRedeploySelfApplicationByAppIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **commit** | **string** |  | 
 **useCache** | **string** |  | 

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


## RemoveEmailAddress

> Message RemoveEmailAddress(ctx, email).Execute()



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
    email := "email_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SelfApi.RemoveEmailAddress(context.Background(), email).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SelfApi.RemoveEmailAddress``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemoveEmailAddress`: Message
    fmt.Fprintf(os.Stdout, "Response from `SelfApi.RemoveEmailAddress`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**email** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveEmailAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## RemoveSelfApplicationEnvByAppIdAndEnvName

> ApplicationView RemoveSelfApplicationEnvByAppIdAndEnvName(ctx, appId, envName).Execute()



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
    envName := "envName_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SelfApi.RemoveSelfApplicationEnvByAppIdAndEnvName(context.Background(), appId, envName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SelfApi.RemoveSelfApplicationEnvByAppIdAndEnvName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemoveSelfApplicationEnvByAppIdAndEnvName`: ApplicationView
    fmt.Fprintf(os.Stdout, "Response from `SelfApi.RemoveSelfApplicationEnvByAppIdAndEnvName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** |  | 
**envName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveSelfApplicationEnvByAppIdAndEnvNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ApplicationView**](ApplicationView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveSelfVhostByAppId

> Message RemoveSelfVhostByAppId(ctx, appId, domain).Execute()



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
    domain := "domain_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SelfApi.RemoveSelfVhostByAppId(context.Background(), appId, domain).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SelfApi.RemoveSelfVhostByAppId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemoveSelfVhostByAppId`: Message
    fmt.Fprintf(os.Stdout, "Response from `SelfApi.RemoveSelfVhostByAppId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** |  | 
**domain** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveSelfVhostByAppIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## RemoveSshKey

> Message RemoveSshKey(ctx, key).Execute()



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SelfApi.RemoveSshKey(context.Background(), key).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SelfApi.RemoveSshKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemoveSshKey`: Message
    fmt.Fprintf(os.Stdout, "Response from `SelfApi.RemoveSshKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveSshKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## RenameAddon

> AddonView RenameAddon(ctx, addonId).WannabeAddonProvision(wannabeAddonProvision).Execute()



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
    addonId := "addonId_example" // string | 
    wannabeAddonProvision := *openapiclient.NewWannabeAddonProvision("ProviderId_example", "Plan_example", "Region_example") // WannabeAddonProvision | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SelfApi.RenameAddon(context.Background(), addonId).WannabeAddonProvision(wannabeAddonProvision).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SelfApi.RenameAddon``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RenameAddon`: AddonView
    fmt.Fprintf(os.Stdout, "Response from `SelfApi.RenameAddon`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**addonId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRenameAddonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **wannabeAddonProvision** | [**WannabeAddonProvision**](WannabeAddonProvision.md) |  | 

### Return type

[**AddonView**](AddonView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RevokeAllTokens

> Message RevokeAllTokens(ctx).Execute()



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
    resp, r, err := api_client.SelfApi.RevokeAllTokens(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SelfApi.RevokeAllTokens``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RevokeAllTokens`: Message
    fmt.Fprintf(os.Stdout, "Response from `SelfApi.RevokeAllTokens`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiRevokeAllTokensRequest struct via the builder pattern


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


## RevokeToken

> Message RevokeToken(ctx, token).Execute()



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
    token := "token_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SelfApi.RevokeToken(context.Background(), token).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SelfApi.RevokeToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RevokeToken`: Message
    fmt.Fprintf(os.Stdout, "Response from `SelfApi.RevokeToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRevokeTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## SetSelfApplicationBranchByAppId

> SetSelfApplicationBranchByAppId(ctx, appId).WannabeBranch(wannabeBranch).Execute()



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
    wannabeBranch := *openapiclient.NewWannabeBranch() // WannabeBranch | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SelfApi.SetSelfApplicationBranchByAppId(context.Background(), appId).WannabeBranch(wannabeBranch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SelfApi.SetSelfApplicationBranchByAppId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetSelfApplicationBranchByAppIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **wannabeBranch** | [**WannabeBranch**](WannabeBranch.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetSelfBuildInstanceFlavorByAppId

> SetSelfBuildInstanceFlavorByAppId(ctx, appId).WannabeBuildFlavor(wannabeBuildFlavor).Execute()



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
    wannabeBuildFlavor := *openapiclient.NewWannabeBuildFlavor() // WannabeBuildFlavor | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SelfApi.SetSelfBuildInstanceFlavorByAppId(context.Background(), appId).WannabeBuildFlavor(wannabeBuildFlavor).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SelfApi.SetSelfBuildInstanceFlavorByAppId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetSelfBuildInstanceFlavorByAppIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **wannabeBuildFlavor** | [**WannabeBuildFlavor**](WannabeBuildFlavor.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetSelfDefaultMethod

> SetSelfDefaultMethod(ctx).PaymentData(paymentData).Execute()



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
    paymentData := *openapiclient.NewPaymentData() // PaymentData | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SelfApi.SetSelfDefaultMethod(context.Background()).PaymentData(paymentData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SelfApi.SetSelfDefaultMethod``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetSelfDefaultMethodRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **paymentData** | [**PaymentData**](PaymentData.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetSelfMaxCreditsPerMonth

> WannabeMaxCredits SetSelfMaxCreditsPerMonth(ctx).WannabeMaxCredits(wannabeMaxCredits).Execute()



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
    wannabeMaxCredits := *openapiclient.NewWannabeMaxCredits() // WannabeMaxCredits | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SelfApi.SetSelfMaxCreditsPerMonth(context.Background()).WannabeMaxCredits(wannabeMaxCredits).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SelfApi.SetSelfMaxCreditsPerMonth``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetSelfMaxCreditsPerMonth`: WannabeMaxCredits
    fmt.Fprintf(os.Stdout, "Response from `SelfApi.SetSelfMaxCreditsPerMonth`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetSelfMaxCreditsPerMonthRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **wannabeMaxCredits** | [**WannabeMaxCredits**](WannabeMaxCredits.md) |  | 

### Return type

[**WannabeMaxCredits**](WannabeMaxCredits.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetUserAvatarFromString

> UrlView SetUserAvatarFromString(ctx).WannabeAvatarSource(wannabeAvatarSource).Execute()



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
    wannabeAvatarSource := *openapiclient.NewWannabeAvatarSource() // WannabeAvatarSource | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SelfApi.SetUserAvatarFromString(context.Background()).WannabeAvatarSource(wannabeAvatarSource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SelfApi.SetUserAvatarFromString``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetUserAvatarFromString`: UrlView
    fmt.Fprintf(os.Stdout, "Response from `SelfApi.SetUserAvatarFromString`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetUserAvatarFromStringRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **wannabeAvatarSource** | [**WannabeAvatarSource**](WannabeAvatarSource.md) |  | 

### Return type

[**UrlView**](UrlView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UndeploySelfApplicationByAppId

> Message UndeploySelfApplicationByAppId(ctx, appId).Execute()



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SelfApi.UndeploySelfApplicationByAppId(context.Background(), appId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SelfApi.UndeploySelfApplicationByAppId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UndeploySelfApplicationByAppId`: Message
    fmt.Fprintf(os.Stdout, "Response from `SelfApi.UndeploySelfApplicationByAppId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUndeploySelfApplicationByAppIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## UnlinkSelfddonFromApplicationByAppAndAddonId

> UnlinkSelfddonFromApplicationByAppAndAddonId(ctx, appId, addonId).Execute()



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
    addonId := "addonId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SelfApi.UnlinkSelfddonFromApplicationByAppAndAddonId(context.Background(), appId, addonId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SelfApi.UnlinkSelfddonFromApplicationByAppAndAddonId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** |  | 
**addonId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnlinkSelfddonFromApplicationByAppAndAddonIdRequest struct via the builder pattern


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


## UnmarkSelfFavouriteVhostByAppId

> UnmarkSelfFavouriteVhostByAppId(ctx, appId).Execute()



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SelfApi.UnmarkSelfFavouriteVhostByAppId(context.Background(), appId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SelfApi.UnmarkSelfFavouriteVhostByAppId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnmarkSelfFavouriteVhostByAppIdRequest struct via the builder pattern


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


## UpdateSelfConsumer

> OAuth1ConsumerView UpdateSelfConsumer(ctx, key).WannabeOAuth1Consumer(wannabeOAuth1Consumer).Execute()



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
    wannabeOAuth1Consumer := *openapiclient.NewWannabeOAuth1Consumer() // WannabeOAuth1Consumer | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SelfApi.UpdateSelfConsumer(context.Background(), key).WannabeOAuth1Consumer(wannabeOAuth1Consumer).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SelfApi.UpdateSelfConsumer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSelfConsumer`: OAuth1ConsumerView
    fmt.Fprintf(os.Stdout, "Response from `SelfApi.UpdateSelfConsumer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSelfConsumerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **wannabeOAuth1Consumer** | [**WannabeOAuth1Consumer**](WannabeOAuth1Consumer.md) |  | 

### Return type

[**OAuth1ConsumerView**](OAuth1ConsumerView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSelfExposedEnvByAppId

> Message UpdateSelfExposedEnvByAppId(ctx, appId).Body(body).Execute()



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SelfApi.UpdateSelfExposedEnvByAppId(context.Background(), appId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SelfApi.UpdateSelfExposedEnvByAppId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSelfExposedEnvByAppId`: Message
    fmt.Fprintf(os.Stdout, "Response from `SelfApi.UpdateSelfExposedEnvByAppId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSelfExposedEnvByAppIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **string** |  | 

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


## ValidateEmail

> ValidateEmail(ctx).ValidationKey(validationKey).Execute()



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
    validationKey := "validationKey_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SelfApi.ValidateEmail(context.Background()).ValidationKey(validationKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SelfApi.ValidateEmail``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidateEmailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **validationKey** | **string** |  | 

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


## ValidateMFA

> ValidateMFA(ctx, kind).WannabeMFACreds(wannabeMFACreds).Execute()



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
    kind := "kind_example" // string | 
    wannabeMFACreds := *openapiclient.NewWannabeMFACreds() // WannabeMFACreds | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SelfApi.ValidateMFA(context.Background(), kind).WannabeMFACreds(wannabeMFACreds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SelfApi.ValidateMFA``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**kind** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiValidateMFARequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **wannabeMFACreds** | [**WannabeMFACreds**](WannabeMFACreds.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

