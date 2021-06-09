# WannabeApplication

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Zone** | Pointer to **string** |  | [optional] 
**Deploy** | Pointer to **string** |  | [optional] 
**Shutdownable** | Pointer to **bool** |  | [optional] 
**InstanceType** | Pointer to **string** |  | [optional] 
**InstanceVersion** | Pointer to **string** |  | [optional] 
**InstanceVariant** | Pointer to **string** |  | [optional] 
**InstanceLifetime** | Pointer to **string** |  | [optional] 
**MinInstances** | Pointer to **int32** |  | [optional] 
**MaxInstances** | Pointer to **int32** |  | [optional] 
**MinFlavor** | Pointer to **string** |  | [optional] 
**MaxFlavor** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**Archived** | Pointer to **bool** |  | [optional] 
**StickySessions** | Pointer to **bool** |  | [optional] 
**Homogeneous** | Pointer to **bool** |  | [optional] 
**Favourite** | Pointer to **bool** |  | [optional] 
**CancelOnPush** | Pointer to **bool** |  | [optional] 
**SeparateBuild** | Pointer to **bool** |  | [optional] 
**BuildFlavor** | Pointer to **string** |  | [optional] 
**OauthService** | Pointer to **string** |  | [optional] 
**OauthAppId** | Pointer to **string** |  | [optional] 
**OauthApp** | Pointer to [**WannabeOauthApp**](WannabeOauthApp.md) |  | [optional] 
**ApplianceId** | Pointer to **string** |  | [optional] 
**Branch** | Pointer to **string** |  | [optional] 
**ForceHttps** | Pointer to **string** |  | [optional] 

## Methods

### NewWannabeApplication

`func NewWannabeApplication() *WannabeApplication`

NewWannabeApplication instantiates a new WannabeApplication object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWannabeApplicationWithDefaults

`func NewWannabeApplicationWithDefaults() *WannabeApplication`

NewWannabeApplicationWithDefaults instantiates a new WannabeApplication object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *WannabeApplication) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WannabeApplication) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WannabeApplication) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WannabeApplication) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *WannabeApplication) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WannabeApplication) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WannabeApplication) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WannabeApplication) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetZone

