# WannabeAddonPlan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Slug** | Pointer to **string** |  | [optional] 
**Price** | Pointer to **float32** |  | [optional] 
**Features** | Pointer to [**[]AddonFeatureInstanceView**](AddonFeatureInstanceView.md) |  | [optional] 

## Methods

### NewWannabeAddonPlan

`func NewWannabeAddonPlan() *WannabeAddonPlan`

NewWannabeAddonPlan instantiates a new WannabeAddonPlan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWannabeAddonPlanWithDefaults

`func NewWannabeAddonPlanWithDefaults() *WannabeAddonPlan`

NewWannabeAddonPlanWithDefaults instantiates a new WannabeAddonPlan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *WannabeAddonPlan) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WannabeAddonPlan) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WannabeAddonPlan) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WannabeAddonPlan) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSlug

`func (o *WannabeAddonPlan) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *WannabeAddonPlan) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *WannabeAddonPlan) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *WannabeAddonPlan) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetPrice

`func (o *WannabeAddonPlan) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *WannabeAddonPlan) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *WannabeAddonPlan) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *WannabeAddonPlan) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetFeatures

`func (o *WannabeAddonPlan) GetFeatures() []AddonFeatureInstanceView`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *WannabeAddonPlan) GetFeaturesOk() (*[]AddonFeatureInstanceView, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *WannabeAddonPlan) SetFeatures(v []AddonFeatureInstanceView)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *WannabeAddonPlan) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


