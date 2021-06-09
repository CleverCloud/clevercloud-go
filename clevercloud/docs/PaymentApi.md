# \PaymentApi

All URIs are relative to *https://api.clever-cloud.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CheckVat**](PaymentApi.md#CheckVat) | **Get** /vat_check | 
[**EndPaymentWithStripe**](PaymentApi.md#EndPaymentWithStripe) | **Post** /payments/{bid}/end/stripe | 
[**GetAvailablePaymentProviders**](PaymentApi.md#GetAvailablePaymentProviders) | **Get** /payments/providers | 
[**GetCoupon**](PaymentApi.md#GetCoupon) | **Get** /payments/coupons/{name} | 
[**GetInvoiceStatusButton**](PaymentApi.md#GetInvoiceStatusButton) | **Get** /payments/assets/pay_button/{token}/button.png | 
[**GetStripeToken**](PaymentApi.md#GetStripeToken) | **Get** /payments/tokens/stripe | 
[**StripeSepaWebhook**](PaymentApi.md#StripeSepaWebhook) | **Post** /payments/webhooks/stripe/sepa | 
[**UpdateStripePayment**](PaymentApi.md#UpdateStripePayment) | **Put** /payments/{bid}/end/stripe | 
[**Validate**](PaymentApi.md#Validate) | **Get** /validation/vat/{key} | 



## CheckVat

> VatResult CheckVat(ctx).Country(country).Vat(vat).Execute()



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
    country := "country_example" // string |  (optional)
    vat := "vat_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PaymentApi.CheckVat(context.Background()).Country(country).Vat(vat).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentApi.CheckVat``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CheckVat`: VatResult
    fmt.Fprintf(os.Stdout, "Response from `PaymentApi.CheckVat`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckVatRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **country** | **string** |  | 
 **vat** | **string** |  | 

### Return type

[**VatResult**](VatResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EndPaymentWithStripe

> InvoiceRendering EndPaymentWithStripe(ctx, bid).PaymentData(paymentData).Execute()



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
    resp, r, err := api_client.PaymentApi.EndPaymentWithStripe(context.Background(), bid).PaymentData(paymentData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentApi.EndPaymentWithStripe``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EndPaymentWithStripe`: InvoiceRendering
    fmt.Fprintf(os.Stdout, "Response from `PaymentApi.EndPaymentWithStripe`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEndPaymentWithStripeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **paymentData** | [**PaymentData**](PaymentData.md) |  | 

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


## GetAvailablePaymentProviders

> []PaymentProviderView GetAvailablePaymentProviders(ctx).Execute()



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
    resp, r, err := api_client.PaymentApi.GetAvailablePaymentProviders(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentApi.GetAvailablePaymentProviders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAvailablePaymentProviders`: []PaymentProviderView
    fmt.Fprintf(os.Stdout, "Response from `PaymentApi.GetAvailablePaymentProviders`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAvailablePaymentProvidersRequest struct via the builder pattern


### Return type

[**[]PaymentProviderView**](PaymentProviderView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCoupon

> CouponView GetCoupon(ctx, name).Execute()



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
    name := "name_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PaymentApi.GetCoupon(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentApi.GetCoupon``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCoupon`: CouponView
    fmt.Fprintf(os.Stdout, "Response from `PaymentApi.GetCoupon`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCouponRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CouponView**](CouponView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInvoiceStatusButton

> GetInvoiceStatusButton(ctx, token).Execute()



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
    resp, r, err := api_client.PaymentApi.GetInvoiceStatusButton(context.Background(), token).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentApi.GetInvoiceStatusButton``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInvoiceStatusButtonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: image/png

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStripeToken

> BraintreeToken GetStripeToken(ctx).Execute()



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
    resp, r, err := api_client.PaymentApi.GetStripeToken(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentApi.GetStripeToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStripeToken`: BraintreeToken
    fmt.Fprintf(os.Stdout, "Response from `PaymentApi.GetStripeToken`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetStripeTokenRequest struct via the builder pattern


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


## StripeSepaWebhook

> StripeSepaWebhook(ctx).StripeSignature(stripeSignature).Body(body).Execute()



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
    stripeSignature := "stripeSignature_example" // string |  (optional)
    body := "body_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PaymentApi.StripeSepaWebhook(context.Background()).StripeSignature(stripeSignature).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentApi.StripeSepaWebhook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStripeSepaWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **stripeSignature** | **string** |  | 
 **body** | **string** |  | 

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


## UpdateStripePayment

> InvoiceRendering UpdateStripePayment(ctx, bid).SetupIntentView(setupIntentView).Execute()



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
    setupIntentView := *openapiclient.NewSetupIntentView() // SetupIntentView | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PaymentApi.UpdateStripePayment(context.Background(), bid).SetupIntentView(setupIntentView).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentApi.UpdateStripePayment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateStripePayment`: InvoiceRendering
    fmt.Fprintf(os.Stdout, "Response from `PaymentApi.UpdateStripePayment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateStripePaymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **setupIntentView** | [**SetupIntentView**](SetupIntentView.md) |  | 

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


## Validate

> Message Validate(ctx, key).Action(action).Execute()



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
    action := "action_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PaymentApi.Validate(context.Background(), key).Action(action).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentApi.Validate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Validate`: Message
    fmt.Fprintf(os.Stdout, "Response from `PaymentApi.Validate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiValidateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **action** | **string** |  | 

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