`func (o *WannabeApplication) GetZone() string`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *WannabeApplication) GetZoneOk() (*string, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *WannabeApplication) SetZone(v string)`

SetZone sets Zone field to given value.

### HasZone

`func (o *WannabeApplication) HasZone() bool`

HasZone returns a boolean if a field has been set.

### GetDeploy

`func (o *WannabeApplication) GetDeploy() string`

GetDeploy returns the Deploy field if non-nil, zero value otherwise.

### GetDeployOk

`func (o *WannabeApplication) GetDeployOk() (*string, bool)`

GetDeployOk returns a tuple with the Deploy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploy

`func (o *WannabeApplication) SetDeploy(v string)`

SetDeploy sets Deploy field to given value.

### HasDeploy

`func (o *WannabeApplication) HasDeploy() bool`

HasDeploy returns a boolean if a field has been set.

### GetShutdownable

`func (o *WannabeApplication) GetShutdownable() bool`

GetShutdownable returns the Shutdownable field if non-nil, zero value otherwise.

### GetShutdownableOk

`func (o *WannabeApplication) GetShutdownableOk() (*bool, bool)`

GetShutdownableOk returns a tuple with the Shutdownable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShutdownable

`func (o *WannabeApplication) SetShutdownable(v bool)`

SetShutdownable sets Shutdownable field to given value.

### HasShutdownable

`func (o *WannabeApplication) HasShutdownable() bool`

HasShutdownable returns a boolean if a field has been set.

### GetInstanceType

`func (o *WannabeApplication) GetInstanceType() string`

GetInstanceType returns the InstanceType field if non-nil, zero value otherwise.

### GetInstanceTypeOk

`func (o *WannabeApplication) GetInstanceTypeOk() (*string, bool)`

GetInstanceTypeOk returns a tuple with the InstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceType

`func (o *WannabeApplication) SetInstanceType(v string)`

SetInstanceType sets InstanceType field to given value.

### HasInstanceType

`func (o *WannabeApplication) HasInstanceType() bool`

HasInstanceType returns a boolean if a field has been set.

### GetInstanceVersion

`func (o *WannabeApplication) GetInstanceVersion() string`

GetInstanceVersion returns the InstanceVersion field if non-nil, zero value otherwise.

### GetInstanceVersionOk

`func (o *WannabeApplication) GetInstanceVersionOk() (*string, bool)`

GetInstanceVersionOk returns a tuple with the InstanceVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceVersion

`func (o *WannabeApplication) SetInstanceVersion(v string)`

SetInstanceVersion sets InstanceVersion field to given value.

### HasInstanceVersion

`func (o *WannabeApplication) HasInstanceVersion() bool`

HasInstanceVersion returns a boolean if a field has been set.

### GetInstanceVariant

`func (o *WannabeApplication) GetInstanceVariant() string`

GetInstanceVariant returns the InstanceVariant field if non-nil, zero value otherwise.

### GetInstanceVariantOk

`func (o *WannabeApplication) GetInstanceVariantOk() (*string, bool)`

GetInstanceVariantOk returns a tuple with the InstanceVariant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceVariant

`func (o *WannabeApplication) SetInstanceVariant(v string)`

SetInstanceVariant sets InstanceVariant field to given value.

### HasInstanceVariant

`func (o *WannabeApplication) HasInstanceVariant() bool`

HasInstanceVariant returns a boolean if a field has been set.

### GetInstanceLifetime

`func (o *WannabeApplication) GetInstanceLifetime() string`

GetInstanceLifetime returns the InstanceLifetime field if non-nil, zero value otherwise.

### GetInstanceLifetimeOk

`func (o *WannabeApplication) GetInstanceLifetimeOk() (*string, bool)`

GetInstanceLifetimeOk returns a tuple with the InstanceLifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceLifetime

`func (o *WannabeApplication) SetInstanceLifetime(v string)`

SetInstanceLifetime sets InstanceLifetime field to given value.

### HasInstanceLifetime

`func (o *WannabeApplication) HasInstanceLifetime() bool`

HasInstanceLifetime returns a boolean if a field has been set.

### GetMinInstances

`func (o *WannabeApplication) GetMinInstances() int32`

GetMinInstances returns the MinInstances field if non-nil, zero value otherwise.

### GetMinInstancesOk

`func (o *WannabeApplication) GetMinInstancesOk() (*int32, bool)`

GetMinInstancesOk returns a tuple with the MinInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinInstances

`func (o *WannabeApplication) SetMinInstances(v int32)`

SetMinInstances sets MinInstances field to given value.

### HasMinInstances

`func (o *WannabeApplication) HasMinInstances() bool`

HasMinInstances returns a boolean if a field has been set.

### GetMaxInstances

`func (o *WannabeApplication) GetMaxInstances() int32`

GetMaxInstances returns the MaxInstances field if non-nil, zero value otherwise.

### GetMaxInstancesOk

`func (o *WannabeApplication) GetMaxInstancesOk() (*int32, bool)`

GetMaxInstancesOk returns a tuple with the MaxInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxInstances

`func (o *WannabeApplication) SetMaxInstances(v int32)`

SetMaxInstances sets MaxInstances field to given value.

### HasMaxInstances

`func (o *WannabeApplication) HasMaxInstances() bool`

HasMaxInstances returns a boolean if a field has been set.

### GetMinFlavor

`func (o *WannabeApplication) GetMinFlavor() string`

GetMinFlavor returns the MinFlavor field if non-nil, zero value otherwise.

### GetMinFlavorOk

`func (o *WannabeApplication) GetMinFlavorOk() (*string, bool)`

GetMinFlavorOk returns a tuple with the MinFlavor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinFlavor

`func (o *WannabeApplication) SetMinFlavor(v string)`

SetMinFlavor sets MinFlavor field to given value.

### HasMinFlavor

`func (o *WannabeApplication) HasMinFlavor() bool`

HasMinFlavor returns a boolean if a field has been set.

### GetMaxFlavor

`func (o *WannabeApplication) GetMaxFlavor() string`

GetMaxFlavor returns the MaxFlavor field if non-nil, zero value otherwise.

### GetMaxFlavorOk

`func (o *WannabeApplication) GetMaxFlavorOk() (*string, bool)`

GetMaxFlavorOk returns a tuple with the MaxFlavor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxFlavor

`func (o *WannabeApplication) SetMaxFlavor(v string)`

SetMaxFlavor sets MaxFlavor field to given value.

### HasMaxFlavor

`func (o *WannabeApplication) HasMaxFlavor() bool`

HasMaxFlavor returns a boolean if a field has been set.

### GetTags

`func (o *WannabeApplication) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *WannabeApplication) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *WannabeApplication) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *WannabeApplication) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetArchived

