# ApplicationView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Zone** | Pointer to **string** |  | [optional] 
**Instance** | Pointer to [**InstanceView**](InstanceView.md) |  | [optional] 
**Deployment** | Pointer to [**DeploymentInfoView**](DeploymentInfoView.md) |  | [optional] 
**Vhosts** | Pointer to [**[]VhostView**](VhostView.md) |  | [optional] 
**CreationDate** | Pointer to **int64** |  | [optional] 
**LastDeploy** | Pointer to **int32** |  | [optional] 
**Archived** | Pointer to **bool** |  | [optional] 
**StickySessions** | Pointer to **bool** |  | [optional] 
**Homogeneous** | Pointer to **bool** |  | [optional] 
**Favourite** | Pointer to **bool** |  | [optional] 
**CancelOnPush** | Pointer to **bool** |  | [optional] 
**WebhookUrl** | Pointer to **string** |  | [optional] 
**WebhookSecret** | Pointer to **string** |  | [optional] 
**SeparateBuild** | Pointer to **bool** |  | [optional] 
**BuildFlavor** | Pointer to [**FlavorView**](FlavorView.md) |  | [optional] 
**OwnerId** | Pointer to **string** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**CommitId** | Pointer to **string** |  | [optional] 
**Appliance** | Pointer to **string** |  | [optional] 
**Branch** | Pointer to **string** |  | [optional] 
**ForceHttps** | Pointer to **string** |  | [optional] 
**DeployUrl** | Pointer to **string** |  | [optional] 

## Methods

### NewApplicationView

`func NewApplicationView() *ApplicationView`

NewApplicationView instantiates a new ApplicationView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationViewWithDefaults

`func NewApplicationViewWithDefaults() *ApplicationView`

NewApplicationViewWithDefaults instantiates a new ApplicationView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ApplicationView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApplicationView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApplicationView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ApplicationView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ApplicationView) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApplicationView) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApplicationView) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApplicationView) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *ApplicationView) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApplicationView) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApplicationView) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ApplicationView) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetZone

`func (o *ApplicationView) GetZone() string`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *ApplicationView) GetZoneOk() (*string, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *ApplicationView) SetZone(v string)`

SetZone sets Zone field to given value.

### HasZone

`func (o *ApplicationView) HasZone() bool`

HasZone returns a boolean if a field has been set.

### GetInstance

`func (o *ApplicationView) GetInstance() InstanceView`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *ApplicationView) GetInstanceOk() (*InstanceView, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *ApplicationView) SetInstance(v InstanceView)`

SetInstance sets Instance field to given value.

### HasInstance

`func (o *ApplicationView) HasInstance() bool`

HasInstance returns a boolean if a field has been set.

### GetDeployment

`func (o *ApplicationView) GetDeployment() DeploymentInfoView`

GetDeployment returns the Deployment field if non-nil, zero value otherwise.

### GetDeploymentOk

`func (o *ApplicationView) GetDeploymentOk() (*DeploymentInfoView, bool)`

GetDeploymentOk returns a tuple with the Deployment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployment

`func (o *ApplicationView) SetDeployment(v DeploymentInfoView)`

SetDeployment sets Deployment field to given value.

### HasDeployment

`func (o *ApplicationView) HasDeployment() bool`

HasDeployment returns a boolean if a field has been set.

### GetVhosts

`func (o *ApplicationView) GetVhosts() []VhostView`

GetVhosts returns the Vhosts field if non-nil, zero value otherwise.

### GetVhostsOk

`func (o *ApplicationView) GetVhostsOk() (*[]VhostView, bool)`

GetVhostsOk returns a tuple with the Vhosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVhosts

`func (o *ApplicationView) SetVhosts(v []VhostView)`

SetVhosts sets Vhosts field to given value.

### HasVhosts

`func (o *ApplicationView) HasVhosts() bool`

HasVhosts returns a boolean if a field has been set.

### GetCreationDate

