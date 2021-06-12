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

> string AuthorizeForm(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AuthorizeFormOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AuthorizeFormOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ccid** | **optional.String**|  | 
 **cctk** | **optional.String**|  | 
 **oauthToken** | **optional.String**|  | 
 **ccid2** | **optional.String**|  | 
 **cliToken** | **optional.String**|  | 
 **fromOauth** | **optional.String**|  | 

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

> AuthorizeToken(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AuthorizeTokenOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AuthorizeTokenOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ccid** | **optional.String**|  | 
 **cctk** | **optional.String**|  | 
 **almighty** | **optional.String**|  | 
 **accessOrganisations** | **optional.String**|  | 
 **manageOrganisations** | **optional.String**|  | 
 **manageOrganisationsServices** | **optional.String**|  | 
 **manageOrganisationsApplications** | **optional.String**|  | 
 **manageOrganisationsMembers** | **optional.String**|  | 
 **accessOrganisationsBills** | **optional.String**|  | 
 **accessOrganisationsCreditCount** | **optional.String**|  | 
 **accessOrganisationsConsumptionStatistics** | **optional.String**|  | 
 **accessPersonalInformation** | **optional.String**|  | 
 **managePersonalInformation** | **optional.String**|  | 
 **manageSshKeys** | **optional.String**|  | 

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

> GetAvailableRights(ctx, )



### Required Parameters

This endpoint does not need any parameter.

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

> GetLoginData(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetLoginDataOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetLoginDataOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oauthKey** | **optional.String**|  | 

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

> PostAccessTokenRequest(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PostAccessTokenRequestOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PostAccessTokenRequestOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oauthConsumerKey** | **optional.String**|  | 
 **oauthToken** | **optional.String**|  | 
 **oauthSignatureMethod** | **optional.String**|  | 
 **oauthSignature** | **optional.String**|  | 
 **oauthTimestamp** | **optional.String**|  | 
 **oauthNonce** | **optional.String**|  | 
 **oauthVersion** | **optional.String**|  | 
 **oauthVerifier** | **optional.String**|  | 
 **oauthCallback** | **optional.String**|  | 
 **oauthTokenSecret** | **optional.String**|  | 
 **oauthCallbackConfirmed** | **optional.String**|  | 

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

> PostAccessTokenRequestQuery(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PostAccessTokenRequestQueryOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PostAccessTokenRequestQueryOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oauthConsumerKey** | **optional.String**|  | 
 **oauthToken** | **optional.String**|  | 
 **oauthSignatureMethod** | **optional.String**|  | 
 **oauthSignature** | **optional.String**|  | 
 **oauthTimestamp** | **optional.String**|  | 
 **oauthNonce** | **optional.String**|  | 
 **oauthVersion** | **optional.String**|  | 
 **oauthVerifier** | **optional.String**|  | 
 **oauthCallback** | **optional.String**|  | 
 **oauthTokenSecret** | **optional.String**|  | 
 **oauthCallbackConfirmed** | **optional.String**|  | 

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

> Message PostAuthorize(ctx, wannabeAuthorization)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**wannabeAuthorization** | [**WannabeAuthorization**](WannabeAuthorization.md)|  | 

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

> string PostReqTokenRequest(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PostReqTokenRequestOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PostReqTokenRequestOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cleverFlavor** | **optional.String**|  | 
 **oauthConsumerKey** | **optional.String**|  | 
 **oauthToken** | **optional.String**|  | 
 **oauthSignatureMethod** | **optional.String**|  | 
 **oauthSignature** | **optional.String**|  | 
 **oauthTimestamp** | **optional.String**|  | 
 **oauthNonce** | **optional.String**|  | 
 **oauthVersion** | **optional.String**|  | 
 **oauthVerifier** | **optional.String**|  | 
 **oauthCallback** | **optional.String**|  | 
 **oauthTokenSecret** | **optional.String**|  | 
 **oauthCallbackConfirmed** | **optional.String**|  | 

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

> string PostReqTokenRequestQueryString(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PostReqTokenRequestQueryStringOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PostReqTokenRequestQueryStringOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cleverFlavor** | **optional.String**|  | 
 **oauthConsumerKey** | **optional.String**|  | 
 **oauthToken** | **optional.String**|  | 
 **oauthSignatureMethod** | **optional.String**|  | 
 **oauthSignature** | **optional.String**|  | 
 **oauthTimestamp** | **optional.String**|  | 
 **oauthNonce** | **optional.String**|  | 
 **oauthVersion** | **optional.String**|  | 
 **oauthVerifier** | **optional.String**|  | 
 **oauthCallback** | **optional.String**|  | 
 **oauthTokenSecret** | **optional.String**|  | 
 **oauthCallbackConfirmed** | **optional.String**|  | 

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

