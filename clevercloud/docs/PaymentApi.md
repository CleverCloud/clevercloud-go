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

> VatResult CheckVat(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CheckVatOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CheckVatOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **country** | **optional.String**|  | 
 **vat** | **optional.String**|  | 

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

> InvoiceRendering EndPaymentWithStripe(ctx, bid, paymentData)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bid** | **string**|  | 
**paymentData** | [**PaymentData**](PaymentData.md)|  | 

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

> []PaymentProviderView GetAvailablePaymentProviders(ctx, )



### Required Parameters

This endpoint does not need any parameter.

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

> CouponView GetCoupon(ctx, name)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string**|  | 

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

> GetInvoiceStatusButton(ctx, token)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string**|  | 

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

> BraintreeToken GetStripeToken(ctx, )



### Required Parameters

This endpoint does not need any parameter.

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

> StripeSepaWebhook(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***StripeSepaWebhookOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a StripeSepaWebhookOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **stripeSignature** | **optional.String**|  | 
 **body** | **optional.String**|  | 

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

> InvoiceRendering UpdateStripePayment(ctx, bid, setupIntentView)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bid** | **string**|  | 
**setupIntentView** | [**SetupIntentView**](SetupIntentView.md)|  | 

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

> Message Validate(ctx, key, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string**|  | 
 **optional** | ***ValidateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ValidateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **action** | **optional.String**|  | 

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

