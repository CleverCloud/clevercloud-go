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

> Message AddEmailAddress(ctx, email, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**email** | **string**|  | 
 **optional** | ***AddEmailAddressOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AddEmailAddressOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **optional.String**|  | 

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

> []string AddSelfAddonTagByAddonId(ctx, addonId, tag)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**addonId** | **string**|  | 
**tag** | **string**|  | 

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

> ApplicationView AddSelfApplication(ctx, wannabeApplication)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**wannabeApplication** | [**WannabeApplication**](WannabeApplication.md)|  | 

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

> Message AddSelfApplicationDependencyByAppId(ctx, appId, dependencyId)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string**|  | 
**dependencyId** | **string**|  | 

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

> []string AddSelfApplicationTagByAppId(ctx, appId, tag)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string**|  | 
**tag** | **string**|  | 

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

> PaymentMethodView AddSelfPaymentMethod(ctx, paymentData)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**paymentData** | [**PaymentData**](PaymentData.md)|  | 

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

> Message AddSelfVhostByAppId(ctx, appId, domain)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string**|  | 
**domain** | **string**|  | 

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

> Message AddSshKey(ctx, key)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string**|  | 

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

> InvoiceRendering BuySelfDrops(ctx, wannaBuyPackage)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**wannaBuyPackage** | [**WannaBuyPackage**](WannaBuyPackage.md)|  | 

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

> Message CancelDeploy(ctx, appId, deploymentId)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string**|  | 
**deploymentId** | **string**|  | 

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

> AddonView ChangeSelfAddonPlanByAddonId(ctx, addonId, wannabePlanChange)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**addonId** | **string**|  | 
**wannabePlanChange** | [**WannabePlanChange**](WannabePlanChange.md)|  | 

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

> Message ChangeUserPassword(ctx, wannabePassword)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**wannabePassword** | [**WannabePassword**](WannabePassword.md)|  | 

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

> NextInPaymentFlow ChooseSelfPaymentProvider(ctx, bid, paymentProviderSelection)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bid** | **string**|  | 
**paymentProviderSelection** | [**PaymentProviderSelection**](PaymentProviderSelection.md)|  | 

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

> CreateMFA(ctx, kind)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**kind** | **string**|  | 

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

> OAuth1ConsumerView CreateSelfConsumer(ctx, wannabeOAuth1Consumer)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**wannabeOAuth1Consumer** | [**WannabeOAuth1Consumer**](WannabeOAuth1Consumer.md)|  | 

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

> DeleteMFA(ctx, kind)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**kind** | **string**|  | 

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

> []string DeleteSelfAddonTagByAddonId(ctx, addonId, tag)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**addonId** | **string**|  | 
**tag** | **string**|  | 

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

> Message DeleteSelfApplicationByAppId(ctx, appId)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string**|  | 

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

> DeleteSelfApplicationDependencyByAppId(ctx, appId, dependencyId)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string**|  | 
**dependencyId** | **string**|  | 

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

> []string DeleteSelfApplicationTagAppId(ctx, appId, tag)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string**|  | 
**tag** | **string**|  | 

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

> DeleteSelfCard(ctx, mId)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mId** | **string**|  | 

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

> DeleteSelfConsumer(ctx, key)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string**|  | 

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

> DeleteSelfPurchaseOrder(ctx, bid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bid** | **string**|  | 

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

> DeleteSelfRecurrentPayment(ctx, )



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


## DeleteUser

> Message DeleteUser(ctx, )



### Required Parameters

This endpoint does not need any parameter.

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

> Message DeprovisionSelfAddonById(ctx, addonId)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**addonId** | **string**|  | 

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

> ApplicationView EditSelfApplicationByAppId(ctx, appId, wannabeApplication)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string**|  | 
**wannabeApplication** | [**WannabeApplication**](WannabeApplication.md)|  | 

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

> ApplicationView EditSelfApplicationEnvByAppId(ctx, appId, body)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string**|  | 
**body** | **string**|  | 

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

> ApplicationView EditSelfApplicationEnvByAppIdAndEnvName(ctx, appId, envName, wannabeValue)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string**|  | 
**envName** | **string**|  | 
**wannabeValue** | [**WannabeValue**](WannabeValue.md)|  | 

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

> UserView EditUser(ctx, wannabeUser)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**wannabeUser** | [**WannabeUser**](WannabeUser.md)|  | 

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

> FavMFA(ctx, kind, wannabeMfaFav)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**kind** | **string**|  | 
**wannabeMfaFav** | [**WannabeMfaFav**](WannabeMfaFav.md)|  | 

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

> AddonSsoData GetAddonSSOData(ctx, addonId)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**addonId** | **string**|  | 

### Return type

[**AddonSsoData**](AddonSSOData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplicationDeployment

> GetApplicationDeployment(ctx, appId, deploymentId)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string**|  | 
**deploymentId** | **string**|  | 

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

> []DeploymentView GetApplicationDeployments(ctx, appId, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string**|  | 
 **optional** | ***GetApplicationDeploymentsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetApplicationDeploymentsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.String**|  | 
 **offset** | **optional.String**|  | 
 **action** | **optional.String**|  | 

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

> []MfaRecoveryCode GetBackupCodes(ctx, kind)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**kind** | **string**|  | 

### Return type

[**[]MfaRecoveryCode**](MFARecoveryCode.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConfirmationEmail

> Message GetConfirmationEmail(ctx, )



### Required Parameters

This endpoint does not need any parameter.

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

> string GetConsumptions(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetConsumptionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetConsumptionsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **appId** | **optional.String**|  | 
 **from** | **optional.String**|  | 
 **to** | **optional.String**|  | 
 **for_** | **optional.String**|  | 

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

> []string GetEmailAddresses(ctx, )



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


## GetId

> string GetId(ctx, )



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


## GetSelfAddonById

> AddonView GetSelfAddonById(ctx, addonId)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**addonId** | **string**|  | 

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

> []AddonEnvironmentView GetSelfAddonEnvByAddonId(ctx, addonId)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**addonId** | **string**|  | 

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

> []string GetSelfAddonTagsByAddonId(ctx, addonId)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**addonId** | **string**|  | 

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

> []AddonView GetSelfAddons(ctx, )



### Required Parameters

This endpoint does not need any parameter.

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

> []AddonView GetSelfAddonsLinkedToApplicationByAppId(ctx, appId)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string**|  | 

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

> DropCountView GetSelfAmount(ctx, )



### Required Parameters

This endpoint does not need any parameter.

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

> []string GetSelfApplicationBranchesByAppId(ctx, appId)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string**|  | 

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

> ApplicationView GetSelfApplicationByAppId(ctx, appId)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string**|  | 

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

> []ApplicationView GetSelfApplicationDependenciesByAppId(ctx, appId)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string**|  | 

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

> GetSelfApplicationDependenciesEnvByAppId(ctx, id, appId)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
**appId** | **string**|  | 

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

> []ApplicationView GetSelfApplicationDependents(ctx, appId)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string**|  | 

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

> []AddonEnvironmentView GetSelfApplicationEnvByAppId(ctx, appId)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string**|  | 

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

> SuperNovaInstanceView GetSelfApplicationInstanceByAppAndInstanceId(ctx, appId, instanceId)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string**|  | 
**instanceId** | **string**|  | 

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

> []SuperNovaInstanceView GetSelfApplicationInstancesByAppId(ctx, appId, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string**|  | 
 **optional** | ***GetSelfApplicationInstancesByAppIdOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetSelfApplicationInstancesByAppIdOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deploymentId** | **optional.String**|  | 
 **withDeleted** | **optional.String**|  | 

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

> []string GetSelfApplicationTagsByAppId(ctx, appId)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string**|  | 

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

> []ApplicationView GetSelfApplications(ctx, )



### Required Parameters

This endpoint does not need any parameter.

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

> []ApplicationView GetSelfApplicationsLinkedToAddonByAddonId(ctx, addonId)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**addonId** | **string**|  | 

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

> []CliTokenView GetSelfCliTokens(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetSelfCliTokensOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetSelfCliTokensOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cliToken** | **optional.String**|  | 

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

> OAuth1ConsumerView GetSelfConsumer(ctx, key)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string**|  | 

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

> SecretView GetSelfConsumerSecret(ctx, key)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string**|  | 

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

> GetSelfConsumers(ctx, )



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


## GetSelfDefaultMethod

> PaymentMethodView GetSelfDefaultMethod(ctx, )



### Required Parameters

This endpoint does not need any parameter.

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

> []LinkedAddonEnvironmentView GetSelfEnvOfAddonsLinkedToApplicationByAppId(ctx, appId)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string**|  | 

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

> string GetSelfExposedEnvByAppId(ctx, appId)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string**|  | 

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

> VhostView GetSelfFavouriteVhostByAppId(ctx, appId)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string**|  | 

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

> SuperNovaInstanceMap GetSelfInstancesForAllApps(ctx, )



### Required Parameters

This endpoint does not need any parameter.

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

> InvoiceRendering GetSelfInvoiceById(ctx, bid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bid** | **string**|  | 

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

> []InvoiceRendering GetSelfInvoices(ctx, )



### Required Parameters

This endpoint does not need any parameter.

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

> string GetSelfMonthlyInvoice(ctx, )



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


## GetSelfPaymentInfo

> PaymentInfoView GetSelfPaymentInfo(ctx, )



### Required Parameters

This endpoint does not need any parameter.

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

> []PaymentMethodView GetSelfPaymentMethods(ctx, )



### Required Parameters

This endpoint does not need any parameter.

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

> GetSelfPdfInvoiceById(ctx, bid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bid** | **string**|  | 
 **optional** | ***GetSelfPdfInvoiceByIdOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetSelfPdfInvoiceByIdOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **token** | **optional.String**|  | 

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

> PriceWithTaxInfo GetSelfPriceWithTax(ctx, price)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**price** | **string**|  | 

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

> RecurrentPaymentView GetSelfRecurrentPayment(ctx, )



### Required Parameters

This endpoint does not need any parameter.

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

> BraintreeToken GetSelfStripeToken(ctx, )



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


## GetSelfTokens

> []OAuth1AccessTokenView GetSelfTokens(ctx, )



### Required Parameters

This endpoint does not need any parameter.

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

> []VhostView GetSelfVhostByAppId(ctx, appId)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string**|  | 

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

> []SshKeyView GetSshKeys(ctx, )



### Required Parameters

This endpoint does not need any parameter.

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

> Summary GetSummary(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetSummaryOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetSummaryOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **full** | **optional.String**|  | 

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

> UserView GetUser(ctx, )



### Required Parameters

This endpoint does not need any parameter.

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

> LinkSelfAddonToApplicationByAppId(ctx, appId)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string**|  | 

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

> VhostView MarkSelfFavouriteVhostByAppId(ctx, appId, vhostView)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string**|  | 
**vhostView** | [**VhostView**](VhostView.md)|  | 

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

> InvoiceRendering PreorderSelfAddon(ctx, wannabeAddonProvision)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**wannabeAddonProvision** | [**WannabeAddonProvision**](WannabeAddonProvision.md)|  | 

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

> AddonView ProvisionSelfAddon(ctx, wannabeAddonProvision)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**wannabeAddonProvision** | [**WannabeAddonProvision**](WannabeAddonProvision.md)|  | 

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

> Message RedeploySelfApplicationByAppId(ctx, appId, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string**|  | 
 **optional** | ***RedeploySelfApplicationByAppIdOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a RedeploySelfApplicationByAppIdOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **commit** | **optional.String**|  | 
 **useCache** | **optional.String**|  | 

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

> Message RemoveEmailAddress(ctx, email)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**email** | **string**|  | 

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

> ApplicationView RemoveSelfApplicationEnvByAppIdAndEnvName(ctx, appId, envName)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string**|  | 
**envName** | **string**|  | 

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

> Message RemoveSelfVhostByAppId(ctx, appId, domain)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string**|  | 
**domain** | **string**|  | 

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

> Message RemoveSshKey(ctx, key)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string**|  | 

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

> AddonView RenameAddon(ctx, addonId, wannabeAddonProvision)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**addonId** | **string**|  | 
**wannabeAddonProvision** | [**WannabeAddonProvision**](WannabeAddonProvision.md)|  | 

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

> Message RevokeAllTokens(ctx, )



### Required Parameters

This endpoint does not need any parameter.

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

> Message RevokeToken(ctx, token)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string**|  | 

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

> SetSelfApplicationBranchByAppId(ctx, appId, wannabeBranch)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string**|  | 
**wannabeBranch** | [**WannabeBranch**](WannabeBranch.md)|  | 

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

> SetSelfBuildInstanceFlavorByAppId(ctx, appId, wannabeBuildFlavor)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string**|  | 
**wannabeBuildFlavor** | [**WannabeBuildFlavor**](WannabeBuildFlavor.md)|  | 

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

> SetSelfDefaultMethod(ctx, paymentData)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**paymentData** | [**PaymentData**](PaymentData.md)|  | 

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

> WannabeMaxCredits SetSelfMaxCreditsPerMonth(ctx, wannabeMaxCredits)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**wannabeMaxCredits** | [**WannabeMaxCredits**](WannabeMaxCredits.md)|  | 

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

> UrlView SetUserAvatarFromString(ctx, wannabeAvatarSource)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**wannabeAvatarSource** | [**WannabeAvatarSource**](WannabeAvatarSource.md)|  | 

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

> Message UndeploySelfApplicationByAppId(ctx, appId)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string**|  | 

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

> UnlinkSelfddonFromApplicationByAppAndAddonId(ctx, appId, addonId)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string**|  | 
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


## UnmarkSelfFavouriteVhostByAppId

> UnmarkSelfFavouriteVhostByAppId(ctx, appId)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string**|  | 

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

> OAuth1ConsumerView UpdateSelfConsumer(ctx, key, wannabeOAuth1Consumer)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string**|  | 
**wannabeOAuth1Consumer** | [**WannabeOAuth1Consumer**](WannabeOAuth1Consumer.md)|  | 

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

> Message UpdateSelfExposedEnvByAppId(ctx, appId, body)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string**|  | 
**body** | **string**|  | 

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

> ValidateEmail(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ValidateEmailOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ValidateEmailOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **validationKey** | **optional.String**|  | 

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

> ValidateMFA(ctx, kind, wannabeMfaCreds)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**kind** | **string**|  | 
**wannabeMfaCreds** | [**WannabeMfaCreds**](WannabeMfaCreds.md)|  | 

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