`func (o *WannabeApplication) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *WannabeApplication) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *WannabeApplication) SetArchived(v bool)`

SetArchived sets Archived field to given value.

### HasArchived

`func (o *WannabeApplication) HasArchived() bool`

HasArchived returns a boolean if a field has been set.

### GetStickySessions

`func (o *WannabeApplication) GetStickySessions() bool`

GetStickySessions returns the StickySessions field if non-nil, zero value otherwise.

### GetStickySessionsOk

`func (o *WannabeApplication) GetStickySessionsOk() (*bool, bool)`

GetStickySessionsOk returns a tuple with the StickySessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStickySessions

`func (o *WannabeApplication) SetStickySessions(v bool)`

SetStickySessions sets StickySessions field to given value.

### HasStickySessions

`func (o *WannabeApplication) HasStickySessions() bool`

HasStickySessions returns a boolean if a field has been set.

### GetHomogeneous

`func (o *WannabeApplication) GetHomogeneous() bool`

GetHomogeneous returns the Homogeneous field if non-nil, zero value otherwise.

### GetHomogeneousOk

`func (o *WannabeApplication) GetHomogeneousOk() (*bool, bool)`

GetHomogeneousOk returns a tuple with the Homogeneous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomogeneous

`func (o *WannabeApplication) SetHomogeneous(v bool)`

SetHomogeneous sets Homogeneous field to given value.

### HasHomogeneous

`func (o *WannabeApplication) HasHomogeneous() bool`

HasHomogeneous returns a boolean if a field has been set.

### GetFavourite

`func (o *WannabeApplication) GetFavourite() bool`

GetFavourite returns the Favourite field if non-nil, zero value otherwise.

### GetFavouriteOk

`func (o *WannabeApplication) GetFavouriteOk() (*bool, bool)`

GetFavouriteOk returns a tuple with the Favourite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFavourite

`func (o *WannabeApplication) SetFavourite(v bool)`

SetFavourite sets Favourite field to given value.

### HasFavourite

`func (o *WannabeApplication) HasFavourite() bool`

HasFavourite returns a boolean if a field has been set.

### GetCancelOnPush

`func (o *WannabeApplication) GetCancelOnPush() bool`

GetCancelOnPush returns the CancelOnPush field if non-nil, zero value otherwise.

### GetCancelOnPushOk

`func (o *WannabeApplication) GetCancelOnPushOk() (*bool, bool)`

GetCancelOnPushOk returns a tuple with the CancelOnPush field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelOnPush

`func (o *WannabeApplication) SetCancelOnPush(v bool)`

SetCancelOnPush sets CancelOnPush field to given value.

### HasCancelOnPush

`func (o *WannabeApplication) HasCancelOnPush() bool`

HasCancelOnPush returns a boolean if a field has been set.

### GetSeparateBuild

`func (o *WannabeApplication) GetSeparateBuild() bool`

GetSeparateBuild returns the SeparateBuild field if non-nil, zero value otherwise.

### GetSeparateBuildOk

`func (o *WannabeApplication) GetSeparateBuildOk() (*bool, bool)`

GetSeparateBuildOk returns a tuple with the SeparateBuild field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeparateBuild

`func (o *WannabeApplication) SetSeparateBuild(v bool)`

SetSeparateBuild sets SeparateBuild field to given value.

### HasSeparateBuild

`func (o *WannabeApplication) HasSeparateBuild() bool`

HasSeparateBuild returns a boolean if a field has been set.

### GetBuildFlavor

`func (o *WannabeApplication) GetBuildFlavor() string`

