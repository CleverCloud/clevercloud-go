# OrganisationSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Avatar** | Pointer to **string** |  | [optional] 
**Applications** | Pointer to [**[]ApplicationSummary**](ApplicationSummary.md) |  | [optional] 
**Addons** | Pointer to [**[]AddonSummary**](AddonSummary.md) |  | [optional] 
**Consumers** | Pointer to [**[]OAuth1ConsumerSummary**](OAuth1ConsumerSummary.md) |  | [optional] 
**Role** | Pointer to **string** |  | [optional] 
**Providers** | Pointer to [**[]ProviderSummary**](ProviderSummary.md) |  | [optional] 
**VatState** | Pointer to **string** |  | [optional] 
**CanPay** | Pointer to **bool** |  | [optional] 
**CanSEPA** | Pointer to **bool** |  | [optional] 
**CleverEnterprise** | Pointer to **bool** |  | [optional] 
**EmergencyNumber** | Pointer to **string** |  | [optional] 

## Methods

### NewOrganisationSummary

`func NewOrganisationSummary() *OrganisationSummary`

NewOrganisationSummary instantiates a new OrganisationSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganisationSummaryWithDefaults

`func NewOrganisationSummaryWithDefaults() *OrganisationSummary`

NewOrganisationSummaryWithDefaults instantiates a new OrganisationSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OrganisationSummary) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrganisationSummary) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrganisationSummary) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OrganisationSummary) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *OrganisationSummary) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OrganisationSummary) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OrganisationSummary) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OrganisationSummary) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAvatar

`func (o *OrganisationSummary) GetAvatar() string`

GetAvatar returns the Avatar field if non-nil, zero value otherwise.

### GetAvatarOk

`func (o *OrganisationSummary) GetAvatarOk() (*string, bool)`

GetAvatarOk returns a tuple with the Avatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatar

`func (o *OrganisationSummary) SetAvatar(v string)`

SetAvatar sets Avatar field to given value.

### HasAvatar

`func (o *OrganisationSummary) HasAvatar() bool`

HasAvatar returns a boolean if a field has been set.

### GetApplications

`func (o *OrganisationSummary) GetApplications() []ApplicationSummary`

GetApplications returns the Applications field if non-nil, zero value otherwise.

### GetApplicationsOk

`func (o *OrganisationSummary) GetApplicationsOk() (*[]ApplicationSummary, bool)`

GetApplicationsOk returns a tuple with the Applications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplications

`func (o *OrganisationSummary) SetApplications(v []ApplicationSummary)`

SetApplications sets Applications field to given value.

### HasApplications

`func (o *OrganisationSummary) HasApplications() bool`

HasApplications returns a boolean if a field has been set.

### GetAddons

`func (o *OrganisationSummary) GetAddons() []AddonSummary`

GetAddons returns the Addons field if non-nil, zero value otherwise.

### GetAddonsOk

`func (o *OrganisationSummary) GetAddonsOk() (*[]AddonSummary, bool)`

GetAddonsOk returns a tuple with the Addons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddons

`func (o *OrganisationSummary) SetAddons(v []AddonSummary)`

SetAddons sets Addons field to given value.

### HasAddons

`func (o *OrganisationSummary) HasAddons() bool`

HasAddons returns a boolean if a field has been set.

### GetConsumers

`func (o *OrganisationSummary) GetConsumers() []OAuth1ConsumerSummary`

GetConsumers returns the Consumers field if non-nil, zero value otherwise.

### GetConsumersOk

`func (o *OrganisationSummary) GetConsumersOk() (*[]OAuth1ConsumerSummary, bool)`

GetConsumersOk returns a tuple with the Consumers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumers

`func (o *OrganisationSummary) SetConsumers(v []OAuth1ConsumerSummary)`

SetConsumers sets Consumers field to given value.

### HasConsumers

`func (o *OrganisationSummary) HasConsumers() bool`

HasConsumers returns a boolean if a field has been set.

