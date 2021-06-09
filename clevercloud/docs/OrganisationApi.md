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

> string AbortAddonMigration(ctx, id, addonId, migrationId).Execute()



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
    addonId := "addonId_example" // string | 
    migrationId := "migrationId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrganisationApi.AbortAddonMigration(context.Background(), id, addonId, migrationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganisationApi.AbortAddonMigration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AbortAddonMigration`: string
    fmt.Fprintf(os.Stdout, "Response from `OrganisationApi.AbortAddonMigration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**addonId** | **string** |  | 
**migrationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAbortAddonMigrationRequest struct via the builder pattern


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


## AddAddonTagByOrgaAndAddonId

> []string AddAddonTagByOrgaAndAddonId(ctx, id, addonId, tag).Execute()



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
    addonId := "addonId_example" // string | 
    tag := "tag_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrganisationApi.AddAddonTagByOrgaAndAddonId(context.Background(), id, addonId, tag).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganisationApi.AddAddonTagByOrgaAndAddonId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddAddonTagByOrgaAndAddonId`: []string
    fmt.Fprintf(os.Stdout, "Response from `OrganisationApi.AddAddonTagByOrgaAndAddonId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**addonId** | **string** |  | 
**tag** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddAddonTagByOrgaAndAddonIdRequest struct via the builder pattern


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


## AddApplicationByOrga

> ApplicationView AddApplicationByOrga(ctx, id).WannabeApplication(wannabeApplication).Execute()



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
    wannabeApplication := *openapiclient.NewWannabeApplication() // WannabeApplication | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrganisationApi.AddApplicationByOrga(context.Background(), id).WannabeApplication(wannabeApplication).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganisationApi.AddApplicationByOrga``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddApplicationByOrga`: ApplicationView
    fmt.Fprintf(os.Stdout, "Response from `OrganisationApi.AddApplicationByOrga`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddApplicationByOrgaRequest struct via the builder pattern


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


## AddApplicationDependencyByOrgaAndAppId

> Message AddApplicationDependencyByOrgaAndAppId(ctx, id, appId, dependencyId).Execute()



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
    dependencyId := "dependencyId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrganisationApi.AddApplicationDependencyByOrgaAndAppId(context.Background(), id, appId, dependencyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganisationApi.AddApplicationDependencyByOrgaAndAppId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddApplicationDependencyByOrgaAndAppId`: Message
    fmt.Fprintf(os.Stdout, "Response from `OrganisationApi.AddApplicationDependencyByOrgaAndAppId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**appId** | **string** |  | 
**dependencyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddApplicationDependencyByOrgaAndAppIdRequest struct via the builder pattern


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


## AddApplicationTagByOrgaAndAppId

> []string AddApplicationTagByOrgaAndAppId(ctx, id, appId, tag).Execute()



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
    tag := "tag_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrganisationApi.AddApplicationTagByOrgaAndAppId(context.Background(), id, appId, tag).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganisationApi.AddApplicationTagByOrgaAndAppId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddApplicationTagByOrgaAndAppId`: []string
    fmt.Fprintf(os.Stdout, "Response from `OrganisationApi.AddApplicationTagByOrgaAndAppId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**appId** | **string** |  | 
**tag** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddApplicationTagByOrgaAndAppIdRequest struct via the builder pattern


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


## AddBetaTester

> Message AddBetaTester(ctx, id, providerId).Execute()



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
    providerId := "providerId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrganisationApi.AddBetaTester(context.Background(), id, providerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganisationApi.AddBetaTester``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddBetaTester`: Message
    fmt.Fprintf(os.Stdout, "Response from `OrganisationApi.AddBetaTester`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**providerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddBetaTesterRequest struct via the builder pattern


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


## AddOrganisationMember

> Message AddOrganisationMember(ctx, id).WannabeMember(wannabeMember).InvitationKey(invitationKey).Execute()



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
    wannabeMember := *openapiclient.NewWannabeMember() // WannabeMember | 
    invitationKey := "invitationKey_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrganisationApi.AddOrganisationMember(context.Background(), id).WannabeMember(wannabeMember).InvitationKey(invitationKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganisationApi.AddOrganisationMember``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddOrganisationMember`: Message
    fmt.Fprintf(os.Stdout, "Response from `OrganisationApi.AddOrganisationMember`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddOrganisationMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **wannabeMember** | [**WannabeMember**](WannabeMember.md) |  | 
 **invitationKey** | **string** |  | 

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

> PaymentMethodView AddPaymentMethodByOrga(ctx, id).PaymentData(paymentData).Execute()



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
    paymentData := *openapiclient.NewPaymentData() // PaymentData | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrganisationApi.AddPaymentMethodByOrga(context.Background(), id).PaymentData(paymentData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganisationApi.AddPaymentMethodByOrga``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddPaymentMethodByOrga`: PaymentMethodView
    fmt.Fprintf(os.Stdout, "Response from `OrganisationApi.AddPaymentMethodByOrga`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddPaymentMethodByOrgaRequest struct via the builder pattern


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


## AddProviderFeature

> AddonFeatureView AddProviderFeature(ctx, id, providerId).WannabeAddonFeature(wannabeAddonFeature).Execute()



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
    providerId := "providerId_example" // string | 
    wannabeAddonFeature := *openapiclient.NewWannabeAddonFeature() // WannabeAddonFeature | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrganisationApi.AddProviderFeature(context.Background(), id, providerId).WannabeAddonFeature(wannabeAddonFeature).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganisationApi.AddProviderFeature``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddProviderFeature`: AddonFeatureView
    fmt.Fprintf(os.Stdout, "Response from `OrganisationApi.AddProviderFeature`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**providerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddProviderFeatureRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **wannabeAddonFeature** | [**WannabeAddonFeature**](WannabeAddonFeature.md) |  | 

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

> AddonPlanView AddProviderPlan(ctx, id, providerId).WannabeAddonPlan(wannabeAddonPlan).Execute()



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
    providerId := "providerId_example" // string | 
    wannabeAddonPlan := *openapiclient.NewWannabeAddonPlan() // WannabeAddonPlan | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrganisationApi.AddProviderPlan(context.Background(), id, providerId).WannabeAddonPlan(wannabeAddonPlan).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganisationApi.AddProviderPlan``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddProviderPlan`: AddonPlanView
    fmt.Fprintf(os.Stdout, "Response from `OrganisationApi.AddProviderPlan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**providerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddProviderPlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **wannabeAddonPlan** | [**WannabeAddonPlan**](WannabeAddonPlan.md) |  | 

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

> TcpRedirView AddTcpRedir(ctx, id, appId).WannabeNamespace(wannabeNamespace).Payment(payment).Execute()



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
    wannabeNamespace := *openapiclient.NewWannabeNamespace() // WannabeNamespace | 
    payment := "payment_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrganisationApi.AddTcpRedir(context.Background(), id, appId).WannabeNamespace(wannabeNamespace).Payment(payment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganisationApi.AddTcpRedir``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddTcpRedir`: TcpRedirView
    fmt.Fprintf(os.Stdout, "Response from `OrganisationApi.AddTcpRedir`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**appId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddTcpRedirRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **wannabeNamespace** | [**WannabeNamespace**](WannabeNamespace.md) |  | 
 **payment** | **string** |  | 

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

> Message AddVhostsByOrgaAndAppId(ctx, id, appId, domain).Execute()



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
    domain := "domain_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrganisationApi.AddVhostsByOrgaAndAppId(context.Background(), id, appId, domain).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganisationApi.AddVhostsByOrgaAndAppId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddVhostsByOrgaAndAppId`: Message
    fmt.Fprintf(os.Stdout, "Response from `OrganisationApi.AddVhostsByOrgaAndAppId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**appId** | **string** |  | 
**domain** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddVhostsByOrgaAndAppIdRequest struct via the builder pattern


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


## BuyDropsByOrga

> InvoiceRendering BuyDropsByOrga(ctx, id).WannaBuyPackage(wannaBuyPackage).Execute()



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
    wannaBuyPackage := *openapiclient.NewWannaBuyPackage() // WannaBuyPackage | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrganisationApi.BuyDropsByOrga(context.Background(), id).WannaBuyPackage(wannaBuyPackage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganisationApi.BuyDropsByOrga``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BuyDropsByOrga`: InvoiceRendering
    fmt.Fprintf(os.Stdout, "Response from `OrganisationApi.BuyDropsByOrga`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBuyDropsByOrgaRequest struct via the builder pattern


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


## CancelApplicationDeploymentForOrga

> Message CancelApplicationDeploymentForOrga(ctx, id, appId, deploymentId).Execute()



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
    deploymentId := "deploymentId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrganisationApi.CancelApplicationDeploymentForOrga(context.Background(), id, appId, deploymentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganisationApi.CancelApplicationDeploymentForOrga``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CancelApplicationDeploymentForOrga`: Message
    fmt.Fprintf(os.Stdout, "Response from `OrganisationApi.CancelApplicationDeploymentForOrga`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**appId** | **string** |  | 
**deploymentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelApplicationDeploymentForOrgaRequest struct via the builder pattern


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


## ChangePlanByOrgaAndAddonId

> string ChangePlanByOrgaAndAddonId(ctx, id, addonId).WannabePlanChange(wannabePlanChange).Execute()



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
    addonId := "addonId_example" // string | 
    wannabePlanChange := *openapiclient.NewWannabePlanChange() // WannabePlanChange | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrganisationApi.ChangePlanByOrgaAndAddonId(context.Background(), id, addonId).WannabePlanChange(wannabePlanChange).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganisationApi.ChangePlanByOrgaAndAddonId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ChangePlanByOrgaAndAddonId`: string
    fmt.Fprintf(os.Stdout, "Response from `OrganisationApi.ChangePlanByOrgaAndAddonId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**addonId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiChangePlanByOrgaAndAddonIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **wannabePlanChange** | [**WannabePlanChange**](WannabePlanChange.md) |  | 

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

> NextInPaymentFlow ChoosePaymentProviderByOrga(ctx, id, bid).PaymentProviderSelection(paymentProviderSelection).Execute()



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
    bid := "bid_example" // string | 
    paymentProviderSelection := *openapiclient.NewPaymentProviderSelection() // PaymentProviderSelection | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrganisationApi.ChoosePaymentProviderByOrga(context.Background(), id, bid).PaymentProviderSelection(paymentProviderSelection).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganisationApi.ChoosePaymentProviderByOrga``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ChoosePaymentProviderByOrga`: NextInPaymentFlow
    fmt.Fprintf(os.Stdout, "Response from `OrganisationApi.ChoosePaymentProviderByOrga`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**bid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiChoosePaymentProviderByOrgaRequest struct via the builder pattern


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


## CreateConsumerByOrga

> OAuth1ConsumerView CreateConsumerByOrga(ctx, id).WannabeOAuth1Consumer(wannabeOAuth1Consumer).Execute()



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
    wannabeOAuth1Consumer := *openapiclient.NewWannabeOAuth1Consumer() // WannabeOAuth1Consumer | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrganisationApi.CreateConsumerByOrga(context.Background(), id).WannabeOAuth1Consumer(wannabeOAuth1Consumer).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganisationApi.CreateConsumerByOrga``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateConsumerByOrga`: OAuth1ConsumerView
    fmt.Fprintf(os.Stdout, "Response from `OrganisationApi.CreateConsumerByOrga`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateConsumerByOrgaRequest struct via the builder pattern


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


## CreateOrganisation

> OrganisationView CreateOrganisation(ctx).WannabeOrganisation(wannabeOrganisation).Execute()



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
    wannabeOrganisation := *openapiclient.NewWannabeOrganisation() // WannabeOrganisation | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrganisationApi.CreateOrganisation(context.Background()).WannabeOrganisation(wannabeOrganisation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganisationApi.CreateOrganisation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrganisation`: OrganisationView
    fmt.Fprintf(os.Stdout, "Response from `OrganisationApi.CreateOrganisation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganisationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **wannabeOrganisation** | [**WannabeOrganisation**](WannabeOrganisation.md) |  | 

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

> AddonProviderInfoFullView CreateProvider(ctx, id).WannabeAddonProvider(wannabeAddonProvider).Execute()



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
    wannabeAddonProvider := *openapiclient.NewWannabeAddonProvider() // WannabeAddonProvider | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrganisationApi.CreateProvider(context.Background(), id).WannabeAddonProvider(wannabeAddonProvider).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganisationApi.CreateProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateProvider`: AddonProviderInfoFullView
    fmt.Fprintf(os.Stdout, "Response from `OrganisationApi.CreateProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **wannabeAddonProvider** | [**WannabeAddonProvider**](WannabeAddonProvider.md) |  | 

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

> []string DeleteAddonTagByOrgaAndAddonId(ctx, id, addonId, tag).Execute()



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
    addonId := "addonId_example" // string | 
    tag := "tag_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrganisationApi.DeleteAddonTagByOrgaAndAddonId(context.Background(), id, addonId, tag).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganisationApi.DeleteAddonTagByOrgaAndAddonId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteAddonTagByOrgaAndAddonId`: []string
    fmt.Fprintf(os.Stdout, "Response from `OrganisationApi.DeleteAddonTagByOrgaAndAddonId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**addonId** | **string** |  | 
**tag** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAddonTagByOrgaAndAddonIdRequest struct via the builder pattern


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


## DeleteApplicationByOrgaAndAppId

> Message DeleteApplicationByOrgaAndAppId(ctx, id, appId).Execute()



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
    resp, r, err := api_client.OrganisationApi.DeleteApplicationByOrgaAndAppId(context.Background(), id, appId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganisationApi.DeleteApplicationByOrgaAndAppId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteApplicationByOrgaAndAppId`: Message
    fmt.Fprintf(os.Stdout, "Response from `OrganisationApi.DeleteApplicationByOrgaAndAppId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**appId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteApplicationByOrgaAndAppIdRequest struct via the builder pattern


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


## DeleteApplicationDependencyByOrgaAndAppId

> DeleteApplicationDependencyByOrgaAndAppId(ctx, id, appId, dependencyId).Execute()



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
    dependencyId := "dependencyId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrganisationApi.DeleteApplicationDependencyByOrgaAndAppId(context.Background(), id, appId, dependencyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganisationApi.DeleteApplicationDependencyByOrgaAndAppId``: %v\n", err)
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
**dependencyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteApplicationDependencyByOrgaAndAppIdRequest struct via the builder pattern


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


## DeleteApplicationTagByOrgaAndAppId

> []string DeleteApplicationTagByOrgaAndAppId(ctx, id, appId, tag).Execute()



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
    tag := "tag_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrganisationApi.DeleteApplicationTagByOrgaAndAppId(context.Background(), id, appId, tag).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganisationApi.DeleteApplicationTagByOrgaAndAppId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteApplicationTagByOrgaAndAppId`: []string
    fmt.Fprintf(os.Stdout, "Response from `OrganisationApi.DeleteApplicationTagByOrgaAndAppId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**appId** | **string** |  | 
**tag** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteApplicationTagByOrgaAndAppIdRequest struct via the builder pattern


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


## DeleteConsumerByOrga

> DeleteConsumerByOrga(ctx, id, key).Execute()



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
    key := "key_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrganisationApi.DeleteConsumerByOrga(context.Background(), id, key).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganisationApi.DeleteConsumerByOrga``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**key** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteConsumerByOrgaRequest struct via the builder pattern


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


## DeleteOrganisation

> Message DeleteOrganisation(ctx, id).Execute()



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
    resp, r, err := api_client.OrganisationApi.DeleteOrganisation(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganisationApi.DeleteOrganisation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteOrganisation`: Message
    fmt.Fprintf(os.Stdout, "Response from `OrganisationApi.DeleteOrganisation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrganisationRequest struct via the builder pattern


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


## DeletePaymentMethodByOrga

> DeletePaymentMethodByOrga(ctx, id, mId).Execute()



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
    mId := "mId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrganisationApi.DeletePaymentMethodByOrga(context.Background(), id, mId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganisationApi.DeletePaymentMethodByOrga``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**mId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePaymentMethodByOrgaRequest struct via the builder pattern


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


## DeleteProvider

> DeleteProvider(ctx, id, providerId).Execute()



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
    providerId := "providerId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrganisationApi.DeleteProvider(context.Background(), id, providerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganisationApi.DeleteProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**providerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteProviderRequest struct via the builder pattern


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


## DeleteProviderFeature

> DeleteProviderFeature(ctx, id, providerId, featureId).Execute()



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
    providerId := "providerId_example" // string | 
    featureId := "featureId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrganisationApi.DeleteProviderFeature(context.Background(), id, providerId, featureId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganisationApi.DeleteProviderFeature``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**providerId** | **string** |  | 
**featureId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteProviderFeatureRequest struct via the builder pattern


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


## DeleteProviderPlan

> DeleteProviderPlan(ctx, id, providerId, planId).Execute()



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
    providerId := "providerId_example" // string | 
    planId := "planId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrganisationApi.DeleteProviderPlan(context.Background(), id, providerId, planId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganisationApi.DeleteProviderPlan``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**providerId** | **string** |  | 
**planId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteProviderPlanRequest struct via the builder pattern


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


## DeleteProviderPlanFeature

> DeleteProviderPlanFeature(ctx, id, providerId, planId, featureName).Execute()



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
    providerId := "providerId_example" // string | 
    planId := "planId_example" // string | 
    featureName := "featureName_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrganisationApi.DeleteProviderPlanFeature(context.Background(), id, providerId, planId, featureName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganisationApi.DeleteProviderPlanFeature``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**providerId** | **string** |  | 
**planId** | **string** |  | 
**featureName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteProviderPlanFeatureRequest struct via the builder pattern


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


## DeletePurchaseOrderByOrga

> DeletePurchaseOrderByOrga(ctx, id, bid).Execute()



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
    bid := "bid_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrganisationApi.DeletePurchaseOrderByOrga(context.Background(), id, bid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganisationApi.DeletePurchaseOrderByOrga``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**bid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePurchaseOrderByOrgaRequest struct via the builder pattern


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


## DeleteRecurrentPaymentByOrga

> DeleteRecurrentPaymentByOrga(ctx, id).Execute()



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
    resp, r, err := api_client.OrganisationApi.DeleteRecurrentPaymentByOrga(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganisationApi.DeleteRecurrentPaymentByOrga``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRecurrentPaymentByOrgaRequest struct via the builder pattern


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


## DeprovisionAddonByOrgaAndAddonId

> Message DeprovisionAddonByOrgaAndAddonId(ctx, id, addonId).Execute()



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
    addonId := "addonId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrganisationApi.DeprovisionAddonByOrgaAndAddonId(context.Background(), id, addonId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganisationApi.DeprovisionAddonByOrgaAndAddonId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeprovisionAddonByOrgaAndAddonId`: Message
    fmt.Fprintf(os.Stdout, "Response from `OrganisationApi.DeprovisionAddonByOrgaAndAddonId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**addonId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeprovisionAddonByOrgaAndAddonIdRequest struct via the builder pattern


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


## EditApplicationByOrgaAndAppId

> ApplicationView EditApplicationByOrgaAndAppId(ctx, id, appId).WannabeApplication(wannabeApplication).Execute()



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
    wannabeApplication := *openapiclient.NewWannabeApplication() // WannabeApplication | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrganisationApi.EditApplicationByOrgaAndAppId(context.Background(), id, appId).WannabeApplication(wannabeApplication).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganisationApi.EditApplicationByOrgaAndAppId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditApplicationByOrgaAndAppId`: ApplicationView
    fmt.Fprintf(os.Stdout, "Response from `OrganisationApi.EditApplicationByOrgaAndAppId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**appId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditApplicationByOrgaAndAppIdRequest struct via the builder pattern


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


## EditApplicationEnvByOrgaAndAppIdAndEnvName

> ApplicationView EditApplicationEnvByOrgaAndAppIdAndEnvName(ctx, id, appId, envName).WannabeValue(wannabeValue).Execute()



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
    envName := "envName_example" // string | 
    wannabeValue := *openapiclient.NewWannabeValue() // WannabeValue | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrganisationApi.EditApplicationEnvByOrgaAndAppIdAndEnvName(context.Background(), id, appId, envName).WannabeValue(wannabeValue).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganisationApi.EditApplicationEnvByOrgaAndAppIdAndEnvName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditApplicationEnvByOrgaAndAppIdAndEnvName`: ApplicationView
    fmt.Fprintf(os.Stdout, "Response from `OrganisationApi.EditApplicationEnvByOrgaAndAppIdAndEnvName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**appId** | **string** |  | 
**envName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditApplicationEnvByOrgaAndAppIdAndEnvNameRequest struct via the builder pattern


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


## EditApplicationEnvironmentByOrgaAndAppId

> ApplicationView EditApplicationEnvironmentByOrgaAndAppId(ctx, id, appId).Body(body).Execute()



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
    body := "body_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrganisationApi.EditApplicationEnvironmentByOrgaAndAppId(context.Background(), id, appId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganisationApi.EditApplicationEnvironmentByOrgaAndAppId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditApplicationEnvironmentByOrgaAndAppId`: ApplicationView
    fmt.Fprintf(os.Stdout, "Response from `OrganisationApi.EditApplicationEnvironmentByOrgaAndAppId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**appId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditApplicationEnvironmentByOrgaAndAppIdRequest struct via the builder pattern


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


## EditOrganisation

> OrganisationView EditOrganisation(ctx, id).WannabeOrganisation(wannabeOrganisation).Execute()



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
    wannabeOrganisation := *openapiclient.NewWannabeOrganisation() // WannabeOrganisation | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrganisationApi.EditOrganisation(context.Background(), id).WannabeOrganisation(wannabeOrganisation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganisationApi.EditOrganisation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditOrganisation`: OrganisationView
    fmt.Fprintf(os.Stdout, "Response from `OrganisationApi.EditOrganisation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditOrganisationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **wannabeOrganisation** | [**WannabeOrganisation**](WannabeOrganisation.md) |  | 

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

> Message EditOrganisationMember(ctx, id, userId).WannabeMember(wannabeMember).Execute()



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
    userId := "userId_example" // string | 
    wannabeMember := *openapiclient.NewWannabeMember() // WannabeMember | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrganisationApi.EditOrganisationMember(context.Background(), id, userId).WannabeMember(wannabeMember).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganisationApi.EditOrganisationMember``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditOrganisationMember`: Message
    fmt.Fprintf(os.Stdout, "Response from `OrganisationApi.EditOrganisationMember`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditOrganisationMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **wannabeMember** | [**WannabeMember**](WannabeMember.md) |  | 

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

> AddonPlanView EditProviderPlan(ctx, id, providerId, planId).WannabeAddonPlan(wannabeAddonPlan).Execute()



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
    providerId := "providerId_example" // string | 
    planId := "planId_example" // string | 
    wannabeAddonPlan := *openapiclient.NewWannabeAddonPlan() // WannabeAddonPlan | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrganisationApi.EditProviderPlan(context.Background(), id, providerId, planId).WannabeAddonPlan(wannabeAddonPlan).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganisationApi.EditProviderPlan``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditProviderPlan`: AddonPlanView
    fmt.Fprintf(os.Stdout, "Response from `OrganisationApi.EditProviderPlan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**providerId** | **string** |  | 
**planId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditProviderPlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **wannabeAddonPlan** | [**WannabeAddonPlan**](WannabeAddonPlan.md) |  | 

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

> AddonPlanView EditProviderPlanFeature(ctx, id, providerId, planId, featureName).AddonFeatureInstanceView(addonFeatureInstanceView).Execute()



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
    providerId := "providerId_example" // string | 
    planId := "planId_example" // string | 
    featureName := "featureName_example" // string | 
    addonFeatureInstanceView := *openapiclient.NewAddonFeatureInstanceView() // AddonFeatureInstanceView | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrganisationApi.EditProviderPlanFeature(context.Background(), id, providerId, planId, featureName).AddonFeatureInstanceView(addonFeatureInstanceView).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganisationApi.EditProviderPlanFeature``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditProviderPlanFeature`: AddonPlanView
    fmt.Fprintf(os.Stdout, "Response from `OrganisationApi.EditProviderPlanFeature`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**providerId** | **string** |  | 
**planId** | **string** |  | 
**featureName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditProviderPlanFeatureRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **addonFeatureInstanceView** | [**AddonFeatureInstanceView**](AddonFeatureInstanceView.md) |  | 

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

> AddonView GetAddonByOrgaAndAddonId(ctx, id, addonId).Execute()



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
    addonId := "addonId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrganisationApi.GetAddonByOrgaAndAddonId(context.Background(), id, addonId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganisationApi.GetAddonByOrgaAndAddonId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAddonByOrgaAndAddonId`: AddonView
    fmt.Fprintf(os.Stdout, "Response from `OrganisationApi.GetAddonByOrgaAndAddonId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**addonId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAddonByOrgaAndAddonIdRequest struct via the builder pattern


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


## GetAddonEnvByOrgaAndAddonId

> []AddonEnvironmentView GetAddonEnvByOrgaAndAddonId(ctx, id, addonId).Execute()



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
    addonId := "addonId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrganisationApi.GetAddonEnvByOrgaAndAddonId(context.Background(), id, addonId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganisationApi.GetAddonEnvByOrgaAndAddonId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAddonEnvByOrgaAndAddonId`: []AddonEnvironmentView
    fmt.Fprintf(os.Stdout, "Response from `OrganisationApi.GetAddonEnvByOrgaAndAddonId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**addonId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAddonEnvByOrgaAndAddonIdRequest struct via the builder pattern


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


## GetAddonInstance

> string GetAddonInstance(ctx, id, addonId, instanceId).Execute()



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
    addonId := "addonId_example" // string | 
    instanceId := "instanceId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrganisationApi.GetAddonInstance(context.Background(), id, addonId, instanceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganisationApi.GetAddonInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAddonInstance`: string
    fmt.Fprintf(os.Stdout, "Response from `OrganisationApi.GetAddonInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**addonId** | **string** |  | 
**instanceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAddonInstanceRequest struct via the builder pattern


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


## GetAddonInstances

> []SuperNovaInstanceView GetAddonInstances(ctx, id, addonId).DeploymentId(deploymentId).WithDeleted(withDeleted).Execute()



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
    addonId := "addonId_example" // string | 
    deploymentId := "deploymentId_example" // string |  (optional)
    withDeleted := "withDeleted_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrganisationApi.GetAddonInstances(context.Background(), id, addonId).DeploymentId(deploymentId).WithDeleted(withDeleted).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganisationApi.GetAddonInstances``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAddonInstances`: []SuperNovaInstanceView
    fmt.Fprintf(os.Stdout, "Response from `OrganisationApi.GetAddonInstances`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**addonId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAddonInstancesRequest struct via the builder pattern


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


## GetAddonMigration

> string GetAddonMigration(ctx, id, addonId, migrationId).Execute()



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
    addonId := "addonId_example" // string | 
    migrationId := "migrationId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrganisationApi.GetAddonMigration(context.Background(), id, addonId, migrationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganisationApi.GetAddonMigration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAddonMigration`: string
    fmt.Fprintf(os.Stdout, "Response from `OrganisationApi.GetAddonMigration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**addonId** | **string** |  | 
**migrationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAddonMigrationRequest struct via the builder pattern


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


## GetAddonMigrations

> string GetAddonMigrations(ctx, id, addonId).Execute()



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
    addonId := "addonId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrganisationApi.GetAddonMigrations(context.Background(), id, addonId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganisationApi.GetAddonMigrations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAddonMigrations`: string
    fmt.Fprintf(os.Stdout, "Response from `OrganisationApi.GetAddonMigrations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**addonId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAddonMigrationsRequest struct via the builder pattern


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


## GetAddonSSODataForOrga

> AddonProviderSSOData GetAddonSSODataForOrga(ctx, id, addonId).Execute()



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
    addonId := "addonId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrganisationApi.GetAddonSSODataForOrga(context.Background(), id, addonId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganisationApi.GetAddonSSODataForOrga``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAddonSSODataForOrga`: AddonProviderSSOData
    fmt.Fprintf(os.Stdout, "Response from `OrganisationApi.GetAddonSSODataForOrga`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**addonId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAddonSSODataForOrgaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**AddonProviderSSOData**](AddonProviderSSOData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAddonTagsByOrgaIdAndAddonId

> []string GetAddonTagsByOrgaIdAndAddonId(ctx, id, addonId).Execute()



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
    addonId := "addonId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrganisationApi.GetAddonTagsByOrgaIdAndAddonId(context.Background(), id, addonId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganisationApi.GetAddonTagsByOrgaIdAndAddonId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAddonTagsByOrgaIdAndAddonId`: []string
    fmt.Fprintf(os.Stdout, "Response from `OrganisationApi.GetAddonTagsByOrgaIdAndAddonId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**addonId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAddonTagsByOrgaIdAndAddonIdRequest struct via the builder pattern


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


## GetAddonsByOrgaId

> []AddonView GetAddonsByOrgaId(ctx, id).Execute()



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
    resp, r, err := api_client.OrganisationApi.GetAddonsByOrgaId(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganisationApi.GetAddonsByOrgaId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAddonsByOrgaId`: []AddonView
    fmt.Fprintf(os.Stdout, "Response from `OrganisationApi.GetAddonsByOrgaId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAddonsByOrgaIdRequest struct via the builder pattern


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


## GetAddonsLinkedToApplicationByOrgaAndAppId

> []AddonView GetAddonsLinkedToApplicationByOrgaAndAppId(ctx, id, appId).Execute()



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
    resp, r, err := api_client.OrganisationApi.GetAddonsLinkedToApplicationByOrgaAndAppId(context.Background(), id, appId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganisationApi.GetAddonsLinkedToApplicationByOrgaAndAppId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAddonsLinkedToApplicationByOrgaAndAppId`: []AddonView
    fmt.Fprintf(os.Stdout, "Response from `OrganisationApi.GetAddonsLinkedToApplicationByOrgaAndAppId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**appId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAddonsLinkedToApplicationByOrgaAndAppIdRequest struct via the builder pattern


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


## GetAllApplicationsByOrga

> []ApplicationView GetAllApplicationsByOrga(ctx, id).InstanceId(instanceId).Execute()



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
    instanceId := "instanceId_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrganisationApi.GetAllApplicationsByOrga(context.Background(), id).InstanceId(instanceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganisationApi.GetAllApplicationsByOrga``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllApplicationsByOrga`: []ApplicationView
    fmt.Fprintf(os.Stdout, "Response from `OrganisationApi.GetAllApplicationsByOrga`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllApplicationsByOrgaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **instanceId** | **string** |  | 

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

> DropCountView GetAmountForOrga(ctx, id).Execute()



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
    resp, r, err := api_client.OrganisationApi.GetAmountForOrga(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganisationApi.GetAmountForOrga``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAmountForOrga`: DropCountView
    fmt.Fprintf(os.Stdout, "Response from `OrganisationApi.GetAmountForOrga`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAmountForOrgaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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

> []string GetApplicationBranchesByOrgaAndAppId(ctx, id, appId).Execute()



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
    resp, r, err := api_client.OrganisationApi.GetApplicationBranchesByOrgaAndAppId(context.Background(), id, appId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganisationApi.GetApplicationBranchesByOrgaAndAppId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApplicationBranchesByOrgaAndAppId`: []string
    fmt.Fprintf(os.Stdout, "Response from `OrganisationApi.GetApplicationBranchesByOrgaAndAppId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**appId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationBranchesByOrgaAndAppIdRequest struct via the builder pattern


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


## GetApplicationByOrgaAndAppId

> ApplicationView GetApplicationByOrgaAndAppId(ctx, id, appId).Execute()



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
    resp, r, err := api_client.OrganisationApi.GetApplicationByOrgaAndAppId(context.Background(), id, appId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganisationApi.GetApplicationByOrgaAndAppId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApplicationByOrgaAndAppId`: ApplicationView
    fmt.Fprintf(os.Stdout, "Response from `OrganisationApi.GetApplicationByOrgaAndAppId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**appId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationByOrgaAndAppIdRequest struct via the builder pattern


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


## GetApplicationDependenciesByOrgaAndAppId

> []ApplicationView GetApplicationDependenciesByOrgaAndAppId(ctx, id, appId).Execute()



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
    resp, r, err := api_client.OrganisationApi.GetApplicationDependenciesByOrgaAndAppId(context.Background(), id, appId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganisationApi.GetApplicationDependenciesByOrgaAndAppId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApplicationDependenciesByOrgaAndAppId`: []ApplicationView
    fmt.Fprintf(os.Stdout, "Response from `OrganisationApi.GetApplicationDependenciesByOrgaAndAppId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**appId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationDependenciesByOrgaAndAppIdRequest struct via the builder pattern


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


## GetApplicationDependenciesEnvByOrgaAndAppId

> GetApplicationDependenciesEnvByOrgaAndAppId(ctx, id, appId).Execute()



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
    resp, r, err := api_client.OrganisationApi.GetApplicationDependenciesEnvByOrgaAndAppId(context.Background(), id, appId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganisationApi.GetApplicationDependenciesEnvByOrgaAndAppId``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetApplicationDependenciesEnvByOrgaAndAppIdRequest struct via the builder pattern


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


## GetApplicationDependentsByOrgaAndAppId

> []ApplicationView GetApplicationDependentsByOrgaAndAppId(ctx, id, appId).Execute()



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
    resp, r, err := api_client.OrganisationApi.GetApplicationDependentsByOrgaAndAppId(context.Background(), id, appId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganisationApi.GetApplicationDependentsByOrgaAndAppId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApplicationDependentsByOrgaAndAppId`: []ApplicationView
    fmt.Fprintf(os.Stdout, "Response from `OrganisationApi.GetApplicationDependentsByOrgaAndAppId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**appId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationDependentsByOrgaAndAppIdRequest struct via the builder pattern


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


## GetApplicationDeploymentForOrga

> GetApplicationDeploymentForOrga(ctx, id, appId, deploymentId).Execute()



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
    deploymentId := "deploymentId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrganisationApi.GetApplicationDeploymentForOrga(context.Background(), id, appId, deploymentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganisationApi.GetApplicationDeploymentForOrga``: %v\n", err)
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
**deploymentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationDeploymentForOrgaRequest struct via the builder pattern


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


## GetApplicationDeploymentsForOrga

> []DeploymentView GetApplicationDeploymentsForOrga(ctx, id, appId).Limit(limit).Offset(offset).Action(action).Execute()



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
    limit := "limit_example" // string |  (optional)
    offset := "offset_example" // string |  (optional)
    action := "action_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrganisationApi.GetApplicationDeploymentsForOrga(context.Background(), id, appId).Limit(limit).Offset(offset).Action(action).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganisationApi.GetApplicationDeploymentsForOrga``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApplicationDeploymentsForOrga`: []DeploymentView
    fmt.Fprintf(os.Stdout, "Response from `OrganisationApi.GetApplicationDeploymentsForOrga`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**appId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationDeploymentsForOrgaRequest struct via the builder pattern


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


## GetApplicationEnvByOrgaAndAppId

> []AddonEnvironmentView GetApplicationEnvByOrgaAndAppId(ctx, id, appId).Execute()



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
    resp, r, err := api_client.OrganisationApi.GetApplicationEnvByOrgaAndAppId(context.Background(), id, appId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganisationApi.GetApplicationEnvByOrgaAndAppId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApplicationEnvByOrgaAndAppId`: []AddonEnvironmentView
    fmt.Fprintf(os.Stdout, "Response from `OrganisationApi.GetApplicationEnvByOrgaAndAppId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**appId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationEnvByOrgaAndAppIdRequest struct via the builder pattern


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


## GetApplicationInstanceByOrgaAndAppAndInstanceId

> string GetApplicationInstanceByOrgaAndAppAndInstanceId(ctx, id, appId, instanceId).Execute()



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
    instanceId := "instanceId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrganisationApi.GetApplicationInstanceByOrgaAndAppAndInstanceId(context.Background(), id, appId, instanceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganisationApi.GetApplicationInstanceByOrgaAndAppAndInstanceId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApplicationInstanceByOrgaAndAppAndInstanceId`: string
    fmt.Fprintf(os.Stdout, "Response from `OrganisationApi.GetApplicationInstanceByOrgaAndAppAndInstanceId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**appId** | **string** |  | 
**instanceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationInstanceByOrgaAndAppAndInstanceIdRequest struct via the builder pattern


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


## GetApplicationInstancesByOrgaAndAppId

> []SuperNovaInstanceView GetApplicationInstancesByOrgaAndAppId(ctx, id, appId).DeploymentId(deploymentId).WithDeleted(withDeleted).Execute()



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
    deploymentId := "deploymentId_example" // string |  (optional)
    withDeleted := "withDeleted_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrganisationApi.GetApplicationInstancesByOrgaAndAppId(context.Background(), id, appId).DeploymentId(deploymentId).WithDeleted(withDeleted).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganisationApi.GetApplicationInstancesByOrgaAndAppId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApplicationInstancesByOrgaAndAppId`: []SuperNovaInstanceView
    fmt.Fprintf(os.Stdout, "Response from `OrganisationApi.GetApplicationInstancesByOrgaAndAppId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**appId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationInstancesByOrgaAndAppIdRequest struct via the builder pattern


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


## GetApplicationTagsByOrgaAndAppId

> []string GetApplicationTagsByOrgaAndAppId(ctx, id, appId).Execute()



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
    resp, r, err := api_client.OrganisationApi.GetApplicationTagsByOrgaAndAppId(context.Background(), id, appId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganisationApi.GetApplicationTagsByOrgaAndAppId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApplicationTagsByOrgaAndAppId`: []string
    fmt.Fprintf(os.Stdout, "Response from `OrganisationApi.GetApplicationTagsByOrgaAndAppId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**appId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationTagsByOrgaAndAppIdRequest struct via the builder pattern


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


## GetApplicationsLinkedToAddonByOrgaAndAddonId

> []ApplicationView GetApplicationsLinkedToAddonByOrgaAndAddonId(ctx, id, addonId).Execute()



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
    addonId := "addonId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrganisationApi.GetApplicationsLinkedToAddonByOrgaAndAddonId(context.Background(), id, addonId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganisationApi.GetApplicationsLinkedToAddonByOrgaAndAddonId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApplicationsLinkedToAddonByOrgaAndAddonId`: []ApplicationView
    fmt.Fprintf(os.Stdout, "Response from `OrganisationApi.GetApplicationsLinkedToAddonByOrgaAndAddonId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**addonId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationsLinkedToAddonByOrgaAndAddonIdRequest struct via the builder pattern


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


## GetConsumerByOrga

> OAuth1ConsumerView GetConsumerByOrga(ctx, id, key).Execute()



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
    key := "key_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrganisationApi.GetConsumerByOrga(context.Background(), id, key).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganisationApi.GetConsumerByOrga``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetConsumerByOrga`: OAuth1ConsumerView
    fmt.Fprintf(os.Stdout, "Response from `OrganisationApi.GetConsumerByOrga`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**key** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConsumerByOrgaRequest struct via the builder pattern


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


## GetConsumerSecretByOrga

> SecretView GetConsumerSecretByOrga(ctx, id, key).Execute()



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
    key := "key_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrganisationApi.GetConsumerSecretByOrga(context.Background(), id, key).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganisationApi.GetConsumerSecretByOrga``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetConsumerSecretByOrga`: SecretView
    fmt.Fprintf(os.Stdout, "Response from `OrganisationApi.GetConsumerSecretByOrga`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**key** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConsumerSecretByOrgaRequest struct via the builder pattern


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


## GetConsumersByOrga

> []OAuth1ConsumerView GetConsumersByOrga(ctx, id).Execute()



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
    resp, r, err := api_client.OrganisationApi.GetConsumersByOrga(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganisationApi.GetConsumersByOrga``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetConsumersByOrga`: []OAuth1ConsumerView
    fmt.Fprintf(os.Stdout, "Response from `OrganisationApi.GetConsumersByOrga`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConsumersByOrgaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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

> string GetConsumptionsForOrga(ctx, id).AppId(appId).From(from).To(to).For_(for_).Execute()



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
    appId := "appId_example" // string |  (optional)
    from := "from_example" // string |  (optional)
    to := "to_example" // string |  (optional)
    for_ := "for__example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrganisationApi.GetConsumptionsForOrga(context.Background(), id).AppId(appId).From(from).To(to).For_(for_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganisationApi.GetConsumptionsForOrga``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetConsumptionsForOrga`: string
    fmt.Fprintf(os.Stdout, "Response from `OrganisationApi.GetConsumptionsForOrga`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConsumptionsForOrgaRequest struct via the builder pattern


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


## GetDefaultMethodByOrga

> PaymentMethodView GetDefaultMethodByOrga(ctx, id).Execute()



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
    resp, r, err := api_client.OrganisationApi.GetDefaultMethodByOrga(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganisationApi.GetDefaultMethodByOrga``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDefaultMethodByOrga`: PaymentMethodView
    fmt.Fprintf(os.Stdout, "Response from `OrganisationApi.GetDefaultMethodByOrga`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDefaultMethodByOrgaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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

> GetDeploymentsForAllApps(ctx, id).Limit(limit).Execute()



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
    limit := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrganisationApi.GetDeploymentsForAllApps(context.Background(), id).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganisationApi.GetDeploymentsForAllApps``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeploymentsForAllAppsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** |  | 

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

> []LinkedAddonEnvironmentView GetEnvOfAddonsLinkedToApplicationByOrgaAndAppId(ctx, id, appId).Execute()



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
    resp, r, err := api_client.OrganisationApi.GetEnvOfAddonsLinkedToApplicationByOrgaAndAppId(context.Background(), id, appId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganisationApi.GetEnvOfAddonsLinkedToApplicationByOrgaAndAppId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEnvOfAddonsLinkedToApplicationByOrgaAndAppId`: []LinkedAddonEnvironmentView
    fmt.Fprintf(os.Stdout, "Response from `OrganisationApi.GetEnvOfAddonsLinkedToApplicationByOrgaAndAppId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**appId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEnvOfAddonsLinkedToApplicationByOrgaAndAppIdRequest struct via the builder pattern


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


## GetExposedEnvByOrgaAndAppId

> string GetExposedEnvByOrgaAndAppId(ctx, id, appId).Execute()



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
    resp, r, err := api_client.OrganisationApi.GetExposedEnvByOrgaAndAppId(context.Background(), id, appId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganisationApi.GetExposedEnvByOrgaAndAppId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetExposedEnvByOrgaAndAppId`: string
    fmt.Fprintf(os.Stdout, "Response from `OrganisationApi.GetExposedEnvByOrgaAndAppId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**appId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetExposedEnvByOrgaAndAppIdRequest struct via the builder pattern


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


## GetFavouriteVhostByOrgaAndAppId

> VhostView GetFavouriteVhostByOrgaAndAppId(ctx, id, appId).Execute()



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
    resp, r, err := api_client.OrganisationApi.GetFavouriteVhostByOrgaAndAppId(context.Background(), id, appId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganisationApi.GetFavouriteVhostByOrgaAndAppId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFavouriteVhostByOrgaAndAppId`: VhostView
    fmt.Fprintf(os.Stdout, "Response from `OrganisationApi.GetFavouriteVhostByOrgaAndAppId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**appId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFavouriteVhostByOrgaAndAppIdRequest struct via the builder pattern


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


## GetInstancesForAllAppsForOrga

> string GetInstancesForAllAppsForOrga(ctx, id).Execute()



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
    resp, r, err := api_client.OrganisationApi.GetInstancesForAllAppsForOrga(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganisationApi.GetInstancesForAllAppsForOrga``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetInstancesForAllAppsForOrga`: string
    fmt.Fprintf(os.Stdout, "Response from `OrganisationApi.GetInstancesForAllAppsForOrga`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInstancesForAllAppsForOrgaRequest struct via the builder pattern


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


## GetInvoiceByOrga

> InvoiceRendering GetInvoiceByOrga(ctx, id, bid).Execute()



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
    bid := "bid_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrganisationApi.GetInvoiceByOrga(context.Background(), id, bid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganisationApi.GetInvoiceByOrga``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetInvoiceByOrga`: InvoiceRendering
    fmt.Fprintf(os.Stdout, "Response from `OrganisationApi.GetInvoiceByOrga`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**bid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInvoiceByOrgaRequest struct via the builder pattern


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


## GetInvoicesByOrga

> []InvoiceRendering GetInvoicesByOrga(ctx, id).Execute()



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
    resp, r, err := api_client.OrganisationApi.GetInvoicesByOrga(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganisationApi.GetInvoicesByOrga``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetInvoicesByOrga`: []InvoiceRendering
    fmt.Fprintf(os.Stdout, "Response from `OrganisationApi.GetInvoicesByOrga`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInvoicesByOrgaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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

> string GetMonthlyInvoiceByOrga(ctx, id).Execute()



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
    resp, r, err := api_client.OrganisationApi.GetMonthlyInvoiceByOrga(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganisationApi.GetMonthlyInvoiceByOrga``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMonthlyInvoiceByOrga`: string
    fmt.Fprintf(os.Stdout, "Response from `OrganisationApi.GetMonthlyInvoiceByOrga`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMonthlyInvoiceByOrgaRequest struct via the builder pattern


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


## GetNamespaces

> []NamespaceView GetNamespaces(ctx, id).Execute()



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
    resp, r, err := api_client.OrganisationApi.GetNamespaces(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganisationApi.GetNamespaces``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNamespaces`: []NamespaceView
    fmt.Fprintf(os.Stdout, "Response from `OrganisationApi.GetNamespaces`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNamespacesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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

> SetupIntentView GetNewSetupIntentByOrga(ctx, id).Type_(type_).Execute()



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
    type_ := "type__example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrganisationApi.GetNewSetupIntentByOrga(context.Background(), id).Type_(type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganisationApi.GetNewSetupIntentByOrga``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNewSetupIntentByOrga`: SetupIntentView
    fmt.Fprintf(os.Stdout, "Response from `OrganisationApi.GetNewSetupIntentByOrga`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNewSetupIntentByOrgaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | **string** |  | 

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

> OrganisationView GetOrganisation(ctx, id).Execute()



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
    resp, r, err := api_client.OrganisationApi.GetOrganisation(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganisationApi.GetOrganisation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganisation`: OrganisationView
    fmt.Fprintf(os.Stdout, "Response from `OrganisationApi.GetOrganisation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganisationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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

> []OrganisationMemberView GetOrganisationMembers(ctx, id).Execute()



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
    resp, r, err := api_client.OrganisationApi.GetOrganisationMembers(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganisationApi.GetOrganisationMembers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganisationMembers`: []OrganisationMemberView
    fmt.Fprintf(os.Stdout, "Response from `OrganisationApi.GetOrganisationMembers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganisationMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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

> PaymentInfoView GetPaymentInfoForOrga(ctx, id).Execute()



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
    resp, r, err := api_client.OrganisationApi.GetPaymentInfoForOrga(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganisationApi.GetPaymentInfoForOrga``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPaymentInfoForOrga`: PaymentInfoView
    fmt.Fprintf(os.Stdout, "Response from `OrganisationApi.GetPaymentInfoForOrga`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPaymentInfoForOrgaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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

> GetPdfInvoiceByOrga(ctx, id, bid).Token(token).Execute()



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
    bid := "bid_example" // string | 
    token := "token_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrganisationApi.GetPdfInvoiceByOrga(context.Background(), id, bid).Token(token).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganisationApi.GetPdfInvoiceByOrga``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**bid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPdfInvoiceByOrgaRequest struct via the builder pattern


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


## GetPriceWithTaxByOrga

> PriceWithTaxInfo GetPriceWithTaxByOrga(ctx, id, price).Execute()



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
    price := "price_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrganisationApi.GetPriceWithTaxByOrga(context.Background(), id, price).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganisationApi.GetPriceWithTaxByOrga``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPriceWithTaxByOrga`: PriceWithTaxInfo
    fmt.Fprintf(os.Stdout, "Response from `OrganisationApi.GetPriceWithTaxByOrga`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**price** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPriceWithTaxByOrgaRequest struct via the builder pattern


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


## GetProviderFeatures

> []AddonFeatureView GetProviderFeatures(ctx, id, providerId).Execute()



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
    providerId := "providerId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrganisationApi.GetProviderFeatures(context.Background(), id, providerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganisationApi.GetProviderFeatures``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProviderFeatures`: []AddonFeatureView
    fmt.Fprintf(os.Stdout, "Response from `OrganisationApi.GetProviderFeatures`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**providerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProviderFeaturesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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

> AddonProviderInfoView GetProviderInfo(ctx, id, providerId).Execute()



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
    providerId := "providerId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrganisationApi.GetProviderInfo(context.Background(), id, providerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganisationApi.GetProviderInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProviderInfo`: AddonProviderInfoView
    fmt.Fprintf(os.Stdout, "Response from `OrganisationApi.GetProviderInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**providerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProviderInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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

> AddonPlanView GetProviderPlan(ctx, id, providerId, planId).Execute()



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
    providerId := "providerId_example" // string | 
    planId := "planId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrganisationApi.GetProviderPlan(context.Background(), id, providerId, planId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganisationApi.GetProviderPlan``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProviderPlan`: AddonPlanView
    fmt.Fprintf(os.Stdout, "Response from `OrganisationApi.GetProviderPlan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**providerId** | **string** |  | 
**planId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProviderPlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




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

> []AddonPlanView GetProviderPlans(ctx, id, providerId).Execute()



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
    providerId := "providerId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrganisationApi.GetProviderPlans(context.Background(), id, providerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganisationApi.GetProviderPlans``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProviderPlans`: []AddonPlanView
    fmt.Fprintf(os.Stdout, "Response from `OrganisationApi.GetProviderPlans`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**providerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProviderPlansRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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

> []string GetProviderTags(ctx, id, providerId).Execute()



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
    providerId := "providerId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrganisationApi.GetProviderTags(context.Background(), id, providerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganisationApi.GetProviderTags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProviderTags`: []string
    fmt.Fprintf(os.Stdout, "Response from `OrganisationApi.GetProviderTags`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**providerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProviderTagsRequest struct via the builder pattern


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


## GetProvidersInfo

> []AddonProviderInfoFullView GetProvidersInfo(ctx, id).Execute()



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
    resp, r, err := api_client.OrganisationApi.GetProvidersInfo(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganisationApi.GetProvidersInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProvidersInfo`: []AddonProviderInfoFullView
    fmt.Fprintf(os.Stdout, "Response from `OrganisationApi.GetProvidersInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProvidersInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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

> RecurrentPaymentView GetRecurrentPaymentByOrga(ctx, id).Execute()



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
    resp, r, err := api_client.OrganisationApi.GetRecurrentPaymentByOrga(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganisationApi.GetRecurrentPaymentByOrga``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRecurrentPaymentByOrga`: RecurrentPaymentView
    fmt.Fprintf(os.Stdout, "Response from `OrganisationApi.GetRecurrentPaymentByOrga`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRecurrentPaymentByOrgaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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

> AddonProviderSSOData GetSSODataForOrga(ctx, id, providerId).Execute()



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
    providerId := "providerId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrganisationApi.GetSSODataForOrga(context.Background(), id, providerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganisationApi.GetSSODataForOrga``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSSODataForOrga`: AddonProviderSSOData
    fmt.Fprintf(os.Stdout, "Response from `OrganisationApi.GetSSODataForOrga`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**providerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSSODataForOrgaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**AddonProviderSSOData**](AddonProviderSSOData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStripeTokenByOrga

> BraintreeToken GetStripeTokenByOrga(ctx, id).Execute()



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
    resp, r, err := api_client.OrganisationApi.GetStripeTokenByOrga(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganisationApi.GetStripeTokenByOrga``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStripeTokenByOrga`: BraintreeToken
    fmt.Fprintf(os.Stdout, "Response from `OrganisationApi.GetStripeTokenByOrga`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStripeTokenByOrgaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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

> []TcpRedirView GetTcpRedirs(ctx, id, appId).Execute()



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
    resp, r, err := api_client.OrganisationApi.GetTcpRedirs(context.Background(), id, appId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganisationApi.GetTcpRedirs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTcpRedirs`: []TcpRedirView
    fmt.Fprintf(os.Stdout, "Response from `OrganisationApi.GetTcpRedirs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**appId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTcpRedirsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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

> []PaymentMethodView GetUnpaidInvoicesByOrga(ctx, id).Execute()



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
    resp, r, err := api_client.OrganisationApi.GetUnpaidInvoicesByOrga(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganisationApi.GetUnpaidInvoicesByOrga``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUnpaidInvoicesByOrga`: []PaymentMethodView
    fmt.Fprintf(os.Stdout, "Response from `OrganisationApi.GetUnpaidInvoicesByOrga`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUnpaidInvoicesByOrgaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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

> []InvoiceRendering GetUnpaidInvoicesByOrga1(ctx, id).Execute()



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
    resp, r, err := api_client.OrganisationApi.GetUnpaidInvoicesByOrga1(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganisationApi.GetUnpaidInvoicesByOrga1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUnpaidInvoicesByOrga1`: []InvoiceRendering
    fmt.Fprintf(os.Stdout, "Response from `OrganisationApi.GetUnpaidInvoicesByOrga1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUnpaidInvoicesByOrga1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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

> []OrganisationView GetUserOrganisationss(ctx).User(user).Execute()



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
    user := "user_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrganisationApi.GetUserOrganisationss(context.Background()).User(user).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganisationApi.GetUserOrganisationss``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUserOrganisationss`: []OrganisationView
    fmt.Fprintf(os.Stdout, "Response from `OrganisationApi.GetUserOrganisationss`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUserOrganisationssRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **user** | **string** |  | 

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

> []VhostView GetVhostsByOrgaAndAppId(ctx, id, appId).Execute()



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
    resp, r, err := api_client.OrganisationApi.GetVhostsByOrgaAndAppId(context.Background(), id, appId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganisationApi.GetVhostsByOrgaAndAppId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVhostsByOrgaAndAppId`: []VhostView
    fmt.Fprintf(os.Stdout, "Response from `OrganisationApi.GetVhostsByOrgaAndAppId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**appId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVhostsByOrgaAndAppIdRequest struct via the builder pattern


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


## LinkAddonToApplicationByOrgaAndAppId

> LinkAddonToApplicationByOrgaAndAppId(ctx, id, appId).Body(body).Execute()



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
    body := "body_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrganisationApi.LinkAddonToApplicationByOrgaAndAppId(context.Background(), id, appId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganisationApi.LinkAddonToApplicationByOrgaAndAppId``: %v\n", err)
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

Other parameters are passed through a pointer to a apiLinkAddonToApplicationByOrgaAndAppIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## MarkFavouriteVhostByOrgaAndAppId

> VhostView MarkFavouriteVhostByOrgaAndAppId(ctx, id, appId).VhostView(vhostView).Execute()



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
    vhostView := *openapiclient.NewVhostView() // VhostView | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrganisationApi.MarkFavouriteVhostByOrgaAndAppId(context.Background(), id, appId).VhostView(vhostView).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganisationApi.MarkFavouriteVhostByOrgaAndAppId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MarkFavouriteVhostByOrgaAndAppId`: VhostView
    fmt.Fprintf(os.Stdout, "Response from `OrganisationApi.MarkFavouriteVhostByOrgaAndAppId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**appId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMarkFavouriteVhostByOrgaAndAppIdRequest struct via the builder pattern


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


## PreorderAddonByOrgaId

> InvoiceRendering PreorderAddonByOrgaId(ctx, id).WannabeAddonProvision(wannabeAddonProvision).Execute()



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
    wannabeAddonProvision := *openapiclient.NewWannabeAddonProvision("ProviderId_example", "Plan_example", "Region_example") // WannabeAddonProvision | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrganisationApi.PreorderAddonByOrgaId(context.Background(), id).WannabeAddonProvision(wannabeAddonProvision).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganisationApi.PreorderAddonByOrgaId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PreorderAddonByOrgaId`: InvoiceRendering
    fmt.Fprintf(os.Stdout, "Response from `OrganisationApi.PreorderAddonByOrgaId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPreorderAddonByOrgaIdRequest struct via the builder pattern


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


## PreorderMigration

> string PreorderMigration(ctx, id, addonId).PlanId(planId).Execute()



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
    addonId := "addonId_example" // string | 
    planId := "planId_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrganisationApi.PreorderMigration(context.Background(), id, addonId).PlanId(planId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganisationApi.PreorderMigration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PreorderMigration`: string
    fmt.Fprintf(os.Stdout, "Response from `OrganisationApi.PreorderMigration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**addonId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPreorderMigrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **planId** | **string** |  | 

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

> AddonView ProvisionAddonByOrgaId(ctx, id).WannabeAddonProvision(wannabeAddonProvision).Execute()



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
    wannabeAddonProvision := *openapiclient.NewWannabeAddonProvision("ProviderId_example", "Plan_example", "Region_example") // WannabeAddonProvision | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrganisationApi.ProvisionAddonByOrgaId(context.Background(), id).WannabeAddonProvision(wannabeAddonProvision).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganisationApi.ProvisionAddonByOrgaId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProvisionAddonByOrgaId`: AddonView
    fmt.Fprintf(os.Stdout, "Response from `OrganisationApi.ProvisionAddonByOrgaId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProvisionAddonByOrgaIdRequest struct via the builder pattern


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


## RedeployApplicationByOrgaAndAppId

> RedeployApplicationByOrgaAndAppId(ctx, id, appId).Commit(commit).UseCache(useCache).Execute()



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
    commit := "commit_example" // string |  (optional)
    useCache := "useCache_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrganisationApi.RedeployApplicationByOrgaAndAppId(context.Background(), id, appId).Commit(commit).UseCache(useCache).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganisationApi.RedeployApplicationByOrgaAndAppId``: %v\n", err)
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

Other parameters are passed through a pointer to a apiRedeployApplicationByOrgaAndAppIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **commit** | **string** |  | 
 **useCache** | **string** |  | 

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

> ApplicationView RemoveApplicationEnvByOrgaAndAppIdAndEnvName(ctx, id, appId, envName).Execute()



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
    envName := "envName_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrganisationApi.RemoveApplicationEnvByOrgaAndAppIdAndEnvName(context.Background(), id, appId, envName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganisationApi.RemoveApplicationEnvByOrgaAndAppIdAndEnvName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemoveApplicationEnvByOrgaAndAppIdAndEnvName`: ApplicationView
    fmt.Fprintf(os.Stdout, "Response from `OrganisationApi.RemoveApplicationEnvByOrgaAndAppIdAndEnvName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**appId** | **string** |  | 
**envName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveApplicationEnvByOrgaAndAppIdAndEnvNameRequest struct via the builder pattern


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


## RemoveOrganisationMember

> Message RemoveOrganisationMember(ctx, id, userId).Execute()



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
    userId := "userId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrganisationApi.RemoveOrganisationMember(context.Background(), id, userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganisationApi.RemoveOrganisationMember``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemoveOrganisationMember`: Message
    fmt.Fprintf(os.Stdout, "Response from `OrganisationApi.RemoveOrganisationMember`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveOrganisationMemberRequest struct via the builder pattern


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


## RemoveTcpRedir

> RemoveTcpRedir(ctx, id, appId, sourcePort).Namespace(namespace).Execute()



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
    sourcePort := int64(789) // int64 | 
    namespace := "namespace_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrganisationApi.RemoveTcpRedir(context.Background(), id, appId, sourcePort).Namespace(namespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganisationApi.RemoveTcpRedir``: %v\n", err)
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
**sourcePort** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveTcpRedirRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **namespace** | **string** |  | 

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

> Message RemoveVhostsByOrgaAndAppId(ctx, id, appId, domain).Execute()



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
    domain := "domain_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrganisationApi.RemoveVhostsByOrgaAndAppId(context.Background(), id, appId, domain).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganisationApi.RemoveVhostsByOrgaAndAppId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemoveVhostsByOrgaAndAppId`: Message
    fmt.Fprintf(os.Stdout, "Response from `OrganisationApi.RemoveVhostsByOrgaAndAppId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**appId** | **string** |  | 
**domain** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveVhostsByOrgaAndAppIdRequest struct via the builder pattern


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


## ReplaceAddonTags

> []string ReplaceAddonTags(ctx, id, addonId).Body(body).Execute()



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
    addonId := "addonId_example" // string | 
    body := "body_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrganisationApi.ReplaceAddonTags(context.Background(), id, addonId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganisationApi.ReplaceAddonTags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplaceAddonTags`: []string
    fmt.Fprintf(os.Stdout, "Response from `OrganisationApi.ReplaceAddonTags`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**addonId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceAddonTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **string** |  | 

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

> []string ReplaceApplicationTags(ctx, id, appId).Body(body).Execute()



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
    body := "body_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrganisationApi.ReplaceApplicationTags(context.Background(), id, appId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganisationApi.ReplaceApplicationTags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplaceApplicationTags`: []string
    fmt.Fprintf(os.Stdout, "Response from `OrganisationApi.ReplaceApplicationTags`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**appId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceApplicationTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **string** |  | 

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

> SetApplicationBranchByOrgaAndAppId(ctx, id, appId).WannabeBranch(wannabeBranch).Execute()



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
    wannabeBranch := *openapiclient.NewWannabeBranch() // WannabeBranch | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrganisationApi.SetApplicationBranchByOrgaAndAppId(context.Background(), id, appId).WannabeBranch(wannabeBranch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganisationApi.SetApplicationBranchByOrgaAndAppId``: %v\n", err)
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

Other parameters are passed through a pointer to a apiSetApplicationBranchByOrgaAndAppIdRequest struct via the builder pattern


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


## SetBuildInstanceFlavorByOrgaAndAppId

> SetBuildInstanceFlavorByOrgaAndAppId(ctx, id, appId).WannabeBuildFlavor(wannabeBuildFlavor).Execute()



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
    wannabeBuildFlavor := *openapiclient.NewWannabeBuildFlavor() // WannabeBuildFlavor | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrganisationApi.SetBuildInstanceFlavorByOrgaAndAppId(context.Background(), id, appId).WannabeBuildFlavor(wannabeBuildFlavor).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganisationApi.SetBuildInstanceFlavorByOrgaAndAppId``: %v\n", err)
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

Other parameters are passed through a pointer to a apiSetBuildInstanceFlavorByOrgaAndAppIdRequest struct via the builder pattern


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


## SetDefaultMethodByOrga

> SetDefaultMethodByOrga(ctx, id).PaymentData(paymentData).Execute()



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
    paymentData := *openapiclient.NewPaymentData() // PaymentData | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrganisationApi.SetDefaultMethodByOrga(context.Background(), id).PaymentData(paymentData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganisationApi.SetDefaultMethodByOrga``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetDefaultMethodByOrgaRequest struct via the builder pattern


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


## SetMaxCreditsPerMonthByOrga

> string SetMaxCreditsPerMonthByOrga(ctx, id).WannabeMaxCredits(wannabeMaxCredits).Execute()



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
    wannabeMaxCredits := *openapiclient.NewWannabeMaxCredits() // WannabeMaxCredits | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrganisationApi.SetMaxCreditsPerMonthByOrga(context.Background(), id).WannabeMaxCredits(wannabeMaxCredits).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganisationApi.SetMaxCreditsPerMonthByOrga``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetMaxCreditsPerMonthByOrga`: string
    fmt.Fprintf(os.Stdout, "Response from `OrganisationApi.SetMaxCreditsPerMonthByOrga`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetMaxCreditsPerMonthByOrgaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **wannabeMaxCredits** | [**WannabeMaxCredits**](WannabeMaxCredits.md) |  | 

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

> UrlView SetOrgaAvatar(ctx, id).Execute()



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
    resp, r, err := api_client.OrganisationApi.SetOrgaAvatar(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganisationApi.SetOrgaAvatar``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetOrgaAvatar`: UrlView
    fmt.Fprintf(os.Stdout, "Response from `OrganisationApi.SetOrgaAvatar`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetOrgaAvatarRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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

> Message UndeployApplicationByOrgaAndAppId(ctx, id, appId).Execute()



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
    resp, r, err := api_client.OrganisationApi.UndeployApplicationByOrgaAndAppId(context.Background(), id, appId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganisationApi.UndeployApplicationByOrgaAndAppId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UndeployApplicationByOrgaAndAppId`: Message
    fmt.Fprintf(os.Stdout, "Response from `OrganisationApi.UndeployApplicationByOrgaAndAppId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**appId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUndeployApplicationByOrgaAndAppIdRequest struct via the builder pattern


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


## UnlinkAddonFromApplicationByOrgaAndAppAnddAddonId

> UnlinkAddonFromApplicationByOrgaAndAppAnddAddonId(ctx, id, appId, addonId).Execute()



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
    addonId := "addonId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrganisationApi.UnlinkAddonFromApplicationByOrgaAndAppAnddAddonId(context.Background(), id, appId, addonId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganisationApi.UnlinkAddonFromApplicationByOrgaAndAppAnddAddonId``: %v\n", err)
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
**addonId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnlinkAddonFromApplicationByOrgaAndAppAnddAddonIdRequest struct via the builder pattern


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


## UnmarkFavouriteVhostByOrgaAndAppId

> UnmarkFavouriteVhostByOrgaAndAppId(ctx, id, appId).Execute()



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
    resp, r, err := api_client.OrganisationApi.UnmarkFavouriteVhostByOrgaAndAppId(context.Background(), id, appId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganisationApi.UnmarkFavouriteVhostByOrgaAndAppId``: %v\n", err)
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

Other parameters are passed through a pointer to a apiUnmarkFavouriteVhostByOrgaAndAppIdRequest struct via the builder pattern


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


## UpdateAddonInfo

> AddonView UpdateAddonInfo(ctx, id, addonId).WannabeAddonProvision(wannabeAddonProvision).Execute()



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
    addonId := "addonId_example" // string | 
    wannabeAddonProvision := *openapiclient.NewWannabeAddonProvision("ProviderId_example", "Plan_example", "Region_example") // WannabeAddonProvision | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrganisationApi.UpdateAddonInfo(context.Background(), id, addonId).WannabeAddonProvision(wannabeAddonProvision).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganisationApi.UpdateAddonInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAddonInfo`: AddonView
    fmt.Fprintf(os.Stdout, "Response from `OrganisationApi.UpdateAddonInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**addonId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAddonInfoRequest struct via the builder pattern


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


## UpdateConsumerByOrga

> OAuth1ConsumerView UpdateConsumerByOrga(ctx, id, key).WannabeOAuth1Consumer(wannabeOAuth1Consumer).Execute()



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
    key := "key_example" // string | 
    wannabeOAuth1Consumer := *openapiclient.NewWannabeOAuth1Consumer() // WannabeOAuth1Consumer | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrganisationApi.UpdateConsumerByOrga(context.Background(), id, key).WannabeOAuth1Consumer(wannabeOAuth1Consumer).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganisationApi.UpdateConsumerByOrga``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateConsumerByOrga`: OAuth1ConsumerView
    fmt.Fprintf(os.Stdout, "Response from `OrganisationApi.UpdateConsumerByOrga`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**key** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateConsumerByOrgaRequest struct via the builder pattern


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


## UpdateExposedEnvByOrgaAndAppId

> Message UpdateExposedEnvByOrgaAndAppId(ctx, id, appId).Body(body).Execute()



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
    body := "body_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrganisationApi.UpdateExposedEnvByOrgaAndAppId(context.Background(), id, appId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganisationApi.UpdateExposedEnvByOrgaAndAppId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateExposedEnvByOrgaAndAppId`: Message
    fmt.Fprintf(os.Stdout, "Response from `OrganisationApi.UpdateExposedEnvByOrgaAndAppId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**appId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateExposedEnvByOrgaAndAppIdRequest struct via the builder pattern


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


## UpdateProviderInfos

> AddonProviderInfoView UpdateProviderInfos(ctx, id, providerId).WannabeAddonProviderInfos(wannabeAddonProviderInfos).Execute()



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
    providerId := "providerId_example" // string | 
    wannabeAddonProviderInfos := *openapiclient.NewWannabeAddonProviderInfos() // WannabeAddonProviderInfos | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrganisationApi.UpdateProviderInfos(context.Background(), id, providerId).WannabeAddonProviderInfos(wannabeAddonProviderInfos).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganisationApi.UpdateProviderInfos``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateProviderInfos`: AddonProviderInfoView
    fmt.Fprintf(os.Stdout, "Response from `OrganisationApi.UpdateProviderInfos`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**providerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateProviderInfosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **wannabeAddonProviderInfos** | [**WannabeAddonProviderInfos**](WannabeAddonProviderInfos.md) |  | 

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