GetBuildFlavor returns the BuildFlavor field if non-nil, zero value otherwise.

### GetBuildFlavorOk

`func (o *WannabeApplication) GetBuildFlavorOk() (*string, bool)`

GetBuildFlavorOk returns a tuple with the BuildFlavor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildFlavor

`func (o *WannabeApplication) SetBuildFlavor(v string)`

SetBuildFlavor sets BuildFlavor field to given value.

### HasBuildFlavor

`func (o *WannabeApplication) HasBuildFlavor() bool`

HasBuildFlavor returns a boolean if a field has been set.

### GetOauthService

`func (o *WannabeApplication) GetOauthService() string`

GetOauthService returns the OauthService field if non-nil, zero value otherwise.

### GetOauthServiceOk

`func (o *WannabeApplication) GetOauthServiceOk() (*string, bool)`

GetOauthServiceOk returns a tuple with the OauthService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauthService

`func (o *WannabeApplication) SetOauthService(v string)`

SetOauthService sets OauthService field to given value.

### HasOauthService

`func (o *WannabeApplication) HasOauthService() bool`

HasOauthService returns a boolean if a field has been set.

### GetOauthAppId

`func (o *WannabeApplication) GetOauthAppId() string`

GetOauthAppId returns the OauthAppId field if non-nil, zero value otherwise.

### GetOauthAppIdOk

`func (o *WannabeApplication) GetOauthAppIdOk() (*string, bool)`

GetOauthAppIdOk returns a tuple with the OauthAppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauthAppId

`func (o *WannabeApplication) SetOauthAppId(v string)`

SetOauthAppId sets OauthAppId field to given value.

### HasOauthAppId

`func (o *WannabeApplication) HasOauthAppId() bool`

HasOauthAppId returns a boolean if a field has been set.

### GetOauthApp

`func (o *WannabeApplication) GetOauthApp() WannabeOauthApp`

GetOauthApp returns the OauthApp field if non-nil, zero value otherwise.

### GetOauthAppOk

`func (o *WannabeApplication) GetOauthAppOk() (*WannabeOauthApp, bool)`

GetOauthAppOk returns a tuple with the OauthApp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauthApp

`func (o *WannabeApplication) SetOauthApp(v WannabeOauthApp)`

SetOauthApp sets OauthApp field to given value.

### HasOauthApp

`func (o *WannabeApplication) HasOauthApp() bool`

HasOauthApp returns a boolean if a field has been set.

### GetApplianceId

`func (o *WannabeApplication) GetApplianceId() string`

GetApplianceId returns the ApplianceId field if non-nil, zero value otherwise.

### GetApplianceIdOk

`func (o *WannabeApplication) GetApplianceIdOk() (*string, bool)`

GetApplianceIdOk returns a tuple with the ApplianceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplianceId

`func (o *WannabeApplication) SetApplianceId(v string)`

SetApplianceId sets ApplianceId field to given value.

### HasApplianceId

`func (o *WannabeApplication) HasApplianceId() bool`

HasApplianceId returns a boolean if a field has been set.

### GetBranch

`func (o *WannabeApplication) GetBranch() string`

GetBranch returns the Branch field if non-nil, zero value otherwise.

### GetBranchOk

`func (o *WannabeApplication) GetBranchOk() (*string, bool)`

GetBranchOk returns a tuple with the Branch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranch

`func (o *WannabeApplication) SetBranch(v string)`

SetBranch sets Branch field to given value.

### HasBranch

`func (o *WannabeApplication) HasBranch() bool`

HasBranch returns a boolean if a field has been set.

### GetForceHttps

`func (o *WannabeApplication) GetForceHttps() string`

GetForceHttps returns the ForceHttps field if non-nil, zero value otherwise.

### GetForceHttpsOk

`func (o *WannabeApplication) GetForceHttpsOk() (*string, bool)`

GetForceHttpsOk returns a tuple with the ForceHttps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceHttps

`func (o *WannabeApplication) SetForceHttps(v string)`

SetForceHttps sets ForceHttps field to given value.

### HasForceHttps

`func (o *WannabeApplication) HasForceHttps() bool`

HasForceHttps returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


