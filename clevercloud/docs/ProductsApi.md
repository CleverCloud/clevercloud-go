# \ProductsApi

All URIs are relative to *https://api.clever-cloud.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BillOwner**](ProductsApi.md#BillOwner) | **Post** /vendor/apps/{addonId}/consumptions | 
[**EditApplicationConfiguration**](ProductsApi.md#EditApplicationConfiguration) | **Put** /vendor/apps/{addonId} | 
[**EndAddonMigration**](ProductsApi.md#EndAddonMigration) | **Put** /vendor/apps/{addonId}/migration_callback | 
[**GetAddonProvider**](ProductsApi.md#GetAddonProvider) | **Get** /products/addonproviders/{provider_id} | 
[**GetAddonProviderInfos**](ProductsApi.md#GetAddonProviderInfos) | **Get** /products/addonproviders/{provider_id}/informations | 
[**GetAddonProviderVersions**](ProductsApi.md#GetAddonProviderVersions) | **Get** /products/addonproviders/{provider_id}/versions | 
[**GetAddonProviders**](ProductsApi.md#GetAddonProviders) | **Get** /products/addonproviders | 
[**GetApplicationInfo**](ProductsApi.md#GetApplicationInfo) | **Get** /vendor/apps/{addonId} | 
[**GetAvailableInstances**](ProductsApi.md#GetAvailableInstances) | **Get** /products/instances | 
[**GetAvailablePackages**](ProductsApi.md#GetAvailablePackages) | **Get** /products/packages | 
[**GetCountries**](ProductsApi.md#GetCountries) | **Get** /products/countries | 
[**GetCountryCodes**](ProductsApi.md#GetCountryCodes) | **Get** /products/countrycodes | 
[**GetExcahngeRates**](ProductsApi.md#GetExcahngeRates) | **Get** /products/prices | 
[**GetFlavors**](ProductsApi.md#GetFlavors) | **Get** /products/flavors | 
[**GetInstance**](ProductsApi.md#GetInstance) | **Get** /products/instances/{type}-{version} | 
[**GetMFAKinds**](ProductsApi.md#GetMFAKinds) | **Get** /products/mfa_kinds | 
[**GetZones**](ProductsApi.md#GetZones) | **Get** /products/zones | 
[**ListApps**](ProductsApi.md#ListApps) | **Get** /vendor/apps | 
[**Logscollector**](ProductsApi.md#Logscollector) | **Get** /vendor/apps/{addonId}/logscollector | 
[**ProvisionOtherAddon**](ProductsApi.md#ProvisionOtherAddon) | **Post** /vendor/addons | 



## BillOwner

> BillOwner(ctx, addonId).WannabeAddonBilling(wannabeAddonBilling).Execute()



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
    wannabeAddonBilling := *openapiclient.NewWannabeAddonBilling() // WannabeAddonBilling | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProductsApi.BillOwner(context.Background(), addonId).WannabeAddonBilling(wannabeAddonBilling).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductsApi.BillOwner``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**addonId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBillOwnerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **wannabeAddonBilling** | [**WannabeAddonBilling**](WannabeAddonBilling.md) |  | 

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


## EditApplicationConfiguration

> AddonView EditApplicationConfiguration(ctx, addonId).WannabeAddonConfig(wannabeAddonConfig).Execute()



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
    wannabeAddonConfig := *openapiclient.NewWannabeAddonConfig() // WannabeAddonConfig | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProductsApi.EditApplicationConfiguration(context.Background(), addonId).WannabeAddonConfig(wannabeAddonConfig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductsApi.EditApplicationConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditApplicationConfiguration`: AddonView
    fmt.Fprintf(os.Stdout, "Response from `ProductsApi.EditApplicationConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**addonId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditApplicationConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **wannabeAddonConfig** | [**WannabeAddonConfig**](WannabeAddonConfig.md) |  | 

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


## EndAddonMigration

> AddonView EndAddonMigration(ctx, addonId).WannabeAddonConfig(wannabeAddonConfig).PlanId(planId).Region(region).Execute()



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
    wannabeAddonConfig := *openapiclient.NewWannabeAddonConfig() // WannabeAddonConfig | 
    planId := "planId_example" // string |  (optional)
    region := "region_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProductsApi.EndAddonMigration(context.Background(), addonId).WannabeAddonConfig(wannabeAddonConfig).PlanId(planId).Region(region).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductsApi.EndAddonMigration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EndAddonMigration`: AddonView
    fmt.Fprintf(os.Stdout, "Response from `ProductsApi.EndAddonMigration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**addonId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEndAddonMigrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **wannabeAddonConfig** | [**WannabeAddonConfig**](WannabeAddonConfig.md) |  | 
 **planId** | **string** |  | 
 **region** | **string** |  | 

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


## GetAddonProvider

> AddonProviderInfoFullView GetAddonProvider(ctx, providerId).OrgaId(orgaId).Execute()



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
    providerId := "providerId_example" // string | 
    orgaId := "orgaId_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProductsApi.GetAddonProvider(context.Background(), providerId).OrgaId(orgaId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductsApi.GetAddonProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAddonProvider`: AddonProviderInfoFullView
    fmt.Fprintf(os.Stdout, "Response from `ProductsApi.GetAddonProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**providerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAddonProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **orgaId** | **string** |  | 

### Return type

[**AddonProviderInfoFullView**](AddonProviderInfoFullView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAddonProviderInfos

> string GetAddonProviderInfos(ctx, providerId).Execute()



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
    providerId := "providerId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProductsApi.GetAddonProviderInfos(context.Background(), providerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductsApi.GetAddonProviderInfos``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAddonProviderInfos`: string
    fmt.Fprintf(os.Stdout, "Response from `ProductsApi.GetAddonProviderInfos`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**providerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAddonProviderInfosRequest struct via the builder pattern


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


## GetAddonProviderVersions

> string GetAddonProviderVersions(ctx, providerId).Execute()



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
    providerId := "providerId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProductsApi.GetAddonProviderVersions(context.Background(), providerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductsApi.GetAddonProviderVersions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAddonProviderVersions`: string
    fmt.Fprintf(os.Stdout, "Response from `ProductsApi.GetAddonProviderVersions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**providerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAddonProviderVersionsRequest struct via the builder pattern


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


## GetAddonProviders

> []AddonProviderInfoFullView GetAddonProviders(ctx).OrgaId(orgaId).Execute()



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
    orgaId := "orgaId_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProductsApi.GetAddonProviders(context.Background()).OrgaId(orgaId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductsApi.GetAddonProviders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAddonProviders`: []AddonProviderInfoFullView
    fmt.Fprintf(os.Stdout, "Response from `ProductsApi.GetAddonProviders`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAddonProvidersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orgaId** | **string** |  | 

### Return type

[**[]AddonProviderInfoFullView**](AddonProviderInfoFullView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplicationInfo

> AddonApplicationInfo GetApplicationInfo(ctx, addonId).Execute()



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
    resp, r, err := api_client.ProductsApi.GetApplicationInfo(context.Background(), addonId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductsApi.GetApplicationInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApplicationInfo`: AddonApplicationInfo
    fmt.Fprintf(os.Stdout, "Response from `ProductsApi.GetApplicationInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**addonId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AddonApplicationInfo**](AddonApplicationInfo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAvailableInstances

> []AvailableInstanceView GetAvailableInstances(ctx).For_(for_).Execute()



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
    for_ := "for__example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProductsApi.GetAvailableInstances(context.Background()).For_(for_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductsApi.GetAvailableInstances``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAvailableInstances`: []AvailableInstanceView
    fmt.Fprintf(os.Stdout, "Response from `ProductsApi.GetAvailableInstances`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAvailableInstancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **for_** | **string** |  | 

### Return type

[**[]AvailableInstanceView**](AvailableInstanceView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAvailablePackages

> []PackageView GetAvailablePackages(ctx).Coupon(coupon).OrgaId(orgaId).Currency(currency).Execute()



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
    coupon := "coupon_example" // string |  (optional)
    orgaId := "orgaId_example" // string |  (optional)
    currency := "currency_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProductsApi.GetAvailablePackages(context.Background()).Coupon(coupon).OrgaId(orgaId).Currency(currency).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductsApi.GetAvailablePackages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAvailablePackages`: []PackageView
    fmt.Fprintf(os.Stdout, "Response from `ProductsApi.GetAvailablePackages`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAvailablePackagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **coupon** | **string** |  | 
 **orgaId** | **string** |  | 
 **currency** | **string** |  | 

### Return type

[**[]PackageView**](PackageView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCountries

> string GetCountries(ctx).Execute()



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
    resp, r, err := api_client.ProductsApi.GetCountries(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductsApi.GetCountries``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCountries`: string
    fmt.Fprintf(os.Stdout, "Response from `ProductsApi.GetCountries`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCountriesRequest struct via the builder pattern


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


## GetCountryCodes

> string GetCountryCodes(ctx).Execute()



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
    resp, r, err := api_client.ProductsApi.GetCountryCodes(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductsApi.GetCountryCodes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCountryCodes`: string
    fmt.Fprintf(os.Stdout, "Response from `ProductsApi.GetCountryCodes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCountryCodesRequest struct via the builder pattern


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


## GetExcahngeRates

> []DropPriceView GetExcahngeRates(ctx).Execute()



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
    resp, r, err := api_client.ProductsApi.GetExcahngeRates(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductsApi.GetExcahngeRates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetExcahngeRates`: []DropPriceView
    fmt.Fprintf(os.Stdout, "Response from `ProductsApi.GetExcahngeRates`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetExcahngeRatesRequest struct via the builder pattern


### Return type

[**[]DropPriceView**](DropPriceView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFlavors

> []FlavorView GetFlavors(ctx).Context(context).Execute()



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
    context := "context_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProductsApi.GetFlavors(context.Background()).Context(context).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductsApi.GetFlavors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFlavors`: []FlavorView
    fmt.Fprintf(os.Stdout, "Response from `ProductsApi.GetFlavors`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFlavorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **context** | **string** |  | 

### Return type

[**[]FlavorView**](FlavorView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInstance

> AvailableInstanceView GetInstance(ctx, type_, version).For_(for_).App(app).Execute()



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
    type_ := "type__example" // string | 
    version := "version_example" // string | 
    for_ := "for__example" // string |  (optional)
    app := "app_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProductsApi.GetInstance(context.Background(), type_, version).For_(for_).App(app).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductsApi.GetInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetInstance`: AvailableInstanceView
    fmt.Fprintf(os.Stdout, "Response from `ProductsApi.GetInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**type_** | **string** |  | 
**version** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **for_** | **string** |  | 
 **app** | **string** |  | 

### Return type

[**AvailableInstanceView**](AvailableInstanceView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMFAKinds

> []string GetMFAKinds(ctx).Execute()



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
    resp, r, err := api_client.ProductsApi.GetMFAKinds(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductsApi.GetMFAKinds``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMFAKinds`: []string
    fmt.Fprintf(os.Stdout, "Response from `ProductsApi.GetMFAKinds`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMFAKindsRequest struct via the builder pattern


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


## GetZones

> []ZoneView GetZones(ctx).Execute()



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
    resp, r, err := api_client.ProductsApi.GetZones(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductsApi.GetZones``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetZones`: []ZoneView
    fmt.Fprintf(os.Stdout, "Response from `ProductsApi.GetZones`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetZonesRequest struct via the builder pattern


### Return type

[**[]ZoneView**](ZoneView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListApps

> []AddonApplicationSummary ListApps(ctx).Offset(offset).Execute()



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
    offset := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProductsApi.ListApps(context.Background()).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductsApi.ListApps``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListApps`: []AddonApplicationSummary
    fmt.Fprintf(os.Stdout, "Response from `ProductsApi.ListApps`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAppsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **int32** |  | 

### Return type

[**[]AddonApplicationSummary**](AddonApplicationSummary.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Logscollector

> Logscollector(ctx, addonId).Execute()



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
    resp, r, err := api_client.ProductsApi.Logscollector(context.Background(), addonId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductsApi.Logscollector``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**addonId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiLogscollectorRequest struct via the builder pattern


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


## ProvisionOtherAddon

> ProvisionOtherAddon(ctx).WannabeInterAddonProvision(wannabeInterAddonProvision).Execute()



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
    wannabeInterAddonProvision := *openapiclient.NewWannabeInterAddonProvision() // WannabeInterAddonProvision | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProductsApi.ProvisionOtherAddon(context.Background()).WannabeInterAddonProvision(wannabeInterAddonProvision).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductsApi.ProvisionOtherAddon``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProvisionOtherAddonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **wannabeInterAddonProvision** | [**WannabeInterAddonProvision**](WannabeInterAddonProvision.md) |  | 

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

