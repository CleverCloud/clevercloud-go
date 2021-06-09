# InvoiceRendering

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Number** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**EmissionDate** | Pointer to **string** |  | [optional] 
**PaymentDate** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Company** | Pointer to **string** |  | [optional] 
**Target** | Pointer to **string** |  | [optional] 
**Address** | Pointer to **string** |  | [optional] 
**ZipCode** | Pointer to **string** |  | [optional] 
**City** | Pointer to **string** |  | [optional] 
**Country** | Pointer to **string** |  | [optional] 
**CountryCode** | Pointer to **string** |  | [optional] 
**VatNumber** | Pointer to **string** |  | [optional] 
**FromSubscription** | Pointer to **bool** |  | [optional] 
**Lines** | Pointer to [**[]InvoiceLineRendering**](InvoiceLineRendering.md) |  | [optional] 
**OriginalTotal** | Pointer to **float32** |  | [optional] 
**TotalHT** | Pointer to **float32** |  | [optional] 
**TotalVAT** | Pointer to **float32** |  | [optional] 
**TotalTTC** | Pointer to **float32** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**PayingUser** | Pointer to [**OrganisationMemberView**](OrganisationMemberView.md) |  | [optional] 
**ErrorCode** | Pointer to **string** |  | [optional] 
**ErrorShortMsg** | Pointer to **string** |  | [optional] 
**ErrorLongMsg** | Pointer to **string** |  | [optional] 
**Token** | Pointer to **string** |  | [optional] 
**TargetId** | Pointer to **string** |  | [optional] 
**ForId** | Pointer to **string** |  | [optional] 
**CustomerOrderId** | Pointer to **string** |  | [optional] 
**CustomerCostCenter** | Pointer to **string** |  | [optional] 
**VatRate** | Pointer to **float32** |  | [optional] 
**PayWhen** | Pointer to **string** |  | [optional] 

## Methods

### NewInvoiceRendering

`func NewInvoiceRendering() *InvoiceRendering`

NewInvoiceRendering instantiates a new InvoiceRendering object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceRenderingWithDefaults

`func NewInvoiceRenderingWithDefaults() *InvoiceRendering`

NewInvoiceRenderingWithDefaults instantiates a new InvoiceRendering object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InvoiceRendering) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InvoiceRendering) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InvoiceRendering) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InvoiceRendering) HasId() bool`

HasId returns a boolean if a field has been set.

### GetNumber

`func (o *InvoiceRendering) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *InvoiceRendering) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *InvoiceRendering) SetNumber(v string)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *InvoiceRendering) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetStatus

`func (o *InvoiceRendering) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InvoiceRendering) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InvoiceRendering) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InvoiceRendering) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetEmissionDate

`func (o *InvoiceRendering) GetEmissionDate() string`

GetEmissionDate returns the EmissionDate field if non-nil, zero value otherwise.

### GetEmissionDateOk

`func (o *InvoiceRendering) GetEmissionDateOk() (*string, bool)`

GetEmissionDateOk returns a tuple with the EmissionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmissionDate

`func (o *InvoiceRendering) SetEmissionDate(v string)`

SetEmissionDate sets EmissionDate field to given value.

### HasEmissionDate

`func (o *InvoiceRendering) HasEmissionDate() bool`

HasEmissionDate returns a boolean if a field has been set.

### GetPaymentDate

`func (o *InvoiceRendering) GetPaymentDate() string`

GetPaymentDate returns the PaymentDate field if non-nil, zero value otherwise.

### GetPaymentDateOk

`func (o *InvoiceRendering) GetPaymentDateOk() (*string, bool)`

GetPaymentDateOk returns a tuple with the PaymentDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentDate

`func (o *InvoiceRendering) SetPaymentDate(v string)`

SetPaymentDate sets PaymentDate field to given value.

### HasPaymentDate

`func (o *InvoiceRendering) HasPaymentDate() bool`

HasPaymentDate returns a boolean if a field has been set.

### GetName

`func (o *InvoiceRendering) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InvoiceRendering) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InvoiceRendering) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InvoiceRendering) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCompany

