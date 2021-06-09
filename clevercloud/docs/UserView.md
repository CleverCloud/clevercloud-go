# UserView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Phone** | Pointer to **string** |  | [optional] 
**Address** | Pointer to **string** |  | [optional] 
**City** | Pointer to **string** |  | [optional] 
**Zipcode** | Pointer to **string** |  | [optional] 
**Country** | Pointer to **string** |  | [optional] 
**Avatar** | Pointer to **string** |  | [optional] 
**CreationDate** | Pointer to **int64** |  | [optional] 
**Lang** | Pointer to **string** |  | [optional] 
**EmailValidated** | Pointer to **bool** |  | [optional] 
**OauthApps** | Pointer to **[]string** |  | [optional] 
**Admin** | Pointer to **bool** |  | [optional] 
**CanPay** | Pointer to **bool** |  | [optional] 
**PreferredMFA** | Pointer to **string** |  | [optional] 
**HasPassword** | Pointer to **bool** |  | [optional] 

## Methods

### NewUserView

`func NewUserView() *UserView`

NewUserView instantiates a new UserView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserViewWithDefaults

`func NewUserViewWithDefaults() *UserView`

NewUserViewWithDefaults instantiates a new UserView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UserView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetEmail

`func (o *UserView) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserView) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserView) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UserView) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetName

`func (o *UserView) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserView) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserView) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UserView) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPhone

`func (o *UserView) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *UserView) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *UserView) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *UserView) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetAddress

`func (o *UserView) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *UserView) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *UserView) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *UserView) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetCity

`func (o *UserView) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *UserView) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *UserView) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *UserView) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetZipcode

`func (o *UserView) GetZipcode() string`

GetZipcode returns the Zipcode field if non-nil, zero value otherwise.

### GetZipcodeOk

`func (o *UserView) GetZipcodeOk() (*string, bool)`

GetZipcodeOk returns a tuple with the Zipcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZipcode

`func (o *UserView) SetZipcode(v string)`

SetZipcode sets Zipcode field to given value.

### HasZipcode

`func (o *UserView) HasZipcode() bool`

HasZipcode returns a boolean if a field has been set.

### GetCountry

`func (o *UserView) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *UserView) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *UserView) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *UserView) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetAvatar

`func (o *UserView) GetAvatar() string`

GetAvatar returns the Avatar field if non-nil, zero value otherwise.

### GetAvatarOk

`func (o *UserView) GetAvatarOk() (*string, bool)`

GetAvatarOk returns a tuple with the Avatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatar

`func (o *UserView) SetAvatar(v string)`

SetAvatar sets Avatar field to given value.

### HasAvatar

`func (o *UserView) HasAvatar() bool`

HasAvatar returns a boolean if a field has been set.

### GetCreationDate

`func (o *UserView) GetCreationDate() int64`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *UserView) GetCreationDateOk() (*int64, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *UserView) SetCreationDate(v int64)`

SetCreationDate sets CreationDate field to given value.

### HasCreationDate

`func (o *UserView) HasCreationDate() bool`

HasCreationDate returns a boolean if a field has been set.

### GetLang

`func (o *UserView) GetLang() string`

GetLang returns the Lang field if non-nil, zero value otherwise.

### GetLangOk

`func (o *UserView) GetLangOk() (*string, bool)`

GetLangOk returns a tuple with the Lang field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLang

`func (o *UserView) SetLang(v string)`

SetLang sets Lang field to given value.

### HasLang

`func (o *UserView) HasLang() bool`

HasLang returns a boolean if a field has been set.

### GetEmailValidated

`func (o *UserView) GetEmailValidated() bool`

GetEmailValidated returns the EmailValidated field if non-nil, zero value otherwise.

### GetEmailValidatedOk

`func (o *UserView) GetEmailValidatedOk() (*bool, bool)`

GetEmailValidatedOk returns a tuple with the EmailValidated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailValidated

`func (o *UserView) SetEmailValidated(v bool)`

SetEmailValidated sets EmailValidated field to given value.

### HasEmailValidated

`func (o *UserView) HasEmailValidated() bool`

HasEmailValidated returns a boolean if a field has been set.

### GetOauthApps

`func (o *UserView) GetOauthApps() []string`

GetOauthApps returns the OauthApps field if non-nil, zero value otherwise.

### GetOauthAppsOk

`func (o *UserView) GetOauthAppsOk() (*[]string, bool)`

GetOauthAppsOk returns a tuple with the OauthApps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauthApps

`func (o *UserView) SetOauthApps(v []string)`

SetOauthApps sets OauthApps field to given value.

### HasOauthApps

`func (o *UserView) HasOauthApps() bool`

HasOauthApps returns a boolean if a field has been set.

### GetAdmin

`func (o *UserView) GetAdmin() bool`

GetAdmin returns the Admin field if non-nil, zero value otherwise.

### GetAdminOk

`func (o *UserView) GetAdminOk() (*bool, bool)`

GetAdminOk returns a tuple with the Admin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdmin

`func (o *UserView) SetAdmin(v bool)`

SetAdmin sets Admin field to given value.

### HasAdmin

`func (o *UserView) HasAdmin() bool`

HasAdmin returns a boolean if a field has been set.

### GetCanPay

`func (o *UserView) GetCanPay() bool`

GetCanPay returns the CanPay field if non-nil, zero value otherwise.

### GetCanPayOk

`func (o *UserView) GetCanPayOk() (*bool, bool)`

GetCanPayOk returns a tuple with the CanPay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanPay

`func (o *UserView) SetCanPay(v bool)`

SetCanPay sets CanPay field to given value.

### HasCanPay

`func (o *UserView) HasCanPay() bool`

HasCanPay returns a boolean if a field has been set.

### GetPreferredMFA

`func (o *UserView) GetPreferredMFA() string`

GetPreferredMFA returns the PreferredMFA field if non-nil, zero value otherwise.

### GetPreferredMFAOk

`func (o *UserView) GetPreferredMFAOk() (*string, bool)`

GetPreferredMFAOk returns a tuple with the PreferredMFA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredMFA

`func (o *UserView) SetPreferredMFA(v string)`

SetPreferredMFA sets PreferredMFA field to given value.

### HasPreferredMFA

`func (o *UserView) HasPreferredMFA() bool`

HasPreferredMFA returns a boolean if a field has been set.

### GetHasPassword

`func (o *UserView) GetHasPassword() bool`

GetHasPassword returns the HasPassword field if non-nil, zero value otherwise.

### GetHasPasswordOk

`func (o *UserView) GetHasPasswordOk() (*bool, bool)`

GetHasPasswordOk returns a tuple with the HasPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPassword

`func (o *UserView) SetHasPassword(v bool)`

SetHasPassword sets HasPassword field to given value.

### HasHasPassword

`func (o *UserView) HasHasPassword() bool`

HasHasPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