`func (o *ApplicationView) GetCreationDate() int64`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *ApplicationView) GetCreationDateOk() (*int64, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *ApplicationView) SetCreationDate(v int64)`

SetCreationDate sets CreationDate field to given value.

### HasCreationDate

`func (o *ApplicationView) HasCreationDate() bool`

HasCreationDate returns a boolean if a field has been set.

### GetLastDeploy

`func (o *ApplicationView) GetLastDeploy() int32`

GetLastDeploy returns the LastDeploy field if non-nil, zero value otherwise.

### GetLastDeployOk

`func (o *ApplicationView) GetLastDeployOk() (*int32, bool)`

GetLastDeployOk returns a tuple with the LastDeploy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastDeploy

`func (o *ApplicationView) SetLastDeploy(v int32)`

SetLastDeploy sets LastDeploy field to given value.

### HasLastDeploy

`func (o *ApplicationView) HasLastDeploy() bool`

HasLastDeploy returns a boolean if a field has been set.

### GetArchived

`func (o *ApplicationView) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *ApplicationView) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *ApplicationView) SetArchived(v bool)`

SetArchived sets Archived field to given value.

### HasArchived

`func (o *ApplicationView) HasArchived() bool`

HasArchived returns a boolean if a field has been set.

### GetStickySessions

`func (o *ApplicationView) GetStickySessions() bool`

GetStickySessions returns the StickySessions field if non-nil, zero value otherwise.

### GetStickySessionsOk

`func (o *ApplicationView) GetStickySessionsOk() (*bool, bool)`

GetStickySessionsOk returns a tuple with the StickySessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStickySessions

`func (o *ApplicationView) SetStickySessions(v bool)`

SetStickySessions sets StickySessions field to given value.

### HasStickySessions

`func (o *ApplicationView) HasStickySessions() bool`

HasStickySessions returns a boolean if a field has been set.

### GetHomogeneous

`func (o *ApplicationView) GetHomogeneous() bool`

GetHomogeneous returns the Homogeneous field if non-nil, zero value otherwise.

### GetHomogeneousOk

`func (o *ApplicationView) GetHomogeneousOk() (*bool, bool)`

GetHomogeneousOk returns a tuple with the Homogeneous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomogeneous

`func (o *ApplicationView) SetHomogeneous(v bool)`

SetHomogeneous sets Homogeneous field to given value.

### HasHomogeneous

`func (o *ApplicationView) HasHomogeneous() bool`

HasHomogeneous returns a boolean if a field has been set.

### GetFavourite

`func (o *ApplicationView) GetFavourite() bool`

GetFavourite returns the Favourite field if non-nil, zero value otherwise.

### GetFavouriteOk

`func (o *ApplicationView) GetFavouriteOk() (*bool, bool)`

GetFavouriteOk returns a tuple with the Favourite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFavourite

`func (o *ApplicationView) SetFavourite(v bool)`

SetFavourite sets Favourite field to given value.

### HasFavourite

`func (o *ApplicationView) HasFavourite() bool`

HasFavourite returns a boolean if a field has been set.

### GetCancelOnPush

`func (o *ApplicationView) GetCancelOnPush() bool`

GetCancelOnPush returns the CancelOnPush field if non-nil, zero value otherwise.

### GetCancelOnPushOk

`func (o *ApplicationView) GetCancelOnPushOk() (*bool, bool)`

GetCancelOnPushOk returns a tuple with the CancelOnPush field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelOnPush

`func (o *ApplicationView) SetCancelOnPush(v bool)`

SetCancelOnPush sets CancelOnPush field to given value.

### HasCancelOnPush

`func (o *ApplicationView) HasCancelOnPush() bool`

HasCancelOnPush returns a boolean if a field has been set.

### GetWebhookUrl

`func (o *ApplicationView) GetWebhookUrl() string`

GetWebhookUrl returns the WebhookUrl field if non-nil, zero value otherwise.

### GetWebhookUrlOk

`func (o *ApplicationView) GetWebhookUrlOk() (*string, bool)`

GetWebhookUrlOk returns a tuple with the WebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookUrl

`func (o *ApplicationView) SetWebhookUrl(v string)`

SetWebhookUrl sets WebhookUrl field to given value.

### HasWebhookUrl

`func (o *ApplicationView) HasWebhookUrl() bool`

HasWebhookUrl returns a boolean if a field has been set.

### GetWebhookSecret

`func (o *ApplicationView) GetWebhookSecret() string`

GetWebhookSecret returns the WebhookSecret field if non-nil, zero value otherwise.

### GetWebhookSecretOk

`func (o *ApplicationView) GetWebhookSecretOk() (*string, bool)`

GetWebhookSecretOk returns a tuple with the WebhookSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookSecret

`func (o *ApplicationView) SetWebhookSecret(v string)`

SetWebhookSecret sets WebhookSecret field to given value.

### HasWebhookSecret

`func (o *ApplicationView) HasWebhookSecret() bool`

HasWebhookSecret returns a boolean if a field has been set.

### GetSeparateBuild

`func (o *ApplicationView) GetSeparateBuild() bool`

GetSeparateBuild returns the SeparateBuild field if non-nil, zero value otherwise.

### GetSeparateBuildOk

`func (o *ApplicationView) GetSeparateBuildOk() (*bool, bool)`

GetSeparateBuildOk returns a tuple with the SeparateBuild field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeparateBuild

`func (o *ApplicationView) SetSeparateBuild(v bool)`

SetSeparateBuild sets SeparateBuild field to given value.

### HasSeparateBuild

`func (o *ApplicationView) HasSeparateBuild() bool`

HasSeparateBuild returns a boolean if a field has been set.

### GetBuildFlavor

`func (o *ApplicationView) GetBuildFlavor() FlavorView`

GetBuildFlavor returns the BuildFlavor field if non-nil, zero value otherwise.

### GetBuildFlavorOk

`func (o *ApplicationView) GetBuildFlavorOk() (*FlavorView, bool)`

GetBuildFlavorOk returns a tuple with the BuildFlavor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildFlavor

`func (o *ApplicationView) SetBuildFlavor(v FlavorView)`

SetBuildFlavor sets BuildFlavor field to given value.

### HasBuildFlavor

`func (o *ApplicationView) HasBuildFlavor() bool`

HasBuildFlavor returns a boolean if a field has been set.

### GetOwnerId

`func (o *ApplicationView) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *ApplicationView) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *ApplicationView) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *ApplicationView) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.

