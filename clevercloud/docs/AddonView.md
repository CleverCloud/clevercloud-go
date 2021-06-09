# AddonView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**RealId** | Pointer to **string** |  | [optional] 
**Region** | Pointer to **string** |  | [optional] 
**Provider** | Pointer to [**AddonProviderInfoView**](AddonProviderInfoView.md) |  | [optional] 
**Plan** | Pointer to [**AddonPlanView**](AddonPlanView.md) |  | [optional] 
**CreationDate** | Pointer to **int64** |  | [optional] 
**ConfigKeys** | Pointer to **[]string** |  | [optional] 

## Methods

### NewAddonView

`func NewAddonView() *AddonView`

NewAddonView instantiates a new AddonView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddonViewWithDefaults

`func NewAddonViewWithDefaults() *AddonView`

NewAddonViewWithDefaults instantiates a new AddonView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AddonView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AddonView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AddonView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AddonView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *AddonView) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddonView) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddonView) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AddonView) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRealId

`func (o *AddonView) GetRealId() string`

GetRealId returns the RealId field if non-nil, zero value otherwise.

### GetRealIdOk

`func (o *AddonView) GetRealIdOk() (*string, bool)`

GetRealIdOk returns a tuple with the RealId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealId

`func (o *AddonView) SetRealId(v string)`

SetRealId sets RealId field to given value.

### HasRealId

`func (o *AddonView) HasRealId() bool`

HasRealId returns a boolean if a field has been set.

### GetRegion

`func (o *AddonView) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *AddonView) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *AddonView) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *AddonView) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetProvider

`func (o *AddonView) GetProvider() AddonProviderInfoView`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *AddonView) GetProviderOk() (*AddonProviderInfoView, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *AddonView) SetProvider(v AddonProviderInfoView)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *AddonView) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetPlan

`func (o *AddonView) GetPlan() AddonPlanView`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *AddonView) GetPlanOk() (*AddonPlanView, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *AddonView) SetPlan(v AddonPlanView)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *AddonView) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetCreationDate

`func (o *AddonView) GetCreationDate() int64`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *AddonView) GetCreationDateOk() (*int64, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *AddonView) SetCreationDate(v int64)`

SetCreationDate sets CreationDate field to given value.

### HasCreationDate

`func (o *AddonView) HasCreationDate() bool`

HasCreationDate returns a boolean if a field has been set.

### GetConfigKeys

`func (o *AddonView) GetConfigKeys() []string`

GetConfigKeys returns the ConfigKeys field if non-nil, zero value otherwise.

### GetConfigKeysOk

`func (o *AddonView) GetConfigKeysOk() (*[]string, bool)`

GetConfigKeysOk returns a tuple with the ConfigKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigKeys

`func (o *AddonView) SetConfigKeys(v []string)`

SetConfigKeys sets ConfigKeys field to given value.

### HasConfigKeys

`func (o *AddonView) HasConfigKeys() bool`

HasConfigKeys returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


