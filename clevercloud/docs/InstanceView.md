# InstanceView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 
**Variant** | Pointer to [**InstanceVariantView**](InstanceVariantView.md) |  | [optional] 
**MinInstances** | Pointer to **int32** |  | [optional] 
**MaxInstances** | Pointer to **int32** |  | [optional] 
**MaxAllowedInstances** | Pointer to **int32** |  | [optional] 
**MinFlavor** | Pointer to [**FlavorView**](FlavorView.md) |  | [optional] 
**MaxFlavor** | Pointer to [**FlavorView**](FlavorView.md) |  | [optional] 
**Flavors** | Pointer to [**[]FlavorView**](FlavorView.md) |  | [optional] 
**DefaultEnv** | Pointer to **map[string]string** |  | [optional] 
**Lifetime** | Pointer to **string** |  | [optional] 
**InstanceAndVersion** | Pointer to **string** |  | [optional] 

## Methods

### NewInstanceView

`func NewInstanceView() *InstanceView`

NewInstanceView instantiates a new InstanceView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceViewWithDefaults

`func NewInstanceViewWithDefaults() *InstanceView`

NewInstanceViewWithDefaults instantiates a new InstanceView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *InstanceView) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InstanceView) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InstanceView) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *InstanceView) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVersion

`func (o *InstanceView) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *InstanceView) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *InstanceView) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *InstanceView) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetVariant

`func (o *InstanceView) GetVariant() InstanceVariantView`

GetVariant returns the Variant field if non-nil, zero value otherwise.

### GetVariantOk

`func (o *InstanceView) GetVariantOk() (*InstanceVariantView, bool)`

GetVariantOk returns a tuple with the Variant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariant

`func (o *InstanceView) SetVariant(v InstanceVariantView)`

SetVariant sets Variant field to given value.

### HasVariant

`func (o *InstanceView) HasVariant() bool`

HasVariant returns a boolean if a field has been set.

### GetMinInstances

`func (o *InstanceView) GetMinInstances() int32`

GetMinInstances returns the MinInstances field if non-nil, zero value otherwise.

### GetMinInstancesOk

`func (o *InstanceView) GetMinInstancesOk() (*int32, bool)`

GetMinInstancesOk returns a tuple with the MinInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinInstances

`func (o *InstanceView) SetMinInstances(v int32)`

SetMinInstances sets MinInstances field to given value.

### HasMinInstances

`func (o *InstanceView) HasMinInstances() bool`

HasMinInstances returns a boolean if a field has been set.

### GetMaxInstances

`func (o *InstanceView) GetMaxInstances() int32`

GetMaxInstances returns the MaxInstances field if non-nil, zero value otherwise.

### GetMaxInstancesOk

`func (o *InstanceView) GetMaxInstancesOk() (*int32, bool)`

GetMaxInstancesOk returns a tuple with the MaxInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxInstances

`func (o *InstanceView) SetMaxInstances(v int32)`

SetMaxInstances sets MaxInstances field to given value.

### HasMaxInstances

`func (o *InstanceView) HasMaxInstances() bool`

HasMaxInstances returns a boolean if a field has been set.

### GetMaxAllowedInstances

`func (o *InstanceView) GetMaxAllowedInstances() int32`

GetMaxAllowedInstances returns the MaxAllowedInstances field if non-nil, zero value otherwise.

### GetMaxAllowedInstancesOk

`func (o *InstanceView) GetMaxAllowedInstancesOk() (*int32, bool)`

GetMaxAllowedInstancesOk returns a tuple with the MaxAllowedInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAllowedInstances

`func (o *InstanceView) SetMaxAllowedInstances(v int32)`

SetMaxAllowedInstances sets MaxAllowedInstances field to given value.

### HasMaxAllowedInstances

`func (o *InstanceView) HasMaxAllowedInstances() bool`

HasMaxAllowedInstances returns a boolean if a field has been set.

### GetMinFlavor

`func (o *InstanceView) GetMinFlavor() FlavorView`

GetMinFlavor returns the MinFlavor field if non-nil, zero value otherwise.

### GetMinFlavorOk

`func (o *InstanceView) GetMinFlavorOk() (*FlavorView, bool)`

GetMinFlavorOk returns a tuple with the MinFlavor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinFlavor

`func (o *InstanceView) SetMinFlavor(v FlavorView)`

SetMinFlavor sets MinFlavor field to given value.

### HasMinFlavor

`func (o *InstanceView) HasMinFlavor() bool`

HasMinFlavor returns a boolean if a field has been set.

### GetMaxFlavor

`func (o *InstanceView) GetMaxFlavor() FlavorView`

GetMaxFlavor returns the MaxFlavor field if non-nil, zero value otherwise.

### GetMaxFlavorOk

`func (o *InstanceView) GetMaxFlavorOk() (*FlavorView, bool)`

GetMaxFlavorOk returns a tuple with the MaxFlavor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxFlavor

`func (o *InstanceView) SetMaxFlavor(v FlavorView)`

SetMaxFlavor sets MaxFlavor field to given value.

### HasMaxFlavor

`func (o *InstanceView) HasMaxFlavor() bool`

HasMaxFlavor returns a boolean if a field has been set.

### GetFlavors

`func (o *InstanceView) GetFlavors() []FlavorView`

GetFlavors returns the Flavors field if non-nil, zero value otherwise.

### GetFlavorsOk

`func (o *InstanceView) GetFlavorsOk() (*[]FlavorView, bool)`

GetFlavorsOk returns a tuple with the Flavors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlavors

`func (o *InstanceView) SetFlavors(v []FlavorView)`

SetFlavors sets Flavors field to given value.

### HasFlavors

`func (o *InstanceView) HasFlavors() bool`

HasFlavors returns a boolean if a field has been set.

### GetDefaultEnv

`func (o *InstanceView) GetDefaultEnv() map[string]string`

GetDefaultEnv returns the DefaultEnv field if non-nil, zero value otherwise.

### GetDefaultEnvOk

`func (o *InstanceView) GetDefaultEnvOk() (*map[string]string, bool)`

GetDefaultEnvOk returns a tuple with the DefaultEnv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultEnv

`func (o *InstanceView) SetDefaultEnv(v map[string]string)`

SetDefaultEnv sets DefaultEnv field to given value.

### HasDefaultEnv

`func (o *InstanceView) HasDefaultEnv() bool`

HasDefaultEnv returns a boolean if a field has been set.

### GetLifetime

`func (o *InstanceView) GetLifetime() string`

GetLifetime returns the Lifetime field if non-nil, zero value otherwise.

### GetLifetimeOk

`func (o *InstanceView) GetLifetimeOk() (*string, bool)`

GetLifetimeOk returns a tuple with the Lifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifetime

`func (o *InstanceView) SetLifetime(v string)`

SetLifetime sets Lifetime field to given value.

### HasLifetime

`func (o *InstanceView) HasLifetime() bool`

HasLifetime returns a boolean if a field has been set.

### GetInstanceAndVersion

`func (o *InstanceView) GetInstanceAndVersion() string`

GetInstanceAndVersion returns the InstanceAndVersion field if non-nil, zero value otherwise.

### GetInstanceAndVersionOk

`func (o *InstanceView) GetInstanceAndVersionOk() (*string, bool)`

GetInstanceAndVersionOk returns a tuple with the InstanceAndVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceAndVersion

`func (o *InstanceView) SetInstanceAndVersion(v string)`

SetInstanceAndVersion sets InstanceAndVersion field to given value.

### HasInstanceAndVersion

`func (o *InstanceView) HasInstanceAndVersion() bool`

HasInstanceAndVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