### GetState

`func (o *ApplicationView) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ApplicationView) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ApplicationView) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *ApplicationView) HasState() bool`

HasState returns a boolean if a field has been set.

### GetCommitId

`func (o *ApplicationView) GetCommitId() string`

GetCommitId returns the CommitId field if non-nil, zero value otherwise.

### GetCommitIdOk

`func (o *ApplicationView) GetCommitIdOk() (*string, bool)`

GetCommitIdOk returns a tuple with the CommitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitId

`func (o *ApplicationView) SetCommitId(v string)`

SetCommitId sets CommitId field to given value.

### HasCommitId

`func (o *ApplicationView) HasCommitId() bool`

HasCommitId returns a boolean if a field has been set.

### GetAppliance

`func (o *ApplicationView) GetAppliance() string`

GetAppliance returns the Appliance field if non-nil, zero value otherwise.

### GetApplianceOk

`func (o *ApplicationView) GetApplianceOk() (*string, bool)`

GetApplianceOk returns a tuple with the Appliance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliance

`func (o *ApplicationView) SetAppliance(v string)`

SetAppliance sets Appliance field to given value.

### HasAppliance

`func (o *ApplicationView) HasAppliance() bool`

HasAppliance returns a boolean if a field has been set.

### GetBranch

`func (o *ApplicationView) GetBranch() string`

GetBranch returns the Branch field if non-nil, zero value otherwise.

### GetBranchOk

`func (o *ApplicationView) GetBranchOk() (*string, bool)`

GetBranchOk returns a tuple with the Branch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranch

`func (o *ApplicationView) SetBranch(v string)`

SetBranch sets Branch field to given value.

### HasBranch

`func (o *ApplicationView) HasBranch() bool`

HasBranch returns a boolean if a field has been set.

### GetForceHttps

`func (o *ApplicationView) GetForceHttps() string`

GetForceHttps returns the ForceHttps field if non-nil, zero value otherwise.

### GetForceHttpsOk

`func (o *ApplicationView) GetForceHttpsOk() (*string, bool)`

GetForceHttpsOk returns a tuple with the ForceHttps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceHttps

`func (o *ApplicationView) SetForceHttps(v string)`

SetForceHttps sets ForceHttps field to given value.

### HasForceHttps

`func (o *ApplicationView) HasForceHttps() bool`

HasForceHttps returns a boolean if a field has been set.

### GetDeployUrl

`func (o *ApplicationView) GetDeployUrl() string`

GetDeployUrl returns the DeployUrl field if non-nil, zero value otherwise.

### GetDeployUrlOk

`func (o *ApplicationView) GetDeployUrlOk() (*string, bool)`

GetDeployUrlOk returns a tuple with the DeployUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployUrl

`func (o *ApplicationView) SetDeployUrl(v string)`

SetDeployUrl sets DeployUrl field to given value.

### HasDeployUrl

`func (o *ApplicationView) HasDeployUrl() bool`

HasDeployUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