`func (o *InvoiceRendering) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *InvoiceRendering) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *InvoiceRendering) SetCompany(v string)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *InvoiceRendering) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetTarget

`func (o *InvoiceRendering) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *InvoiceRendering) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *InvoiceRendering) SetTarget(v string)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *InvoiceRendering) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### GetAddress

`func (o *InvoiceRendering) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *InvoiceRendering) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *InvoiceRendering) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *InvoiceRendering) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetZipCode

`func (o *InvoiceRendering) GetZipCode() string`

GetZipCode returns the ZipCode field if non-nil, zero value otherwise.

### GetZipCodeOk

`func (o *InvoiceRendering) GetZipCodeOk() (*string, bool)`

GetZipCodeOk returns a tuple with the ZipCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZipCode

`func (o *InvoiceRendering) SetZipCode(v string)`

SetZipCode sets ZipCode field to given value.

### HasZipCode

`func (o *InvoiceRendering) HasZipCode() bool`

HasZipCode returns a boolean if a field has been set.

### GetCity

`func (o *InvoiceRendering) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *InvoiceRendering) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *InvoiceRendering) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *InvoiceRendering) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetCountry

`func (o *InvoiceRendering) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *InvoiceRendering) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *InvoiceRendering) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *InvoiceRendering) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetCountryCode

`func (o *InvoiceRendering) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *InvoiceRendering) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *InvoiceRendering) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *InvoiceRendering) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetVatNumber

`func (o *InvoiceRendering) GetVatNumber() string`

GetVatNumber returns the VatNumber field if non-nil, zero value otherwise.

### GetVatNumberOk

`func (o *InvoiceRendering) GetVatNumberOk() (*string, bool)`

GetVatNumberOk returns a tuple with the VatNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVatNumber

`func (o *InvoiceRendering) SetVatNumber(v string)`

SetVatNumber sets VatNumber field to given value.

### HasVatNumber

`func (o *InvoiceRendering) HasVatNumber() bool`

HasVatNumber returns a boolean if a field has been set.

### GetFromSubscription

`func (o *InvoiceRendering) GetFromSubscription() bool`

GetFromSubscription returns the FromSubscription field if non-nil, zero value otherwise.

### GetFromSubscriptionOk

`func (o *InvoiceRendering) GetFromSubscriptionOk() (*bool, bool)`

GetFromSubscriptionOk returns a tuple with the FromSubscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromSubscription

`func (o *InvoiceRendering) SetFromSubscription(v bool)`

SetFromSubscription sets FromSubscription field to given value.

### HasFromSubscription

`func (o *InvoiceRendering) HasFromSubscription() bool`

HasFromSubscription returns a boolean if a field has been set.

### GetLines

`func (o *InvoiceRendering) GetLines() []InvoiceLineRendering`

GetLines returns the Lines field if non-nil, zero value otherwise.

### GetLinesOk

`func (o *InvoiceRendering) GetLinesOk() (*[]InvoiceLineRendering, bool)`

GetLinesOk returns a tuple with the Lines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLines

`func (o *InvoiceRendering) SetLines(v []InvoiceLineRendering)`

SetLines sets Lines field to given value.

### HasLines

`func (o *InvoiceRendering) HasLines() bool`

HasLines returns a boolean if a field has been set.

### GetOriginalTotal

`func (o *InvoiceRendering) GetOriginalTotal() float32`

GetOriginalTotal returns the OriginalTotal field if non-nil, zero value otherwise.

### GetOriginalTotalOk

`func (o *InvoiceRendering) GetOriginalTotalOk() (*float32, bool)`

GetOriginalTotalOk returns a tuple with the OriginalTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalTotal

`func (o *InvoiceRendering) SetOriginalTotal(v float32)`

SetOriginalTotal sets OriginalTotal field to given value.

### HasOriginalTotal

`func (o *InvoiceRendering) HasOriginalTotal() bool`

HasOriginalTotal returns a boolean if a field has been set.

### GetTotalHT

`func (o *InvoiceRendering) GetTotalHT() float32`

