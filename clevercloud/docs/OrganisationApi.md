# \OrganisationApi

All URIs are relative to *https://api.clever-cloud.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AbortAddonMigration**](OrganisationApi.md#AbortAddonMigration) | **Delete** /organisations/{id}/addons/{addonId}/migrations/{migrationId} | 
[**AddAddonTagByOrgaAndAddonId**](OrganisationApi.md#AddAddonTagByOrgaAndAddonId) | **Put** /organisations/{id}/addons/{addonId}/tags/{tag} | 
[**AddApplicationByOrga**](OrganisationApi.md#AddApplicationByOrga) | **Post** /organisations/{id}/applications | 
[**AddApplicationDependencyByOrgaAndAppId**](OrganisationApi.md#AddApplicationDependencyByOrgaAndAppId) | **Put** /organisations/{id}/applications/{appId}/dependencies/{dependencyId} | 
[**AddApplicationTagByOrgaAndAppId**](OrganisationApi.md#AddApplicationTagByOrgaAndAppId) | **Put** /organisations/{id}/applications/{appId}/tags/{tag} | 
[**AddBetaTester**](OrganisationApi.md#AddBetaTester) | **Post** /organisations/{id}/addonproviders/{providerId}/testers | 
[**AddOrganisationMember**](OrganisationApi.md#AddOrganisationMember) | **Post** /organisations/{id}/members | 
[**AddPaymentMethodByOrga**](OrganisationApi.md#AddPaymentMethodByOrga) | **Post** /organisations/{id}/payments/methods | 
[**AddProviderFeature**](OrganisationApi.md#AddProviderFeature) | **Post** /organisations/{id}/addonproviders/{providerId}/features | 
[**AddProviderPlan**](OrganisationApi.md#AddProviderPlan) | **Post** /organisations/{id}/addonproviders/{providerId}/plans | 
[**AddTcpRedir**](OrganisationApi.md#AddTcpRedir) | **Post** /organisations/{id}/applications/{appId}/tcpRedirs | 
[**AddVhostsByOrgaAndAppId**](OrganisationApi.md#AddVhostsByOrgaAndAppId) | **Put** /organisations/{id}/applications/{appId}/vhosts/{domain} | 
[**BuyDropsByOrga**](OrganisationApi.md#BuyDropsByOrga) | **Post** /organisations/{id}/payments/billings | 
[**CancelApplicationDeploymentForOrga**](OrganisationApi.md#CancelApplicationDeploymentForOrga) | **Delete** /organisations/{id}/applications/{appId}/deployments/{deploymentId}/instances | 
[**ChangePlanByOrgaAndAddonId**](OrganisationApi.md#ChangePlanByOrgaAndAddonId) | **Post** /organisations/{id}/addons/{addonId}/migrations | 
[**ChoosePaymentProviderByOrga**](OrganisationApi.md#ChoosePaymentProviderByOrga) | **Put** /organisations/{id}/payments/billings/{bid} | 
[**CreateConsumerByOrga**](OrganisationApi.md#CreateConsumerByOrga) | **Post** /organisations/{id}/consumers | 
[**CreateOrganisation**](OrganisationApi.md#CreateOrganisation) | **Post** /organisations | 
[**CreateProvider**](OrganisationApi.md#CreateProvider) | **Post** /organisations/{id}/addonproviders | 
[**DeleteAddonTagByOrgaAndAddonId**](OrganisationApi.md#DeleteAddonTagByOrgaAndAddonId) | **Delete** /organisations/{id}/addons/{addonId}/tags/{tag} | 
[**DeleteApplicationByOrgaAndAppId**](OrganisationApi.md#DeleteApplicationByOrgaAndAppId) | **Delete** /organisations/{id}/applications/{appId} | 
[**DeleteApplicationDependencyByOrgaAndAppId**](OrganisationApi.md#DeleteApplicationDependencyByOrgaAndAppId) | **Delete** /organisations/{id}/applications/{appId}/dependencies/{dependencyId} | 
[**DeleteApplicationTagByOrgaAndAppId**](OrganisationApi.md#DeleteApplicationTagByOrgaAndAppId) | **Delete** /organisations/{id}/applications/{appId}/tags/{tag} | 
[**DeleteConsumerByOrga**](OrganisationApi.md#DeleteConsumerByOrga) | **Delete** /organisations/{id}/consumers/{key} | 
[**DeleteOrganisation**](OrganisationApi.md#DeleteOrganisation) | **Delete** /organisations/{id} | 
[**DeletePaymentMethodByOrga**](OrganisationApi.md#DeletePaymentMethodByOrga) | **Delete** /organisations/{id}/payments/methods/{mId} | 
[**DeleteProvider**](OrganisationApi.md#DeleteProvider) | **Delete** /organisations/{id}/addonproviders/{providerId} | 
[**DeleteProviderFeature**](OrganisationApi.md#DeleteProviderFeature) | **Delete** /organisations/{id}/addonproviders/{providerId}/features/{featureId} | 
[**DeleteProviderPlan**](OrganisationApi.md#DeleteProviderPlan) | **Delete** /organisations/{id}/addonproviders/{providerId}/plans/{planId} | 
[**DeleteProviderPlanFeature**](OrganisationApi.md#DeleteProviderPlanFeature) | **Delete** /organisations/{id}/addonproviders/{providerId}/plans/{planId}/features/{featureName} | 
[**DeletePurchaseOrderByOrga**](OrganisationApi.md#DeletePurchaseOrderByOrga) | **Delete** /organisations/{id}/payments/billings/{bid} | 
[**DeleteRecurrentPaymentByOrga**](OrganisationApi.md#DeleteRecurrentPaymentByOrga) | **Delete** /organisations/{id}/payments/recurring | 
[**DeprovisionAddonByOrgaAndAddonId**](OrganisationApi.md#DeprovisionAddonByOrgaAndAddonId) | **Delete** /organisations/{id}/addons/{addonId} | 
[**EditApplicationByOrgaAndAppId**](OrganisationApi.md#EditApplicationByOrgaAndAppId) | **Put** /organisations/{id}/applications/{appId} | 
[**EditApplicationEnvByOrgaAndAppIdAndEnvName**](OrganisationApi.md#EditApplicationEnvByOrgaAndAppIdAndEnvName) | **Put** /organisations/{id}/applications/{appId}/env/{envName} | 
[**EditApplicationEnvironmentByOrgaAndAppId**](OrganisationApi.md#EditApplicationEnvironmentByOrgaAndAppId) | **Put** /organisations/{id}/applications/{appId}/env | 
[**EditOrganisation**](OrganisationApi.md#EditOrganisation) | **Put** /organisations/{id} | 
[**EditOrganisationMember**](OrganisationApi.md#EditOrganisationMember) | **Put** /organisations/{id}/members/{userId} | 
[**EditProviderPlan**](OrganisationApi.md#EditProviderPlan) | **Put** /organisations/{id}/addonproviders/{providerId}/plans/{planId} | 
[**EditProviderPlanFeature**](OrganisationApi.md#EditProviderPlanFeature) | **Put** /organisations/{id}/addonproviders/{providerId}/plans/{planId}/features/{featureName} | 
[**GetAddonByOrgaAndAddonId**](OrganisationApi.md#GetAddonByOrgaAndAddonId) | **Get** /organisations/{id}/addons/{addonId} | 
[**GetAddonEnvByOrgaAndAddonId**](OrganisationApi.md#GetAddonEnvByOrgaAndAddonId) | **Get** /organisations/{id}/addons/{addonId}/env | 
[**GetAddonInstance**](OrganisationApi.md#GetAddonInstance) | **Get** /organisations/{id}/addons/{addonId}/instances/{instanceId} | 
[**GetAddonInstances**](OrganisationApi.md#GetAddonInstances) | **Get** /organisations/{id}/addons/{addonId}/instances | 
[**GetAddonMigration**](OrganisationApi.md#GetAddonMigration) | **Get** /organisations/{id}/addons/{addonId}/migrations/{migrationId} | 
[**GetAddonMigrations**](OrganisationApi.md#GetAddonMigrations) | **Get** /organisations/{id}/addons/{addonId}/migrations | 
[**GetAddonSSODataForOrga**](OrganisationApi.md#GetAddonSSODataForOrga) | **Get** /organisations/{id}/addons/{addonId}/sso | 
[**GetAddonTagsByOrgaIdAndAddonId**](OrganisationApi.md#GetAddonTagsByOrgaIdAndAddonId) | **Get** /organisations/{id}/addons/{addonId}/tags | 
[**GetAddonsByOrgaId**](OrganisationApi.md#GetAddonsByOrgaId) | **Get** /organisations/{id}/addons | 
[**GetAddonsLinkedToApplicationByOrgaAndAppId**](OrganisationApi.md#GetAddonsLinkedToApplicationByOrgaAndAppId) | **Get** /organisations/{id}/applications/{appId}/addons | 
[**GetAllApplicationsByOrga**](OrganisationApi.md#GetAllApplicationsByOrga) | **Get** /organisations/{id}/applications | 
[**GetAmountForOrga**](OrganisationApi.md#GetAmountForOrga) | **Get** /organisations/{id}/credits | 
[**GetApplicationBranchesByOrgaAndAppId**](OrganisationApi.md#GetApplicationBranchesByOrgaAndAppId) | **Get** /organisations/{id}/applications/{appId}/branches | 
[**GetApplicationByOrgaAndAppId**](OrganisationApi.md#GetApplicationByOrgaAndAppId) | **Get** /organisations/{id}/applications/{appId} | 
[**GetApplicationDependenciesByOrgaAndAppId**](OrganisationApi.md#GetApplicationDependenciesByOrgaAndAppId) | **Get** /organisations/{id}/applications/{appId}/dependencies | 
[**GetApplicationDependenciesEnvByOrgaAndAppId**](OrganisationApi.md#GetApplicationDependenciesEnvByOrgaAndAppId) | **Get** /organisations/{id}/applications/{appId}/dependencies/env | 
[**GetApplicationDependentsByOrgaAndAppId**](OrganisationApi.md#GetApplicationDependentsByOrgaAndAppId) | **Get** /organisations/{id}/applications/{appId}/dependents | 
[**GetApplicationDeploymentForOrga**](OrganisationApi.md#GetApplicationDeploymentForOrga) | **Get** /organisations/{id}/applications/{appId}/deployments/{deploymentId} | 
[**GetApplicationDeploymentsForOrga**](OrganisationApi.md#GetApplicationDeploymentsForOrga) | **Get** /organisations/{id}/applications/{appId}/deployments | 
[**GetApplicationEnvByOrgaAndAppId**](OrganisationApi.md#GetApplicationEnvByOrgaAndAppId) | **Get** /organisations/{id}/applications/{appId}/env | 
[**GetApplicationInstanceByOrgaAndAppAndInstanceId**](OrganisationApi.md#GetApplicationInstanceByOrgaAndAppAndInstanceId) | **Get** /organisations/{id}/applications/{appId}/instances/{instanceId} | 
[**GetApplicationInstancesByOrgaAndAppId**](OrganisationApi.md#GetApplicationInstancesByOrgaAndAppId) | **Get** /organisations/{id}/applications/{appId}/instances | 
[**GetApplicationTagsByOrgaAndAppId**](OrganisationApi.md#GetApplicationTagsByOrgaAndAppId) | **Get** /organisations/{id}/applications/{appId}/tags | 
[**GetApplicationsLinkedToAddonByOrgaAndAddonId**](OrganisationApi.md#GetApplicationsLinkedToAddonByOrgaAndAddonId) | **Get** /organisations/{id}/addons/{addonId}/applications | 
[**GetConsumerByOrga**](OrganisationApi.md#GetConsumerByOrga) | **Get** /organisations/{id}/consumers/{key} | 
[**GetConsumerSecretByOrga**](OrganisationApi.md#GetConsumerSecretByOrga) | **Get** /organisations/{id}/consumers/{key}/secret | 
[**GetConsumersByOrga**](OrganisationApi.md#GetConsumersByOrga) | **Get** /organisations/{id}/consumers | 
[**GetConsumptionsForOrga**](OrganisationApi.md#GetConsumptionsForOrga) | **Get** /organisations/{id}/consumptions | 
[**GetDefaultMethodByOrga**](OrganisationApi.md#GetDefaultMethodByOrga) | **Get** /organisations/{id}/payments/methods/default | 
[**GetDeploymentsForAllApps**](OrganisationApi.md#GetDeploymentsForAllApps) | **Get** /organisations/{id}/deployments | 
[**GetEnvOfAddonsLinkedToApplicationByOrgaAndAppId**](OrganisationApi.md#GetEnvOfAddonsLinkedToApplicationByOrgaAndAppId) | **Get** /organisations/{id}/applications/{appId}/addons/env | 
[**GetExposedEnvByOrgaAndAppId**](OrganisationApi.md#GetExposedEnvByOrgaAndAppId) | **Get** /organisations/{id}/applications/{appId}/exposed_env | 
[**GetFavouriteVhostByOrgaAndAppId**](OrganisationApi.md#GetFavouriteVhostByOrgaAndAppId) | **Get** /organisations/{id}/applications/{appId}/vhosts/favourite | 
[**GetInstancesForAllAppsForOrga**](OrganisationApi.md#GetInstancesForAllAppsForOrga) | **Get** /organisations/{id}/instances | 
[**GetInvoiceByOrga**](OrganisationApi.md#GetInvoiceByOrga) | **Get** /organisations/{id}/payments/billings/{bid} | 
[**GetInvoicesByOrga**](OrganisationApi.md#GetInvoicesByOrga) | **Get** /organisations/{id}/payments/billings | 
[**GetMonthlyInvoiceByOrga**](OrganisationApi.md#GetMonthlyInvoiceByOrga) | **Get** /organisations/{id}/payments/monthlyinvoice | 
[**GetNamespaces**](OrganisationApi.md#GetNamespaces) | **Get** /organisations/{id}/namespaces | 
[**GetNewSetupIntentByOrga**](OrganisationApi.md#GetNewSetupIntentByOrga) | **Get** /organisations/{id}/payments/methods/newintent | 
[**GetOrganisation**](OrganisationApi.md#GetOrganisation) | **Get** /organisations/{id} | 
[**GetOrganisationMembers**](OrganisationApi.md#GetOrganisationMembers) | **Get** /organisations/{id}/members | 
[**GetPaymentInfoForOrga**](OrganisationApi.md#GetPaymentInfoForOrga) | **Get** /organisations/{id}/payment-info | 
[**GetPdfInvoiceByOrga**](OrganisationApi.md#GetPdfInvoiceByOrga) | **Get** /organisations/{id}/payments/billings/{bid}.pdf | 
[**GetPriceWithTaxByOrga**](OrganisationApi.md#GetPriceWithTaxByOrga) | **Get** /organisations/{id}/payments/fullprice/{price} | 
[**GetProviderFeatures**](OrganisationApi.md#GetProviderFeatures) | **Get** /organisations/{id}/addonproviders/{providerId}/features | 
[**GetProviderInfo**](OrganisationApi.md#GetProviderInfo) | **Get** /organisations/{id}/addonproviders/{providerId} | 
[**GetProviderPlan**](OrganisationApi.md#GetProviderPlan) | **Get** /organisations/{id}/addonproviders/{providerId}/plans/{planId} | 
[**GetProviderPlans**](OrganisationApi.md#GetProviderPlans) | **Get** /organisations/{id}/addonproviders/{providerId}/plans | 
[**GetProviderTags**](OrganisationApi.md#GetProviderTags) | **Get** /organisations/{id}/addonproviders/{providerId}/tags | 
[**GetProvidersInfo**](OrganisationApi.md#GetProvidersInfo) | **Get** /organisations/{id}/addonproviders | 
[**GetRecurrentPaymentByOrga**](OrganisationApi.md#GetRecurrentPaymentByOrga) | **Get** /organisations/{id}/payments/recurring | 
[**GetSSODataForOrga**](OrganisationApi.md#GetSSODataForOrga) | **Get** /organisations/{id}/addonproviders/{providerId}/sso | 
[**GetStripeTokenByOrga**](OrganisationApi.md#GetStripeTokenByOrga) | **Get** /organisations/{id}/payments/tokens/stripe | 
[**GetTcpRedirs**](OrganisationApi.md#GetTcpRedirs) | **Get** /organisations/{id}/applications/{appId}/tcpRedirs | 
[**GetUnpaidInvoicesByOrga**](OrganisationApi.md#GetUnpaidInvoicesByOrga) | **Get** /organisations/{id}/payments/methods | 
[**GetUnpaidInvoicesByOrga1**](OrganisationApi.md#GetUnpaidInvoicesByOrga1) | **Get** /organisations/{id}/payments/billings/unpaid | 
[**GetUserOrganisationss**](OrganisationApi.md#GetUserOrganisationss) | **Get** /organisations | 
[**GetVhostsByOrgaAndAppId**](OrganisationApi.md#GetVhostsByOrgaAndAppId) | **Get** /organisations/{id}/applications/{appId}/vhosts | 
[**LinkAddonToApplicationByOrgaAndAppId**](OrganisationApi.md#LinkAddonToApplicationByOrgaAndAppId) | **Post** /organisations/{id}/applications/{appId}/addons | 
[**MarkFavouriteVhostByOrgaAndAppId**](OrganisationApi.md#MarkFavouriteVhostByOrgaAndAppId) | **Put** /organisations/{id}/applications/{appId}/vhosts/favourite | 
[**PreorderAddonByOrgaId**](OrganisationApi.md#PreorderAddonByOrgaId) | **Post** /organisations/{id}/addons/preorders | 
[**PreorderMigration**](OrganisationApi.md#PreorderMigration) | **Get** /organisations/{id}/addons/{addonId}/migrations/preorders | 
[**ProvisionAddonByOrgaId**](OrganisationApi.md#ProvisionAddonByOrgaId) | **Post** /organisations/{id}/addons | 
[**RedeployApplicationByOrgaAndAppId**](OrganisationApi.md#RedeployApplicationByOrgaAndAppId) | **Post** /organisations/{id}/applications/{appId}/instances | 
[**RemoveApplicationEnvByOrgaAndAppIdAndEnvName**](OrganisationApi.md#RemoveApplicationEnvByOrgaAndAppIdAndEnvName) | **Delete** /organisations/{id}/applications/{appId}/env/{envName} | 
[**RemoveOrganisationMember**](OrganisationApi.md#RemoveOrganisationMember) | **Delete** /organisations/{id}/members/{userId} | 
[**RemoveTcpRedir**](OrganisationApi.md#RemoveTcpRedir) | **Delete** /organisations/{id}/applications/{appId}/tcpRedirs/{sourcePort} | 
[**RemoveVhostsByOrgaAndAppId**](OrganisationApi.md#RemoveVhostsByOrgaAndAppId) | **Delete** /organisations/{id}/applications/{appId}/vhosts/{domain} | 
[**ReplaceAddonTags**](OrganisationApi.md#ReplaceAddonTags) | **Put** /organisations/{id}/addons/{addonId}/tags | 
[**ReplaceApplicationTags**](OrganisationApi.md#ReplaceApplicationTags) | **Put** /organisations/{id}/applications/{appId}/tags | 
[**SetApplicationBranchByOrgaAndAppId**](OrganisationApi.md#SetApplicationBranchByOrgaAndAppId) | **Put** /organisations/{id}/applications/{appId}/branch | 
[**SetBuildInstanceFlavorByOrgaAndAppId**](OrganisationApi.md#SetBuildInstanceFlavorByOrgaAndAppId) | **Put** /organisations/{id}/applications/{appId}/buildflavor | 
[**SetDefaultMethodByOrga**](OrganisationApi.md#SetDefaultMethodByOrga) | **Put** /organisations/{id}/payments/methods/default | 
[**SetMaxCreditsPerMonthByOrga**](OrganisationApi.md#SetMaxCreditsPerMonthByOrga) | **Put** /organisations/{id}/payments/monthlyinvoice/maxcredit | 
[**SetOrgaAvatar**](OrganisationApi.md#SetOrgaAvatar) | **Put** /organisations/{id}/avatar | 
[**UndeployApplicationByOrgaAndAppId**](OrganisationApi.md#UndeployApplicationByOrgaAndAppId) | **Delete** /organisations/{id}/applications/{appId}/instances | 
[**UnlinkAddonFromApplicationByOrgaAndAppAnddAddonId**](OrganisationApi.md#UnlinkAddonFromApplicationByOrgaAndAppAnddAddonId) | **Delete** /organisations/{id}/applications/{appId}/addons/{addonId} | 
[**UnmarkFavouriteVhostByOrgaAndAppId**](OrganisationApi.md#UnmarkFavouriteVhostByOrgaAndAppId) | **Delete** /organisations/{id}/applications/{appId}/vhosts/favourite | 
[**UpdateAddonInfo**](OrganisationApi.md#UpdateAddonInfo) | **Put** /organisations/{id}/addons/{addonId} | 
[**UpdateConsumerByOrga**](OrganisationApi.md#UpdateConsumerByOrga) | **Put** /organisations/{id}/consumers/{key} | 
[**UpdateExposedEnvByOrgaAndAppId**](OrganisationApi.md#UpdateExposedEnvByOrgaAndAppId) | **Put** /organisations/{id}/applications/{appId}/exposed_env | 
[**UpdateProviderInfos**](OrganisationApi.md#UpdateProviderInfos) | **Put** /organisations/{id}/addonproviders/{providerId} | 



## AbortAddonMigration

> string AbortAddonMigration(ctx, id, addonId, migrationId)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
**addonId** | **string**|  | 
**migrationId** | **string**|  | 

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


## AddAddonTagByOrgaAndAddonId

> []string AddAddonTagByOrgaAndAddonId(ctx, id, addonId, tag)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
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


## AddApplicationByOrga

> ApplicationView AddApplicationByOrga(ctx, id, wannabeApplication)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
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


## AddApplicationDependencyByOrgaAndAppId

> Message AddApplicationDependencyByOrgaAndAppId(ctx, id, appId, dependencyId)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
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


## AddApplicationTagByOrgaAndAppId

> []string AddApplicationTagByOrgaAndAppId(ctx, id, appId, tag)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
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


## AddBetaTester

> Message AddBetaTester(ctx, id, providerId)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
**providerId** | **string**|  | 

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


## AddOrganisationMember

> Message AddOrganisationMember(ctx, id, wannabeMember, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
**wannabeMember** | [**WannabeMember**](WannabeMember.md)|  | 
 **optional** | ***AddOrganisationMemberOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AddOrganisationMemberOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **invitationKey** | **optional.String**|  | 

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


## AddPaymentMethodByOrga

> PaymentMethodView AddPaymentMethodByOrga(ctx, id, paymentData)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
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


## AddProviderFeature

> AddonFeatureView AddProviderFeature(ctx, id, providerId, wannabeAddonFeature)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
**providerId** | **string**|  | 
**wannabeAddonFeature** | [**WannabeAddonFeature**](WannabeAddonFeature.md)|  | 

### Return type

[**AddonFeatureView**](AddonFeatureView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddProviderPlan

> AddonPlanView AddProviderPlan(ctx, id, providerId, wannabeAddonPlan)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
**providerId** | **string**|  | 
**wannabeAddonPlan** | [**WannabeAddonPlan**](WannabeAddonPlan.md)|  | 

### Return type

[**AddonPlanView**](AddonPlanView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddTcpRedir

> TcpRedirView AddTcpRedir(ctx, id, appId, wannabeNamespace, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
**appId** | **string**|  | 
**wannabeNamespace** | [**WannabeNamespace**](WannabeNamespace.md)|  | 
 **optional** | ***AddTcpRedirOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AddTcpRedirOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **payment** | **optional.String**|  | 

### Return type

[**TcpRedirView**](TcpRedirView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddVhostsByOrgaAndAppId

> Message AddVhostsByOrgaAndAppId(ctx, id, appId, domain)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
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


## BuyDropsByOrga

> InvoiceRendering BuyDropsByOrga(ctx, id, wannaBuyPackage)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
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


## CancelApplicationDeploymentForOrga

> Message CancelApplicationDeploymentForOrga(ctx, id, appId, deploymentId)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
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


## ChangePlanByOrgaAndAddonId

> string ChangePlanByOrgaAndAddonId(ctx, id, addonId, wannabePlanChange)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
**addonId** | **string**|  | 
**wannabePlanChange** | [**WannabePlanChange**](WannabePlanChange.md)|  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChoosePaymentProviderByOrga

> NextInPaymentFlow ChoosePaymentProviderByOrga(ctx, id, bid, paymentProviderSelection)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
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


## CreateConsumerByOrga

> OAuth1ConsumerView CreateConsumerByOrga(ctx, id, wannabeOAuth1Consumer)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
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


## CreateOrganisation

> OrganisationView CreateOrganisation(ctx, wannabeOrganisation)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**wannabeOrganisation** | [**WannabeOrganisation**](WannabeOrganisation.md)|  | 

### Return type

[**OrganisationView**](OrganisationView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateProvider

> AddonProviderInfoFullView CreateProvider(ctx, id, wannabeAddonProvider)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
**wannabeAddonProvider** | [**WannabeAddonProvider**](WannabeAddonProvider.md)|  | 

### Return type

[**AddonProviderInfoFullView**](AddonProviderInfoFullView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAddonTagByOrgaAndAddonId

> []string DeleteAddonTagByOrgaAndAddonId(ctx, id, addonId, tag)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
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


## DeleteApplicationByOrgaAndAppId

> Message DeleteApplicationByOrgaAndAppId(ctx, id, appId)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
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


## DeleteApplicationDependencyByOrgaAndAppId

> DeleteApplicationDependencyByOrgaAndAppId(ctx, id, appId, dependencyId)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
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


## DeleteApplicationTagByOrgaAndAppId

> []string DeleteApplicationTagByOrgaAndAppId(ctx, id, appId, tag)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
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


## DeleteConsumerByOrga

> DeleteConsumerByOrga(ctx, id, key)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
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


## DeleteOrganisation

> Message DeleteOrganisation(ctx, id)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 

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


## DeletePaymentMethodByOrga

> DeletePaymentMethodByOrga(ctx, id, mId)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
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


## DeleteProvider

> DeleteProvider(ctx, id, providerId)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
**providerId** | **string**|  | 

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


## DeleteProviderFeature

> DeleteProviderFeature(ctx, id, providerId, featureId)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
**providerId** | **string**|  | 
**featureId** | **string**|  | 

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


## DeleteProviderPlan

> DeleteProviderPlan(ctx, id, providerId, planId)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
**providerId** | **string**|  | 
**planId** | **string**|  | 

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


## DeleteProviderPlanFeature

> DeleteProviderPlanFeature(ctx, id, providerId, planId, featureName)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
**providerId** | **string**|  | 
**planId** | **string**|  | 
**featureName** | **string**|  | 

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


## DeletePurchaseOrderByOrga

> DeletePurchaseOrderByOrga(ctx, id, bid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
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


## DeleteRecurrentPaymentByOrga

> DeleteRecurrentPaymentByOrga(ctx, id)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 

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


## DeprovisionAddonByOrgaAndAddonId

> Message DeprovisionAddonByOrgaAndAddonId(ctx, id, addonId)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
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


## EditApplicationByOrgaAndAppId

> ApplicationView EditApplicationByOrgaAndAppId(ctx, id, appId, wannabeApplication)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
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


## EditApplicationEnvByOrgaAndAppIdAndEnvName

> ApplicationView EditApplicationEnvByOrgaAndAppIdAndEnvName(ctx, id, appId, envName, wannabeValue)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
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


## EditApplicationEnvironmentByOrgaAndAppId

> ApplicationView EditApplicationEnvironmentByOrgaAndAppId(ctx, id, appId, body)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
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


## EditOrganisation

> OrganisationView EditOrganisation(ctx, id, wannabeOrganisation)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
**wannabeOrganisation** | [**WannabeOrganisation**](WannabeOrganisation.md)|  | 

### Return type

[**OrganisationView**](OrganisationView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditOrganisationMember

> Message EditOrganisationMember(ctx, id, userId, wannabeMember)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
**userId** | **string**|  | 
**wannabeMember** | [**WannabeMember**](WannabeMember.md)|  | 

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


## EditProviderPlan

> AddonPlanView EditProviderPlan(ctx, id, providerId, planId, wannabeAddonPlan)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
**providerId** | **string**|  | 
**planId** | **string**|  | 
**wannabeAddonPlan** | [**WannabeAddonPlan**](WannabeAddonPlan.md)|  | 

### Return type

[**AddonPlanView**](AddonPlanView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditProviderPlanFeature

> AddonPlanView EditProviderPlanFeature(ctx, id, providerId, planId, featureName, addonFeatureInstanceView)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
**providerId** | **string**|  | 
**planId** | **string**|  | 
**featureName** | **string**|  | 
**addonFeatureInstanceView** | [**AddonFeatureInstanceView**](AddonFeatureInstanceView.md)|  | 

### Return type

[**AddonPlanView**](AddonPlanView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAddonByOrgaAndAddonId

> AddonView GetAddonByOrgaAndAddonId(ctx, id, addonId)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
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


## GetAddonEnvByOrgaAndAddonId

> []AddonEnvironmentView GetAddonEnvByOrgaAndAddonId(ctx, id, addonId)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
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


## GetAddonInstance

> string GetAddonInstance(ctx, id, addonId, instanceId)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
**addonId** | **string**|  | 
**instanceId** | **string**|  | 

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


## GetAddonInstances

> []SuperNovaInstanceView GetAddonInstances(ctx, id, addonId, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
**addonId** | **string**|  | 
 **optional** | ***GetAddonInstancesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetAddonInstancesOpts struct


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


## GetAddonMigration

> string GetAddonMigration(ctx, id, addonId, migrationId)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
**addonId** | **string**|  | 
**migrationId** | **string**|  | 

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


## GetAddonMigrations

> string GetAddonMigrations(ctx, id, addonId)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
**addonId** | **string**|  | 

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


## GetAddonSSODataForOrga

> AddonProviderSsoData GetAddonSSODataForOrga(ctx, id, addonId)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
**addonId** | **string**|  | 

### Return type

[**AddonProviderSsoData**](AddonProviderSSOData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAddonTagsByOrgaIdAndAddonId

> []string GetAddonTagsByOrgaIdAndAddonId(ctx, id, addonId)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
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


## GetAddonsByOrgaId

> []AddonView GetAddonsByOrgaId(ctx, id)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 

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


## GetAddonsLinkedToApplicationByOrgaAndAppId

> []AddonView GetAddonsLinkedToApplicationByOrgaAndAppId(ctx, id, appId)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
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


## GetAllApplicationsByOrga

> []ApplicationView GetAllApplicationsByOrga(ctx, id, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
 **optional** | ***GetAllApplicationsByOrgaOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetAllApplicationsByOrgaOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **instanceId** | **optional.String**|  | 

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


## GetAmountForOrga

> DropCountView GetAmountForOrga(ctx, id)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 

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


## GetApplicationBranchesByOrgaAndAppId

> []string GetApplicationBranchesByOrgaAndAppId(ctx, id, appId)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
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


## GetApplicationByOrgaAndAppId

> ApplicationView GetApplicationByOrgaAndAppId(ctx, id, appId)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
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


## GetApplicationDependenciesByOrgaAndAppId

> []ApplicationView GetApplicationDependenciesByOrgaAndAppId(ctx, id, appId)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
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


## GetApplicationDependenciesEnvByOrgaAndAppId

> GetApplicationDependenciesEnvByOrgaAndAppId(ctx, id, appId)



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


## GetApplicationDependentsByOrgaAndAppId

> []ApplicationView GetApplicationDependentsByOrgaAndAppId(ctx, id, appId)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
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


## GetApplicationDeploymentForOrga

> GetApplicationDeploymentForOrga(ctx, id, appId, deploymentId)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
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


## GetApplicationDeploymentsForOrga

> []DeploymentView GetApplicationDeploymentsForOrga(ctx, id, appId, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
**appId** | **string**|  | 
 **optional** | ***GetApplicationDeploymentsForOrgaOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetApplicationDeploymentsForOrgaOpts struct


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


## GetApplicationEnvByOrgaAndAppId

> []AddonEnvironmentView GetApplicationEnvByOrgaAndAppId(ctx, id, appId)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
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


## GetApplicationInstanceByOrgaAndAppAndInstanceId

> string GetApplicationInstanceByOrgaAndAppAndInstanceId(ctx, id, appId, instanceId)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
**appId** | **string**|  | 
**instanceId** | **string**|  | 

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


## GetApplicationInstancesByOrgaAndAppId

> []SuperNovaInstanceView GetApplicationInstancesByOrgaAndAppId(ctx, id, appId, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
**appId** | **string**|  | 
 **optional** | ***GetApplicationInstancesByOrgaAndAppIdOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetApplicationInstancesByOrgaAndAppIdOpts struct


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


## GetApplicationTagsByOrgaAndAppId

> []string GetApplicationTagsByOrgaAndAppId(ctx, id, appId)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
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


## GetApplicationsLinkedToAddonByOrgaAndAddonId

> []ApplicationView GetApplicationsLinkedToAddonByOrgaAndAddonId(ctx, id, addonId)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
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


## GetConsumerByOrga

> OAuth1ConsumerView GetConsumerByOrga(ctx, id, key)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
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


## GetConsumerSecretByOrga

> SecretView GetConsumerSecretByOrga(ctx, id, key)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
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


## GetConsumersByOrga

> []OAuth1ConsumerView GetConsumersByOrga(ctx, id)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 

### Return type

[**[]OAuth1ConsumerView**](OAuth1ConsumerView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConsumptionsForOrga

> string GetConsumptionsForOrga(ctx, id, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
 **optional** | ***GetConsumptionsForOrgaOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetConsumptionsForOrgaOpts struct


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


## GetDefaultMethodByOrga

> PaymentMethodView GetDefaultMethodByOrga(ctx, id)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 

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


## GetDeploymentsForAllApps

> GetDeploymentsForAllApps(ctx, id, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
 **optional** | ***GetDeploymentsForAllAppsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetDeploymentsForAllAppsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int32**|  | 

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


## GetEnvOfAddonsLinkedToApplicationByOrgaAndAppId

> []LinkedAddonEnvironmentView GetEnvOfAddonsLinkedToApplicationByOrgaAndAppId(ctx, id, appId)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
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


## GetExposedEnvByOrgaAndAppId

> string GetExposedEnvByOrgaAndAppId(ctx, id, appId)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
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


## GetFavouriteVhostByOrgaAndAppId

> VhostView GetFavouriteVhostByOrgaAndAppId(ctx, id, appId)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
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


## GetInstancesForAllAppsForOrga

> string GetInstancesForAllAppsForOrga(ctx, id)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 

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


## GetInvoiceByOrga

> InvoiceRendering GetInvoiceByOrga(ctx, id, bid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
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


## GetInvoicesByOrga

> []InvoiceRendering GetInvoicesByOrga(ctx, id)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 

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


## GetMonthlyInvoiceByOrga

> string GetMonthlyInvoiceByOrga(ctx, id)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 

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


## GetNamespaces

> []NamespaceView GetNamespaces(ctx, id)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 

### Return type

[**[]NamespaceView**](NamespaceView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNewSetupIntentByOrga

> SetupIntentView GetNewSetupIntentByOrga(ctx, id, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
 **optional** | ***GetNewSetupIntentByOrgaOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetNewSetupIntentByOrgaOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | **optional.String**|  | 

### Return type

[**SetupIntentView**](SetupIntentView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganisation

> OrganisationView GetOrganisation(ctx, id)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 

### Return type

[**OrganisationView**](OrganisationView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganisationMembers

> []OrganisationMemberView GetOrganisationMembers(ctx, id)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 

### Return type

[**[]OrganisationMemberView**](OrganisationMemberView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPaymentInfoForOrga

> PaymentInfoView GetPaymentInfoForOrga(ctx, id)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 

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


## GetPdfInvoiceByOrga

> GetPdfInvoiceByOrga(ctx, id, bid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
**bid** | **string**|  | 
 **optional** | ***GetPdfInvoiceByOrgaOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetPdfInvoiceByOrgaOpts struct


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


## GetPriceWithTaxByOrga

> PriceWithTaxInfo GetPriceWithTaxByOrga(ctx, id, price)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
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


## GetProviderFeatures

> []AddonFeatureView GetProviderFeatures(ctx, id, providerId)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
**providerId** | **string**|  | 

### Return type

[**[]AddonFeatureView**](AddonFeatureView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProviderInfo

> AddonProviderInfoView GetProviderInfo(ctx, id, providerId)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
**providerId** | **string**|  | 

### Return type

[**AddonProviderInfoView**](AddonProviderInfoView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProviderPlan

> AddonPlanView GetProviderPlan(ctx, id, providerId, planId)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
**providerId** | **string**|  | 
**planId** | **string**|  | 

### Return type

[**AddonPlanView**](AddonPlanView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProviderPlans

> []AddonPlanView GetProviderPlans(ctx, id, providerId)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
**providerId** | **string**|  | 

### Return type

[**[]AddonPlanView**](AddonPlanView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProviderTags

> []string GetProviderTags(ctx, id, providerId)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
**providerId** | **string**|  | 

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


## GetProvidersInfo

> []AddonProviderInfoFullView GetProvidersInfo(ctx, id)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 

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


## GetRecurrentPaymentByOrga

> RecurrentPaymentView GetRecurrentPaymentByOrga(ctx, id)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 

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


## GetSSODataForOrga

> AddonProviderSsoData GetSSODataForOrga(ctx, id, providerId)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
**providerId** | **string**|  | 

### Return type

[**AddonProviderSsoData**](AddonProviderSSOData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStripeTokenByOrga

> BraintreeToken GetStripeTokenByOrga(ctx, id)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 

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


## GetTcpRedirs

> []TcpRedirView GetTcpRedirs(ctx, id, appId)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
**appId** | **string**|  | 

### Return type

[**[]TcpRedirView**](TcpRedirView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUnpaidInvoicesByOrga

> []PaymentMethodView GetUnpaidInvoicesByOrga(ctx, id)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 

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


## GetUnpaidInvoicesByOrga1

> []InvoiceRendering GetUnpaidInvoicesByOrga1(ctx, id)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 

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


## GetUserOrganisationss

> []OrganisationView GetUserOrganisationss(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetUserOrganisationssOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetUserOrganisationssOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **user** | **optional.String**|  | 

### Return type

[**[]OrganisationView**](OrganisationView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVhostsByOrgaAndAppId

> []VhostView GetVhostsByOrgaAndAppId(ctx, id, appId)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
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


## LinkAddonToApplicationByOrgaAndAppId

> LinkAddonToApplicationByOrgaAndAppId(ctx, id, appId, body)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
**appId** | **string**|  | 
**body** | **string**|  | 

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


## MarkFavouriteVhostByOrgaAndAppId

> VhostView MarkFavouriteVhostByOrgaAndAppId(ctx, id, appId, vhostView)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
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


## PreorderAddonByOrgaId

> InvoiceRendering PreorderAddonByOrgaId(ctx, id, wannabeAddonProvision)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
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


## PreorderMigration

> string PreorderMigration(ctx, id, addonId, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
**addonId** | **string**|  | 
 **optional** | ***PreorderMigrationOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PreorderMigrationOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **planId** | **optional.String**|  | 

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


## ProvisionAddonByOrgaId

> AddonView ProvisionAddonByOrgaId(ctx, id, wannabeAddonProvision)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
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


## RedeployApplicationByOrgaAndAppId

> RedeployApplicationByOrgaAndAppId(ctx, id, appId, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
**appId** | **string**|  | 
 **optional** | ***RedeployApplicationByOrgaAndAppIdOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a RedeployApplicationByOrgaAndAppIdOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **commit** | **optional.String**|  | 
 **useCache** | **optional.String**|  | 

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


## RemoveApplicationEnvByOrgaAndAppIdAndEnvName

> ApplicationView RemoveApplicationEnvByOrgaAndAppIdAndEnvName(ctx, id, appId, envName)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
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


## RemoveOrganisationMember

> Message RemoveOrganisationMember(ctx, id, userId)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
**userId** | **string**|  | 

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


## RemoveTcpRedir

> RemoveTcpRedir(ctx, id, appId, sourcePort, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
**appId** | **string**|  | 
**sourcePort** | **int64**|  | 
 **optional** | ***RemoveTcpRedirOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a RemoveTcpRedirOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **namespace** | **optional.String**|  | 

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


## RemoveVhostsByOrgaAndAppId

> Message RemoveVhostsByOrgaAndAppId(ctx, id, appId, domain)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
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


## ReplaceAddonTags

> []string ReplaceAddonTags(ctx, id, addonId, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
**addonId** | **string**|  | 
 **optional** | ***ReplaceAddonTagsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ReplaceAddonTagsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **optional.String**|  | 

### Return type

**[]string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceApplicationTags

> []string ReplaceApplicationTags(ctx, id, appId, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
**appId** | **string**|  | 
 **optional** | ***ReplaceApplicationTagsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ReplaceApplicationTagsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **optional.String**|  | 

### Return type

**[]string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetApplicationBranchByOrgaAndAppId

> SetApplicationBranchByOrgaAndAppId(ctx, id, appId, wannabeBranch)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
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


## SetBuildInstanceFlavorByOrgaAndAppId

> SetBuildInstanceFlavorByOrgaAndAppId(ctx, id, appId, wannabeBuildFlavor)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
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


## SetDefaultMethodByOrga

> SetDefaultMethodByOrga(ctx, id, paymentData)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
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


## SetMaxCreditsPerMonthByOrga

> string SetMaxCreditsPerMonthByOrga(ctx, id, wannabeMaxCredits)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
**wannabeMaxCredits** | [**WannabeMaxCredits**](WannabeMaxCredits.md)|  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetOrgaAvatar

> UrlView SetOrgaAvatar(ctx, id)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 

### Return type

[**UrlView**](UrlView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UndeployApplicationByOrgaAndAppId

> Message UndeployApplicationByOrgaAndAppId(ctx, id, appId)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
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


## UnlinkAddonFromApplicationByOrgaAndAppAnddAddonId

> UnlinkAddonFromApplicationByOrgaAndAppAnddAddonId(ctx, id, appId, addonId)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
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


## UnmarkFavouriteVhostByOrgaAndAppId

> UnmarkFavouriteVhostByOrgaAndAppId(ctx, id, appId)



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


## UpdateAddonInfo

> AddonView UpdateAddonInfo(ctx, id, addonId, wannabeAddonProvision)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
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


## UpdateConsumerByOrga

> OAuth1ConsumerView UpdateConsumerByOrga(ctx, id, key, wannabeOAuth1Consumer)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
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


## UpdateExposedEnvByOrgaAndAppId

> Message UpdateExposedEnvByOrgaAndAppId(ctx, id, appId, body)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
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


## UpdateProviderInfos

> AddonProviderInfoView UpdateProviderInfos(ctx, id, providerId, wannabeAddonProviderInfos)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
**providerId** | **string**|  | 
**wannabeAddonProviderInfos** | [**WannabeAddonProviderInfos**](WannabeAddonProviderInfos.md)|  | 

### Return type

[**AddonProviderInfoView**](AddonProviderInfoView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

