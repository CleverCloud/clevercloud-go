# AddonPlanView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Slug** | Pointer to **string** |  | [optional] 
**Price** | Pointer to **float32** |  | [optional] 
**PriceId** | Pointer to **string** |  | [optional] 
**Features** | Pointer to [**[]AddonFeatureInstanceView**](AddonFeatureInstanceView.md) |  | [optional] 
**Zones** | Pointer to **[]string** |  | [optional] 

## Methods

### NewAddonPlanView

`func NewAddonPlanView() *AddonPlanView`

NewAddonPlanView instantiates a new AddonPlanView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddonPlanViewWithDefaults

`func NewAddonPlanViewWithDefaults() *AddonPlanView`

NewAddonPlanViewWithDefaults instantiates a new AddonPlanView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AddonPlanView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AddonPlanView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AddonPlanView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AddonPlanView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *AddonPlanView) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddonPlanView) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddonPlanView) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AddonPlanView) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSlug

`func (o *AddonPlanView) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *AddonPlanView) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *AddonPlanView) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *AddonPlanView) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetPrice

`func (o *AddonPlanView) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *AddonPlanView) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *AddonPlanView) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *AddonPlanView) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetPriceId

`func (o *AddonPlanView) GetPriceId() string`

GetPriceId returns the PriceId field if non-nil, zero value otherwise.

### GetPriceIdOk

`func (o *AddonPlanView) GetPriceIdOk() (*string, bool)`

GetPriceIdOk returns a tuple with the PriceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceId

`func (o *AddonPlanView) SetPriceId(v string)`

SetPriceId sets PriceId field to given value.

### HasPriceId

`func (o *AddonPlanView) HasPriceId() bool`

HasPriceId returns a boolean if a field has been set.

### GetFeatures

`func (o *AddonPlanView) GetFeatures() []AddonFeatureInstanceView`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *AddonPlanView) GetFeaturesOk() (*[]AddonFeatureInstanceView, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *AddonPlanView) SetFeatures(v []AddonFeatureInstanceView)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *AddonPlanView) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetZones

`func (o *AddonPlanView) GetZones() []string`

GetZones returns the Zones field if non-nil, zero value otherwise.

### GetZonesOk

`func (o *AddonPlanView) GetZonesOk() (*[]string, bool)`

GetZonesOk returns a tuple with the Zones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZones

`func (o *AddonPlanView) SetZones(v []string)`

SetZones sets Zones field to given value.

### HasZones

`func (o *AddonPlanView) HasZones() bool`

HasZones returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