GetTotalHT returns the TotalHT field if non-nil, zero value otherwise.

### GetTotalHTOk

`func (o *InvoiceRendering) GetTotalHTOk() (*float32, bool)`

GetTotalHTOk returns a tuple with the TotalHT field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalHT

`func (o *InvoiceRendering) SetTotalHT(v float32)`

SetTotalHT sets TotalHT field to given value.

### HasTotalHT

`func (o *InvoiceRendering) HasTotalHT() bool`

HasTotalHT returns a boolean if a field has been set.

### GetTotalVAT

`func (o *InvoiceRendering) GetTotalVAT() float32`

GetTotalVAT returns the TotalVAT field if non-nil, zero value otherwise.

### GetTotalVATOk

`func (o *InvoiceRendering) GetTotalVATOk() (*float32, bool)`

GetTotalVATOk returns a tuple with the TotalVAT field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalVAT

`func (o *InvoiceRendering) SetTotalVAT(v float32)`

SetTotalVAT sets TotalVAT field to given value.

### HasTotalVAT

`func (o *InvoiceRendering) HasTotalVAT() bool`

HasTotalVAT returns a boolean if a field has been set.

### GetTotalTTC

`func (o *InvoiceRendering) GetTotalTTC() float32`

GetTotalTTC returns the TotalTTC field if non-nil, zero value otherwise.

### GetTotalTTCOk

`func (o *InvoiceRendering) GetTotalTTCOk() (*float32, bool)`

GetTotalTTCOk returns a tuple with the TotalTTC field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTTC

`func (o *InvoiceRendering) SetTotalTTC(v float32)`

SetTotalTTC sets TotalTTC field to given value.

### HasTotalTTC

`func (o *InvoiceRendering) HasTotalTTC() bool`

HasTotalTTC returns a boolean if a field has been set.

### GetType

`func (o *InvoiceRendering) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InvoiceRendering) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InvoiceRendering) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *InvoiceRendering) HasType() bool`

HasType returns a boolean if a field has been set.

### GetPayingUser

`func (o *InvoiceRendering) GetPayingUser() OrganisationMemberView`

GetPayingUser returns the PayingUser field if non-nil, zero value otherwise.

### GetPayingUserOk

`func (o *InvoiceRendering) GetPayingUserOk() (*OrganisationMemberView, bool)`

GetPayingUserOk returns a tuple with the PayingUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayingUser

`func (o *InvoiceRendering) SetPayingUser(v OrganisationMemberView)`

SetPayingUser sets PayingUser field to given value.

### HasPayingUser

`func (o *InvoiceRendering) HasPayingUser() bool`

HasPayingUser returns a boolean if a field has been set.

### GetErrorCode

`func (o *InvoiceRendering) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *InvoiceRendering) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *InvoiceRendering) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *InvoiceRendering) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetErrorShortMsg

`func (o *InvoiceRendering) GetErrorShortMsg() string`

GetErrorShortMsg returns the ErrorShortMsg field if non-nil, zero value otherwise.

### GetErrorShortMsgOk

`func (o *InvoiceRendering) GetErrorShortMsgOk() (*string, bool)`

GetErrorShortMsgOk returns a tuple with the ErrorShortMsg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorShortMsg

`func (o *InvoiceRendering) SetErrorShortMsg(v string)`

SetErrorShortMsg sets ErrorShortMsg field to given value.

### HasErrorShortMsg

`func (o *InvoiceRendering) HasErrorShortMsg() bool`

HasErrorShortMsg returns a boolean if a field has been set.

### GetErrorLongMsg

`func (o *InvoiceRendering) GetErrorLongMsg() string`

GetErrorLongMsg returns the ErrorLongMsg field if non-nil, zero value otherwise.

### GetErrorLongMsgOk

`func (o *InvoiceRendering) GetErrorLongMsgOk() (*string, bool)`

GetErrorLongMsgOk returns a tuple with the ErrorLongMsg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorLongMsg

`func (o *InvoiceRendering) SetErrorLongMsg(v string)`

SetErrorLongMsg sets ErrorLongMsg field to given value.

### HasErrorLongMsg

`func (o *InvoiceRendering) HasErrorLongMsg() bool`

HasErrorLongMsg returns a boolean if a field has been set.

### GetToken

`func (o *InvoiceRendering) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *InvoiceRendering) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *InvoiceRendering) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *InvoiceRendering) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetTargetId

