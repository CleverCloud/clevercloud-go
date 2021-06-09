# OrganisationView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**BillingEmail** | Pointer to **string** |  | [optional] 
**Address** | Pointer to **string** |  | [optional] 
**City** | Pointer to **string** |  | [optional] 
**Zipcode** | Pointer to **string** |  | [optional] 
**Country** | Pointer to **string** |  | [optional] 
**Company** | Pointer to **string** |  | [optional] 
**VAT** | Pointer to **string** |  | [optional] 
**Avatar** | Pointer to **string** |  | [optional] 
**VatState** | Pointer to **string** |  | [optional] 
**CustomerFullName** | Pointer to **string** |  | [optional] 
**CanPay** | Pointer to **bool** |  | [optional] 
**CleverEnterprise** | Pointer to **bool** |  | [optional] 
**EmergencyNumber** | Pointer to **string** |  | [optional] 
**CanSEPA** | Pointer to **bool** |  | [optional] 

## Methods

### NewOrganisationView

`func NewOrganisationView() *OrganisationView`

NewOrganisationView instantiates a new OrganisationView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganisationViewWithDefaults

`func NewOrganisationViewWithDefaults() *OrganisationView`

NewOrganisationViewWithDefaults instantiates a new OrganisationView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OrganisationView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrganisationView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrganisationView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OrganisationView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *OrganisationView) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OrganisationView) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OrganisationView) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OrganisationView) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *OrganisationView) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OrganisationView) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OrganisationView) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *OrganisationView) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetBillingEmail

`func (o *OrganisationView) GetBillingEmail() string`

GetBillingEmail returns the BillingEmail field if non-nil, zero value otherwise.

### GetBillingEmailOk

`func (o *OrganisationView) GetBillingEmailOk() (*string, bool)`

GetBillingEmailOk returns a tuple with the BillingEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingEmail

`func (o *OrganisationView) SetBillingEmail(v string)`

SetBillingEmail sets BillingEmail field to given value.

### HasBillingEmail

`func (o *OrganisationView) HasBillingEmail() bool`

HasBillingEmail returns a boolean if a field has been set.

### GetAddress

`func (o *OrganisationView) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *OrganisationView) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *OrganisationView) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *OrganisationView) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetCity

`func (o *OrganisationView) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *OrganisationView) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *OrganisationView) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *OrganisationView) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetZipcode

`func (o *OrganisationView) GetZipcode() string`

GetZipcode returns the Zipcode field if non-nil, zero value otherwise.

### GetZipcodeOk

`func (o *OrganisationView) GetZipcodeOk() (*string, bool)`

GetZipcodeOk returns a tuple with the Zipcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZipcode

`func (o *OrganisationView) SetZipcode(v string)`

SetZipcode sets Zipcode field to given value.

### HasZipcode

`func (o *OrganisationView) HasZipcode() bool`

HasZipcode returns a boolean if a field has been set.

### GetCountry

`func (o *OrganisationView) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *OrganisationView) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *OrganisationView) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *OrganisationView) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetCompany

`func (o *OrganisationView) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *OrganisationView) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *OrganisationView) SetCompany(v string)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *OrganisationView) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetVAT

`func (o *OrganisationView) GetVAT() string`

GetVAT returns the VAT field if non-nil, zero value otherwise.

### GetVATOk

`func (o *OrganisationView) GetVATOk() (*string, bool)`

GetVATOk returns a tuple with the VAT field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVAT

`func (o *OrganisationView) SetVAT(v string)`

SetVAT sets VAT field to given value.

### HasVAT

`func (o *OrganisationView) HasVAT() bool`

HasVAT returns a boolean if a field has been set.

### GetAvatar

`func (o *OrganisationView) GetAvatar() string`

GetAvatar returns the Avatar field if non-nil, zero value otherwise.

### GetAvatarOk

`func (o *OrganisationView) GetAvatarOk() (*string, bool)`

GetAvatarOk returns a tuple with the Avatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatar

`func (o *OrganisationView) SetAvatar(v string)`

SetAvatar sets Avatar field to given value.

### HasAvatar

`func (o *OrganisationView) HasAvatar() bool`

HasAvatar returns a boolean if a field has been set.

### GetVatState

`func (o *OrganisationView) GetVatState() string`

GetVatState returns the VatState field if non-nil, zero value otherwise.

### GetVatStateOk

`func (o *OrganisationView) GetVatStateOk() (*string, bool)`

GetVatStateOk returns a tuple with the VatState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVatState

`func (o *OrganisationView) SetVatState(v string)`

SetVatState sets VatState field to given value.

### HasVatState

`func (o *OrganisationView) HasVatState() bool`

HasVatState returns a boolean if a field has been set.

### GetCustomerFullName

`func (o *OrganisationView) GetCustomerFullName() string`

GetCustomerFullName returns the CustomerFullName field if non-nil, zero value otherwise.

### GetCustomerFullNameOk

`func (o *OrganisationView) GetCustomerFullNameOk() (*string, bool)`

GetCustomerFullNameOk returns a tuple with the CustomerFullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerFullName

`func (o *OrganisationView) SetCustomerFullName(v string)`

SetCustomerFullName sets CustomerFullName field to given value.

### HasCustomerFullName

`func (o *OrganisationView) HasCustomerFullName() bool`

HasCustomerFullName returns a boolean if a field has been set.

### GetCanPay

`func (o *OrganisationView) GetCanPay() bool`

GetCanPay returns the CanPay field if non-nil, zero value otherwise.

### GetCanPayOk

`func (o *OrganisationView) GetCanPayOk() (*bool, bool)`

GetCanPayOk returns a tuple with the CanPay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanPay

`func (o *OrganisationView) SetCanPay(v bool)`

SetCanPay sets CanPay field to given value.

### HasCanPay

`func (o *OrganisationView) HasCanPay() bool`

HasCanPay returns a boolean if a field has been set.

### GetCleverEnterprise

`func (o *OrganisationView) GetCleverEnterprise() bool`

GetCleverEnterprise returns the CleverEnterprise field if non-nil, zero value otherwise.

### GetCleverEnterpriseOk

`func (o *OrganisationView) GetCleverEnterpriseOk() (*bool, bool)`

GetCleverEnterpriseOk returns a tuple with the CleverEnterprise field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCleverEnterprise

`func (o *OrganisationView) SetCleverEnterprise(v bool)`

SetCleverEnterprise sets CleverEnterprise field to given value.

### HasCleverEnterprise

`func (o *OrganisationView) HasCleverEnterprise() bool`

HasCleverEnterprise returns a boolean if a field has been set.

### GetEmergencyNumber

`func (o *OrganisationView) GetEmergencyNumber() string`

GetEmergencyNumber returns the EmergencyNumber field if non-nil, zero value otherwise.

### GetEmergencyNumberOk

`func (o *OrganisationView) GetEmergencyNumberOk() (*string, bool)`

GetEmergencyNumberOk returns a tuple with the EmergencyNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmergencyNumber

`func (o *OrganisationView) SetEmergencyNumber(v string)`

SetEmergencyNumber sets EmergencyNumber field to given value.

### HasEmergencyNumber

`func (o *OrganisationView) HasEmergencyNumber() bool`

HasEmergencyNumber returns a boolean if a field has been set.

### GetCanSEPA

`func (o *OrganisationView) GetCanSEPA() bool`

GetCanSEPA returns the CanSEPA field if non-nil, zero value otherwise.

### GetCanSEPAOk

`func (o *OrganisationView) GetCanSEPAOk() (*bool, bool)`

GetCanSEPAOk returns a tuple with the CanSEPA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanSEPA

`func (o *OrganisationView) SetCanSEPA(v bool)`

SetCanSEPA sets CanSEPA field to given value.

### HasCanSEPA

`func (o *OrganisationView) HasCanSEPA() bool`

HasCanSEPA returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


