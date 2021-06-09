# PaymentMethodView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OwnerId** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Token** | Pointer to **string** |  | [optional] 
**ImageUrl** | Pointer to **string** |  | [optional] 
**IsDefault** | Pointer to **bool** |  | [optional] 
**HolderName** | Pointer to **string** |  | [optional] 
**Number** | Pointer to **string** |  | [optional] 
**ExpirationDate** | Pointer to **string** |  | [optional] 
**IsExpired** | Pointer to **bool** |  | [optional] 
**CardType** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**BankCode** | Pointer to **string** |  | [optional] 
**BranchCode** | Pointer to **string** |  | [optional] 
**Country** | Pointer to **string** |  | [optional] 
**Fingerprint** | Pointer to **string** |  | [optional] 

## Methods

### NewPaymentMethodView

`func NewPaymentMethodView() *PaymentMethodView`

NewPaymentMethodView instantiates a new PaymentMethodView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentMethodViewWithDefaults

`func NewPaymentMethodViewWithDefaults() *PaymentMethodView`

NewPaymentMethodViewWithDefaults instantiates a new PaymentMethodView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOwnerId

`func (o *PaymentMethodView) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *PaymentMethodView) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *PaymentMethodView) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *PaymentMethodView) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.

### GetType

`func (o *PaymentMethodView) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PaymentMethodView) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PaymentMethodView) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PaymentMethodView) HasType() bool`

HasType returns a boolean if a field has been set.

### GetToken

`func (o *PaymentMethodView) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *PaymentMethodView) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *PaymentMethodView) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *PaymentMethodView) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetImageUrl

`func (o *PaymentMethodView) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *PaymentMethodView) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *PaymentMethodView) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.

### HasImageUrl

`func (o *PaymentMethodView) HasImageUrl() bool`

HasImageUrl returns a boolean if a field has been set.

### GetIsDefault

`func (o *PaymentMethodView) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *PaymentMethodView) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *PaymentMethodView) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *PaymentMethodView) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetHolderName

`func (o *PaymentMethodView) GetHolderName() string`

GetHolderName returns the HolderName field if non-nil, zero value otherwise.

### GetHolderNameOk

`func (o *PaymentMethodView) GetHolderNameOk() (*string, bool)`

GetHolderNameOk returns a tuple with the HolderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHolderName

`func (o *PaymentMethodView) SetHolderName(v string)`

SetHolderName sets HolderName field to given value.

### HasHolderName

`func (o *PaymentMethodView) HasHolderName() bool`

HasHolderName returns a boolean if a field has been set.

### GetNumber

`func (o *PaymentMethodView) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *PaymentMethodView) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *PaymentMethodView) SetNumber(v string)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *PaymentMethodView) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetExpirationDate

`func (o *PaymentMethodView) GetExpirationDate() string`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *PaymentMethodView) GetExpirationDateOk() (*string, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *PaymentMethodView) SetExpirationDate(v string)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *PaymentMethodView) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### GetIsExpired

`func (o *PaymentMethodView) GetIsExpired() bool`

GetIsExpired returns the IsExpired field if non-nil, zero value otherwise.

### GetIsExpiredOk

`func (o *PaymentMethodView) GetIsExpiredOk() (*bool, bool)`

GetIsExpiredOk returns a tuple with the IsExpired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsExpired

`func (o *PaymentMethodView) SetIsExpired(v bool)`

SetIsExpired sets IsExpired field to given value.

### HasIsExpired

`func (o *PaymentMethodView) HasIsExpired() bool`

HasIsExpired returns a boolean if a field has been set.

### GetCardType

`func (o *PaymentMethodView) GetCardType() string`

GetCardType returns the CardType field if non-nil, zero value otherwise.

### GetCardTypeOk

`func (o *PaymentMethodView) GetCardTypeOk() (*string, bool)`

GetCardTypeOk returns a tuple with the CardType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardType

`func (o *PaymentMethodView) SetCardType(v string)`

SetCardType sets CardType field to given value.

### HasCardType

`func (o *PaymentMethodView) HasCardType() bool`

HasCardType returns a boolean if a field has been set.

### GetEmail

`func (o *PaymentMethodView) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *PaymentMethodView) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *PaymentMethodView) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *PaymentMethodView) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetBankCode

`func (o *PaymentMethodView) GetBankCode() string`

GetBankCode returns the BankCode field if non-nil, zero value otherwise.

### GetBankCodeOk

`func (o *PaymentMethodView) GetBankCodeOk() (*string, bool)`

GetBankCodeOk returns a tuple with the BankCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankCode

`func (o *PaymentMethodView) SetBankCode(v string)`

SetBankCode sets BankCode field to given value.

### HasBankCode

`func (o *PaymentMethodView) HasBankCode() bool`

HasBankCode returns a boolean if a field has been set.

### GetBranchCode

`func (o *PaymentMethodView) GetBranchCode() string`

GetBranchCode returns the BranchCode field if non-nil, zero value otherwise.

### GetBranchCodeOk

`func (o *PaymentMethodView) GetBranchCodeOk() (*string, bool)`

GetBranchCodeOk returns a tuple with the BranchCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranchCode

`func (o *PaymentMethodView) SetBranchCode(v string)`

SetBranchCode sets BranchCode field to given value.

### HasBranchCode

`func (o *PaymentMethodView) HasBranchCode() bool`

HasBranchCode returns a boolean if a field has been set.

### GetCountry

`func (o *PaymentMethodView) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *PaymentMethodView) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *PaymentMethodView) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *PaymentMethodView) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetFingerprint

`func (o *PaymentMethodView) GetFingerprint() string`

GetFingerprint returns the Fingerprint field if non-nil, zero value otherwise.

### GetFingerprintOk

`func (o *PaymentMethodView) GetFingerprintOk() (*string, bool)`

GetFingerprintOk returns a tuple with the Fingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprint

`func (o *PaymentMethodView) SetFingerprint(v string)`

SetFingerprint sets Fingerprint field to given value.

### HasFingerprint

`func (o *PaymentMethodView) HasFingerprint() bool`

HasFingerprint returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


