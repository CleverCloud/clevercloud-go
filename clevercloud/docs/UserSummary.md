# UserSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Avatar** | Pointer to **string** |  | [optional] 
**Applications** | Pointer to [**[]ApplicationSummary**](ApplicationSummary.md) |  | [optional] 
**Addons** | Pointer to [**[]AddonSummary**](AddonSummary.md) |  | [optional] 
**Consumers** | Pointer to [**[]OAuth1ConsumerSummary**](OAuth1ConsumerSummary.md) |  | [optional] 
**Lang** | Pointer to **string** |  | [optional] 
**Admin** | Pointer to **bool** |  | [optional] 
**CanSEPA** | Pointer to **bool** |  | [optional] 

## Methods

### NewUserSummary

`func NewUserSummary() *UserSummary`

NewUserSummary instantiates a new UserSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserSummaryWithDefaults

`func NewUserSummaryWithDefaults() *UserSummary`

NewUserSummaryWithDefaults instantiates a new UserSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserSummary) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserSummary) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserSummary) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UserSummary) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *UserSummary) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserSummary) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserSummary) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UserSummary) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAvatar

`func (o *UserSummary) GetAvatar() string`

GetAvatar returns the Avatar field if non-nil, zero value otherwise.

### GetAvatarOk

`func (o *UserSummary) GetAvatarOk() (*string, bool)`

GetAvatarOk returns a tuple with the Avatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatar

`func (o *UserSummary) SetAvatar(v string)`

SetAvatar sets Avatar field to given value.

### HasAvatar

`func (o *UserSummary) HasAvatar() bool`

HasAvatar returns a boolean if a field has been set.

### GetApplications

`func (o *UserSummary) GetApplications() []ApplicationSummary`

GetApplications returns the Applications field if non-nil, zero value otherwise.

### GetApplicationsOk

`func (o *UserSummary) GetApplicationsOk() (*[]ApplicationSummary, bool)`

GetApplicationsOk returns a tuple with the Applications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplications

`func (o *UserSummary) SetApplications(v []ApplicationSummary)`

SetApplications sets Applications field to given value.

### HasApplications

`func (o *UserSummary) HasApplications() bool`

HasApplications returns a boolean if a field has been set.

### GetAddons

`func (o *UserSummary) GetAddons() []AddonSummary`

GetAddons returns the Addons field if non-nil, zero value otherwise.

### GetAddonsOk

`func (o *UserSummary) GetAddonsOk() (*[]AddonSummary, bool)`

GetAddonsOk returns a tuple with the Addons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddons

`func (o *UserSummary) SetAddons(v []AddonSummary)`

SetAddons sets Addons field to given value.

### HasAddons

`func (o *UserSummary) HasAddons() bool`

HasAddons returns a boolean if a field has been set.

### GetConsumers

`func (o *UserSummary) GetConsumers() []OAuth1ConsumerSummary`

GetConsumers returns the Consumers field if non-nil, zero value otherwise.

### GetConsumersOk

`func (o *UserSummary) GetConsumersOk() (*[]OAuth1ConsumerSummary, bool)`

GetConsumersOk returns a tuple with the Consumers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumers

`func (o *UserSummary) SetConsumers(v []OAuth1ConsumerSummary)`

SetConsumers sets Consumers field to given value.

### HasConsumers

`func (o *UserSummary) HasConsumers() bool`

HasConsumers returns a boolean if a field has been set.

### GetLang

`func (o *UserSummary) GetLang() string`

GetLang returns the Lang field if non-nil, zero value otherwise.

### GetLangOk

`func (o *UserSummary) GetLangOk() (*string, bool)`

GetLangOk returns a tuple with the Lang field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLang

`func (o *UserSummary) SetLang(v string)`

SetLang sets Lang field to given value.

### HasLang

`func (o *UserSummary) HasLang() bool`

HasLang returns a boolean if a field has been set.

### GetAdmin

`func (o *UserSummary) GetAdmin() bool`

GetAdmin returns the Admin field if non-nil, zero value otherwise.

### GetAdminOk

`func (o *UserSummary) GetAdminOk() (*bool, bool)`

GetAdminOk returns a tuple with the Admin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdmin

`func (o *UserSummary) SetAdmin(v bool)`

SetAdmin sets Admin field to given value.

### HasAdmin

`func (o *UserSummary) HasAdmin() bool`

HasAdmin returns a boolean if a field has been set.

### GetCanSEPA

`func (o *UserSummary) GetCanSEPA() bool`

GetCanSEPA returns the CanSEPA field if non-nil, zero value otherwise.

### GetCanSEPAOk

`func (o *UserSummary) GetCanSEPAOk() (*bool, bool)`

GetCanSEPAOk returns a tuple with the CanSEPA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanSEPA

`func (o *UserSummary) SetCanSEPA(v bool)`

SetCanSEPA sets CanSEPA field to given value.

### HasCanSEPA

`func (o *UserSummary) HasCanSEPA() bool`

HasCanSEPA returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


