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

> BillOwner(ctx, addonId, wannabeAddonBilling)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**addonId** | **string**|  | 
**wannabeAddonBilling** | [**WannabeAddonBilling**](WannabeAddonBilling.md)|  | 

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

> AddonView EditApplicationConfiguration(ctx, addonId, wannabeAddonConfig)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**addonId** | **string**|  | 
**wannabeAddonConfig** | [**WannabeAddonConfig**](WannabeAddonConfig.md)|  | 

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

> AddonView EndAddonMigration(ctx, addonId, wannabeAddonConfig, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**addonId** | **string**|  | 
**wannabeAddonConfig** | [**WannabeAddonConfig**](WannabeAddonConfig.md)|  | 
 **optional** | ***EndAddonMigrationOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a EndAddonMigrationOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **planId** | **optional.String**|  | 
 **region** | **optional.String**|  | 

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

> AddonProviderInfoFullView GetAddonProvider(ctx, providerId, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**providerId** | **string**|  | 
 **optional** | ***GetAddonProviderOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetAddonProviderOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **orgaId** | **optional.String**|  | 

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

> string GetAddonProviderInfos(ctx, providerId)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**providerId** | **string**|  | 

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

> string GetAddonProviderVersions(ctx, providerId)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**providerId** | **string**|  | 

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

> []AddonProviderInfoFullView GetAddonProviders(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetAddonProvidersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetAddonProvidersOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orgaId** | **optional.String**|  | 

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

> AddonApplicationInfo GetApplicationInfo(ctx, addonId)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**addonId** | **string**|  | 

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

> []AvailableInstanceView GetAvailableInstances(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetAvailableInstancesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetAvailableInstancesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **for_** | **optional.String**|  | 

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

> []PackageView GetAvailablePackages(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetAvailablePackagesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetAvailablePackagesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **coupon** | **optional.String**|  | 
 **orgaId** | **optional.String**|  | 
 **currency** | **optional.String**|  | 

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

> string GetCountries(ctx, )



### Required Parameters

This endpoint does not need any parameter.

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

> string GetCountryCodes(ctx, )



### Required Parameters

This endpoint does not need any parameter.

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

> []DropPriceView GetExcahngeRates(ctx, )



### Required Parameters

This endpoint does not need any parameter.

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

> []FlavorView GetFlavors(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetFlavorsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetFlavorsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **context** | **optional.String**|  | 

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

> AvailableInstanceView GetInstance(ctx, type_, version, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**type_** | **string**|  | 
**version** | **string**|  | 
 **optional** | ***GetInstanceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetInstanceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **for_** | **optional.String**|  | 
 **app** | **optional.String**|  | 

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

> []string GetMFAKinds(ctx, )



### Required Parameters

This endpoint does not need any parameter.

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

> []ZoneView GetZones(ctx, )



### Required Parameters

This endpoint does not need any parameter.

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

> []AddonApplicationSummary ListApps(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListAppsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListAppsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **optional.Int32**|  | 

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

> Logscollector(ctx, addonId)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**addonId** | **string**|  | 

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

> ProvisionOtherAddon(ctx, wannabeInterAddonProvision)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**wannabeInterAddonProvision** | [**WannabeInterAddonProvision**](WannabeInterAddonProvision.md)|  | 

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