### GetRole

`func (o *OrganisationSummary) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *OrganisationSummary) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *OrganisationSummary) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *OrganisationSummary) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetProviders

`func (o *OrganisationSummary) GetProviders() []ProviderSummary`

GetProviders returns the Providers field if non-nil, zero value otherwise.

### GetProvidersOk

`func (o *OrganisationSummary) GetProvidersOk() (*[]ProviderSummary, bool)`

GetProvidersOk returns a tuple with the Providers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviders

`func (o *OrganisationSummary) SetProviders(v []ProviderSummary)`

SetProviders sets Providers field to given value.

### HasProviders

`func (o *OrganisationSummary) HasProviders() bool`

HasProviders returns a boolean if a field has been set.

### GetVatState

`func (o *OrganisationSummary) GetVatState() string`

GetVatState returns the VatState field if non-nil, zero value otherwise.

### GetVatStateOk

`func (o *OrganisationSummary) GetVatStateOk() (*string, bool)`

GetVatStateOk returns a tuple with the VatState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVatState

`func (o *OrganisationSummary) SetVatState(v string)`

SetVatState sets VatState field to given value.

### HasVatState

`func (o *OrganisationSummary) HasVatState() bool`

HasVatState returns a boolean if a field has been set.

### GetCanPay

`func (o *OrganisationSummary) GetCanPay() bool`

GetCanPay returns the CanPay field if non-nil, zero value otherwise.

### GetCanPayOk

`func (o *OrganisationSummary) GetCanPayOk() (*bool, bool)`

GetCanPayOk returns a tuple with the CanPay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanPay

`func (o *OrganisationSummary) SetCanPay(v bool)`

SetCanPay sets CanPay field to given value.

### HasCanPay

`func (o *OrganisationSummary) HasCanPay() bool`

HasCanPay returns a boolean if a field has been set.

### GetCanSEPA

`func (o *OrganisationSummary) GetCanSEPA() bool`

GetCanSEPA returns the CanSEPA field if non-nil, zero value otherwise.

### GetCanSEPAOk

`func (o *OrganisationSummary) GetCanSEPAOk() (*bool, bool)`

GetCanSEPAOk returns a tuple with the CanSEPA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanSEPA

`func (o *OrganisationSummary) SetCanSEPA(v bool)`

SetCanSEPA sets CanSEPA field to given value.

### HasCanSEPA

`func (o *OrganisationSummary) HasCanSEPA() bool`

HasCanSEPA returns a boolean if a field has been set.

### GetCleverEnterprise

`func (o *OrganisationSummary) GetCleverEnterprise() bool`

GetCleverEnterprise returns the CleverEnterprise field if non-nil, zero value otherwise.

### GetCleverEnterpriseOk

`func (o *OrganisationSummary) GetCleverEnterpriseOk() (*bool, bool)`

GetCleverEnterpriseOk returns a tuple with the CleverEnterprise field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCleverEnterprise

`func (o *OrganisationSummary) SetCleverEnterprise(v bool)`

SetCleverEnterprise sets CleverEnterprise field to given value.

### HasCleverEnterprise

`func (o *OrganisationSummary) HasCleverEnterprise() bool`

HasCleverEnterprise returns a boolean if a field has been set.

### GetEmergencyNumber

`func (o *OrganisationSummary) GetEmergencyNumber() string`

GetEmergencyNumber returns the EmergencyNumber field if non-nil, zero value otherwise.

### GetEmergencyNumberOk

`func (o *OrganisationSummary) GetEmergencyNumberOk() (*string, bool)`

GetEmergencyNumberOk returns a tuple with the EmergencyNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmergencyNumber

`func (o *OrganisationSummary) SetEmergencyNumber(v string)`

SetEmergencyNumber sets EmergencyNumber field to given value.

### HasEmergencyNumber

`func (o *OrganisationSummary) HasEmergencyNumber() bool`

HasEmergencyNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


