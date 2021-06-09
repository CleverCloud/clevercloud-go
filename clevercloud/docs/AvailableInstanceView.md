# AvailableInstanceView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Variant** | Pointer to [**InstanceVariantView**](InstanceVariantView.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**ComingSoon** | Pointer to **bool** |  | [optional] 
**MaxInstances** | Pointer to **int32** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**Deployments** | Pointer to **[]string** |  | [optional] 
**Flavors** | Pointer to [**[]FlavorView**](FlavorView.md) |  | [optional] 
**DefaultFlavor** | Pointer to [**FlavorView**](FlavorView.md) |  | [optional] 
**BuildFlavor** | Pointer to [**FlavorView**](FlavorView.md) |  | [optional] 

## Methods

### NewAvailableInstanceView

`func NewAvailableInstanceView() *AvailableInstanceView`

NewAvailableInstanceView instantiates a new AvailableInstanceView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAvailableInstanceViewWithDefaults

`func NewAvailableInstanceViewWithDefaults() *AvailableInstanceView`

NewAvailableInstanceViewWithDefaults instantiates a new AvailableInstanceView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AvailableInstanceView) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AvailableInstanceView) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AvailableInstanceView) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AvailableInstanceView) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVersion

`func (o *AvailableInstanceView) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *AvailableInstanceView) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *AvailableInstanceView) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *AvailableInstanceView) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetName

`func (o *AvailableInstanceView) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AvailableInstanceView) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AvailableInstanceView) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AvailableInstanceView) HasName() bool`

HasName returns a boolean if a field has been set.

### GetVariant

`func (o *AvailableInstanceView) GetVariant() InstanceVariantView`

GetVariant returns the Variant field if non-nil, zero value otherwise.

### GetVariantOk

`func (o *AvailableInstanceView) GetVariantOk() (*InstanceVariantView, bool)`

GetVariantOk returns a tuple with the Variant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariant

`func (o *AvailableInstanceView) SetVariant(v InstanceVariantView)`

SetVariant sets Variant field to given value.

### HasVariant

`func (o *AvailableInstanceView) HasVariant() bool`

HasVariant returns a boolean if a field has been set.

### GetDescription

`func (o *AvailableInstanceView) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AvailableInstanceView) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AvailableInstanceView) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AvailableInstanceView) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AvailableInstanceView) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AvailableInstanceView) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AvailableInstanceView) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *AvailableInstanceView) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetComingSoon

`func (o *AvailableInstanceView) GetComingSoon() bool`

GetComingSoon returns the ComingSoon field if non-nil, zero value otherwise.

### GetComingSoonOk

`func (o *AvailableInstanceView) GetComingSoonOk() (*bool, bool)`

GetComingSoonOk returns a tuple with the ComingSoon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComingSoon

`func (o *AvailableInstanceView) SetComingSoon(v bool)`

SetComingSoon sets ComingSoon field to given value.

### HasComingSoon

`func (o *AvailableInstanceView) HasComingSoon() bool`

HasComingSoon returns a boolean if a field has been set.

### GetMaxInstances

`func (o *AvailableInstanceView) GetMaxInstances() int32`

GetMaxInstances returns the MaxInstances field if non-nil, zero value otherwise.

### GetMaxInstancesOk

`func (o *AvailableInstanceView) GetMaxInstancesOk() (*int32, bool)`

GetMaxInstancesOk returns a tuple with the MaxInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxInstances

`func (o *AvailableInstanceView) SetMaxInstances(v int32)`

SetMaxInstances sets MaxInstances field to given value.

### HasMaxInstances

`func (o *AvailableInstanceView) HasMaxInstances() bool`

HasMaxInstances returns a boolean if a field has been set.

### GetTags

`func (o *AvailableInstanceView) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *AvailableInstanceView) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *AvailableInstanceView) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *AvailableInstanceView) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetDeployments

`func (o *AvailableInstanceView) GetDeployments() []string`

GetDeployments returns the Deployments field if non-nil, zero value otherwise.

### GetDeploymentsOk

`func (o *AvailableInstanceView) GetDeploymentsOk() (*[]string, bool)`

GetDeploymentsOk returns a tuple with the Deployments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployments

`func (o *AvailableInstanceView) SetDeployments(v []string)`

SetDeployments sets Deployments field to given value.

### HasDeployments

`func (o *AvailableInstanceView) HasDeployments() bool`

HasDeployments returns a boolean if a field has been set.

### GetFlavors

`func (o *AvailableInstanceView) GetFlavors() []FlavorView`

GetFlavors returns the Flavors field if non-nil, zero value otherwise.

### GetFlavorsOk

`func (o *AvailableInstanceView) GetFlavorsOk() (*[]FlavorView, bool)`

GetFlavorsOk returns a tuple with the Flavors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlavors

`func (o *AvailableInstanceView) SetFlavors(v []FlavorView)`

SetFlavors sets Flavors field to given value.

### HasFlavors

`func (o *AvailableInstanceView) HasFlavors() bool`

HasFlavors returns a boolean if a field has been set.

### GetDefaultFlavor

`func (o *AvailableInstanceView) GetDefaultFlavor() FlavorView`

GetDefaultFlavor returns the DefaultFlavor field if non-nil, zero value otherwise.

### GetDefaultFlavorOk

`func (o *AvailableInstanceView) GetDefaultFlavorOk() (*FlavorView, bool)`

GetDefaultFlavorOk returns a tuple with the DefaultFlavor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultFlavor

`func (o *AvailableInstanceView) SetDefaultFlavor(v FlavorView)`

SetDefaultFlavor sets DefaultFlavor field to given value.

### HasDefaultFlavor

`func (o *AvailableInstanceView) HasDefaultFlavor() bool`

HasDefaultFlavor returns a boolean if a field has been set.

### GetBuildFlavor

`func (o *AvailableInstanceView) GetBuildFlavor() FlavorView`

GetBuildFlavor returns the BuildFlavor field if non-nil, zero value otherwise.

### GetBuildFlavorOk

`func (o *AvailableInstanceView) GetBuildFlavorOk() (*FlavorView, bool)`

GetBuildFlavorOk returns a tuple with the BuildFlavor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildFlavor

`func (o *AvailableInstanceView) SetBuildFlavor(v FlavorView)`

SetBuildFlavor sets BuildFlavor field to given value.

### HasBuildFlavor

`func (o *AvailableInstanceView) HasBuildFlavor() bool`

HasBuildFlavor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


