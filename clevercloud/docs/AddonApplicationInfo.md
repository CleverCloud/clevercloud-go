# AddonApplicationInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Config** | Pointer to **map[string]string** |  | [optional] 
**CallbackUrl** | Pointer to **string** |  | [optional] 
**OwnerEmail** | Pointer to **string** |  | [optional] 
**OwnerId** | Pointer to **string** |  | [optional] 
**OwnerName** | Pointer to **string** |  | [optional] 
**OwnerEmails** | Pointer to **[]string** |  | [optional] 
**Region** | Pointer to **string** |  | [optional] 
**Domains** | Pointer to **[]string** |  | [optional] 

## Methods

### NewAddonApplicationInfo

`func NewAddonApplicationInfo() *AddonApplicationInfo`

NewAddonApplicationInfo instantiates a new AddonApplicationInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddonApplicationInfoWithDefaults

`func NewAddonApplicationInfoWithDefaults() *AddonApplicationInfo`

NewAddonApplicationInfoWithDefaults instantiates a new AddonApplicationInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AddonApplicationInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AddonApplicationInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AddonApplicationInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AddonApplicationInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *AddonApplicationInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddonApplicationInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddonApplicationInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AddonApplicationInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetConfig

`func (o *AddonApplicationInfo) GetConfig() map[string]string`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *AddonApplicationInfo) GetConfigOk() (*map[string]string, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *AddonApplicationInfo) SetConfig(v map[string]string)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *AddonApplicationInfo) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetCallbackUrl

`func (o *AddonApplicationInfo) GetCallbackUrl() string`

GetCallbackUrl returns the CallbackUrl field if non-nil, zero value otherwise.

### GetCallbackUrlOk

`func (o *AddonApplicationInfo) GetCallbackUrlOk() (*string, bool)`

GetCallbackUrlOk returns a tuple with the CallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackUrl

`func (o *AddonApplicationInfo) SetCallbackUrl(v string)`

SetCallbackUrl sets CallbackUrl field to given value.

### HasCallbackUrl

`func (o *AddonApplicationInfo) HasCallbackUrl() bool`

HasCallbackUrl returns a boolean if a field has been set.

### GetOwnerEmail

`func (o *AddonApplicationInfo) GetOwnerEmail() string`

GetOwnerEmail returns the OwnerEmail field if non-nil, zero value otherwise.

### GetOwnerEmailOk

`func (o *AddonApplicationInfo) GetOwnerEmailOk() (*string, bool)`

GetOwnerEmailOk returns a tuple with the OwnerEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerEmail

`func (o *AddonApplicationInfo) SetOwnerEmail(v string)`

SetOwnerEmail sets OwnerEmail field to given value.

### HasOwnerEmail

`func (o *AddonApplicationInfo) HasOwnerEmail() bool`

HasOwnerEmail returns a boolean if a field has been set.

### GetOwnerId

`func (o *AddonApplicationInfo) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *AddonApplicationInfo) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *AddonApplicationInfo) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *AddonApplicationInfo) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.

### GetOwnerName

`func (o *AddonApplicationInfo) GetOwnerName() string`

GetOwnerName returns the OwnerName field if non-nil, zero value otherwise.

### GetOwnerNameOk

`func (o *AddonApplicationInfo) GetOwnerNameOk() (*string, bool)`

GetOwnerNameOk returns a tuple with the OwnerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerName

`func (o *AddonApplicationInfo) SetOwnerName(v string)`

SetOwnerName sets OwnerName field to given value.

### HasOwnerName

`func (o *AddonApplicationInfo) HasOwnerName() bool`

HasOwnerName returns a boolean if a field has been set.

### GetOwnerEmails

`func (o *AddonApplicationInfo) GetOwnerEmails() []string`

GetOwnerEmails returns the OwnerEmails field if non-nil, zero value otherwise.

### GetOwnerEmailsOk

`func (o *AddonApplicationInfo) GetOwnerEmailsOk() (*[]string, bool)`

GetOwnerEmailsOk returns a tuple with the OwnerEmails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerEmails

`func (o *AddonApplicationInfo) SetOwnerEmails(v []string)`

SetOwnerEmails sets OwnerEmails field to given value.

### HasOwnerEmails

`func (o *AddonApplicationInfo) HasOwnerEmails() bool`

HasOwnerEmails returns a boolean if a field has been set.

### GetRegion

`func (o *AddonApplicationInfo) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *AddonApplicationInfo) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *AddonApplicationInfo) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *AddonApplicationInfo) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetDomains

`func (o *AddonApplicationInfo) GetDomains() []string`

GetDomains returns the Domains field if non-nil, zero value otherwise.

### GetDomainsOk

`func (o *AddonApplicationInfo) GetDomainsOk() (*[]string, bool)`

GetDomainsOk returns a tuple with the Domains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomains

`func (o *AddonApplicationInfo) SetDomains(v []string)`

SetDomains sets Domains field to given value.

### HasDomains

`func (o *AddonApplicationInfo) HasDomains() bool`

HasDomains returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