`func (o *InvoiceRendering) GetTargetId() string`

GetTargetId returns the TargetId field if non-nil, zero value otherwise.

### GetTargetIdOk

`func (o *InvoiceRendering) GetTargetIdOk() (*string, bool)`

GetTargetIdOk returns a tuple with the TargetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetId

`func (o *InvoiceRendering) SetTargetId(v string)`

SetTargetId sets TargetId field to given value.

### HasTargetId

`func (o *InvoiceRendering) HasTargetId() bool`

HasTargetId returns a boolean if a field has been set.

### GetForId

`func (o *InvoiceRendering) GetForId() string`

GetForId returns the ForId field if non-nil, zero value otherwise.

### GetForIdOk

`func (o *InvoiceRendering) GetForIdOk() (*string, bool)`

GetForIdOk returns a tuple with the ForId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForId

`func (o *InvoiceRendering) SetForId(v string)`

SetForId sets ForId field to given value.

### HasForId

`func (o *InvoiceRendering) HasForId() bool`

HasForId returns a boolean if a field has been set.

### GetCustomerOrderId

`func (o *InvoiceRendering) GetCustomerOrderId() string`

GetCustomerOrderId returns the CustomerOrderId field if non-nil, zero value otherwise.

### GetCustomerOrderIdOk

`func (o *InvoiceRendering) GetCustomerOrderIdOk() (*string, bool)`

GetCustomerOrderIdOk returns a tuple with the CustomerOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerOrderId

`func (o *InvoiceRendering) SetCustomerOrderId(v string)`

SetCustomerOrderId sets CustomerOrderId field to given value.

### HasCustomerOrderId

`func (o *InvoiceRendering) HasCustomerOrderId() bool`

HasCustomerOrderId returns a boolean if a field has been set.

### GetCustomerCostCenter

`func (o *InvoiceRendering) GetCustomerCostCenter() string`

GetCustomerCostCenter returns the CustomerCostCenter field if non-nil, zero value otherwise.

### GetCustomerCostCenterOk

`func (o *InvoiceRendering) GetCustomerCostCenterOk() (*string, bool)`

GetCustomerCostCenterOk returns a tuple with the CustomerCostCenter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerCostCenter

`func (o *InvoiceRendering) SetCustomerCostCenter(v string)`

SetCustomerCostCenter sets CustomerCostCenter field to given value.

### HasCustomerCostCenter

`func (o *InvoiceRendering) HasCustomerCostCenter() bool`

HasCustomerCostCenter returns a boolean if a field has been set.

### GetVatRate

`func (o *InvoiceRendering) GetVatRate() float32`

GetVatRate returns the VatRate field if non-nil, zero value otherwise.

### GetVatRateOk

`func (o *InvoiceRendering) GetVatRateOk() (*float32, bool)`

GetVatRateOk returns a tuple with the VatRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVatRate

`func (o *InvoiceRendering) SetVatRate(v float32)`

SetVatRate sets VatRate field to given value.

### HasVatRate

`func (o *InvoiceRendering) HasVatRate() bool`

HasVatRate returns a boolean if a field has been set.

### GetPayWhen

`func (o *InvoiceRendering) GetPayWhen() string`

GetPayWhen returns the PayWhen field if non-nil, zero value otherwise.

### GetPayWhenOk

`func (o *InvoiceRendering) GetPayWhenOk() (*string, bool)`

GetPayWhenOk returns a tuple with the PayWhen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayWhen

`func (o *InvoiceRendering) SetPayWhen(v string)`

SetPayWhen sets PayWhen field to given value.

### HasPayWhen

`func (o *InvoiceRendering) HasPayWhen() bool`

HasPayWhen returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


