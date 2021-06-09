# PaymentInfoView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FromOrga** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Company** | Pointer to **string** |  | [optional] 
**Address** | Pointer to **string** |  | [optional] 
**ZipCode** | Pointer to **string** |  | [optional] 
**City** | Pointer to **string** |  | [optional] 
**Country** | Pointer to **string** |  | [optional] 
**VatNumber** | Pointer to **string** |  | [optional] 
**Errors** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewPaymentInfoView

`func NewPaymentInfoView() *PaymentInfoView`

NewPaymentInfoView instantiates a new PaymentInfoView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentInfoViewWithDefaults

`func NewPaymentInfoViewWithDefaults() *PaymentInfoView`

NewPaymentInfoViewWithDefaults instantiates a new PaymentInfoView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFromOrga

`func (o *PaymentInfoView) GetFromOrga() bool`

GetFromOrga returns the FromOrga field if non-nil, zero value otherwise.

### GetFromOrgaOk

`func (o *PaymentInfoView) GetFromOrgaOk() (*bool, bool)`

GetFromOrgaOk returns a tuple with the FromOrga field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromOrga

`func (o *PaymentInfoView) SetFromOrga(v bool)`

SetFromOrga sets FromOrga field to given value.

### HasFromOrga

`func (o *PaymentInfoView) HasFromOrga() bool`

HasFromOrga returns a boolean if a field has been set.

### GetName

`func (o *PaymentInfoView) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PaymentInfoView) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PaymentInfoView) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PaymentInfoView) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCompany

`func (o *PaymentInfoView) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *PaymentInfoView) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *PaymentInfoView) SetCompany(v string)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *PaymentInfoView) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetAddress

`func (o *PaymentInfoView) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *PaymentInfoView) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *PaymentInfoView) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *PaymentInfoView) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetZipCode

`func (o *PaymentInfoView) GetZipCode() string`

GetZipCode returns the ZipCode field if non-nil, zero value otherwise.

### GetZipCodeOk

`func (o *PaymentInfoView) GetZipCodeOk() (*string, bool)`

GetZipCodeOk returns a tuple with the ZipCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZipCode

`func (o *PaymentInfoView) SetZipCode(v string)`

SetZipCode sets ZipCode field to given value.

### HasZipCode

`func (o *PaymentInfoView) HasZipCode() bool`

HasZipCode returns a boolean if a field has been set.

### GetCity

`func (o *PaymentInfoView) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *PaymentInfoView) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *PaymentInfoView) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *PaymentInfoView) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetCountry

`func (o *PaymentInfoView) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *PaymentInfoView) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *PaymentInfoView) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *PaymentInfoView) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetVatNumber

`func (o *PaymentInfoView) GetVatNumber() string`

GetVatNumber returns the VatNumber field if non-nil, zero value otherwise.

### GetVatNumberOk

`func (o *PaymentInfoView) GetVatNumberOk() (*string, bool)`

GetVatNumberOk returns a tuple with the VatNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVatNumber

`func (o *PaymentInfoView) SetVatNumber(v string)`

SetVatNumber sets VatNumber field to given value.

### HasVatNumber

`func (o *PaymentInfoView) HasVatNumber() bool`

HasVatNumber returns a boolean if a field has been set.

### GetErrors

`func (o *PaymentInfoView) GetErrors() map[string]string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *PaymentInfoView) GetErrorsOk() (*map[string]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *PaymentInfoView) SetErrors(v map[string]string)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *PaymentInfoView) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


